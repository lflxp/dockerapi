package main

/*
docker服务自动注册
*/
import (
	"flag"
	"fmt"
	"github.com/lflxp/dbui/etcd"
	"github.com/lflxp/dockerapi/register"
	"time"
)

var etcdServer *string = flag.String("host", "127.0.0.1:2379", "etcd服务器地址")
var serverPath *string = flag.String("path", "/ams/main/services", "docker 服务注册路径")
var interval *int64 = flag.Int64("t", 5, "服务注册刷新时间")
var Ip *string = flag.String("ip", "127.0.0.1", "宿主机ip")
var Conn *etcd.EtcdUi

func init() {
	flag.Parse()
	//初始化etcd连接
	//获取ip
	//etcd 连接
	//etcd 服务器地址由中控机提供
	st := &etcd.EtcdUi{Endpoints: []string{*etcdServer}}
	st.InitClientConn()
	Conn = st
	// defer st.Close()
}

func Go() {
	err := register.Register(Conn, *serverPath, *Ip, *interval)
	if err != nil {
		fmt.Println("Go",err.Error())
	}
}

func WatchDog() {
	Go()
	gcInterval, _ := time.ParseDuration(fmt.Sprintf("%ds", *interval-1))
	ticker := time.NewTicker(gcInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				Go()
			}
		}
	}()
}

func main() {
	wait := make(chan int)
	WatchDog()
	<-wait
}
