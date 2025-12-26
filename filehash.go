package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type FileHash struct {
	FilePath    string
	FileSize    int64
	FileName    string
	FileValue   []byte
	FileModTime time.Time
	mu          sync.Mutex
}

// NewFileHash creates a new FileHash instance
func NewFileHash() *FileHash {
	return &FileHash{}
}

// Read File content
func ReadTextFileInfo(file *os.File) {
	scanner := bufio.NewScanner(file)

	scanner.Buffer(make([]byte, 1024), 10*1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func (fh *FileHash) CreateFileHash(inp os.DirEntry, dir string, idx int, ar *[]FileHash) {
	file, err := inp.Info()
	fmt.Println("[Goroutine %d] Starting to process entry %d", idx, idx)

	if err != nil {
		log.Fatal(err)
	}

	// If file is not directory read the content of full file
	if !file.IsDir() {
		filePath := dir + "/" + file.Name()
		hashValue, err := HashFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		fh.FileValue = hashValue
	}

	fh.FileName = file.Name()
	fh.FileSize = file.Size()
	fh.FileModTime = file.ModTime()

	result := *fh
	fmt.Println("[Goroutine %d] finishing process entry %d", idx, idx)
	*ar = append(*ar, result)
}

// hash file
// This needs to work in parallel

func HashFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	h := sha256.New()
	_, err = io.Copy(h, file) // reads whole file safely
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func (fh *FileHash) printInfoFile() {
	fmt.Println("--------------------------------")
	fmt.Println("File name: ", fh.FileName)
	fmt.Println("File size: ", fh.FileSize)
	fmt.Println("File path: ", fh.FilePath)
	fmt.Println("File value: ", fh.FileValue)
}
