package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type PostGob struct {
	Id 		int
	Content string
	Author 	string
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename,buffer.Bytes(),0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := PostGob{1, "golang", "wby"}
	store(post, "postq")
	var postRead PostGob
	load(&postRead, "postq")
	fmt.Println(postRead)
}