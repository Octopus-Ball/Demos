# 概述：
> ## 简介：
> 跨语言、跨平台的数据格式  
> 很适合用作数据存储和作为不同应用不同语言之间的通信的数据交换格式  

> ## 优点：
> 更小，更简单，解析速度更快  
> 只需使用Protobuf对数据结构进行一次描述，  
> 即可利用各种语言或从各种不同数据流中对你的结构化数据轻松读写  
> 使用Protobuf无需学习复杂的文档对象模型，简单易学  

> ## 缺点：
> 功能简单，无法用来表示复杂的概念  
> 相比xml，在通用性上差很多  
> 不适合对基于文本的标记文档(如HTML)建模  
> 以二进制形式存储，除非有.proto定义，否则无法读出其内容  
---

# Protobuf安装：
## 1. 下载：  
> git clone https://github.com/protocolbuffers/protobuf.git
## 2. 安装依赖库：  
> sudo apt install autoconf automake libtool curl make g++ unzip libffi-dev -y  
## 3. 检查环境和文件：  
> cd protobuf  
> ./autogen.sh  
> ./configure  
## 4. 编译：  
> make
## 5. 安装：  
> sudo make install  
> sudo ldconfig     // 刷新共享库(不然重启才可以使用)
## 6. 测试是否安装成功：
> protoc -h
## 7. 安装protoc-gen-go：  
> (用来将.proto文件转换为Go代码,proto会调用protoc-gen-go)  
> go get -v -u github.com/golang/protobuf/protoc-gen-go
---

# ProtoBuf使用流程：
> ## 1. proto文件示例，(定义消息类型和服务):
> ```proto3
> // 标明使用proto3版本
> syntax = "proto3"
>
> // 声明包名(防止消息类型有命名冲突)
> // 该项为可选项
> package pkg1
>
> // message关键字定义消息类型(一个.proto文件中可写多个消息类型)
> // (定义名为Student的消息)
> message Student {
>   // 字段类型 字段名 = 标识符
>   // (每个字段都必须提供一个唯一标识符，用来在二进制格式中识别各字段，取值[1,2^29-1])
>   string name = 1;
>   bool male = 2;
>   // repeated表示该字段可重复(对应Go语言中的切片)
>   repeated int32 scores = 3;
> }
>
>
> // 定义服务
> // 若想将消息类型用在RPC系统中，可以在.proto文件中定义一个服务接口
> // 编译器会根据所选不同语言生成服务接口代码存根
>
> // service关键字定义一个服务(会被转化成go里的一个接口)
> service MyService {
>   // rpc 服务的函数名 (传入参数) 返回 (返回参数)
>   rpc Search (int) return (Student)
> }
> ```
> ## 2.编译.proto文件，生成读写接口：
> ```
> protoc --proto_path=存放.proto文件的路径 --go_out=插件:所生成go文件的存放路径 要编译的.proto文件的路径  
>       --go_out=路径：生成go代码的存放路径   
>       --python_out=路径：生成python代码的存放路径  
> 
> eg:
>    protoc --proto_path=./protos --go_out=plugins=grpc:./msg ./protos/*.proto
> ```
> ## 3.go调用接口实现序列化、反序列化以及读写  
> ```go
> import "github.com/golang/protobuf/proto"
> s := &Student{
>   Name: "geektutu",
>   Male:  true,
>   Scores: []int32{98, 85, 88},
> }
> // Marshal将数据序列化为二进制格式
> data, err := proto.Marshal(test)
> // Unmarshal将二进制数据反序列化为结构体实例
> newS := new(Student)
> err = proto.Unmarshal(data, newS)
> ```
---

# proto文件编写规范：
> 1. message采用驼峰命名
> 2. message的字段命名采用小写字母加下划线分隔方式
> 3. enums类型采用驼峰命名
> 4. enums的字段采用大写字母加下划线分隔方式
> 5. service与rpc方法统一采用驼峰式命名
> ## 定义字段的规则:
> > 限定修饰符 数据类型 字段名称 = 字段编码 [字段默认值]
> ## 字段限定修饰符：
> 1. Required:表示一个必选字段  
> 2. Optional:表示一个可选字段
> 3. Repeated:可包含多个值(可看做数组)

# protobuf与Go字段类型的对应：
> |proto类型|go类型|备注|
> |:-------:|:---:|:--:|
> |double|float64|
> |float|float32|
> |int32|int32|
> |int64|int64|
> |uint32|uint32|
> |uint64|uint64|
> |sint32|int32|适合负数|
> |sint64|int64|适合负数|
> |fixed32|uint32|固定长编码,适合大于2^28的值|
> |fixed64|uint64|固定长编码，适合大于2^56的值
> |sfixed32|int32|固定长编码|
> |sfixed64|int64|固定长编码|
> |bool|bool|
> |string|string|UTF8编码，长度不超过2^32|
> |bytes|[]byte|任意字节序列,长度不超2^32|
---