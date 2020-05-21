package main

import (
	"sync"
)

type Singleton struct{}

var ins *Singleton
var mx sync.Mutex

func GetInstance() *Singleton {
	if ins == nil {
		mx.Lock()
		defer mx.Unlock()
		if ins == nil {
			ins = &Singleton{}
		}
	}
	return ins
}

func main() {}
