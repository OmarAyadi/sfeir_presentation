package test

import (
	"reflect"
	"testing"
)

func ErrorShouldBeEqualOrFail(t *testing.T, expectedError error, actualError error) {
	equal := expectedError != nil && actualError != nil && expectedError.Error() == actualError.Error()
	if equal {
		return
	}

	t.Fatalf(ExpectedSameValuesErr(expectedError, actualError))
}

func ShouldBeEqualOrFail(t *testing.T, expected any, actual any) {
	if expected == nil && actual == nil {
		return
	}

	if (expected != nil && actual != nil) &&
		(reflect.DeepEqual(expected, actual) || expected == actual) {
		return
	}

	t.Fatalf(ExpectedSameValuesErr(expected, actual))
}

func ShouldNotBeEqualOrFail(t *testing.T, expected any, actual any) {
	if !reflect.DeepEqual(actual, expected) {
		return
	}

	t.Fatalf(ExpectedDifferentValuesErr(actual))
}

func ShouldBeNullOrFail(t *testing.T, value any) {

	if value == nil || reflect.ValueOf(value).IsNil() {
		return
	}

	t.Fatalf(ExpectedNullValueErr(value))
}

func ShouldBeTrueOrFail(t *testing.T, value bool) {
	ShouldBeEqualOrFail(t, true, value)
}

func ShouldBeFalseOrFail(t *testing.T, value bool) {
	ShouldBeEqualOrFail(t, false, value)
}

func ShouldNotBeNullOrFail(t *testing.T, value any) {
	if value != nil {
		return
	}

	t.Fatalf(ExpectedNotNullValueErr(value))
}

func ShouldNotBeEmptyOrFail[T any](t *testing.T, values []T) {
	if len(values) != 0 {
		return
	}
	t.Errorf(ExpectedValuesToNotBeEmptyErr(values))
}

func ShouldBeEmptyOrFail[T any](t *testing.T, values []T) {
	if len(values) == 0 {
		return
	}
	t.Errorf(ExpectedValuesToBeEmptyErr(values))
}
