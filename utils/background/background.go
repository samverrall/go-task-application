package background

import (
	"log"
)

func Go(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()

		fn()
	}()
}
