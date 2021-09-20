# go语言实战流媒体视频网站笔记4：流媒体网站架构以及API模块的实现2


# API设计

![API设计](https://cdn.jsdelivr.net/gh/sailaoda/sai_img//img/3/image-20210916184022858.png)

用户、资源和评论之间的关系。



## API设计：用户

- 创建(注册)用户：URL:/user Method:POST,SC:201,400,500      

user 代表URI里面的资源，201创建成功，200 OK（GET时），400 请求错误，500 内部错误

- 用户登录： URL:/user/:username Method:POST,SC:200,400,500

username 是他的参数，用200 是因为只提交了一个表单

- 获取用户的基本信息： URL:/user/:username Method:GET,SC:200,400,401,403,500

并没有往后台写任何东西 ，加401（并没有验证）,403（通过验证了但是不具备操作资源的权限） 

- 用户注销： URL:/user/:username Method:DELETE,SC:204,400,401,403,500

成功204，不用返回任何东西



## 设计用户API并实现

### 一、测试创建用户

主函数

```go
//main.go
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http" //http 的 server
	//自动的把RESTful的API按照他请求的方式以及他的一些参数和URL的格式自动路由到我们想要处理的Handler上
)

func RegisterHandlers() (*httprouter.Router) {
	router := httprouter.New()   //初始化一个router

	router.POST("/user",CreateUser)

	return router
}

func main()  {
	r := RegisterHandlers()
	http.ListenAndServe(":8000",r)

}
//main 函数里面放一些比较简单的定义型的东西，把逻辑处理的放到别的文件里面。
//http分层的处理方式
```

利用包`"github.com/julienschmidt/httprouter"` ，[一个可扩展的高性能 HTTP 请求路由器](https://github.com/julienschmidt/httprouter)。



另外创建一个handlers函数

```go
//用于创建api包函数的使用
package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	io.WriteString(w,"Create User Handler")
}

```

以上就完成了对创建用户API的函数，运行函数，在浏览器中使用谷歌插件不返回数据Talend API Tester（API测试工具）。

**在上面的输入模式选择POST，链接选择`http://localhost:8000/user`，点击发送send**



![成功返回值](https://cdn.jsdelivr.net/gh/sailaoda/sai_img//img/3/image-20210919113228782.png)



### 二、设计用户登录API

1. 在主函数创建函数`RegisterHandlers()`中添加登录路由，

```go
router.POST("/user/:user_name",Login)
```



2. 而后在我们创建的`handlers.go`文件中，添加登录函数`Login()`如下：

```go
func Login(w http.ResponseWriter,r *http.Request,p httprouter.Param) {
    uname := p.ByName("user_name")
    io.WriteString(w,uname)
}
```



3. 在谷歌API测试插件中POST模式输入`http://localhost:8000/user/sailaoda`，返回如下：

   ![登录API返回结果](https://cdn.jsdelivr.net/gh/sailaoda/sai_img//img/3/image-20210919120429084.png)



该过程源代码如下：

主函数：

```go 
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http" //http 的 server
	//自动的把RESTful的API按照他请求的方式以及他的一些参数和URL的格式自动路由到我们想要处理的Handler上
)

func RegisterHandlers() (*httprouter.Router) {
	router := httprouter.New()   //初始化一个router

	router.POST("/user",CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

func main()  {
	r := RegisterHandlers()
	http.ListenAndServe(":8000",r)

}
//main 函数里面放一些比较简单的定义型的东西，把逻辑处理的放到别的文件里面。
//http分层的处理方式
```



handlers函数：

```go
package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	io.WriteString(w,"Create User Handler")
}

func Login(w http.ResponseWriter, r *http.Request,p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w,uname)
}
```



### 总结Golang处理HTTP请求

```go
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http" //http 的 server
	//自动的把RESTful的API按照他请求的方式以及他的一些参数和URL的格式自动路由到我们想要处理的Handler上
)

func RegisterHandlers() (*httprouter.Router) {
	router := httprouter.New()   //初始化一个router

	router.POST("/user",CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

func main()  {
	r := RegisterHandlers()             
	http.ListenAndServe(":8000",r)     //是一个类似于注册handler的函数，阻塞在这里调用r

}
//main 函数里面放一些比较简单的定义型的东西，把逻辑处理的放到别的文件里面。
//http分层的处理方式
```



```go
http.ListenAndServe(":8000",r)     
//是一个类似于注册handler的函数，阻塞在这里会注册调用r ，即RegisterHandlers

当一个request从http.ListenAndServer()函数进来之后

http.ListenAndServe()是block形式

会把整个main Goruntine 完全block住

所以
listen  ->  RegisterHandlers  ->  handlers

每一个handler映射到例如CreateUser,Login,等上面，都是用不同的Goruntine来处理的。
一瞬间可以创建几千几百个Goruntine，Golang并发能力非常强，
天然处理http请求作为http server上有得天独厚的优势
真正把多核特性利用起来
原生的http server框架会自动用Goruntine 方式来调用它。
```

