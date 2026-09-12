package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func wrappingError(e error) error {
	return fmt.Errorf("error occured during %w", e)
}

func Run22ErrorWrappingDivide() {
	_, err := divide(10, 0)
	if err != nil {
		wrapped := wrappingError(err)
		fmt.Println(wrapped)

		if unWrappedError := errors.Unwrap(wrapped); unWrappedError != nil {
			fmt.Println("Unwrapped:", unWrappedError)
		}
	}
}
