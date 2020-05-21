package main

type Singleton struct{}

var ins *Singleton

func init() {
	ins = &Singleton{}
}

func GetInstance() *Singleton {
	return ins
}

func main() {}
