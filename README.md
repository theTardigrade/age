# golang-age

This is a simple Go package for determining age.

The most important public function is `CalculateToNow`, which determines an age based on the duration between a given date and the present moment. There is also a more general function named `Calculate`, which determines an age based on the duration between any two given dates.

Finally, a helper function named `IsLeapYear` is also available to determine whether a given year contains a leap day or not. 

[![Go Reference](https://pkg.go.dev/badge/github.com/theTardigrade/golang-age.svg)](https://pkg.go.dev/github.com/theTardigrade/golang-age) [![Go Report Card](https://goreportcard.com/badge/github.com/thetardigrade/golang-age)](https://goreportcard.com/report/github.com/thetardigrade/golang-age)

## Example

```golang
package main

import (
	"time"
	"fmt"

	age "github.com/theTardigrade/golang-age"
)

func main() {
	const day, month, year = 27, 5, 1960

	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	dateAge := age.CalculateToNow(date)

	fmt.Println(dateAge)

	wasLeapYear := age.IsLeapYear(year)

	fmt.Println(wasLeapYear) // true
}
```

## Support

If you use this package, or find any value in it, please consider donating:

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S2EIRL0)