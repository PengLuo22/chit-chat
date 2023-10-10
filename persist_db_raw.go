package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strings"
)

// 数据库配置
const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "gwp"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	Db, err = sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from post where id = ?", post.Id)
	return
}

func Posts() (posts []Post, err error) {
	rows, err := Db.Query("select id, content,author from post")
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func (post *Post) Update() (err error) {
	Db.Exec("update post set content = ?,author = ? where id = ?", post.Content, post.Author, post.Id)
	return
}

// 根据主键查询
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id,content,author from post where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return

}

func (post *Post) Create() (err error) {

	statement := "insert into post (id,content,author) values (?,?,?)"

	_, err = Db.Exec(statement, post.Id, post.Content, post.Author)
	if err != nil {
		return
	}
	return
}

func main() {

	// 实例化一条帖子
	post := Post{Id: rand.Intn(11), Content: "add one record by orm tool", Author: "gorm"}

	// 新增
	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	// 根据主键查询
	readPost, _ := GetPost(post.Id)

	// 修改
	readPost.Content = "content had updated"
	readPost.Author = "gorm-update"
	readPost.Update()

	// 查所有
	posts, _ := Posts()
	fmt.Println(posts)

	// 根据主键删除
	readPost.Delete()

}
