package format

import (
	"strings"
)

func FormatPrice(price *string) {

	r := strings.Count(*price, "£")

	if r >= 1 {
		splitStr := strings.Split(*price, "£")
		*price = "£" + splitStr[1]
	}

}

func FormatStars(stars *string) {
	if len(*stars) >= 3 {
		*stars = (*stars)[0:3]
	} else {
		*stars = "Unknown"
	}
}
