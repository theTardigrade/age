# golang-age

This is a simple Go package for determining age. `Calculate` is the main exposed function, which calculates an based on the duration between a given date and the present moment, but `IsLeapYear` can also be called to determine whether a given year contains a leap day or not. 

[![Go Reference](https://pkg.go.dev/badge/github.com/theTardigrade/golang-age.svg)](https://pkg.go.dev/github.com/theTardigrade/golang-age) [![Go Report Card](https://goreportcard.com/badge/github.com/thetardigrade/golang-age)](https://goreportcard.com/report/github.com/thetardigrade/golang-age)

## Example

```golang
package example

import (
	"time"
	"fmt"

	age "github.com/theTardigrade/golang-age"
)

func main() {
	const day, month, year = 27, 5, 1960

	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	dateAge := age.Calculate(date)

	fmt.Println(dateAge)
}
```

## Support

If you use this package, or find any value in it, please consider donating:

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S2EIRL0)