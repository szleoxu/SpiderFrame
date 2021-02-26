package common

import (
	"strings"
	"time"
)

//Get current date add and subtract
func GetCustomTime(format int,duration int64, durationTime time.Duration) string {
	now := time.Now().Add(time.Duration(duration) * durationTime)
	nowTime := "0"
	if format == 1 {
		nowTime = now.Format("2006-01-02")
	} else if format == 2 {
		nowTime = now.Format("2006-01-02 15:04:05")
	}
	return nowTime
}

//Get current date time
//Support only format:2006-01-02 and 2006-01-02 15:04:05
func GetTime(format int) string {
	nowTime := "0"
	nowTime = GetCustomTime(format, 0, time.Second)
	return nowTime
}

//Date convert to unix time
//Support only format:2006-01-02 and 2006-01-02 15:04:05
func GetUnixTime(date string) int64 {
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	dateA := strings.Split(date, " ")
	if len(dateA) == 1 {
		timeLayout = "2006-01-02"
	}
	loc, _ := time.LoadLocation("Local") //获取时区
	tm, _ := time.ParseInLocation(timeLayout, date, loc)
	return tm.Unix()
}

//Get current week day
//PS:sunday is 0
func GetWeekDay() int{
	t := time.Now()
	weekNum:=int(t.Weekday())
	return weekNum
}
