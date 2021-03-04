package future

import (
	"errors"
	"testing"
)

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}

	t.Run("Success result", func(t *testing.T) {
		future.Success(func(s string) {
			t.Log(s)
		}).Fail(func(e error) {
			t.Fail()
		})
		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})
	})

	t.Run("Error result", func(t *testing.T) {
		future.Success(func(s string) {
			t.Fail()
		}).Fail(func(e error) {
			t.Log(e.Error())
		})
		future.Execute(func() (string, error) {
			return "", errors.New("Error ocurred")
		})
	})
}