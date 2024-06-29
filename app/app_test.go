package app_test

import (
	"errors"
	"testing"

	"github.com/ahmedsameha1/ccjsonparser/app"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	var tests = []struct {
		fileName    string
		fileContent string
		result      string
		err         error
	}{
		{fileName: "valid.json", fileContent: "{}", result: "This is a valid JSON", err: nil},
		{fileName: "valid.json", fileContent: `""`, result: "This is a valid JSON", err: nil},
		{fileName: "valid.json", fileContent: `null`, result: "This is a valid JSON", err: nil},
		{fileName: "valid.json", fileContent: `false`, result: "This is a valid JSON", err: nil},
		{fileName: "valid.json", fileContent: `true`, result: "This is a valid JSON", err: nil},
		{fileName: "valid.json", fileContent: `-74.23`, result: "This is a valid JSON", err: nil},
		{fileName: "valid.json", fileContent: `0`, result: "This is a valid JSON", err: nil},
		{fileName: "invalid.json", fileContent: "", result: "",
			err: errors.New("This is an invalid JSON\nMUST be an object, array, number, or string, or false or null or true")},
		{fileName: "invalid.json", fileContent: "+83", result: "",
			err: errors.New("This is an invalid JSON\nAn invalid number, there is a leading +")},
		{fileName: "invalid.json", fileContent: "Null", result: "",
			err: errors.New("This is an invalid JSON\nShould be \"null\"")},
		{fileName: "invalid.json", fileContent: "False", result: "",
			err: errors.New("This is an invalid JSON\nShould be \"false\"")},
		{fileName: "invalid.json", fileContent: "True", result: "",
			err: errors.New("This is an invalid JSON\nShould be \"true\"")},
		{fileName: "invalid.json", fileContent: "078", result: "",
			err: errors.New("This is an invalid JSON\nAn invalid number, there is a leading zero")},
		{fileName: "invalid.json", fileContent: `"string1", "string 2"`, result: "",
			err: errors.New("This is an invalid JSON\nMultiple values outside of an array")},
		{fileName: "invalid.json", fileContent: `"string1" "string 2"`, result: "",
			err: errors.New("This is an invalid JSON\nMultiple values outside of an array")},
		{fileName: "invalid.json", fileContent: "\"string1\",\n \"string 2\"", result: "",
			err: errors.New("This is an invalid JSON\nMultiple values outside of an array")},
		{fileName: "invalid.json", fileContent: `"str\074b"`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an invalid string")},
		{fileName: "valid.json", fileContent: `{"key": "value"}`, result: "This is a valid JSON",
			err: nil},
		{fileName: "valid.json", fileContent: `{
			"key": "value",
			"key2": "value"
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileName: "valid.json", fileContent: `{
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
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileName: "invalid.json", fileContent: `["key", "value"`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an unclosed array")},
		{fileName: "invalid.json", fileContent: `["key", "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is closed as an object")},
		{fileName: "invalid.json", fileContent: `["key", "value",]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that contains extra tail comma(s)")},
		{fileName: "invalid.json", fileContent: `["key", "value",,]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that contains extra tail comma(s)")},
		{fileName: "invalid.json", fileContent: `[,"key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that contains extra advancing comma(s)")},
		{fileName: "invalid.json", fileContent: `[,,"key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that contains extra advancing comma(s)")},
		{fileName: "invalid.json", fileContent: `[,"key", "value",]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that contains extra advancing comma(s)")},
		{fileName: "invalid.json", fileContent: `["key", "value"]]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `[["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `]["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `{["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `{ ["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `{ { ["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `}["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `}["key", "value"] ]]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `["key", "value"][`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `["key", "value"]{`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `["key", "value"]}`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `["key", "value"]} ]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `[["key", "value"]{`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `{["key", "value"]]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `]["key", "value"][`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `}["key", "value"]{`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `[["key", "value"]}`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `{["key", "value"]]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileName: "invalid.json", fileContent: `,["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileName: "invalid.json", fileContent: `["key", "value"],`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileName: "invalid.json", fileContent: `, ["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileName: "invalid.json", fileContent: `["key", "value"] ,`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileName: "invalid.json", fileContent: `, ["key", "value"] ,`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileName: "invalid.json", fileContent: `,, ["key", "value"] ,,`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileName: "invalid.json", fileContent: `{"key": "value",}`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			key2: "value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{"key":value","key":"value"}`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{"key":value","key":
		"value"}`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{"key":"	tab	character	in	string	"}`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `["tab\   character\   in\  string\  "]`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
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
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
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
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
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
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
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
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key": "Illegal backslash escape: \x15"
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"Illegal backslash escape: \x15": "value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "valid.json", fileContent: `{
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
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileName: "valid.json", fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 642
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileName: "valid.json", fileContent: `[
			0,
			-0,
			0e-0,
			0e0,
			0e+0,
			-0e-0,
			-0e0,
			-0e+0,
			0e-03,
			0e03,
			0e+03,
			-0e-03,
			-0e03,
			-0e+03,
			0e-7,
			0e7,
			0e+7,
			-0e-7,
			-0e7,
			-0e+7
			]`, result: "This is a valid JSON",
			err: nil},
		{fileName: "invalid.json", fileContent: `{
			"key1": True,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 101
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key1": true,
			"key2": false,
			"key3": nulll,
			"key4": "value",
			"key5": 101
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key5": 0101
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key5": 01.01
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key5": 023456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key5": 02.3456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key5": -0101
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key5": -01.01
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key5": -023456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key5": -02.3456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 101true
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key1": true,
			"key2": false,
			"key3": nulll,
			"key4": "value",
			"key5": 101true
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "valid.json", fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": -101
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileName: "valid.json", fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": -1.1
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileName: "valid.json", fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": -1
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": ["list value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": "list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": 
			  "inner key": "inner value"
			},
			"key-l": ["list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			,
			"key-l": ["list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "valid.json", fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {},
			"key-l": []
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileName: "valid.json", fileContent: `[
			"value",
			 101,
			 {},
			 []
		  ]`, result: "This is a valid JSON",
			err: nil},
		{fileName: "valid.json", fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value",
			  "inner key2": [1, true, "hi"]
			},
			"key-l": ["list value", {"key3": "value 3"}]
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
				"inner key: "inner value",
				"key-l": [
					"list value"
				]
			}
		}`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"array": [
			  n, 1, 2, 3, 4, 5, 6, 7
			]
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `[
			[
			 n, 1, 2, 3
		   ]
		 ]`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `[
			[
			 "value1, "value2", [1, 2]
		   ]
		 ]`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileName: "invalid.json", fileContent: `{
			"key": [
			 "value1, "value2", [1, 2]
			 ]
		   }`, result: "",
			err: errors.New("This is an invalid JSON")},
	}

	for _, test := range tests {
		result, err := app.App(func(name string) ([]byte, error) {
			if name != test.fileName {
				panic("error")
			}
			return []byte(test.fileContent), nil
		}, []string{"ccjsonparser", test.fileName})
		if !assert.Equal(t, test.err, err) || !assert.Equal(t, test.result, result) {
			t.Log(test.fileContent)
		}
	}

	result, err := app.App(func(name string) ([]byte, error) {
		if name != "doesntexist.json" {
			panic("error")
		}
		return nil, errors.New("no such file")
	}, []string{"ccjsonparser", "doesntexist.json"})
	assert.Contains(t, err.Error(), "no such file")
	assert.Equal(t, "", result)
}
