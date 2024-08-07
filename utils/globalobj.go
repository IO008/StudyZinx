package utils

import (
	"StudyZinx/ziface"
	"encoding/json"
	"io/ioutil"
)

type GlobalObj struct {
	TCPServer    ziface.Iserver
	Host         string
	TcpPort      int
	Name         string
	Version      string
	MaxPacktSize uint32
	MaxConn      int
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:         "ZinxServerApp",
		Version:      "v0.4",
		TcpPort:      7777,
		Host:         "0.0.0.0",
		MaxConn:      12000,
		MaxPacktSize: 4096,
	}
	GlobalObject.Reload()
}
