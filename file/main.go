package main

import (
	"fmt"
	"os" //os means operating system, and it provides a platform-independent interface to operating system functionality.
)

func main() {
	/*
		file, err := os.Create("example.txt") //os returns file type data and an error if any occurs.
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close() //defer is used to ensure that the file is closed after all operations are done, even if an error occurs.
		//resource management is important to prevent memory leaks and ensure that file handles are released.

		content := "Hello, this is a sample file content."
		byte, err2 := io.WriteString(file, content+"\n") //io.WriteString writes the string content to the file.
		fmt.Printf("Wrote %d bytes to file.\n", byte)    //Prints the number of bytes written to the file.
		if err2 != nil {
			fmt.Println("Error writing to file:", err2)
			return
		}
		fmt.Println("File created successfully with content.")
	*/

	/*
		file, err := os.Open("example.txt") //os.Open opens the file for reading, but does not create it if it doesn't exist.
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		//create a buffer to hold the file content
		buffer := make([]byte, 1024) //1024 bytes buffer to read the file

		// Read the file content into the buffer
		for {
			n, err := file.Read(buffer) //Read reads up to len(buffer) bytes from the file.
			if err == io.EOF {          //io.EOF indicates the end of the file.
				break
			}
			if err != nil { //If no bytes were read, we reached the end of the file.
				fmt.Println("Error reading file:", err)
				return
			}
			fmt.Println(string(buffer[:n])) //Print the content read from the file.
		}
	*/

	content, err := os.ReadFile("example.txt") //os.ReadFile reads the entire content of the file into memory.
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:")
	fmt.Println(string(content)) //Print the content of the file as a string.
}

/*
The number of bytes written (38) comes from the length of the string you are writing to the file, including the newline character (\n).

Let's break it down:
"Hello, this is a sample file content." has 37 characters.
Adding \n (newline) makes it 38 characters.
Each character (including spaces and punctuation) is 1 byte in UTF-8 for standard ASCII characters, which your string uses.

So, io.WriteString(file, content+"\n") writes 38 bytes to the file.
*/
