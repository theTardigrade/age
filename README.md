# age

__If you use this package, please consider donating at PayPal:__ [https://www.paypal.me/jismithpp](https://www.paypal.me/jismithpp)

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