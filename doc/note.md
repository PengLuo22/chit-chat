
## 任务1 跑 Hello World



创建Hello.go 文件，键入如下1-1代码，保存

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```

命令行执行 `go run .\Hello.go`，输出结果：`Hello world`。



## 任务2 创建一个简单的 RestFul API

能顺利完成任务1，那么go的开发和运行环境基本就没问题了。接着，我们来创建一个最简单的 RestFul API

创建server.go 文件，键入如下1-2代码，保存

```go
package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world %s!", request.URL.Path[1:])
}

func main() {

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return 
	}

}

```

命令行执行 `go run .\server.go`，在浏览器地址栏

输入`http://localhost:8080/`，返回页面显示：`Hello world !`。

输入`http://localhost:8080/first_restful_api`，返回页面显示：`Hello world first_restful_api!`。

## 任务3 数据的持久化（一）：内存

**需求**：创建一些帖子Post，并使用map容器存储这些帖子，可以根据帖子的Id和Author去查询到关联的帖子

**实现**：创建`persistence_memory.go`文件，键入如下代码

```go
package main

import "fmt"

// Post 帖子
type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostByAuthor map[string][]*Post

func main() {

	// 初始化容器
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	// 创建帖子
	post1 := Post{Id: 1, Content: "Hello world", Author: "go"}
	post2 := Post{Id: 2, Content: "你好，go 开发者", Author: "go"}
	post3 := Post{Id: 3, Content: "世界那么大，我好想去看看", Author: "pengluo"}
	post4 := Post{Id: 4, Content: "我喜欢旅行和健身呀", Author: "pengluo"}

	// 持久化到内存
	store(post1)
	store(post2)
	store(post3)
	store(post4)

	// 打印输出
	fmt.Println("====================")
	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	fmt.Println("====================")
	for _, post := range PostByAuthor["go"] {
		fmt.Println(post)
	}
	for _, post := range PostByAuthor["pengluo"] {
		fmt.Println(post)
	}

}

// 持久化方法
func store(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

```

程序执行结果：

```te
====================
&{1 Hello world go}
&{2 你好，go 开发者 go}
====================
&{1 Hello world go}
&{2 你好，go 开发者 go}
&{3 世界那么大，我好想去看看 pengluo}
&{4 我喜欢旅行和健身呀 pengluo}
```



## 任务4 数据的持久化（二）：标准库操作数据库

需求：能够通过go的标准库完成CRUD操作。没有特殊说明，数据库使用的是 MySQL

第1步：完成post表的数据初始化

```sql
-- 第1步：创建数据库
CREATE DATABASE gwp;
-- 第2步：使用数据库
USE gwp;
-- 第3步：创建表
CREATE TABLE post (
    id INT PRIMARY KEY,
    content TEXT,
    author VARCHAR(255)
);
-- 第4步：添加测试数据
INSERT INTO post (id, content, author)
VALUES
    (1, 'Content 1', 'Author 1'),
    (2, 'Content 2', 'Author 2'),
    (3, 'Content 3', 'Author 3'),
    (4, 'Content 4', 'Author 4'),
    (5, 'Content 5', 'Author 5'),
    (6, 'Content 6', 'Author 6'),
    (7, 'Content 7', 'Author 7'),
    (8, 'Content 8', 'Author 8'),
    (9, 'Content 9', 'Author 9'),
    (10, 'Content 10', 'Author 10');
```

第2步：创建`persist_db_raw.go`文件，键入如下代码

```go
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
```

第3步：执行程序，数据库中查看post表的记录



## 任务5 数据的持久化（三）：ORM框架操作数据库

需求：能够通过开源的ORM框架GORM完成CRUD操作。

在任务4，我们已经完成了数据库、post表的初始化，接下来只需要关注[GORM](https://gorm.io/docs/index.html)框架的使用

第1步：安装依赖

```go
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

第2步：创建`persist_db_gorm.go`文件，键入如下代码

注意实现Tabler接口  指定表名

```go
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

```

第3步：执行程序，数据库中查看post表的记录



