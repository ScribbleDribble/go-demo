package daytime

import (
	"time"
	"fmt"
)


func PrintTime() {
	fmt.Println(time.Now())
}

// not publicly accessible 
func printYo() {
	fmt.Println("yo")
}