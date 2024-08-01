package app_test

import (
	"errors"
	"testing"

	"github.com/ahmedsameha1/ccjsonparser/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	var tests = []struct {
		fileContent string
		result      string
		err         error
	}{
		{fileContent: "{}", result: "This is a valid JSON", err: nil},
		{fileContent: `""`, result: "This is a valid JSON", err: nil},
		{fileContent: `null`, result: "This is a valid JSON", err: nil},
		{fileContent: `false`, result: "This is a valid JSON", err: nil},
		{fileContent: `true`, result: "This is a valid JSON", err: nil},
		{fileContent: `-74.23`, result: "This is a valid JSON", err: nil},
		{fileContent: `0`, result: "This is a valid JSON", err: nil},
		{fileContent: "", result: "",
			err: errors.New("This is an invalid JSON\nMUST be an object, array, number, or string, or false or null or true")},
		{fileContent: "+83", result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading +")},
		{fileContent: "+083", result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading +")},
		{fileContent: "Null", result: "",
			err: errors.New("This is an invalid JSON\nThere is a wrongly written \"null\"")},
		{fileContent: "False", result: "",
			err: errors.New("This is an invalid JSON\nThere is a wrongly written \"false\"")},
		{fileContent: "True", result: "",
			err: errors.New("This is an invalid JSON\nThere is a wrongly written \"true\"")},
		{fileContent: "078", result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileContent: `"string1", "string 2"`, result: "",
			err: errors.New("This is an invalid JSON\nMultiple values outside of an object or array")},
		{fileContent: `"string1" "string 2"`, result: "",
			err: errors.New("This is an invalid JSON\nMultiple values outside of an object or array")},
		{fileContent: "\"string1\",\n \"string 2\"", result: "",
			err: errors.New("This is an invalid JSON\nMultiple values outside of an object or array")},
		{fileContent: `"str\074b"`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
		{fileContent: `{"key": "value"}`, result: "This is a valid JSON",
			err: nil},
		{fileContent: `{
			"key": "value",
			"key2": "value"
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileContent: `{
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
		{fileContent: `["key", "value"`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an unclosed array")},
		{fileContent: `["key", "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that is closed as an object")},
		{fileContent: `{"key": "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that is closed as an array")},
		{fileContent: `["key", "value",]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that contains an extra tail comma(s)")},
		{fileContent: `["key", "value",,]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that contains an extra tail comma(s)")},
		{fileContent: `[,"key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that contains an extra advancing comma(s)")},
		{fileContent: `[,,"key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that contains an extra advancing comma(s)")},
		{fileContent: `[,"key", "value",]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that contains an extra advancing comma(s)")},
		{fileContent: `{,"key": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that contains an extra advancing comma(s)")},
		{fileContent: `{,,"key": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that contains an extra advancing comma(s)")},
		{fileContent: `{,"key": "value",}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that contains an extra advancing comma(s)")},
		{fileContent: `{"key" "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a missing (:)")},
		{fileContent: `{"key" "value", "key2" "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a missing (:)")},
		{fileContent: `{"key": "value", "key2" "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a missing (:)")},
		{fileContent: `{"key" "value", "key2": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a missing (:)")},
		{fileContent: `{"key": "value",, "key2": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an extra comma(s) between pairs of key:value")},
		{fileContent: `{"key": "value", "key2" "value", "key3": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a missing (:)")},
		{fileContent: `{"key": :"value", "key2": "value", "key3": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileContent: `{"key": :"value", "key2": "value", "key3": :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileContent: `{"key": "value", "key2": :"value", "key3": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileContent: `{"key": "value", "key2": "value", "key3": :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileContent: `{:"key": "value", "key2": "value", "key3" :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileContent: `{:"key": "value", "key2": "value", "key3" :"value":}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileContent: `{"key": "value", "key2": "value", "key3" :"value":}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileContent: `{"key": "value", "key2": "value",: "key3" :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileContent: `{"key": "value", :"key2": "value", "key3" :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileContent: `{"key": "value":, "key2": "value",: "key3" :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileContent: `{"key": "value":, "key2": "value", "key3" :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileContent: `["key",, "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that has an extra comma(s) between some values")},
		{fileContent: `["key" "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that has a missing comma between two values")},
		{fileContent: `["key", "value"]]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `[["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `]["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `{["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `{ ["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `{ { ["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `}["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `}["key", "value"] ]]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `["key", "value"][`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `["key", "value"]{`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `["key", "value"]}`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `["key", "value"]} ]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `[["key", "value"]{`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `{["key", "value"]]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `]["key", "value"][`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `}["key", "value"]{`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `[["key", "value"]}`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `{["key", "value"]]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid \"][}{\"")},
		{fileContent: `,["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileContent: `["key", "value"],`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileContent: `, ["key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileContent: `["key", "value"] ,`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileContent: `, ["key", "value"] ,`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileContent: `,, ["key", "value"] ,,`, result: "",
			err: errors.New("This is an invalid JSON\nThis is an array that is surrounded by invalid commas")},
		{fileContent: `{"key": "value", ,,}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that contains an extra tail comma(s)")},
		{fileContent: `{
			"key": "value",
			key2: "value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `{"key":value","key":"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `{"key":value","key":
		"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `{"key":"	tab	character	in	string	"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
		{fileContent: `["tab\   character\   in\  string\  "]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
		{fileContent: `{
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
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
		{fileContent: `{
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
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
		{fileContent: `{
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
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `{
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
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `{
			"key": "value",
			"key 5 ": " value 5",
		  `, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that is closed with a comma(s)")},
		{fileContent: `{
			"key": "value",
			"key 5 ": " value 5",
		  ,`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that is closed with a comma(s)")},
		{fileContent: `{
			"key": "value",
			"key 5 ": " value 5"
		  `, result: "",
			err: errors.New("This is an invalid JSON\nThere is an unclosed object")},
		{fileContent: `[
			"key", "value",
			"key 5 ", " value 5",
		  `, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that is closed with a comma(s)")},
		{fileContent: `[
			"key", "value",
			"key 5 ", " value 5",
		  ,`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that is closed with a comma(s)")},
		{fileContent: `{
			"key": "Illegal backslash escape: \x15"
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
		{fileContent: `{
			"Illegal backslash escape: \x15": "value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
		{fileContent: `{
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
		{fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 642
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileContent: `[
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
		{fileContent: `{
			"key1": True,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 101
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a wrongly written \"true\"")},
		{fileContent: `{
			"key1": true,
			"key2": false,
			"key3": nulll,
			"key4": "value",
			"key5": 101
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `{
			"key5": +0101
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading +")},
		{fileContent: `{
			"key5": +0101,
			"key4": "value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading +")},
		{fileContent: `{
			"key5": 0101
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileContent: `{
			"key5": 01.01
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileContent: `{
			"key5": 023456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileContent: `{
			"key5": 023456789012E66,
			"key6": "value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileContent: `{
			"key5": 02.3456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileContent: `{
			"key5": -0101
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileContent: `{
			"key5": -01.01
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileContent: `{
			"key5": -023456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileContent: `{
			"key5": -02.3456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 101true
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `{
			"key1": true,
			"key2": false,
			"key3": nulll,
			"key4": "value",
			"key5": 101true
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": true,
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that contains an extra tail comma(s)")},
		{fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": -101
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": -1.1
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": -1
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": ["list value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": 0x9
			},
			"key-l": ["list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, hexadecimal numbers are not allowed")},
		{fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": [0xc]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, hexadecimal numbers are not allowed")},
		{fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": 0X9
			},
			"key-l": ["list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, hexadecimal numbers are not allowed")},
		{fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": [0Xc]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, hexadecimal numbers are not allowed")},
		{fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": "list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": 
			  "inner key": "inner value"
			},
			"key-l": ["list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			,
			"key-l": ["list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
		{fileContent: `{
			"key": "value",
			"key-n", 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": ["list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a comma instead of a colon")},
		{fileContent: `[[[[[[[[[[[[[[77]]]]]]]]]`, result: "",
			err: errors.New("This is an invalid JSON\nThere are ([{)s that are more than (]})s")},
		{fileContent: `[[[[[[[[["hi"]]]]]]]]]]]]`, result: "",
			err: errors.New("This is an invalid JSON\nThere are ([{)s that are fewer than (]})s")},
		{fileContent: `[
			"key", "value",
			"key-n": 101,
			"key-o", {},
			"key-l", []
		  ]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that has a (:) instead of a (,)")},
		{fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {},
			"key-l": []
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileContent: `[
			"value",
			 101,
			 {},
			 []
		  ]`, result: "This is a valid JSON",
			err: nil},
		{fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value",
			  "inner key2": [1, true, "hi"]
			},
			"key-l": ["list value", {"key3": "value 3"}]
		  }`, result: "This is a valid JSON",
			err: nil},
		{fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
				"inner key: "inner value",
				"key-l": [
					"list value"
				]
			}
		}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `{
			"array": [
			  n, 1, 2, 3, 4, 5, 6, 7
			]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `[
			[
			 n, 1, 2, 3
		   ]
		 ]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `[
			[
			 "value1, "value2", [1, 2]
		   ]
		 ]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `{
			"key": [
			 "value1, "value2", [1, 2]
			 ]
		   }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileContent: `{
			"batter":
			  [
				{ "id": "1001", "type": "Regular" },
				{ "id": "1002", "type": "Chocolate" },
				{ "id": "1003", "type": "Blueberry" },
				{ "id": "1004" "type": "Devil's Food" }
			  ]
		  }`, result: "",
			err: errors.New("This is an invalid JSON")},
	}

	for _, test := range tests {
		err := app.Validate(test.fileContent)
		if !assert.Equal(t, test.err, err) {
			t.Log(test.fileContent)
		}
	}
}
