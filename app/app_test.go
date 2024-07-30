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
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading +")},
		{fileName: "invalid.json", fileContent: "+083", result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading +")},
		{fileName: "invalid.json", fileContent: "Null", result: "",
			err: errors.New("This is an invalid JSON\nThere is a wrongly written \"null\"")},
		{fileName: "invalid.json", fileContent: "False", result: "",
			err: errors.New("This is an invalid JSON\nThere is a wrongly written \"false\"")},
		{fileName: "invalid.json", fileContent: "True", result: "",
			err: errors.New("This is an invalid JSON\nThere is a wrongly written \"true\"")},
		{fileName: "invalid.json", fileContent: "078", result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileName: "invalid.json", fileContent: `"string1", "string 2"`, result: "",
			err: errors.New("This is an invalid JSON\nMultiple values outside of an object or array")},
		{fileName: "invalid.json", fileContent: `"string1" "string 2"`, result: "",
			err: errors.New("This is an invalid JSON\nMultiple values outside of an object or array")},
		{fileName: "invalid.json", fileContent: "\"string1\",\n \"string 2\"", result: "",
			err: errors.New("This is an invalid JSON\nMultiple values outside of an object or array")},
		{fileName: "invalid.json", fileContent: `"str\074b"`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
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
			err: errors.New("This is an invalid JSON\nThere is an unclosed array")},
		{fileName: "invalid.json", fileContent: `["key", "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that is closed as an object")},
		{fileName: "invalid.json", fileContent: `{"key": "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that is closed as an array")},
		{fileName: "invalid.json", fileContent: `["key", "value",]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that contains an extra tail comma(s)")},
		{fileName: "invalid.json", fileContent: `["key", "value",,]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that contains an extra tail comma(s)")},
		{fileName: "invalid.json", fileContent: `[,"key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that contains an extra advancing comma(s)")},
		{fileName: "invalid.json", fileContent: `[,,"key", "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that contains an extra advancing comma(s)")},
		{fileName: "invalid.json", fileContent: `[,"key", "value",]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that contains an extra advancing comma(s)")},
		{fileName: "invalid.json", fileContent: `{,"key": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that contains an extra advancing comma(s)")},
		{fileName: "invalid.json", fileContent: `{,,"key": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that contains an extra advancing comma(s)")},
		{fileName: "invalid.json", fileContent: `{,"key": "value",}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that contains an extra advancing comma(s)")},
		{fileName: "invalid.json", fileContent: `{"key" "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a missing (:)")},
		{fileName: "invalid.json", fileContent: `{"key" "value", "key2" "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a missing (:)")},
		{fileName: "invalid.json", fileContent: `{"key": "value", "key2" "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a missing (:)")},
		{fileName: "invalid.json", fileContent: `{"key" "value", "key2": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a missing (:)")},
		{fileName: "invalid.json", fileContent: `{"key": "value",, "key2": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an extra comma(s) between pairs of key:value")},
		{fileName: "invalid.json", fileContent: `{"key": "value", "key2" "value", "key3": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a missing (:)")},
		{fileName: "invalid.json", fileContent: `{"key": :"value", "key2": "value", "key3": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileName: "invalid.json", fileContent: `{"key": :"value", "key2": "value", "key3": :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileName: "invalid.json", fileContent: `{"key": "value", "key2": :"value", "key3": "value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileName: "invalid.json", fileContent: `{"key": "value", "key2": "value", "key3": :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileName: "invalid.json", fileContent: `{:"key": "value", "key2": "value", "key3" :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileName: "invalid.json", fileContent: `{:"key": "value", "key2": "value", "key3" :"value":}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileName: "invalid.json", fileContent: `{"key": "value", "key2": "value", "key3" :"value":}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileName: "invalid.json", fileContent: `{"key": "value", "key2": "value",: "key3" :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileName: "invalid.json", fileContent: `{"key": "value", :"key2": "value", "key3" :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileName: "invalid.json", fileContent: `{"key": "value":, "key2": "value",: "key3" :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileName: "invalid.json", fileContent: `{"key": "value":, "key2": "value", "key3" :"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has an invalid (:)s")},
		{fileName: "invalid.json", fileContent: `["key",, "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that has an extra comma(s) between some values")},
		{fileName: "invalid.json", fileContent: `["key" "value"]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that has a missing comma between two values")},
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
		{fileName: "invalid.json", fileContent: `{"key": "value", ,,}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that contains an extra tail comma(s)")},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			key2: "value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `{"key":value","key":"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `{"key":value","key":
		"value"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `{"key":"	tab	character	in	string	"}`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
		{fileName: "invalid.json", fileContent: `["tab\   character\   in\  string\  "]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
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
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
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
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
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
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
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
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key 5 ": " value 5",
		  `, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that is closed with a comma(s)")},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key 5 ": " value 5",
		  ,`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that is closed with a comma(s)")},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key 5 ": " value 5"
		  `, result: "",
			err: errors.New("This is an invalid JSON\nThere is an unclosed object")},
		{fileName: "invalid.json", fileContent: `[
			"key", "value",
			"key 5 ", " value 5",
		  `, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that is closed with a comma(s)")},
		{fileName: "invalid.json", fileContent: `[
			"key", "value",
			"key 5 ", " value 5",
		  ,`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that is closed with a comma(s)")},
		{fileName: "invalid.json", fileContent: `{
			"key": "Illegal backslash escape: \x15"
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
		{fileName: "invalid.json", fileContent: `{
			"Illegal backslash escape: \x15": "value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that contains tabs or new lines or Illegal backslash escapes")},
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
			err: errors.New("This is an invalid JSON\nThere is a wrongly written \"true\"")},
		{fileName: "invalid.json", fileContent: `{
			"key1": true,
			"key2": false,
			"key3": nulll,
			"key4": "value",
			"key5": 101
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `{
			"key5": +0101
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading +")},
		{fileName: "invalid.json", fileContent: `{
			"key5": +0101,
			"key4": "value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading +")},
		{fileName: "invalid.json", fileContent: `{
			"key5": 0101
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileName: "invalid.json", fileContent: `{
			"key5": 01.01
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileName: "invalid.json", fileContent: `{
			"key5": 023456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileName: "invalid.json", fileContent: `{
			"key5": 023456789012E66,
			"key6": "value"
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileName: "invalid.json", fileContent: `{
			"key5": 02.3456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileName: "invalid.json", fileContent: `{
			"key5": -0101
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileName: "invalid.json", fileContent: `{
			"key5": -01.01
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileName: "invalid.json", fileContent: `{
			"key5": -023456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileName: "invalid.json", fileContent: `{
			"key5": -02.3456789012E66
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, there is a leading zero")},
		{fileName: "invalid.json", fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": 101true
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `{
			"key1": true,
			"key2": false,
			"key3": nulll,
			"key4": "value",
			"key5": 101true
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `{
			"key1": true,
			"key2": false,
			"key3": null,
			"key4": "value",
			"key5": true,
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that contains an extra tail comma(s)")},
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
			  "inner key": 0x9
			},
			"key-l": ["list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, hexadecimal numbers are not allowed")},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": [0xc]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, hexadecimal numbers are not allowed")},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": 0X9
			},
			"key-l": ["list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, hexadecimal numbers are not allowed")},
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key-n": 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": [0Xc]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an invalid number, hexadecimal numbers are not allowed")},
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
		{fileName: "invalid.json", fileContent: `{
			"key": "value",
			"key-n", 101,
			"key-o": {
			  "inner key": "inner value"
			},
			"key-l": ["list value"]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an object that has a comma instead of a colon")},
		{fileName: "invalid.json", fileContent: `[[[[[[[[[[[[[[77]]]]]]]]]`, result: "",
			err: errors.New("This is an invalid JSON\nThere are ([{)s that are more than (]})s")},
		{fileName: "invalid.json", fileContent: `[[[[[[[[["hi"]]]]]]]]]]]]`, result: "",
			err: errors.New("This is an invalid JSON\nThere are ([{)s that are fewer than (]})s")},
		{fileName: "invalid.json", fileContent: `[
			"key", "value",
			"key-n": 101,
			"key-o", {},
			"key-l", []
		  ]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is an array that has a (:) instead of a (,)")},
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
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `{
			"array": [
			  n, 1, 2, 3, 4, 5, 6, 7
			]
		  }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `[
			[
			 n, 1, 2, 3
		   ]
		 ]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `[
			[
			 "value1, "value2", [1, 2]
		   ]
		 ]`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `{
			"key": [
			 "value1, "value2", [1, 2]
			 ]
		   }`, result: "",
			err: errors.New("This is an invalid JSON\nThere is a string that is not surrounded correctly by (\"\")")},
		{fileName: "invalid.json", fileContent: `{
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

	result, err = app.App(func(name string) ([]byte, error) {
		return nil, nil
	}, []string{"ccjsonparser"})
	assert.Equal(t, err.Error(), "There is a missing file name")
	assert.Equal(t, "", result)
}
