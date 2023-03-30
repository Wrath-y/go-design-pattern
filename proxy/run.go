package proxy

import "fmt"

func Run() {
	var sub Subject = &Proxy{}
	fmt.Println(sub.Do())
}
