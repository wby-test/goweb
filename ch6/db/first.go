package main

import (
	_"github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"fmt"
)

type Post struct {
	Id 		int
	Content string
	Author 	string
}

var Db *sql.DB
func init() {
	var err error
	//Db, err = sql.Open("mysql", "user=root, dbname=goweb,password=070122, sslmode=disable")
	Db, err = sql.Open("mysql", "root:070122@tcp(127.0.0.1:3306)/goweb?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id , content, author from posts limit $1", limit)
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
	err = Db.QueryRow("select id , content, author from posts where id = $1" ,id).Scan(&post.Id, &post.Content, &post.Author)
	fmt.Println("getpost: ", post)
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post)Update() (err error) {
	fmt.Println("update: ",post)
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1",
		post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {
	post := Post{Content:"golang", Author:"wby"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "php"
	readPost.Author = "wyr"
	readPost.Update()

	posts, _ := Posts(post.Id)
	fmt.Println(posts)

	post1,_ := GetPost(101)
	fmt.Println(post1.Id)
	a,_ := GetPost(101);
	fmt.Println(a.Author)

	//readPost.Delete()
}