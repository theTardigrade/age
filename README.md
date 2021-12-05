# golang-age

This is a simple Golang package for determining age from a given date until the present moment. `Calculate` is the only exposed function.

## Example

```golang
package example

import (
	"time"
	"fmt"

	"github.com/theTardigrade/age"
)

func main() {
	const day, month, year = 27, 5, 1960

	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	dateAge := age.Calculate(date)

	fmt.Println(dateAge)
}
```