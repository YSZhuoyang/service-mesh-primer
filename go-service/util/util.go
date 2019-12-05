package util

import "log"

// CheckError - Panic if error is not nil
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckWarning - Log error if its not nil
func CheckWarning(err error) bool {
	if err != nil {
		log.Printf("Warning: %s\n", err.Error())
		return true
	}

	return false
}
