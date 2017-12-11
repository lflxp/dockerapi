package main
/*
docker服务自动注册
*/
import (
	"fmt"
	"flag"
	"time"
	"github.com/lflxp/dockerapi/register"
)

var etcdServer *string = flag.String("host","127.0.0.1:2379","etcd服务器地址")
var serverPath *string = flag.String("path","/ams/main/services","docker 服务注册路径")
var interval *int64 = flag.Int64("t",5,"服务注册刷新时间")

func Go() {
	err := register.Register([]string{"127.0.0.1:2379"},*serverPath,*etcdServer,*interval)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func WatchDog() {
	Go()
	gcInterval,_ := time.ParseDuration(fmt.Sprintf("%ds",*interval-1))
	ticker := time.NewTicker(gcInterval)
	go func() {
		for {
			select {
			case <- ticker.C:
				Go()
			}
		}
	}()
}

func main() {
	flag.Parse()
	wait := make(chan int)
	WatchDog()
	<- wait
}