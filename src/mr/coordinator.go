package mr

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//import "os"

type Coordinator struct {
	// Your definitions here.
	Allfiles      []string // 所有文件名
	NReduce       int      // 用于安排Reduce任务的个数
	Havetaskindex int      //进行安排的任务个数

}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
//func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
//	reply.Y = args.X + 1
//	return nil
//}
func (c *Coordinator) SendFile(args *MyArgs, reply *MyReply) error { //参数：命令类型 回复：尚未执行map操作的文件名
	if args == "map" {
		if c.Havetaskindex < len(c.Allfiles) {
			reply = c.Allfiles[c.Havetaskindex]
			c.Havetaskindex++
		}
	}
	return nil
}

//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	//sockname := coordinatorSock()
	//os.Remove(sockname)
	//l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.
	//如果任务全部完成,就返回真
	if c.Havetaskindex == len(c.Allfiles) {
		ret = true
	}
	return ret
}

//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}
	// Your code here.
	c.Havetaskindex = 0
	c.Allfiles = files
	c.NReduce = nReduce
	c.server()
	return &c
}
