---
title: go语言实战流媒体视频网站笔记1：要点介绍
cover:     img/cover/goliu.jpg
tags:
  - - golang实战
categories:
  - - golang
    - go语言实战流媒体视频网站
abstract: '有东西被加密了, 请输入密码查看.'
message: '您好, 这里需要密码.'
wrong_pass_message: '抱歉, 这个密码看着不太对, 请再试试.'
wrong_hash_message: '抱歉, 这个文章不能被校验, 不过您还是能看看解密后的内容.'
date: 2021-09-14 23:44:12
keywords:
description:
sticky:
password:
swiper_index:
---



# 为什么选择golang

- 开发效率高（完整的开发工具链tools,test,benchmark,builtin.etc）

- 部署简单（Compile once,run everywhere）

- 良好的native http 库以及模板引擎（无需任何第三方框架）

- 优秀的并发模型



# 课程概述

- golang简单介绍，通过简单的web server了解Go的工具链以及玩法
- 用golang完成一个可以使用的流媒体视频网站
- 网站上云



# golang简单介绍

- Go的工具链的使用，build,install,fmt等
- Go test和Go bench



# golang搭建视频网站

- 前后端分离的系统架构设计
- RESTful风格API的设计与实现
- golang实现webservice
- 系统的服务化解耦
- golang的channel和并发模型的实践应用
- 利用golang原生的template完成web UI 的实现



# 网站上云

- 用ECS，ELB等云服务重构业务工程架构
- 在阿里云上实现cloud native架构下的业务部署



```
课程名称: Go语言开发高性能网站项目实战 实战流媒体视频网站 
课程简介:
go是很多php程序员的首选 php在web现在还有一席之地 而go面向未来 在微服务 云计算 直播 区块链等热门领域都有不错的发展
----------------------课程目录------------------------------
第1章 课程介绍
介绍这门课程大纲，技术堆栈以及环境
1-2 课程介绍及知识要点预习
第2章 一个例子了解golang工具链
通过一个简单的webservice具体从golang的工具链，到test，全面介绍golang在工程项目里需要掌握的知识点。
2-1 一个例子了解golang常用工具链
2-2 golang项目中test的写法
2-3 golang项目中benchmark的写法
2-4 章节总结
第3章 流媒体网站架构以及API模块的实现
本章通过实战演练，从网站的整体架构设计，到服务划分，数据库设计，到api模块的实现，全面讲述golang对webservice的实现以及代码分层架构的思想，同时辅以test cases的全程编写与指导，全面了解工程化golang项目的实现。
3-1 流媒体网站整体介绍与架构梳理
3-2 api设计与架构
3-3 api实现之详细设计(上)
3-4 api实现之详细设计(中)
3-5 api实现之详细设计(下)
3-6 api之http handler层
3-7 api之数据库层设计
3-8 api之数据库层实现_数据库连接
3-9 api之数据库层实现_实现User
3-10 api之数据库层实现_编写User Test Case
3-11 api之数据库层实现_User部分代码优化
3-12 api之数据库层实现_实现和验证Video
3-13 api之数据库层实现_实现Comments
3-14 api之数据库层实现_Comments Test Case
3-15 api之session处理与实现(上)
3-16 api之session处理与实现(下)
3-17 api之http middleware的实现与handler收尾(上)
3-18 api之http middleware的实现与handler收尾(下)
第4章 stream模块
通过stream server的实现过程，着重讲述通过golang实现流式播放，上传文件，以及利用channel实现流控等实用知识点，进一步加深对golang的掌握。
4-2 streaming的架构搭建
4-4 流控模块的实现
4-5 在http middleware中嵌入流控
4-6 streamHandler实现
4-7 验证streamHandler
4-8 uploadHandler实现
4-9 验证uploadHandler
第5章 scheduler模块
通过对生产者消费者模型在scheduler中的实现，全面了解golang是如何处理并发场景，以及如何在并发场景下通过channel实现消息同步。
5-1 scheduler介绍
5-2 代码架构搭建
5-3 runner的生产消费者模型实现
5-4 runner的使用与测试
5-5 task示例的实现
5-6 timer的实现
5-7 api实现以及scheduler完成
第6章 前端服务和模版引擎渲染
讲述如何使用golang的模版引擎来渲染html文件，如何通过原生proxy和api两种模式实现后端服务接口透传并避免跨域访问，以及整个前台在实现业务上的js逻辑代码。
6-1 大前端和golang模版引擎介绍
6-2 前端代码架构搭建
6-3 静态页面渲染
6-4 build脚本和homeHandler
6-6 api透传模块实现
6-7 proxy转发的实现
6-8 UI部分的预览
6-9 API service补全与讲解
6-10 UI之html讲解
6-11 js部分实现
第7章 网站上云
通过对网站部分架构的改造和代码重构，使之更符合cloud native架构，辅以阿里云计算存储网络等服务，最终实现网站上云，打通网站上线最后一公里。
7-1 云原生讲解
7-2 云存储改造之OSS方案分析
7-3 云存储改造之OSS适配
7-4 公共配置实现
7-5 用vendor处理公共配置包
7-6 SLB讲解与配置
7-7 SLB之添加session容错
7-8 ECS云主机和安全组配置
7-9 scheduler的改造
7-10 部署脚本以及db初始化
7-11 部署演示以及完成效果展示
7-12 课程总结（回顾，延伸和优化）
```

