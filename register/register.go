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
		id := strings.Join(strings.Split(c.ID,"")[:12],"")
		//注册地址格式
		//服务名-> 镜像名
		serverName := strings.Split(c.Image,"/")[len(strings.Split(c.Image,"/"))-1]
		//协议类型
		err = conn.AddLease(fmt.Sprintf("%s/%s",serverPath,serverName),c.Image,t+30)	
		for _,port := range c.Ports {
			dcommand := strings.Join(strings.Split(c.Command,"")[:12],"")
			//如果没有对外publicPort 则不注册
			//如果无publicPort 数据设置ip:ID
			//有publicPort 数据设置tcp@ip:publicPort
			if port.PublicPort != 0 {
				protocol := port.Type
				//注册
				//0.0.0.0:32769->389/tcp
				err =conn.AddLease(fmt.Sprintf("%s/%s/%s@%s:%d",serverPath,serverName,protocol,ip,port.PublicPort),fmt.Sprintf("%s::%s::%s::%s::%s::%s::%s",id,c.Image,dcommand,c.State,c.Status,fmt.Sprintf("%s:%d->%d/%s",port.IP,port.PrivatePort,port.PublicPort,port.Type),c.Names),t)
				if err != nil {
					return err
				}
			} else {
				err =conn.AddLease(fmt.Sprintf("%s/%s/%s:%d",serverPath,serverName,ip,id),fmt.Sprintf("%s::%s::%s::%s::%s::%s::%s",id,c.Image,dcommand,c.State,c.Status,"",c.Names),t)
				if err != nil {
					return err
				}	
			}
		}
	}
	
	return nil
}