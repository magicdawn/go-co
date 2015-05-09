package demo

import "fmt"
import "time"
import "github.com/magicdawn/go-co/main"

func main() {
	logTime()
}

func logTime() {
	t := time.Now()
	fmt.Println("now : ", t)
}
