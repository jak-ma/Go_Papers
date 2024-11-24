package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.OpenFile("D:\\GO\\gopath\\src\\go_papers\\paper4\\lv2\\content.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Printf("Open file error :%v ", err)
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	events_date := make(map[time.Time]string)
	for {
		lineData, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(lineData) > 0 {
					lineData = strings.Trim(lineData, "\n")
					s := strings.Split(lineData, " ")
					parsedtime, _ := time.Parse("2006-01-02", s[0])
					if parsedtime.Year() == 2024 {
						events_date[parsedtime] = s[1]
					} else {
						events_date[parsedtime.AddDate(1, 0, 0)] = s[1]
					}
				}
			}
			break
		} else {
			lineData = strings.Trim(lineData, "\n")
			s := strings.Split(lineData, " ")
			parsedtime, _ := time.Parse("2006-01-02", s[0])
			if parsedtime.Year() == 2024 {
				events_date[parsedtime] = s[1]
			} else {
				events_date[parsedtime.AddDate(1, 0, 0)] = s[1]
			}
		}
	}
	var nearestdate time.Time
	var farestdate time.Time
	var nearestdays int64 = 365
	var farestdays int64 = 0
	for k := range events_date {
		sub := k.Sub(time.Now())
		subdays := int64(sub.Hours() / 24)
		if subdays == 0 {
			fmt.Printf("今天过节:%v", events_date[k])
			return
		}
		if sub > 0 {
			if nearestdays > subdays {
				nearestdays = subdays
				nearestdate = k
			}
		} else {
			if farestdays > subdays {
				farestdays = subdays
				farestdate = k
			}
		}
	}
	if nearestdays != 365 {
		fmt.Printf("最近的一个节日:%v\n还有%d天", events_date[nearestdate], nearestdays)
	} else {
		fmt.Printf("最近的一个节日:%v\n还有%d天", events_date[farestdate], farestdays+365)
	}
}
