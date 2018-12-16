package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//data := []byte(time.Now().String())
	data := []byte("ddddddd")
	err := ioutil.WriteFile("tmp.txt", data, 0777)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("tmp.txt")
	fmt.Println(string(read1))

	file1, _ := os.Create("temp.txt")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("temp.txt")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("read %d bytes from file\n" ,bytes)
	fmt.Println(string(read2))
}