# 简介：
> Go Micro是一个插件化的基础框架  
> 它提供了微服务实施所需的基础框架(rpc、web、api网关等工具)
> 它隐藏了分布式系统的复杂性，为开发人员提供了更简洁的概念  
> Micro的设计哲学是可插拔的插件化架构  
> 在架构之外，它默认实现了consul作为服务发现  
> 通过http进行通信，通过protobuf和json进行编码  
---

# go-micro通信流程：
1. Server监听客户端的调用和Brocker推送过来的信息进行处理  
2. Server端需要向Register注册自己的存在或消亡，以便Client得知其状态  
3. Client端从Register中得到Server的信息  
   然后每次调用都根据算法选择一个Server进行通信
4. 如果有需要通知所有的Server端，可用使用Brocker进行信息的推送
---

# go-micro核心接口：
> ## 简介：
> go-micro之所以高度定制，和它的框架结构是分不开的  
> go-micro由8个关键的interface组成  
> 每个interface都可用根据自己的需求重新实现  
> ## Transort 通信接口：
> 服务之间通信的接口  
> 也就是服务发送和接收的最终实现方式，是由这些接口定制的
> ## Code 编码接口：
> go-micro有很多种编解码方式  
> 默认的实现方式是protobuf  
> 当然也有其他实现方式json、jsonrpc、mercury等  
> ## Registry 注册接口：
> 服务的注册和发现  
> 目前实现的有consul、mdns、etcd、zookeeper、kubemetes等  
> ## Selector 负载均衡：
> 以Register为基础，Selector是客户端级别的负载均衡  
> 当有客户端向服务发送请求时，  
> selector根据不同的算法从Registery中的主机列表得到可用Service节点进行通信  
> 目前实现的有循环算法和随机算法，默认是随机算法  
> ## Client 客户端接口：
> Client是请求服务的接口  
> 它封装Transort和Code进行rpc调用  
> 也封装了Brocker进行信息的发布  
> ## Server 服务端接口：
> Server监听等待rpc请求  
> 监听broker的订阅信息，等待信息队列的推送等  
> ## Service 接口：
> Service是Client和Server的封装  
> 它包含了一系列的方法使用初始值去初始化Server和Client  
> 使我们可以简单的创建一个rpc服务  
> 
---

# go-micro的主要功能：
## 1.服务发现：
> 自动服务注册和名称解析  
> 服务发现是微服务开发的核心  
## 2.负载均衡：
> 基于服务发现构建客户端负载均衡  
> 我们使用随机散列负载均衡来提供跨服务的均匀分布，并在出现问题时重试不同节点  
## 3.消息编码：
> 基于内容类型的动态消息编码  
> 客户端和服务端将使用编解码器和内容类型无缝编码和解码Go类型  
## 请求/响应：
> 基于RPC的请求响应，支持双向流  
> 提供了同步通信的抽象  
> 对服务的请求将自动解决负载均衡、拨号、流式传输  
## Async Messaging:
> PubSub是异步通信和事件驱动架构的一流公民  
> 事件通知是微服务开发的核心模式  
## 可插拔接口：
> Go Micro为每个分布式系统抽象使用Go接口  
> 因此这些接口是可插拔的  
> 并允许Go Micro与运行时无关，可插入任何基础技术  
> [插件地址](https://github.com/micro/go-plugins)
---

# Micro工具集组件：
## API:
> ### 功能：
> 1. 将HTTP请求映射到API接口
> 2. 将HTTP请求映射到RPC服务
> 3. 将HTTP请求广播到订阅者
## Web:
> web反向代理与管理控制台
## Proxy:
> 代理Micro风格的请求  
> 支持异构系统只需要瘦客户端便可调用Micro服务  
## CLI: 
> 以命令行操控Micro服务  
> (执行micro help了解更多)  
## Bot:
> 与常见的通信软件对接  
> 负责传送信息，远程指令操作  
---

# Go-micro框架模块：
## Service：
> 具体实例化的服务  
> 包含两个重要组件(Server、Client)
## Server：
> 接受RPC请求与广播消息  
## Client：
> 发送RPC请求与广播消息  
## Codec：
> 数据编解码组件  
## Registry：
> 服务注册组件  
## Selector：
> 负载均衡器  
## Transport：
> 同步通信组件  
## Broker：
> 异步通信组件  

# go-Micro环境搭建：
## 安装protobuf：
> sudo apt install protobuf-compiler
## 安装protobuf相关依赖:
> `go get -v -u github.com/golang/protobuf/protoc-gen-go`
> `go get -v -u github.com/micro/protoc-gen-micro/v2`
## 安装Micro：
> Micro是运行时工具集  
> 用来管理go-Micro编写的微服务  
> `go get -v github.com/micro/micro/v2`
## 安装go-micro：
> go-micro是用来编写微服务的框架、库  
> `go get -v -u github.com/micro/go-micro/v2`
---