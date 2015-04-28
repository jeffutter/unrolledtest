package utils

import (
	"log"
)

func PanicIf(err error) {
	if err != nil {
		log.Println("FATAL", err)
		panic(err)
	}
}
