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
