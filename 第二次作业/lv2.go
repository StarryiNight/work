package main

import "fmt"

func Receiver(v interface{}){//接收者函数
	switch v.(type){//用接口断言来判断传入空接口的变量类型
	case string:
		fmt.Println("这个是string")
	case int:
		fmt.Println("这个是int")
	case bool:
		fmt.Println("这个是bool")

	}
}

func main(){
	Receiver("你好吗")//传入string
	Receiver(32)//传入int
	Receiver(true)//传入bool
}
