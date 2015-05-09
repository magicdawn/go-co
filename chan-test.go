package chanTest

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	logTime()

	go func() {
		time.Sleep(time.Second * 2)
		c <- 10
	}()

	println(<-c)
	logTime()
}

func logTime() {
	t := time.Now()
	fmt.Println("now : ", t)
}
