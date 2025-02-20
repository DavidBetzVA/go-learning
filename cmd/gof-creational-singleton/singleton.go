package main

import "sync"

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	if s1 == s2 {
		println("Both instances are the same")
	} else {
		println("Instances are different")
	}
}
