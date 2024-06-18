/**
 * Author 李艺
 * Email 9830131@qq.com
 * 2018/03/23
 */
 package main

 import (
		 "github.com/kataras/iris/v12"
 )
 
 // 使用 go-bindata 嵌入了静态资源
 // go-bindata -o assets.go assets/
 
 func main() {
		 app := iris.New()
 
		 // 小程序 https://api.douban.com/v2/movie
		 api := app.Party("/v2/movie")
		 {
				 // 模拟 https://api.douban.com/v2/movie/coming_soon?start=0&count=3
				 // http://localhost:8080/v2/movie/coming_soon?start=0&count=3
				 api.Get("/coming_soon", func(ctx iris.Context) {
						 data, err := Asset("assets/coming_soon.json")
						 if err != nil {
								 ctx.StatusCode(iris.StatusNotFound)
								 ctx.WriteString("Resource not found")
								 return
						 }
						 ctx.ContentType("application/json")
						 ctx.WriteString(string(data))
				 })
 
				 // https://api.douban.com/v2/movie/search?q=${this.data.searchWords}&start=${start}&count=${this.data.size}
				 // http://localhost:8080/v2/movie/search?q=和平&start=0&count=10
				 api.Get("/search", func(ctx iris.Context) {
						 data, err := Asset("assets/search.json")
						 if err != nil {
								 ctx.StatusCode(iris.StatusNotFound)
								 ctx.WriteString("Resource not found")
								 return
						 }
						 ctx.ContentType("application/json")
						 ctx.WriteString(string(data))
				 })
 
				 // https://api.douban.com/v2/movie/${this.data.type}?start=${start}&count=${this.data.size}
				 // http://localhost:8080/v2/movie/in_theaters?start=0&count=10
				 api.Get("/in_theaters", func(ctx iris.Context) {
						 data, err := Asset("assets/in_theaters.json")
						 if err != nil {
								 ctx.StatusCode(iris.StatusNotFound)
								 ctx.WriteString("Resource not found")
								 return
						 }
						 ctx.ContentType("application/json")
						 ctx.WriteString(string(data))
				 })
 
				 // https://api.douban.com/v2/movie/subject/${options.id}
				 // http://localhost:8080/v2/movie/subject/3267549
				 api.Get("/subject/{id:string regexp(^[0-9]+$)}", func(ctx iris.Context) {
						 data, err := Asset("assets/subject_3267549.json")
						 if err != nil {
								 ctx.StatusCode(iris.StatusNotFound)
								 ctx.WriteString("Resource not found")
								 return
						 }
						 ctx.ContentType("application/json")
						 ctx.WriteString(string(data))
				 })
 
				 // https://api.douban.com/v2/movie/bang/${board.key}?start=0&count=10
				 // http://localhost:8080/v2/movie/top250?start=0&count=10
				 api.Get("/top250", func(ctx iris.Context) {
						 data, err := Asset("assets/top250.json")
						 if err != nil {
								 ctx.StatusCode(iris.StatusNotFound)
								 ctx.WriteString("Resource not found")
								 return
						 }
						 ctx.ContentType("application/json")
						 ctx.WriteString(string(data))
				 })
		 }
 
		 // app.Get("/user/{id:string regexp(^[0-9]+$)}")
		 app.Get("/", func(ctx iris.Context) {
				 data, err := Asset("assets/index.html")
				 if err != nil {
						 ctx.StatusCode(iris.StatusNotFound)
						 ctx.WriteString("Resource not found")
						 return
				 }
				 ctx.ContentType("text/html")
				 ctx.WriteString(string(data))
		 })
 
		 // app.Run(iris.Addr(":8081"))
		 app.Run(iris.Addr(":8080"))
 }
 