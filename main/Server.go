package main

import "StudyZinx/znet"

func main() {
	s := znet.NewServer("[zinx v0.1]")

	s.Serve()
}
