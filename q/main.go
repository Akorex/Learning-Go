package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)

	if err != nil {
		return fmt.Errorf("in filechecker: %w", err)
	}

	f.Close()
	return nil
}

func main() {
	err := fileChecker("nothere.txt")

	if err != nil {
		fmt.Println(err)

		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}

}
