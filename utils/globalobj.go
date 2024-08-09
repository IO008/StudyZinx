package utils

import (
	"StudyZinx/ziface"
	"encoding/json"
	"io/ioutil"
)

type GlobalObj struct {
	/*
		server
	*/
	TCPServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	/*
		zinx
	*/
	Version          string
	MaxPacktSize     uint32
	MaxConn          int
	WorkerPoolSize   uint32
	MaxWorkerTaskLen uint32
	MaxMsgChanLen    uint32

	/*
		config file path
	*/
	ConfFilePath string
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
		Name:             "ZinxServerApp",
		Version:          "v0.4",
		TcpPort:          7777,
		Host:             "0.0.0.0",
		MaxConn:          12000,
		MaxPacktSize:     4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    10,
	}
	GlobalObject.Reload()
}
