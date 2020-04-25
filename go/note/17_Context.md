# 番外：
> 每个Goroutine在执行之前，都要先知道程序当前的执行状态  
> 通常将这些执行状态封装在一个Context变量中，传递给要执行的Goroutne中  
---

# context作用：
> 用于控制goroutine的生命周期  
> context提供了一种管理多个goroutine的机制  
---

# context包：
> context包不仅实现了在程序单元之间共享变量的方法  
> 同时能通过简单的方法使我们在被调用程序的外部通过设置ctx变量值  
> 将过期或撤销这些信号传递给被调用的程序单元  
> 