package src

/*


// return type of os.Create is *File

file, err := os.Create("data.txt")
if err != nil {
	log.Fatal(err)
}
defer file.Close()

// Now printing the context of file
file.WriteString("Hello this is harsh\n")

// user buffer writer for the file

writer := bufio.NewWriter(file)
writer.WriteString("This is written by buffered\n")
writer.Flush()

// To Read small file
data, err := os.ReadFile("data.txt")
if err != nil {
	log.Fatal(err)
}

fmt.Println("content of file is : ", (string(data)))

// We can't just read the file like this, For that we have to use some file reader


*/
