package time_pro

import "time"

//获取指定时间时间戳
func GetDateLocation(date string) int64 {
	loc, _ := time.LoadLocation("PRC")
	timeLayout := "2006-01-02 15:04:05"
	times, _ := time.ParseInLocation(timeLayout, date, loc)
	timeUnix := times.Unix()

	return timeUnix
}

//获取新保单的起止时间
func GetTime(endTime string) (string, string) {
	//获取当前时间戳
	nowStamp := time.Now().Unix()

	//获取截止日期时间戳
	endStamp := GetDateLocation(endTime)
	beginStamp := endStamp + 1

	if nowStamp >= beginStamp {
		beginStamp = nowStamp
	}

	//获取开始时间
	beginDate := time.Unix(beginStamp, 0).Format("2006-01-02")
	beginDateRe := beginDate + " 00:00:00"

	//截止日期 +1 year
	endTimeCalc := time.Unix(GetDateLocation(beginDateRe), 0).AddDate(1, 0, 0)
	//获取截止日期时间戳 - 1 秒，
	endStampResult := GetDateLocation(endTimeCalc.Format("2006-01-02 00:00:00")) - 1
	//时间戳转时间
	endTimeResult := time.Unix(endStampResult, 0).Format("2006-01-02 15:04:05")

	return beginDateRe, endTimeResult
}
