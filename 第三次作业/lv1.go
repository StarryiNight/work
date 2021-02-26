package main

import "fmt"

func putNum(numChan chan int) { //把3-50000的奇数放入numChan通道
	for i := 3; i <= 50000; i += 2 {
		numChan <- i
	}

	close(numChan)
}

func prime(numChan chan int, primeChan chan int, exitChan chan bool) { //判断是否是素数并放入primeChan
	primeChan <- 2 //2也放入素数通道
	var flag bool  //判断素数的标志
	for {
		num, ok := <-numChan //从numchan取数
		if !ok {
			break //取完就退出
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false //标记不是素数
				break
			}
		}
		if flag { //是素数就把数放入primeChan
			primeChan <- num
		}
	}
	exitChan <- true //标志一个协程已经完成

}
func main() {
	numChan := make(chan int, 5000)    //存放3-50000的奇数的通道
	primeChan := make(chan int, 10000) //存放素数的通道
	exitChan := make(chan bool, 8)     //退出通道

	go putNum(numChan)
	for i := 0; i < 8; i++ {
		go prime(numChan, primeChan, exitChan) //启动8个协程并发
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan //等待8个协程都完成
		}

		close(primeChan) //关闭素数通道
	}()

	for { //打印素有存入primeChan的素数
		primeNum, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(primeNum)
	}

}
