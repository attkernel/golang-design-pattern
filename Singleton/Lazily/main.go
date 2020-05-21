package main

type Singleton struct{}

var ins *Singleton

func GetInstance() *Singleton {
	if ins == nil {
		ins = &Singleton{}
	}
	return ins
}

func main() {}
