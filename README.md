[![Coverage Status](https://coveralls.io/repos/github/Hexilee/unhtml/badge.svg)](https://coveralls.io/github/Hexilee/unhtml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Hexilee/unhtml)](https://goreportcard.com/report/github.com/Hexilee/unhtml)
[![Build Status](https://travis-ci.org/Hexilee/unhtml.svg?branch=master)](https://travis-ci.org/Hexilee/unhtml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Hexilee/unhtml/blob/master/LICENSE)
[![Documentation](https://godoc.org/github.com/Hexilee/unhtml?status.svg)](https://godoc.org/github.com/Hexilee/unhtml)

### Example

A HTML file

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <div id="test">
        <ul>
            <li>0</li>
            <li>1</li>
            <li>2</li>
            <li>3</li>
        </ul>
        <div>
            <p>Hexilee</p>
            <p>20</p>
            <p>true</p>
        </div>
        <p>Hello World!</p>
        <p>10</p>
        <p>3.14</p>
        <p>true</p>
    </div>
</body>
</html>
```

Unmarshal it to a struct

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/Hexilee/unhtml"
	"io/ioutil"
)

type (
	AllTypeTest struct {
		Slice   []int    `html:"ul > li"`
		Struct  TestUser `html:"div"`
		String  string   `html:"#test > p:nth-child(3)"`
		Int     int      `html:"#test > p:nth-child(4)"`
		Int8    int8     `html:"#test > p:nth-child(4)"`
		Int16   int16    `html:"#test > p:nth-child(4)"`
		Int32   int32    `html:"#test > p:nth-child(4)"`
		Int64   int64    `html:"#test > p:nth-child(4)"`
		Uint    uint     `html:"#test > p:nth-child(4)"`
		Uint8   uint8    `html:"#test > p:nth-child(4)"`
		Uint16  uint16   `html:"#test > p:nth-child(4)"`
		Uint32  uint32   `html:"#test > p:nth-child(4)"`
		Uint64  uint64   `html:"#test > p:nth-child(4)"`
		Float32 float32  `html:"#test > p:nth-child(5)"`
		Float64 float64  `html:"#test > p:nth-child(5)"`
		Bool    bool     `html:"#test > p:nth-child(6)"`
	}

	TestUser struct {
		Name      string `html:"p:nth-child(1)"`
		Age       uint   `html:"p:nth-child(2)"`
		LikeLemon bool   `html:"p:nth-child(3)"`
	}
)

func (AllTypeTest) Root() string {
	return "#test"
}

func main() {
	AllTypeHTML, _ := ioutil.ReadFile("testHTML/all-type.html")
	allTypes := AllTypeTest{}
	err := unhtml.Unmarshal(AllTypeHTML, &allTypes)
	if err == nil {
		result, err := json.Marshal(&allTypes)
		if err == nil {
			fmt.Println(result)
		}
	}
}
```

得到

```json
{
  "Slice": [
    0,
    1,
    2,
    3
  ],
  "Struct": {
    "Name": "Hexilee",
    "Age": 20,
    "LikeLemon": true
  },
  "String": "Hello World!",
  "Int": 10,
  "Int8": 10,
  "Int16": 10,
  "Int32": 10,
  "Int64": 10,
  "Uint": 10,
  "Uint8": 10,
  "Uint16": 10,
  "Uint32": 10,
  "Uint64": 10,
  "Float32": 3.14,
  "Float64": 3.14,
  "Bool": true
}
```