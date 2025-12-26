package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Please provide the Directory: ")
	var dir string

	fmt.Scanf("%s", &dir)

	// Now Read this directory and read metadata of all the files
	entries, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	var ar []FileHash
	for idx, value := range entries {
		//wg.Add(1)
		fh := NewFileHash()
		value := value
		go fh.CreateFileHash(value, dir, idx, &ar)
	}

	time.Sleep(5 * time.Second)
	for _, value := range ar {
		value.printInfoFile()
	}
	//wg.Wait()
}
