package main

import "fmt"

type video struct{
	author  string//作者
	name  string//视频名称
	likes int //点赞数
	coin  int //投币数
	collections int //收藏数
	Comments   int //评论数量
}

func (v *video)Like(){//方法：点赞
	v.likes++
}

func (v *video)Coin(){//方法：投币
	v.coin++
}

func (v *video)Collect(){//方法：收藏
	v.collections++
}

func (v *video)Sanlian(){//方法：一键三连
	v.likes++
	v.collections++
	v.coin++
}
func (v *video)status(){//打印当前视频状态
	fmt.Println("视频作者：",v.author)
	fmt.Println("视频名称：",v.name)
	fmt.Println("点赞数：",v.likes)
	fmt.Println("收藏数：",v.collections)
	fmt.Println("投币数：",v.coin)
}

func publish(authorName ,videoName string)video{//发布视频的函数
	v:=video{
		authorName,
		videoName,
		0,
		0,
		0,
		0,
	}
	return v
}
func main(){
	video:=publish("作者名","视频名")//调用发布函数初始化视频结构体
	video.status()

	video.Like()
	video.Collect()
	video.Coin()
	video.status()//分别点赞、收藏、投币

	video.Sanlian()
	video.status()//一键三连
}
