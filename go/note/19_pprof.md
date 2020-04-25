# 简介：
> PProf是go程序的性能剖析工具  
> 可以捕捉到多维度的运行状态的数据  
> golang在语言层面集成了profile采样工具  
> 在程序运行过程中可以获取cpu、heap、block、traces等执行信息
---

# 内置工具：
> ## runtime/pprof:
> > 用来采集工具型应用的运行数据进行分析  
> > (即程序运行一段时间就结束退出类型的程序)  
> > 最好在应用退出的时候把profiling的报告保存到文件中，以进行分析  
> ## net/http/pprof:
> > 用来采集服务型应用的运行时数据进行分析
> > (即程序是一直运行的，比如web应用)
> ## github.com/DeanThompson/ginpprof:
> > 若使用的是gin框架，则推荐使用该工具  
---  

# 工作形式：
> pprof开启后，每隔一段时间(10ms)就会收集当前的堆栈信息  
> 获取各个函数占用CPU及内存资源的情况  
> 最后通过对这些采样数据进行分析，形成一个性能分析报告  
> `注意：应该只在性能测试时才引入pprof，不要在工作环境引入`  
---

# gin框架中使用示例:
```go
import (
    "github.com/gin-contrib/pprof"
    "github.com/gin-gonic/gin"
)

func main() {
    app := gin.Default()
    pprof.Register(app)         // 引入性能分析
    app.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
    app.Run()
}

/*
运行代码后，会发现自动增加了很多/debug/pprof的API
通过这些API我们可以看到需要的数据
(浏览器访问ip:port/debug/pprof)
*/
```
---

# 火焰图：
## 简介：
> 火焰图是一种性能分析图表  
> 因为它的样子近似火焰而得名  
## 生成火焰图的方式：
> ### go-torch(go版本1.10之前)
> ### 原生的pprof(go版本1.10以后)