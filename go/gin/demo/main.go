// gin框架使用示例

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始实例化Gin
func initGin() *gin.Engine {
	r := gin.Default()
	return r
}

// 绑定路由规则
func initRouter(r *gin.Engine) {
	// Any可注册一个匹配所有HTTP方法的路由
	r.Any("/any", func(c *gin.Context) { // gin.Context封装了request和response
		c.String(200, "any") // c.String是返回给客户端一个字符串
	})
	// GET可以注册一个匹配GET请求的路由
	// POST可以注册一个匹配POST请求的路由
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})
	// 范绑定
	// 所有url以test/开头的GET请求都走这里
	r.GET("/test/*action", func(c *gin.Context) {
		c.String(200, "范绑定示例")
	})
}

// 路由组routes group
func routesGroup(r *gin.Engine) {
	tmp1Handler := func(c *gin.Context) {
		c.String(200, "this is tmp1")
	}
	tmp2Handler := func(c *gin.Context) {
		c.String(200, "this is tmp2")
	}
	// 声明一个路由组来管理一些相同的URL
	// (路由组也是可以嵌套的)
	g1 := r.Group("/group1")
	{ // 放到大括号里是书写规范，不放大括号也可以
		g1.GET("/tmp1", tmp1Handler) // 127.0.0.1:8080/group1/tmp1走这里
		g1.GET("/tmp2", tmp2Handler)
	}
}

// 设置静态文件的路由
func staticRouter(r *gin.Engine) {
	// 指定静态文件夹(路由，文件夹路径)
	r.Static("/static", "./static")
	// 指定单个静态文件
	r.StaticFile("favicon.ico", "./favicon.ico")
}

// 路由参数解析(c.Param)
func routerParameter(r *gin.Engine) {
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"ID": id,
		})
	})
}

// 获取url里的参数(url里？后面的参数)(c.Query、c.DefaultQuery)
func urlParameter(r *gin.Engine) {
	// POST、GET请求都可以
	// http://127.0.0.1:8080/url_p?first_name=zz&last_name=ii
	r.Any("/url_p", func(c *gin.Context) {
		firstName := c.Query("first_name")            // 若该参数不存在，则返回空字符串
		lastName := c.DefaultQuery("last_name", "yy") // 若该参数不存在，则返回默认值(第二个参数)
		c.String(200, "first_name:%s,last_name:%s", firstName, lastName)
	})
}

// 获取form表单参数(c.PostForm、c.DefaultPostForm)
func bodyParameter(r *gin.Engine) {
	// body里的内容格式通常为四种：
	// application/json						(json格式)
	// application/x-www-form-urlencoded	(form表单)
	// application/xml						(xml格式)
	// multipart/form-data					(文件、图片)
	// PostForm方法默认解析的是 form表单或文件图片格式的参数
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")             // 用PostForm解析，若无该值，则返回空字符串
		email := c.DefaultPostForm("email", "123.com") // 用DefaultPostForm解析，若无该值，则返回默认值
		hobbys := c.PostFormArray("hobby")             // 复选框可用PostFormArray解析为数组
		c.String(200, "username is %s, password is %s, email is %s hobby is %v", username, password, email, hobbys)
	})
}

// 获取表单上传的文件
func upload(r *gin.Engine) {
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file") //my_file为表单里input标签指定的name
		if err != nil {
			c.String(200, "err: %#v", err)
			return
		}
		err = c.SaveUploadedFile(file, "./static/"+file.Filename) // 第二个参数为文件存储的路径
		if err != nil {
			c.String(200, "faile: %#v", err)
			return
		} else {
			c.String(200, "upload success")
		}
	})
}

// 数据解析绑定
func data2struct(r *gin.Engine) {
	// 建立接收数据的结构体
	// 若绑定body里的json数据则声明json的tag
	// 若绑定body里的表单数据或url里的参数，则声明form的tag
	// (若body里的表单参数与url里参数重名，则最终获得的绑定是body的参数)
	type User struct {
		// 若有binding: "required"的tag修饰,则该字段为必选字段，前端数据必须有该字段
		Name      string `form:"username" json:"username" binding:"required"`
		Passworld string `form:"passworld" json:"passworld" binding:"required"`
	}
	// 将接收到的json数据绑定到结构体
	r.POST("/json2struct", func(c *gin.Context) {
		var u User
		// ShouldBind会根据Content-Type的值自行选择对应的绑定引擎
		err := c.ShouldBind(&u) // 将数据解析到结构体实例u里
		if err != nil {
			// c.JSON给客户端返回json格式消息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Printf("%#v\n", u)
		c.JSON(200, gin.H{
			"name":      u.Name,
			"passworld": u.Passworld,
		})
	})
}

// 验证请求参数*******(待解决)
func verify(r *gin.Engine) {
	// gin使用的验证规则是基于Validator.v8
	// https://godoc.org/gopkg.in/go-playground/validator.v8

	//结构体验证(将验证规则定义到结构体的tag上)
	//自定义验证
	//升级验证
}

// 构造响应
func makeResponse(r *gin.Engine) {
	// 构造字符串响应
	r.GET("stringResponse", func(c *gin.Context) {
		c.String(200, "字符串响应")
	})

	// 构造json响应(c.JSON)
	r.GET("/jsonResponse", func(c *gin.Context) {
		c.JSON(200, gin.H{ // 若要将struct实例转为json返回,把gin.H{}替换为struct实例即可
			"key1": "value1",
			"key2": "value2",
		})
	})
	// 构造xml响应(c.XML())
	r.GET("/xmlResponse", func(c *gin.Context) {
		c.XML(200, gin.H{
			"key": "value",
		})
	})
	// 构造YAML响应(c.YAML)
	r.GET("/YAMLResponse", func(c *gin.Context) {
		c.YAML(200, gin.H{
			"key": "value",
		})
	})
	// 构造protobuf格式(c.ProtoBuf)****(待解决)
	r.GET("/protobufResponse", func(c *gin.Context) {

	})
}

// 重定向
func redirect(r *gin.Engine) {
	r.GET("/redirect", func(c *gin.Context) {
		// 支持内部和外部重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})
}

// 异步
func async(r *gin.Engine) {
	// 异步处理请求时，不能直接使用gin.Context,需要取副本使用
	r.GET("/async", func(c *gin.Context) {
		// 在异步处理请求时,后续使用的Context需要是当前Context的副本
		copyContext := c.Copy() // Copy方法返回当前上下文的一个副本，该副本可以安全地在请求范围外使用
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			name := copyContext.Query("name")
			fmt.Printf("请求完成，name=%s\n", name)
		}()
	})
}

// gin中间件
// 中间件作用：
// 在接收到request后到响应response前,可通过中间件做一系列数据处理
// 这样的钩子函数(Hook)就叫做中间件
// 中间件适合处理一些公共的业务逻辑(登陆认证、权限校验、数据分页、记录日志、耗时统计)

// 使用gin自带中间件：
// gin.Default()方法默认启用Logger、Recovery这两个中间件
// gin.New()方法不启用中间件
// Use(中间件)方法，用来启用中间件
func useMiddleware() {
	r := gin.New()
	// 该方式是在全局作用域内使用中间件
	r.Use(gin.Logger())   // 使用日志中间件(默认日志打印到控制台)
	r.Use(gin.Recovery()) // 使用该中间件后，若触发panic，不会导致整个进程崩溃
	r.GET("/useMiddleware", func(c *gin.Context) {
		name := c.DefaultQuery("name", "zy")
		c.String(200, "name is %s\n", name)
	})
}

// 组作用域内使用中间件和单个路由作用域内使用中间件
// 上述方法是在全局作用域内使用中间件
func useMiddlewareInGroup() {
	r := gin.New()
	// 声明一个路由组，并对该组路由启用相应的中间件
	g := r.Group("/middlewareGroup", gin.Logger(), gin.Recovery())
	{
		g.GET("test1", func(c *gin.Context) { c.String(200, "ok") })
		g.GET("test2", func(c *gin.Context) { c.String(200, "ok") })
	}
	// 声明单个路由，并对该路由启用相应中间件
	r.GET("/singleMiddleware", gin.Recovery(), gin.Logger(), func(c *gin.Context) {
		c.String(200, "ok")
	})
}

// 自定义中间件
// Gin中的中间件必须是一个gin.HandlerFunc类型
// 个人认为中间件和普通的路由处理函数没什么区别
// 声明一个路由后，本来就可以传入多个处理函数,中间件不过是被默认传入到其他处理函数的前面
func makeMiddleware(r *gin.Engine) {
	// 自定义一个统计请求耗时的中间件
	myMiddleware := func(c *gin.Context) {
		// 在中间件里可以调用
		// c.Next()		用来调用后续路由处理函数
		// c.Abort()	用来阻止调用后续的处理函数
		startTime := time.Now()
		c.Next() // 调用后续处理函数
		cost := time.Since(startTime)
		fmt.Printf("该请求持续时间为:%v\n", cost)
	}
	// 将自定义的中间件应用到全局
	// (若添加多个中间件，则后加入的中间件先执行)
	r.Use(myMiddleware)
	r.GET("/makeMiddleware", func(c *gin.Context) {
		c.String(302, "ok")
	})
}

// cookie和session
// cookie是解决HTTP协议无状态的方案之一
// cookie实际上就是服务器保存在浏览器上的一段信息
// 浏览器有了cookie之后，每次向服务器发起请求时都会同时将该信息发送给服务器
// 服务器便可根据该信息判断发起请求方的身份
func cookieAndSession(r *gin.Engine) {
	r.GET("cookieTest", func(c *gin.Context) {
		// 先查看客户端请求是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			// 给客户端设置cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/", )
		}
	})
}


func main() {
	r := initGin()
	// makeResponse(r)
	// redirect(r)
	// async(r)
	// makeMiddleware(r)
	cookieAndSession(r)
	r.Run()
}
