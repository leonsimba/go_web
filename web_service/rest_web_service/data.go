package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

var (
	user="gwp"
	dbname="gwp"
	passwd="qwe123"
	dbhostip="127.0.0.1"
)

// connect to the Db
func init() {
	var err error
	Db, err = sql.Open("mysql", user+":"+passwd+"@tcp("+dbhostip+")/"+dbname+"?charset=utf8")
	if err != nil {
		panic(err)
	}

	//Db.Query("drop database if exists gwp")
	//Db.Query("create database gwp")
	//Db.Query("use gwp")
	//Db.Query("create table posts(id int primary key, )database gwp")
}

// Get a single post
func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// insert a new post
func (post *Post) create() (err error) {
	// 准备一条insert sql语句
	statement := "insert into posts (content, author) values (?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// 执行前面准备的预处理语句，即执行插入操作 
	_, err = stmt.Exec(post.Content, post.Author) //.Scan(&post.Id)
	return
}

// Update a post
func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id = ?", post.Content, post.Author, post.Id)
	return
}

// Delete a post
func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)
	return
}
