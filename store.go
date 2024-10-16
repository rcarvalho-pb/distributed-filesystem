package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLen := len(hashStr) / blockSize
	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		Pathname: filepath.Join(paths...),
		Filename: hashStr,
	}
}

type PathKey struct {
	Pathname string
	Filename string
}

func (p PathKey) FullPathToFile() string {
	return fmt.Sprintf("%s/%s", p.Pathname, p.Filename)
}

type PathTransformFunc func(string) PathKey

var DefaultPathTransformFunc = func(key string) string {
	return key
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) has(key string) bool {
	pathKey := s.PathTransformFunc(key)

	_, err := os.Stat(pathKey.FullPathToFile())
	if err != nil {
		return false
	}

	return true
}

func (s *Store) Delete(key string) error {
	pathKey := s.PathTransformFunc(key)

	defer func() {
		log.Printf("deleted [%s] from disk\n", pathKey.Filename)
	}()

	return os.RemoveAll(pathKey.FullPathToFile()[:6])
}

func (s *Store) Read(key string) (io.Reader, error) {
	f, err := s.readStream(key)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := new(bytes.Buffer)

	_, err = io.Copy(buf, f)

	return buf, err
}

func (s *Store) readStream(key string) (io.ReadCloser, error) {
	pathKey := s.PathTransformFunc(key)

	return os.Open(pathKey.FullPathToFile())
}

func (s *Store) WriteStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathKey.Pathname, os.ModePerm); err != nil {
		return err
	}

	fullPath := pathKey.FullPathToFile()

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	log.Printf("written (%d) bytes to disk: %s", n, fullPath)

	return nil
}
