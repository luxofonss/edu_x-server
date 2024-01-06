package common

import "time"

func CurrentTime() string {
	return time.Now().String()
}

func CompareTimeNow(t string) bool {
	return time.Now().String() > t
}
