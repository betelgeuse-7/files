// Package files provides create/delete/update operations on files.
package files

import (
	"fmt"
	"os"
)

func BulkCreate(prefix, extension string, howMany int) {
	var counter int
	for i := 0; i < howMany; i++ {
		os.Create(prefix + fmt.Sprint(counter) + extension)
		counter++
	}
}
