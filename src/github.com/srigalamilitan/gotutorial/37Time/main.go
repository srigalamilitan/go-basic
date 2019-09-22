package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())

	fmt.Println("===========Convert String to Time")
	timeString := "September 29,2019"
	form := "January 2,2006"
	t1, err := time.Parse(form, timeString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t1)

	fmt.Println("===========Created Date")
	t2 := time.Date(1987, 2, 21, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(1987, 2, 21, 0, 0, 0, 0, time.UTC)
	eq := t2.Equal(t3)

	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(eq)
	fmt.Println("===========Konversi Time To String")

	layout := "2006-01-02"
	timeString2 := t2.Format(layout)
	fmt.Print("Hasil Format : ", timeString2)
}
