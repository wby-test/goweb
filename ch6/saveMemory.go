package main

import "fmt"

type Post struct {
	Id 		int
	Content string
	Author 	string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{1, "C++", "wby"}
	post2 := Post{2, "golang", "wbyPlus"}
	post3 := Post{3, "python", "wbyPython"}
	post4 := Post{4, "PHP", "wby"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for author, post := range PostsByAuthor["wby"] {
		fmt.Println(author, " ",  post)
	}

	for author, post := range PostsByAuthor["wbyPlus"] {
		fmt.Println(author, "   ",  post)
	}
}