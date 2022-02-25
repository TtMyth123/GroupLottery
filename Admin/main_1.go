package main

import (
	"fmt"
	"time"
	"ttmyth123/kit/timeKit"
)

func main() {
	t := time.Now()
	aaa := t.Format(timeKit.DateTimeLayout)
	fmt.Println("aaa:", aaa)
	Local_Shanghai := time.FixedZone("Shanghai", 800)
	fmt.Println("Local_Shanghai:", Local_Shanghai)

	t1, e1 := time.ParseInLocation(timeKit.DateTimeLayout, aaa, time.Local)
	t2, e2 := time.ParseInLocation(timeKit.DateTimeLayout, aaa, Local_Shanghai)
	fmt.Println("t1:", t1, e1)
	fmt.Println("t2:", t2, e2)
}
