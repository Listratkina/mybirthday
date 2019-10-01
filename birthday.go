package main

import (
	"fmt"
	"time"
)
func main() {
    d:=time.Date(1988, 02, 12, 22, 10, 0, 0, time.UTC)
	weekday:=d.Weekday()
    for i := 0; i < 32; i++ {
	d= d.AddDate(1, 0, 0)
	weekday=d.Weekday()
	fmt.Println(d, weekday)
}
}
