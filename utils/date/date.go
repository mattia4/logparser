package date

import (
	"fmt"
	"time"
)

func ParseApacheTime(s string) string {
	t, err := time.Parse("02/Jan/2006:15:04:05 -0700", s)
	if err != nil {
		fmt.Println("Errore parsing:", err)
		return ""
	}

	hour := fmt.Sprintf("%02d", t.Hour())
	minute := fmt.Sprintf("%02d", t.Minute())
	second := fmt.Sprintf("%02d", t.Second())

	time := hour + ":" + minute + ":" + second

	return time
}

func ParseApacheDate(s string) string {
	t, err := time.Parse("02/Jan/2006:15:04:05 -0700", s)
	if err != nil {
		fmt.Println("Errore parsing:", err)
		return ""
	}

	date := t.Format("02-01-2006")

	return date
}
