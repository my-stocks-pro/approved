package redis

import "fmt"

func Post(data []byte, done chan bool) {
	fmt.Println(string(data))


	done <- true
	return
}
