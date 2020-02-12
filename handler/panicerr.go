package handler

import "log"

// PanicErr : if err, panic(err)
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

// LogErr : if err, log.Fatal(mgs, err)
func LogErr(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
