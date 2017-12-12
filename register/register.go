package register

import (
	"fmt"
	"strings"
	"github.com/lflxp/curl"
	"github.com/lflxp/dbui/etcd"
	"github.com/lflxp/dockerapi/api"
)

//获取docker ps运行信息
//注册etcd v3
func Register(endpoint []string,serverPath,server string,t int64) error {
	data,err := api.Ps()
	if err != nil {
		return err
	}
	
	//获取ip
	//etcd 连接 
	//etcd 服务器地址由中控机提供
	st := &etcd.EtcdUi{Endpoints:[]string{server}}
	st.InitClientConn()
	defer st.Close()
	//验证ip
	resp := st.Get("/ams/main/ansible/ip")
	ip := curl.HttpGet(string(resp.Kvs[0].Value))
	// fmt.Println(ip)
	for _,c := range data {
		//注册地址格式
		//服务名-> 镜像名
		serverName := strings.Split(c.Image,"/")[len(strings.Split(c.Image,"/"))-1]
		//协议类型
		err = st.AddLease(fmt.Sprintf("%s/%s",serverPath,serverName),c.Image,t+30)	
		for _,port := range c.Ports {
			//如果没有对外publicPort 则不注册
			if port.PublicPort != 0 {
				protocol := port.Type
				//注册
				err = st.AddLease(fmt.Sprintf("%s/%s/%s@%s:%d",serverPath,serverName,protocol,ip,port.PublicPort),fmt.Sprintf("Names:%s,ID:%s,Status:%s,State:%s",c.Names,c.ID,c.Status,c.State),t)
				if err != nil {
					return err
				}
			}
		}
	}
	
	return nil
}