package app_test

import (
	"errors"
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
		if name != "valid.json" {
			panic("error")
		}
		return []byte(`""`), nil
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
	assert.Equal(t, "MUST be an object, array, number, or string, or false or null or true", result)

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
			"\"": "",
			"": "\"",
			"": "\" \"",
			"\" \"": "",
			"backslash": "\\",
  			"\\": "backslash",
			"slash": "/ & \/",
  			"/ & \/": "slash", 
			"controls": "\b\f\n\r\t",
    		"\b\f\n\r\t": "controls",
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
		return []byte(`{"key":"	tab	character	in	string	"}`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`["tab\   character\   in\  string\  "]`), nil
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
			""
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
			"key2": "value",
			"": "",
			" ": " ",
			" ": "",
			"": " ",
			"key 3" : "value 3",
			" key 4" : "value 4 ",
			"key 5 ": " value 5",
			"":
			"
			"
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
			"key2": "value",
			"": "",
			" ": " ",
			" ": "",
			"": " ",
			"key 3" : "value 3",
			" key 4" : "value 4 ",
			"key 5 ": " value 5",
			"\"  wief"gbi": "gwoeh"
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
			"key2": "value",
			"": "",
			" ": " ",
			" ": "",
			"": " ",
			"key 3" : "value 3",
			" key 4" : "value 4 ",
			"key 5 ": " value 5",
			"gbi": "\"  wief"gbi"
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "Illegal backslash escape: \x15"
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"Illegal backslash escape: \x15": "value"
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
			"key5": 2.2,
			"key6": 1.234567890E+34,
			"key7": 0.123456789e-12,
			"key8": 0.1e1
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
			"key5": 5,
			"key6": 1e1,
  			"key7": 1e-1,
			"key8": 1e00,
  			"key9": 2e+00,
  			"key10": 2e-00,
			"key11": 23456789012E66
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
			"key5": 0101
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key5": 01.01
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key5": 023456789012E66
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key5": 02.3456789012E66
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key5": -0101
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key5": -01.01
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key5": -023456789012E66
		  }`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "invalid.json" {
			panic("error")
		}
		return []byte(`{
			"key5": -02.3456789012E66
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
		if name != "valid.json" {
			panic("error")
		}
		return []byte(`{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": -101
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
			"key5": -1.1
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
			"key5": -1
		  }`), nil
	}, []string{"ccjsonparser", "valid.json"})
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
		if name != "valid.json" {
			panic("error")
		}
		return []byte(`{
			"key": "value",
			"key-n": 101,
			"key-o": {},
			"key-l": []
		  }`), nil
	}, []string{"ccjsonparser", "valid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "valid.json" {
			panic("error")
		}
		return []byte(`[
			"value",
			 101,
			 {},
			 []
		  ]`), nil
	}, []string{"ccjsonparser", "valid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is a valid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "valid.json" {
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
	}, []string{"ccjsonparser", "valid.json"})
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
				"inner key: "inner value",
				"key-l": [
					"list value"
				]
			}
		}`), nil
	}, []string{"ccjsonparser", "invalid.json"})
	assert.NoError(t, err)
	assert.Equal(t, "This is an invalid JSON", result)

	result, err = app.App(func(name string) ([]byte, error) {
		if name != "doesntexist.json" {
			panic("error")
		}
		return nil, errors.New("no such file")
	}, []string{"ccjsonparser", "doesntexist.json"})
	assert.Contains(t, err.Error(), "no such file")
	assert.Equal(t, "", result)
}
