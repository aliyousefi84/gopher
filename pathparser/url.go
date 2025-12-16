package pathparser

import (
	"net/url"
	"strings"
)

func Urlstr (Url string) string {
	url , _ := url.Parse(Url)
	validurl := strings.Trim(Url , url.RawQuery)
	cutstring := strings.Trim(validurl , "https://")
	slicestr := strings.Split(cutstring , "/")
	return slicestr[len(slicestr)-1]
}
