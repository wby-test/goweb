package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id 			int
	Content 	string
	Author 		string
	Commnets 	[]Comment
}

type Comment struct {
	Id 			int
	Content 	string
	Author 		string
	Post 		*Post
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "wby:070122@/goweb")
	if err != nil {
		panic(err)
	}
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values (?, ?, ?)", comment.Content, comment.Author,comment.Post.Id).Scan(&comment.Author)
	Db.QueryRow("select id from comments where content = ? and author = ? and post_id = ?", comment.Content, comment.Author,comment.Post.Id).Scan(&comment.Id)
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Commnets = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id=?", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("select id, content, author from comments")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err := rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			panic(err)
		}

		post.Commnets = append(post.Commnets, comment)
	}

	rows.Close()
	return
}

func (post *Post) Create()  {
	Db.QueryRow("insert into posts (content, author) values (?, ?)", post.Content, post.Author)
	Db.QueryRow("select id from posts where content=? and author=?", post.Content, post.Author).Scan(&post.Id)
}

func (post *Post) GetId() {
	Db.QueryRow("select id from posts where content=? and author=?", post.Content, post.Author).Scan(&post.Id)
	//if err != nil {
	//	panic(err)
	//}
}
func main() {
	post := Post{Content:"i miss you", Author: "wby"}
	post.Create()

	commnet := Comment{Content:"you miss your mama, sb", Author: "zhx", Post:&post}
	commnet.Create()

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)
	fmt.Println(readPost.Commnets)
	fmt.Println(readPost.Commnets[0].Post)
}
