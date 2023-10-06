package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

// ctrl + i 实现Tabler接口  指定表名
func (p Post) TableName() string {
	return "post"
}

var Db *gorm.DB

func init() {
	var err error
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	dsn := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func main() {

	// 实例化一条帖子
	post := Post{Id: rand.Intn(11), Content: "add one record by gorm", Author: "gorm"}

	// 新增
	result := Db.Create(&post)
	fmt.Println(result)

	// 根据主键查询 SELECT * FROM post WHERE id = 10;
	readPost := Post{}
	Db.First(&readPost, 2)

	// 修改
	readPost.Content = "content had updated"
	readPost.Author = "gorm-update"
	Db.Updates(readPost)

	// 查所有
	posts := make([]Post, 10)
	Db.Find(&posts)
	for _, p := range posts {
		fmt.Println(p)
	}

	// 根据主键删除
	Db.Delete(readPost, 8)

}
