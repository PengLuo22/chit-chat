
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





