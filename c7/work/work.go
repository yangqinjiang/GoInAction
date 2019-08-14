//work包管理一个goroutine池来完成工作
package work

import "sync"

//Worker必须满足接口类型
//才能使用工作池
type Worker interface {
	Task()
}

//Pool提供一个goroutine池,这个池可以完成任何已经提交的Worker任务
type Pool struct {
	work chan Worker
	wg sync.WaitGroup
}

//New创建一个新工作池
func New (maxGoroutines int) *Pool{
	p := Pool{
		work:make (chan Worker),
	}

	p.wg.Add(maxGoroutines)

	//工作池
	for i:=0;i<maxGoroutines;i++ {
		go func() {
			//for range循环会一直阻塞,直到从work通道收到一个Worker接口值
			for w:= range p.work{
				w.Task()
			}
			//一旦work通道被关闭,for range结束,调用 Done,然后goroutine终止
			p.wg.Done()
		}()
	}
	return  &p
}

//Run提交工作到工作池
func (p *Pool)Run(w Worker)  {
	p.work <- w
}
//shutdown等待所有goroutine停止工作
func (p *Pool)Shutdown()  {
	close(p.work)//首先,关闭work通道,这会导致所有池里的goroutine停止工作
	//等待所有goroutine停止工作
	p.wg.Wait()
}
