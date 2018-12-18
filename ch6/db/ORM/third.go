package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Post struct {
	Id int
	Content string
	AuthorName string `db: author`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("mysql", "wby:070122@/goweb")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id , content, author from posts where id=?", id).StructScan(&post)
	if err != nil {
		panic(err)
	}
	return
}

func (post *Post) Create() (err error) {
	err = Db.QueryRowx("insert into posts (content, author) values (?, ?)", post.Content,post.AuthorName).Scan(&post.Id)
	return
}


func main() {
	post := Post{Content:"i miss you sss", AuthorName:"wby"}
	post.Create()
	fmt.Println(post)
}
