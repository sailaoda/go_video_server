---
title: go语言实战流媒体视频网站笔记3：流媒体网站架构以及API模块的实现1
cover:  /img/cover/goliu.jpg
tags:
  - - golang实战
categories:
  - - golang
    - go语言实战流媒体视频网站
abstract: '有东西被加密了, 请输入密码查看.'
message: '您好, 这里需要密码.'
wrong_pass_message: '抱歉, 这个密码看着不太对, 请再试试.'
wrong_hash_message: '抱歉, 这个文章不能被校验, 不过您还是能看看解密后的内容.'
date: 2021-09-16 14:22:48
keywords:
description:  本章通过实战演练，从网站的整体架构设计，到服务划分，数据库设计，到api模块的实现，全面讲述golang对webservice的实现以及代码分层架构的思想，同时辅以test cases的全程编写与指导，全面了解工程化golang项目的实现。
sticky: 
password:
swiper_index: 4
---





# 流媒体点播网站

- GO是一门网络编程语言
- 视频网站包含Go在实战项目中的绝大部分技能要点
- 优良的native http库以及模板引擎（无需任何第三方框架）



## 总体架构

![总体架构](https://cdn.jsdelivr.net/gh/sailaoda/sai_img//img/3/image-20210916181324756.png)

典型的前后端分离的服务，后端用API来连接前端。

API会将一些业务的数据往DB里面写，同时会将一些业务的处理比如说视频的播放，上传下载请求送到Streaming模块里面，Scheduler会处理一些删除，软删除，定期清理一些。

Streaming和Scheduler都会同时访问DB，会直接去文件系统里面找视频文件做相应的处理。



## 什么是前后端解耦

- 前后端解耦是时下流行的web网站架构
- 前端页面和服务通过普通的web引擎渲染
- 后端数据通过渲染后的页面脚本调用后处理和呈现





### 前后端解耦的优势

- 解放生产力，提高合作效率
- 松耦合的架构更灵活，部署更方便，更符合微服务的设计特征
- 性能的提升，可靠性的提升



### 前后端耦合的缺点

- 工作量大
- 前后端分离带来的团队成本以及学习成本
- 系统更复杂



# API

- REST API
- REST是一种设计风格，不是任何架构标准
- 当今RESTful API通常使用**HTTP**作为通信协议，**JSON**作为数据格式



## API特点

- 统一接口(Uniform Interface)
- 无状态(Stateless)    
- 可缓存(Cacheable)     减少后端服务的压力
- 分层(Layered System)  将一个API的servers 分成很多层很多个服务，每一层次负责一部分功能
- CS模式(Client-server Atchitecture)



## API设计原则

- 以URL(统一资源定位符) 风格设计API
- 通过不同的METHOD(GET,POST,PUT,DELETE)来区分对资源的CRUD。用METHOD来区分对资源的不同操作
- 返回码(Status Code ) 符合HTTP资源描述的规定   （ 404 找不到页面等）可读性













