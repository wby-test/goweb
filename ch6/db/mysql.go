package main

import (
	"database/sql"
	"fmt"
	_"github.com/Go-SQL-Driver/MySQL"
)
type Post1 struct {
	Id 		int
	Content string
	Author 	string
}

func Get(db *sql.DB) (post []Post1){
	rows, err := db.Query("select * from posts ")
	if err != nil {
		fmt.Println(err)
	}


	for rows.Next() {
		posts := Post1{}
		err = rows.Scan(&posts.Id, &posts.Content, &posts.Author)
		if err != nil {
			return
		}
		post = append(post, posts)
	}
	rows.Close()
	return
}

func main () {
	db, err := sql.Open("mysql", "wby:070122@/goweb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if db.Ping() != nil {
		panic(err)
	}
	a := Get(db)
	fmt.Println(a)
	_,err = db.Exec("create table if not exists  goweb.postsbf (id int(10) not null primary key auto_increment,content varchar(255) not null,author varchar(20) not null)")
	if err != nil {
		panic(err)
	}

	rs, err := db.Exec("insert into goweb.postsbf(id, content, author) values (2,'goland','54188')")
	if err != nil {
		panic(err)
	}
	rowConut, err := rs.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("rowCount:  ", rowConut)



}