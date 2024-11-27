package lib

import "time"

func GetTimeStamp(t time.Time) string {
	return t.Format("01/02/06")
}
