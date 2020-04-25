# RPC简介：
> Remote Procedure Call Protocol(远程过程调用协议)  
> 它是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络技术的协议  
> 它跟远程访问或者web请求差不多  
> 但减少了信息的包装，加快了处理速度
---

# gRPC简介：
> gRPC是一个高性能的、开源通用的RPC框架  
> 面向移动和HTTP/2设计  
> 支持：双向流、流控、头部压缩、单TCP连接上的多路复用请求等待  
> gRPC默认使用protoBuf进行序列化和反序列化  
> 在gRPC里客户端应用可以像调用本地一样调用另一台机器上服务端的方法  
> 使得我们能够更容易的创建分布式应用和服务
---

# gRPC使用流程：
> ## 1.定义proto文件：
> ### 命名规范：
> 1. message采用驼峰命名
> 2. message的字段命名采用小写字母加下划线分隔(翻译成Go后，首字母会被大写)
> 3. enums类型采用驼峰命名
> 4. enums的字段采用大写字母加下划线分隔方式
> 5. service与rpc方法统一采用驼峰式命名
> ### 定义message字段的规则:
> > 限定修饰符 数据类型 字段名称 = 字段编码 [字段默认值]
> ### 限定修饰符：
> 1. Required:表示一个必选字段  
> 2. Optional:表示一个可选字段
> 3. Repeated:可包含多个值(可看做数组)
> ### 示例：
> ```proto
> // 指明版本为proto3
> syntax = "proto3";
> 
> // 指定包名
> package proto;
>
> // 请求消息
> message NumRequest {
>   int32 n = 1;
> }
> // 响应消息
> message NumResponse {
>   int32 n = 1;
> }
>
> // 定义服务
> service UserInfoService {
>   // 单一请求，单一应答
>   rpc GetNum (NumRequest) returns (NumResponse) {}
>   // 服务端流式应答(可用于下载)
>   rpc GetNums (NumRequest) returns (stream NumResponse) {}
>   // 客户端流式请求(可用于上传)
>   rpc PutNums (stream NumRequest) returns (NumResponse) {}
>   // 双向流式请求应答
>   rpc LoopNums(stream NumRequest) returns (stream NumResponse) {}
> }
> ```
> ---
> ## 2.编译proto文件为xx.pb.go文件：
> ### 编译命令：
> ```shell
> protoc --proto_path=存放.proto文件的路径 --go_out=插件:所生成go文件的存放路径 存放.proto文件的路径
> eg：
>   protoc --proto_path=./protos --go_out=plugins=grpc:./msg ./protos/*.proto
> ```
> ### 编译后的.pb.go文件解析：
> > #### New服务名Client
> > > 该方法用来实例化grpc客户端
> > #### Register服务名Server
> > > 该方法用来将定义的服务注册到grpc
> > #### 服务名Server
> > > 用来实现服务端逻辑的接口
> > #### Unimplemented服务名Server
> > > 实现了服务端逻辑接口的一个结构体
---
> ## 3.实现服务端：
> ### 实现服务端逻辑：
> > 实现 `服务名Server` 接口  
> > `Unimplemented服务名Server`就是对该接口的实现  
> > 可以仿照这个写，或者补全它  
> ```go
> type UnimplementedUserInfoServiceServer struct {
> }
>
> func (*UnimplementedUserInfoServiceServer) GetNum(ctx context.Context, req *NumRequest) (*NumResponse, error) {
>   fmt.Printf("GetNum：%d\n", req.N)
>   return &NumResponse{N: req.N}, nil
> }
> func (*UnimplementedUserInfoServiceServer) GetNums(req *NumRequest, srv UserInfoService_GetNumsServer) error {
>   fmt.Printf("GetNums：%d\n", req.N)
>   for i := 0; i < int(req.N); i++ {
>       srv.Send(&NumResponse{N: int32(i)})
>   }
>   return nil
> }
> func (*UnimplementedUserInfoServiceServer) PutNums(srv UserInfoService_PutNumsServer) error {
>   sum := 0
>   for {
>       req, err := srv.Recv()
>       if err != nil {
>           if err == io.EOF {
>               fmt.Printf("请求数据传输完毕")
>               return srv.SendAndClose(&NumResponse{N: int32(sum)})
>           } else {
>               fmt.Printf("接收错误: %v\n", err)
>           }
>       }
>       fmt.Printf("PutNums: %d\n", req.N)
>       sum++
>   }
> }
> func (*UnimplementedUserInfoServiceServer) LoopNums(srv UserInfoService_LoopNumsServer) error {
>   for {
>       req, err := srv.Recv()
>       fmt.Printf("LoopNums: %d\n", req.N)
>       if err != nil {
>           if err == io.EOF {
>               return nil
>           } else {
>               return err
>           }
>       }
>       err = srv.Send(&NumResponse{N: req.N})
>       if err != nil {
>           return err
>       }
>   }
> }
> ```
> ### 开启服务： 
> ```go
> func main() {
>   addr := "127.0.0.1:8080"
>   listener, err := net.Listen("tcp", addr)
>   if err != nil {
>       fmt.Printf("端口监听异常:%v\n", err)
>   }
>   // 实例化gRPC
>   s := grpc.NewServer()
>   // 将服务注册到gRPC(第二个参数应该为实现了对应接口的结构体)
>   Register服务名Server(s, new(Unimplemented服务名Server))
>   // 启动服务端
>   s.Serve(listener)
> }
> ```
---
> ## 4.实例化客户端 & 调用服务：
> ```go
> func main() {
>   // 与服务端建立连接
>   conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
>   if err != nil {
>       fmt.Printf("连接异常:%s", err)
>   }
>   defer conn.Close()
>   // 实例化grpc客户端
>   client := New服务名Client(conn)
>   // 测试调用服务
>   GetNum_test(client)
>   GetNums_test(client)
>   PutNums_test(client)
>   LoopNums_test(client)
> }
>
> func GetNum_test(client proto.UserInfoServiceClient) {
>   fmt.Println("*************GetNum测试**************")
>   response, _ := client.GetNum(context.Background(), &proto.NumRequest{N: 5})
>   fmt.Println(response)
>   fmt.Println("***********GetNum测试完毕************")
> }
>
> func GetNums_test(client proto.UserInfoServiceClient) {
>   fmt.Println("*************GetNums测试**************")
>   stream, _ := client.GetNums(context.Background(), &proto.NumRequest{N: 5})
>   for {
>       response, err := stream.Recv()
>       if err != nil {
>           if err == io.EOF {
>               fmt.Println("接收完毕")
>           }
>           break
>       }
>       fmt.Printf("rsp=%d\n", response.N)
>   }
>   fmt.Println("***********GetNums测试完毕************")
> }
>
> func PutNums_test(client proto.UserInfoServiceClient) {
>   fmt.Println("*************PutNums测试**************")
>   stream, _ := client.PutNums(context.Background(), grpc.EmptyCallOption{})
>   for i := 0; i < 5; i++ {
>       stream.Send(&proto.NumRequest{N: int32(i)})
>   }
>   response, _ := stream.CloseAndRecv()
>   fmt.Printf("rsp=%d\n", response.N)
>   fmt.Println("***********PutNums测试完毕************")
> }
>
> func LoopNums_test(client proto.UserInfoServiceClient) {
>   fmt.Println("*************LoopNums测试**************")
>   stream, _ := client.LoopNums(context.Background(), grpc.EmptyCallOption{})
>   for i:=0; i<5; i++ {
>       stream.Send(&proto.NumRequest{N: int32(i)})
>   }
>   for {
>       response, err := stream.Recv()
>       if err != nil {
>           if err == io.EOF {
>               fmt.Println("接收完毕")
>           }
>           break
>       }
>       fmt.Printf("rsp=%d\n", response.N)
>   }
>   fmt.Println("***********LoopNums测试完毕**************")
> }
> ```