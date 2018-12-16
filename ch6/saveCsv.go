package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type PostCsv struct {
	Id 		int
	Content string
	Author 	string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []PostCsv {
		PostCsv{1, "golang", "wby"},
		{2, "C++", "wbyC++"},
		{3, "python", "wbyPython"},
		{4, "php", "wbyPhp"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []PostCsv
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := PostCsv{int(id), item[1], item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
	fmt.Println(posts)
}