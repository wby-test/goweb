package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id 		int
	Content string
	Author 	string
}

var Db *sql.DB
func init() {
	var err error
	//Db, err = sql.Open("mysql", "user=wby, dbname=goweb,password=070122")
	Db, err = sql.Open("mysql", "wby:070122@/goweb")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id , content, author from posts limit ?", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		fmt.Println("getPost:...",err)
	}
	return
}
func (post *Post)GetId(){
	err := Db.QueryRow("select id from posts where content=? and author=?", post.Content, post.Author).Scan(&post.Id)
	if err != nil {
		panic(err)
	}
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values (?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(post.Content)
	if err != nil {
//		panic(err)
		return
	}
	return
}


func (post *Post)Update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id = ? ",
		 post.Content, post.Author,post.Id)
	if err != nil {
		panic(err)
	}
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)
	return
}

func main() {
	post := Post{Content:"golang", Author:"wby"}

	fmt.Println(post)
	//post.Create()
	post.GetId()

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "php"
	readPost.Author = "wyr"
	fmt.Println(readPost)
	readPost.Update()

	posts, _ := Posts(10)
	fmt.Println(posts)

	post.Delete()

	//Db.Exec("delete from posts")
}