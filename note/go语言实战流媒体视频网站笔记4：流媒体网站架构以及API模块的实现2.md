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

