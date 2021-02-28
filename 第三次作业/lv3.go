package main

import (
	"fmt"
	"sync"
)
//因为没有弄明白contxt库,所以模仿通道退出写的简易contxt
//init创建退出通道,done返回一个结构体通道,cancel会发送一个空结构体让done可以输出,然后关闭通道

type myContext struct{//创建新的结构体
	exitChan chan struct{}
}
var wg sync.WaitGroup

func(m myContext)init( )chan struct{}{//创建一个无缓存的退出通道并返回赋值
	exitChan:=make(chan struct{})
	return exitChan
}

func(m myContext)done()chan struct{}{//返回结构体里的退出通道
	return m.exitChan
}

func(m myContext)cancel(){//发送空结构体并关闭通道
	m.exitChan<- struct{}{}
	close(m.exitChan)
}

func worker(ctx myContext) {//测试协程
LOOP:
	for {
		fmt.Println("worker")
		select {
		case <-ctx.done()://cancel发送数据后使通道可以输出 执行braek
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main(){
	var ctx myContext
	ctx.exitChan=ctx.init()
	wg.Add(1)
	go worker(ctx)
	ctx.cancel()
	wg.Wait()
	fmt.Println("over")
}
