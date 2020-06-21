package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		1年=365天 day
		1天=24小时 hour
		1小时=3600秒 second
		1秒=1000ms millisecond
		1ms = 1000us微秒 microsecond
		1us = 1000ns纳秒 nanosecond
		1ns = 1000ps皮秒 picosecond
	*/

	t1 := time.Now()
	fmt.Printf("%T\n", t1)

	t2 := time.Date(2008, 7, 6, 5, 4, 3, 0, time.Local)
	fmt.Println(t2)

	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)

	s2 := t1.Format("2006/1/2")
	fmt.Println(s2)

	s3 := "1999年10月10日"
	t3, err := time.Parse("2006年01月02日", s3)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(t3)
	fmt.Printf("%T\n", t3)

	fmt.Println(t1.String())

	year, month, day := t1.Date()
	fmt.Println(year, month, day)

	hour, min, sec := t1.Clock()
	fmt.Println(hour, min, sec)

	year2 := t1.Year()
	fmt.Println("年：", year2)
	fmt.Println(t1.YearDay())

	month2 := t1.Month()
	fmt.Println("月：", month2)
	fmt.Println("日：", t1.Day())
	fmt.Println("时：", t1.Hour())
	fmt.Println("分：", t1.Minute())
	fmt.Println("秒：", t1.Second())
	fmt.Println("纳秒：", t1.Nanosecond())

	fmt.Println(t1.Weekday())

	t4 := time.Date(1970, 1, 1, 1, 1, 0, 0, time.UTC)
	timeStamp1 := t4.Unix()
	fmt.Println(timeStamp1)
	timeStamp2 := t1.Unix()
	fmt.Println(timeStamp2)

	timeStamp3 := t4.UnixNano()
	fmt.Println(timeStamp3)
	timeStamp4 := t1.UnixNano()
	fmt.Println(timeStamp4)

	t5 := t1.Add(time.Minute)
	fmt.Println(t1)
	fmt.Println(t5)
	fmt.Println(t1.Add(24 * time.Hour))

	t6 := t1.AddDate(1, 0, 0)
	fmt.Println(t6)

	d1 := t5.Sub(t1)
	fmt.Println(d1)

	time.Sleep(3 * time.Second)
	fmt.Println("main ending...")

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1
	fmt.Println(randNum)
	time.Sleep(time.Duration(randNum) * time.Second)

	fmt.Println("main ending111...")

}
