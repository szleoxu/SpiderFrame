package common

import (
	"time"
)

func GetTime(format string) time.Time{
	now := time.Now()
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	t, _ := time.ParseDuration(format)
	newTime := now.Add(t)
	return newTime
}

func GetUnixTime(date string) time.Time{
	//日期转化为时间戳
	timeLayout := "2006-01-02"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	tm, _ := time.ParseInLocation(timeLayout, date, loc)
	return tm
}