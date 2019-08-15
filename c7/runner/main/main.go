package main

import (
	"time"
	"log"
	"github.com/yangqinjiang/GoInAction/c7/runner"
	"os"
)
//timeout 规定了必须在多少秒内处理完成
const timeout  = 3 * time.Second
func main() {
	log.Println("Starting work.")
	//为本次执行分配超时时间
	r := runner.New(timeout)

	//加入要执行的任务
	r.Add(createTask(),createTask(),createTask())

	//执行任务并处理结果
	if err := r.Start();err != nil{
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}
	log.Println("Process ended")
}
/**
2019/08/15 17:04:55 Starting work.
2019/08/15 17:04:55 Processor - Task #0. doing
2019/08/15 17:04:55 Processor - Task #0. DONE
2019/08/15 17:04:55 Processor - Task #1. doing
2019/08/15 17:04:56 Processor - Task #1. DONE
2019/08/15 17:04:56 Processor - Task #2. doing
2019/08/15 17:04:58 Terminating due to timeout
 */

//createTask返回一个根据id休眠指定秒数的示例任务
func createTask() func(int)  {
	return func(id int) {
		log.Printf("Processor - Task #%d. doing",id)
		time.Sleep(time.Duration(id)*time.Second)
		log.Printf("Processor - Task #%d. DONE",id)
	}
}
