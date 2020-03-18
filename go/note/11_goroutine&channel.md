# Goroutine:
## 简介：
> Goroutine是Go中最基本的执行单元  
> 每个Go程序至少有一个Goroutine  
> 程序启动时的main函数本身就是一个Goroutine(它会自动创建)  
## Coroutine与协程：
> Goroutine可以理解为一种Go语言的协程  
> 不同的是Go在runtime、系统调用等多方面对goroutine调度进行了封装和处理  
> 即goroutine不完全是用户控制，一定程度上是由go运行时(runtime)管理  
## Goroutine什么时候结束：
1. goroutine对应的函数结束了，该goroutine就结束了
2. main函数结束了，由main函数创建的goroutine就都结束了
---

.  
# channel：
## 引言：
> 协程与协程之间需要交换数据才能体现并发执行的意义  
> 若使用共享内存进行数据交换，为避免竟态问题，必须对互斥量加锁  
> 而这种做法势必会造成性能问题  
> Go语言的CSP并发模型提倡通过通信共享内存而不是通过共享内存通信  
## 简介：
> goroutine是Go程序并发的执行体，channel是它们之间的连接  
> 类似于Unix上的管道(可在进程之间传递消息)  
> 其实就是在做goroutine之间的内存共享  
> channel类型本身就是并发安全的，这也是Go语言自带的唯一可以满足并发安全性的类型  
## 类型相关：
> Go中的channel是一种特殊的类型，像一个队列，遵循先进先出的规则  
> 每一个channel都是一个具体类型的通道(也就是声明channel时要为其指定元素类型)
> 一个channel只能传递一种类型的值  
## 注意：
1. 通道关闭后，依然可以读取未读数据，但是不能继续写数据(会引发panic)  
   因此除非有特殊保障，我们应该让发送方关闭通道，而不是接收方  
2. 若我们试图关闭一个已经关闭的通道会引发panic
## 有缓存通道与无缓存通道的区别：
> ### 无缓存通道：
> > 读写协程在无缓存通道间是同步的  
> > 写端使用无缓存通道时不仅仅是将数据写入，而是要一直等到有读端接手，写端程序才会继续执行下去  
> > 如果只有读端，读端阻塞  
> > 如果只有写端，写端阻塞  
> ### 有缓存通道：
> > 读写协程在无缓存通道间是异步的  
> > 直到缓存区被填满，写端才阻塞  
> > 直到缓存区被清空，读端才阻塞  
## 单向通道：
> 有时我们会将通道作为参数在多个任务函数间传递  
> 很多时候，我们在不同的任务函数中使用通道都会对其进行限制(比如只能发送或只能接收)  
> Go语言中提供了单向通道来处理这种情况
```go
// 只写通道
func funName(ch chan<- int) {
    ...
}
// 只读通道
func funName(ch <-chan int>) {
    ...
}
```
## 操作示例：
```go
// 声明channel
// var 名称 chan 类型
var a chan int          // 声明一个可传递int型数据的channel a

// 初始化channel
b := make(chan int)     // 初始化一个可传递int型数据的channel b
// 初始化带缓存的通道
c := make(chan int 10)  // 初始化一个缓存大小为10的通道

// 读写channel
c <- 1          // 将数据1写入channel c
data := <- c    // 从channel c读取数到data

// 遍历
// 当通道被关闭并且内部数据都已经读完的情况下，再次读取该通道时
// data, ok := <- c     data值为对应类型的零值，ok的值为false
for {
    data, ok := <- c {
    if !ok {
        break
    }
    fmt.Println(data)
    }
}
// channel也可以使用range取值，会一直从channel中读取数据，直到该channel被执行close
for data := range c {
    fmt.Println(data)
}

// 关闭channel
close(c)
```
## 优雅的关闭channel：
> ### Channel的关闭原则：
> 1. 不要在消费端关闭channel  
> (消费端关掉后，若生产端还写入数据，会引发panic)
> 2. 不要在有多个并行的生产者时对channel执行关闭操作  
> (若一个生产者关闭了通道，可其他生产者还有数据需要写入,会引发panic，且数据无法写入而丢失)
> 3. 应该只在唯一的最后剩下的生产者协程中关闭channel
> ### 暴力关闭channel的正确方法：
> > 若想在消费端关闭channel，或在多个生产端关闭channel，  
> > 可以使用recover机制来上保险,避免程序因为panic而崩溃  
> > 这样可以避免panic导致崩溃，但数据还是会丢失  
> ### 1.防止关闭一个已经关闭的通道会引发panic：
> > 在关闭channel时可使用sync.Once或者sync.Mutex来避免多次关闭通道  
> > 同一个sync.Once的Do方法只会执行第一次被调用时传入的参数函数  
> > `注意：声明的一个sync.One，只能用于一个channel，否则会让其他channel关闭失败`  
> > `Do方法接收的参数应该是一个函数类型`  


> ### 2.单发送者多接收者：
> > 发送者去关闭通道  
> ### 3.单接收者多发送者：
> > 接收者通过额外信号通道告知发送者停止发送
> ### 4.多接收者多发送者：
> > 任何一个通知哨兵关闭信号通道  
---

.  
# WaitGroup:
## 简介：
> WaitGroup是Go应用开发过程中经常使用的并发控制技术  
> WaitGroup里有一个counter计数器，用来记录需要等待的协程数量  
> `开箱即用，并发安全`
## 相关接口：
> 1. Add(delta int)  
> 把delta值累加到counter中(delta可以为负值)  
> 2. Wait()
> 调用该方法会阻塞等待  
> 3. Done()
> 把counter减一  
> 实际上是调用了Add(-1)
## 注意：
> `Add操作必须早于Wait，否则会引发panic`  
> `Add设置的值必须与实际等待的goroutine个数一致，否则会引发panic`  
## 番外：
> 信号量是Unix系统提供的一种保护共享资源的机制  
> 用于防止多个线程同时访问某个资源  
> WaitGroup实现中使用了信号量  
## 推荐操作：
> 在使用WaitGroup值得时候，最好先统一Add，再并发Done，最后Wait
> `若在调用Wait方法的同时再调用Add方法可能引发panic`  
---

.  
# select:
## 简介：
> 听到select，很容易想到基于select、poll、epoll系统调用构建的IO多路复用模型  
> 而Go语言中的select有着比较相似的功能  
> 系统调用select可以同时监听多个文件描述符的可读写状态  
> Go语言中的select关键字也能够让Goroutine同时等待多个Channel的读写  
> 在多个文件或Channel发生状态改变前select会一直阻塞当前Goroutine  
## select与switch：
> select是一种与switch相似的控制结构  
> 与switch不同的是，select中虽然也有多个case  
> 但这些case中的表达式必须都是channel的收发操作  
## 执行规律：
1. select会阻塞当前goroutine，等待channel中的一个达到可收发状态
2. 哪个case的channel读写操作触发了，则立即执行该case中的代码
3. 当多个case同时被触发时，就会随机选择一个case执行  
   (为了避免顺序执行导致后面的一直得不到执行，所以用随机)  
4. 若select控制语句中包含default语句那么：
   当不存在可收发的channel时，不是阻塞，而是执行default中的语句
---

.
# 编程示例：
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// 工厂
type factory struct {
	wg  sync.WaitGroup // 用于并发控制
	one sync.Once      // 用于避免多消费者情况下通道被多次关闭而引发panic
	// 原料
	part string
	// 成品
	product string
	// 原料通道
	partCh chan string
	// 成品通道
	productCh chan string
	// 成品数量
	productSum int
}

// 原料生产者
func (f *factory) producer(partName string, partSum int) {
	defer f.wg.Done()
	defer close(f.partCh)
	for i := 0; i < partSum; i++ {
		fmt.Printf("提供原料:%s,%d\n", partName, i+1)
		f.partCh <- partName
	}
}

// 原料消费者
func (f *factory) consumer() {
	defer f.wg.Done()
	defer close(f.productCh)
	for _ = range f.partCh {
		time.Sleep(time.Second)
		product := f.product
		fmt.Printf("生产线制作%s\n", product)
		f.productCh <- product
	}
}

// 成品检查者
func (f *factory) examiner() {
	defer f.wg.Done()
	for product := range f.productCh {
		fmt.Printf("检查%s\n", product)
		f.productSum++
	}
}

// 运行工厂
func newFactory(part, product string, partSum int) int {
	f := new(factory)
	// 初始化属性
	f.part = part
	f.product = product
	f.productSum = 0
	// 初始化通道
	f.partCh = make(chan string, 100)
	f.productCh = make(chan string, 100)
	// 开始制作
	f.wg.Add(3)
	go f.producer(f.part, partSum)
	go f.examiner()
	go f.consumer()
	f.wg.Wait()
	return f.productSum
}

func main() {
	fmt.Println("*********开始")
	sum := newFactory("面粉", "面包", 10)
	fmt.Println("*********结束")
	fmt.Printf("生产面包%d个\n", sum)
}  
```