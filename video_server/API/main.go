package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http" //http 的 server
	//自动的把RESTful的API按照他请求的方式以及他的一些参数和URL的格式自动路由到我们想要处理的Handler上
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New() //初始化一个router

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)

}

//main 函数里面放一些比较简单的定义型的东西，把逻辑处理的放到别的文件里面。
//http分层的处理方式 结构

//handler -> validation{1.request,2.user} -> business logic -> response .
