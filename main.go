package goutil


import (
	"fmt"
	"time"
)

func PrintYestoday(){
	fmt.Println(time.Now().Add(10000))
}