package register

import (
	"fmt"
	"strings"
	"github.com/lflxp/dbui/etcd"
	"github.com/lflxp/dockerapi/api"
)

//获取docker ps运行信息
//注册etcd v3
func Register(conn *etcd.EtcdUi,serverPath,ip string,t int64) error {
	data,err := api.Ps()
	if err != nil {
		return err
	}
	
	
	// fmt.Println(ip)
	for _,c := range data {
		//注册地址格式
		//服务名-> 镜像名
		serverName := strings.Split(c.Image,"/")[len(strings.Split(c.Image,"/"))-1]
		//协议类型
		err = conn.AddLease(fmt.Sprintf("%s/%s",serverPath,serverName),c.Image,t+30)	
		for _,port := range c.Ports {
			//如果没有对外publicPort 则不注册
			if port.PublicPort != 0 {
				protocol := port.Type
				//注册
				err =conn.AddLease(fmt.Sprintf("%s/%s/%s@%s:%d",serverPath,serverName,protocol,ip,port.PublicPort),fmt.Sprintf("Names:%s,ID:%s,Status:%s,State:%s",c.Names,c.ID,c.Status,c.State),t)
				if err != nil {
					return err
				}
			}
		}
	}
	
	return nil
}