
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
	fmt.Fprintf(writer, "Hello world.%s!", request.URL.Path[1:])
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

```

命令行执行 `go run .\server.go`，在浏览器地址栏

输入`http://localhost:8080/`，返回页面显示：`Hello world !`。

输入`http://localhost:8080/first_restful_api`，返回页面显示：`Hello world first_restful_api!`。





