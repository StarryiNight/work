package main

import "fmt"

type dove interface{//鸽子精接口
	gugugu()
}

type repeater interface {//复读机接口
	repeat(string)
}

type ningmeng interface {//柠檬精接口
	acid()
}

type zhenxiang interface {//真香怪接口
	xiang()
}

type person struct{//定义person结构体
	name string//名字
	age   int//年龄
	gender string//性别
}

func (p person)gugugu(){//用person完成gugugu的方法
	fmt.Println("gugugu...")
}

func (p person) repeat(word string){//用person完成复读的方法
	fmt.Println(word)
	fmt.Println(word)
	fmt.Println(word)
}

func (p person)acid(){//用person结构体完成柠檬精的方法
	fmt.Println("真酸")
}

func (p person)xiang(){//用person完成真香怪的方法
	fmt.Println("真香！")
}


func main(){
	p:=person{//定义一个新的person结构体
		name:"小明",
		age:18,
		gender:"男",
	}
	p.gugugu()
	p.repeat("Hello!")
	p.acid()
	p.xiang()//依次调用鸽子精、复读机、真香怪、柠檬精的接口
}
