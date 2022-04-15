package test

import "fmt"

func ExpectedSameValuesErr[T any, R any](expected T, found R) string {
	return fmt.Sprintf("Expected '%v', found '%v'", expected, found)
}

func ExpectedDifferentValuesErr[R any](found R) string {
	return fmt.Sprintf("Expected: not equal, but found '%v'", found)
}

func ExpectedNotNullValueErr[T any](found T) string {
	return fmt.Sprintf("Expected '%v', to be not null", found)
}

func ExpectedNullValueErr[T any](found T) string {
	return fmt.Sprintf("Expected '%v', to be not null", found)
}

func ExpectedValuesToBeEmptyErr[T any](found []T) string {
	return fmt.Sprintf("Expected '%v', to be empty", found)
}

func ExpectedValuesToNotBeEmptyErr[T any](found []T) string {
	return fmt.Sprintf("Expected '%v', to not be  empty", found)
}
