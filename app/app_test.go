package app_test

import (
	"testing"

	"github.com/ahmedsameha1/ccjsonparser/app"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	result, err := app.App(func(name string) ([]byte, error) {
		if name != "valid.json" {
			panic("error")
		}
		return []byte("{}"), nil
	}, []string{"ccjsonparser", "valid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(""), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)
}
