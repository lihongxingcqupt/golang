package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	//n1, name, n3 := 100, "tom", 88
	//fmt.Print(n1)
	//fmt.Print(name)
	//fmt.Print(n3)
	p5 := person{
		name: "tom",
		city: "beijing",
		age:  16,
	}
	fmt.Printf("p5=%#v\n", &p5.age)
}
