package convert_tool

import "time"

func TimeFormat(time time.Time) string {
	return time.Format("2006-01-02T15:04:05.000")
}
