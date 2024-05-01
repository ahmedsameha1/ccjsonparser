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

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "valid.json" {
			panic("error")
		}
		return []byte(`{"key": "value"}`), nil
	}, []string{"ccjsonparser", "valid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "valid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "value",
			"key2": "value"
		  }`), nil
	}, []string{"ccjsonparser", "valid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "valid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "value",
			"key2": "value",
			"": "",
			" ": " ",
			" ": "",
			"": " ",
			"key 3" : "value 3",
			" key 4" : "value 4 ",
			"key 5 ": " value 5"
		  }`), nil
	}, []string{"ccjsonparser", "valid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{"key": "value",}`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "value",
			key2: "value"
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{"key":value","key":"value"}`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{"key":value","key":
		"value"}`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "value",
			"key2": "value",
			"": "",
			" ": " ",
			" ": "",
			"": " ",
			"key 3" : "value 3",
			" key 4" : "value 4 ",
			"key 5 ": " value 5",
			"
			":
			"
			"
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "valid.json" {
			panic("error")
		}
		return []byte(`{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 2.2
		  }`), nil
	}, []string{"ccjsonparser", "valid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "valid.json" {
			panic("error")
		}
		return []byte(`{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 5
		  }`), nil
	}, []string{"ccjsonparser", "valid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "valid.json" {
			panic("error")
		}
		return []byte(`{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 642
		  }`), nil
	}, []string{"ccjsonparser", "valid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key1": True,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 101
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key1": true,
			"key2": false,
			"key3": nulll,
			"key4": "value",
			"key5": 101
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 101true
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key1": true,
			"key2": false,
			"key3": nulll,
			"key4": "value",
			"key5": 101true
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": -101
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": -1.1
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": -1
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": ["list value"
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": "list value"]
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "value",
			"key-n": 101,
			"key-o": 
			  "inner key": "inner value"
			},
			"key-l": ["list value"]
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			,
			"key-l": ["list value"]
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "value",
			"key-n": 101,
			"key-o": {},
			"key-l": []
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`[
			"value",
			 101,
			 {},
			 []
		  ]`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value",
			  "inner key2": [1, true, "hi"]
			},
			"key-l": ["list value", {"key3": "value 3"}]
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)
}
