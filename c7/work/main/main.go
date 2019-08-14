package main

import (
	"github.com/yangqinjiang/GoInAction/c7/work"
	"sync"
	"log"
	"time"
)
//names 提供了一组用来显示的名字
var names = []string{
"steve","bob","mary","therese","jason",
}
//namePrinter使用特定方式打印名字
type namePrinter struct {
	name string
}

//Task实现Worker接口
func (m *namePrinter)Task(){
	log.Println(m.name,":任务准备开始处理...")
	time.Sleep(time.Second)//睡眠1s
	log.Println(m.name,":任务已经处理完成")
}

func main()  {
	//使用两个goroutine来创建工作池
	p := work.New(2)

	times := 100
	var wg sync.WaitGroup
	wg.Add(times *len(names))//等待wg的次数

	for i:=0;i<times;i++{
		//迭代names切片
		log.Println("开始迭代names切片",i)
		for _,name := range names{
			//创建一个namePrinter并提供指定的名字
			np := namePrinter{
				name:name,
			}

			go func(){
				//将任务提交执行,当Run返回时,
				//我们就知道任务已经处理完成
				log.Println("将任务",np.name,"提交执行")
				p.Run(&np)

				wg.Done()
			}()
		}
		log.Println("完成迭代names切片",i)
	}
	wg.Wait()

	//让工作池停止工作,等待所有现有的工作完成
	log.Println("让工作池停止工作,等待所有现有的工作完成")
	p.Shutdown()

}
