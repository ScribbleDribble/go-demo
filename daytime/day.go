package daytime

import (
	"time"
	"fmt"
)

func main() {
	PrintTime();
}

func PrintTime() {
	fmt.Println(time.Now())
}