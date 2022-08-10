# spg-go

## 项目介绍

此项目为spg项目组开发go语言后端项目时用于快速构建开发目录的脚手架,只针对于windows系统开发，ios和linux需要手动创建目录

> #### 该项目生成的开发目录仅为spg开发组通过多个项目的开发和总结，结合gin框架和 [project-layout](https://github.com/golang-standards/project-layout)结构化目录，设计的一套可维护、可扩展的代码目录

## 项目源码

> #### https://github.com/zhangshuai268/spg-go-framework

## 项目使用

1.克隆此项目

>克隆此项目到gopath/src目录下，克隆时注意需要指定项目名称！
````
$ git clone https://github.com/zhangshuai268/spg-go.git project-name
````

2.通过编译工具(goland)进入项目,构建配置文件

>根据config_example.json文件模板构建项目需要的配置文件，可增加其他配置只需要注意json语法格式即可

>注意: 配置文件的名称必须为config.json

3.编译器命令行执行命令构建目录

````
$ .\framework_create.exe
````

4.命令执行过程

(1)修改git源

>输入线上代码仓库地址
````
请输入git仓库地址：
https://github.****.com/test/testfranmework.git
初始化go.mod,下载相关依赖
生成目录结构成功    
生成配置文件成功    
生成数据库目录成功  
````

(2)构建所需系统的目录结构

>系统名称只能包含字母、'-'、'_'

````
请输入开发系统数量：
2
请输入开发系统名称：
api_admin
请输入开发系统运行端口号：
8088
请输入开发系统名称：
api_wx
请输入开发系统运行端口号：
8087
````

## 项目目录介绍

1.目录结构

````
├── api
│   ├── dokcer
│   │   └── api_admin
│   │       └── Dockerfile
│   └── swagger
│       └── api_admin
│           └── doc
│              └── doc.go 
├── cmd
│   └── api_admin
│       └── main.go
└── internal
│   ├── api
│   │   ├── api_admin
│   │   │   ├──auth
│   │   │   │  └──auth.go
│   │   │   ├──controller 
│   │   │   ├──service
│   │   │   │  └──service.go
│   │   │   └──route.go
│   │   └── store
│   │       ├──store.go
│   │       └──factory.go
│   ├── config
│   │   ├──config.go
│   │   └──config_init.go
│   ├── crontab
│   ├── model
│   └── pkg
│       └──middle
│          └──middle.go
└──config_update.exe

````
2.目录作用

* /api/docker: 根据系统存放系统的Dockerfile文件，可用于cicd部署或者生成容器
* /api/swagger: swagger接口文档存放地址
* /cmd: main文件存放地址
* /internal/api: 项目代码存放目录
* /internal/api/api_admin: 根据输入的系统名称所生成的目录用于区分不同系统的代码
* /internal/api/api_admin/auth/auth.go: 接口权限访问拦截器，根据session、token等自行编写
* /internal/api/api_admin/controller: 控制层代码目录
* /internal/api/api_admin/service: 服务层代码目录
* /internal/api/api_admin/route.go: 路由
* /internal/store/store.go: 数据层dao生产代码
* /internal/store/factory.go: 数据层dao工厂代码
* /internal/config/config.go: 根据config.json文件生成的配置结构体
* /internal/config/config_init.go: 初始化配置结构体
* /internal/crontab: 定时任务
* /internal/model: 数据库表结构对应结构体存放目录
* /internal/pkg: 项目内部公共方法，可自行根据功能扩展，已包含异常处理和跨域
* /config_update.exe: 开发过程中若配置文件有修改，可在命令行执行./config_update.exe -g 同步config.go
* /config.json: 配置文件

## 其他

1.通过系统变量，选择不同的配置文件

>设置环境变量STAGE,则初始化config.go结构体时，会根据环境变量来选择config_${STAGE}.json文件来初始化结构体，详情见/internal/config/config_init.go

2.项目运行端口

>项目运行端口可调整为配置文件控制，只需要将cmd中main方法的port字段改为配置结构体控制即可，主要配置文件的运行端口要与docker，swagger文件一致，否则上线或测试会出现端口号不一致的问题

3.go.mod文件报错

>可能原因未设置编译器启动go.mod, goland调整方法File->Settings->Go->Go modules->勾中Enable Go modules integration即可