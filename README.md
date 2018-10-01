[![Coverage Status](https://coveralls.io/repos/github/Hexilee/unhtml/badge.svg)](https://coveralls.io/github/Hexilee/unhtml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Hexilee/unhtml)](https://goreportcard.com/report/github.com/Hexilee/unhtml)
[![Build Status](https://travis-ci.org/Hexilee/unhtml.svg?branch=master)](https://travis-ci.org/Hexilee/unhtml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Hexilee/unhtml/blob/master/LICENSE)
[![Documentation](https://godoc.org/github.com/Hexilee/unhtml?status.svg)](https://godoc.org/github.com/Hexilee/unhtml)

Table of Contents
=================

* [Example &amp; Performance](#example--performance)
* [Tips &amp; Features](#tips--features)
  * [Types](#types)
  * [Root](#root)
  * [Selector](#selector)
     * [Struct](#struct)
     * [Slice](#slice)
  * [Tags](#tags)
     * [html](#html)
     * [attr](#attr)
     * [converter](#converter)


### Example & Performance

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

Read it

```go
AllTypeHTML, _ := ioutil.ReadFile("testHTML/all-type.html")
```

If we want to parse it and get the values we want, like follow structs, how should we do?


```go
type (
	AllTypeTest struct {
		Slice   []int    
		Struct  TestUser 
		String  string   
		Int     int      
		Int8    int8     
		Int16   int16    
		Int32   int32    
		Int64   int64    
		Uint    uint     
		Uint8   uint8    
		Uint16  uint16   
		Uint32  uint32   
		Uint64  uint64   

		Float32 float32  
		Float64 float64  
		Bool    bool     
	}
	TestUser struct {
		Name      string 
		Age       uint   
		LikeLemon bool   
	}
)
```

In traditional way, we should do like this

```go
package example

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

func parseAllTypesLogically() (AllTypeTest, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(AllTypeHTML))
	allTypes := AllTypeTest{}
	if err == nil {
		selection := doc.Find(allTypes.Root())
		allTypes.Slice = make([]int, 0)
		selection.Find(`ul > li`).Each(func(i int, selection *goquery.Selection) {
			Int, parseErr := strconv.Atoi(selection.Text())
			if parseErr != nil {
				err = parseErr
			}
			allTypes.Slice = append(allTypes.Slice, Int)
		})
		if err == nil {
			allTypes.Struct.Name = selection.Find(`#test > div > p:nth-child(1)`).Text()
			Int, parseErr := strconv.Atoi(selection.Find(`#test > div > p:nth-child(2)`).Text())
			if err = parseErr; err == nil {
				allTypes.Struct.Age = uint(Int)
				Bool, parseErr := strconv.ParseBool(selection.Find(`#test > div > p:nth-child(3)`).Text())
				if err = parseErr; err == nil {
					allTypes.Struct.LikeLemon = Bool

					String := selection.Find(`#test > p:nth-child(3)"`).Text()
					IntStr := selection.Find(`#test > p:nth-child(4)`).Text()
					Int, parseErr := strconv.Atoi(IntStr)
					if err = parseErr; err != nil {
						return allTypes, err
					}

					Uint64, parseErr := strconv.ParseUint(IntStr, 0, 0)
					if err = parseErr; err != nil {
						return allTypes, err
					}

					FloatStr := selection.Find(`#test > p:nth-child(5)`).Text()
					Float64, parseErr := strconv.ParseFloat(FloatStr, 0)
					if err = parseErr; err != nil {
						return allTypes, err
					}

					Bool, parseErr := strconv.ParseBool(selection.Find(`#test > p:nth-child(6)`).Text())
					if err = parseErr; err != nil {
						return allTypes, err
					}
					allTypes.String = String
					allTypes.Int = Int
					allTypes.Int8 = int8(Int)
					allTypes.Int16 = int16(Int)
					allTypes.Int32 = int32(Int)
					allTypes.Int64 = int64(Int)
					allTypes.Uint = uint(Uint64)
					allTypes.Uint8 = uint8(Uint64)
					allTypes.Uint16 = uint16(Uint64)
					allTypes.Uint32 = uint32(Uint64)
					allTypes.Uint64 = uint64(Uint64)
					allTypes.Float32 = float32(Float64)
					allTypes.Float64 = Float64
					allTypes.Bool = Bool
				}
			}
		}
	}

	return allTypes, err
}

```

It works pretty good, but is boring. And now, you can do like this:

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
		Struct  TestUser `html:"#test > div"`
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
	allTypes := AllTypeTest{}
	_ := unhtml.Unmarshal(AllTypeHTML, &allTypes)
	result, _ := json.Marshal(&allTypes)
	fmt.Println(string(result))
}
```

Result: 

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

I think it can improve much efficiency of my development, however, what about its performance?

There are two benchmarks

```go
func BenchmarkUnmarshalAllTypes(b *testing.B) {
	assert.NotNil(b, AllTypeHTML)
	for i := 0; i < b.N; i++ {
		allTypes := AllTypeTest{}
		assert.Nil(b, Unmarshal(AllTypeHTML, &allTypes))
	}
}

func BenchmarkParseAllTypesLogically(b *testing.B) {
	assert.NotNil(b, AllTypeHTML)
	for i := 0; i < b.N; i++ {
		_, err := parseAllTypesLogically()
		assert.Nil(b, err)
	}
}
```

Test it:

```bash
> go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/Hexilee/unhtml
BenchmarkUnmarshalAllTypes-4        	   20000	     83271 ns/op
BenchmarkParseAllTypesLogically-4   	   30000	     45934 ns/op
PASS
ok  	github.com/Hexilee/unhtml	4.621s
```

Not very bad, in consideration of the small size of the demo HTML. In true development with more complicated HTML, their efficiency are almost the same.

### Tips & Features

The only API this package exposed is the function, 

```go
func Unmarshal(data []byte, v interface{}) error
```

which is compatible with the standard libraries `json` and `xml`. However, you can do some jobs with the data types in your code.

#### Types

This package supports part kinds of type, the all kinds of type in the `reflect` package except `Ptr/Uintptr/Interface/Chan/Func`.

Follow fields are invalid and will cause `UnmarshalerItemKindError`.

```go
type WrongFieldsStruct struct {
    Ptr *int
    Uintptr uintptr
    Interface io.Reader
    Chan chan int
    Func func()
}
```

However, when you call the function `Unmarshal`, you **MUST** pass a pointer otherwise you will get an `UnmarshaledKindMustBePtrError`.

```go
a := 1

// Wrong
Unmarshal([]byte(""), a)

// Right
Unmarshal([]byte(""), &a)
```

#### Root

Return the root selector.

You are only supported to define a `Root() string` method for the root type, like

```go
func (AllTypeTest) Root() string {
	return "#test"
}
```

If you define it for a field type, such as `TestUser`

```go
func (TestUser) Root() string {
	return "#test"
}
```

In this case, in `AllTypeTest`, the field selector will be covered.

```go
type (
	AllTypeTest struct {
		...
		Struct  TestUser `html:"#test > div"`
		...
	}
)

// real
type (
	AllTypeTest struct {
		...
		Struct  TestUser `html:"#test"`
		...
	}
)
```



#### Selector

This package base on `github.com/PuerkitoBio/goquery` and supports standard css selector.

You can define selector of a field in tag, like this

```go
type (
	AllTypeTest struct {
	   ...
		Int     int      `html:"#test > p:nth-child(4)"`
		...
	}
)
```

In most cases, this package will find the `#test > p:nth-child(4)` element and try to parse its `innerText` as int.

However, when the field type is `Struct` or `Slice`, something will be more complex.

##### Struct

```go
type (
	AllTypeTest struct {
		...
		Struct  TestUser `html:"#test > div"`
		...
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
```

First, it will call `*goquery.Selection.Find("#test")`, we get:

```html
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
```

Then, it will call `*goquery.Selection.Find("#test > div")`, we get

```html
<div>
    <p>Hexilee</p>
    <p>20</p>
    <p>true</p>
</div>
```

Then, in `TestUser`, it will call

```go
*goquery.Selection.Find("p:nth-child(1)") // as Name
*goquery.Selection.Find("p:nth-child(2)") // as Age
*goquery.Selection.Find("p:nth-child(3)") // as LikeLemon
```

##### Slice

```go
type (
	AllTypeTest struct {
		Slice   []int    `html:"ul > li"`		...
	}
)

func (AllTypeTest) Root() string {
	return "#test"
}
```

As above, we get

```html
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
```

Then it will call `*goquery.Selection.Find("ul > li")`, we get

```html
  <li>0</li>
  <li>1</li>
  <li>2</li>
  <li>3</li>
```

Then, it will call `*goquery.Selection.Each(func(int, *goquery.Selection))`, iterate the list and parse values for slice.

#### Tags

This package supports three tags, `html`, `attr` and `converter`

##### html

Provide the `css selector` of this field.

##### attr

By default, this package regard the `innerText` of a element as its `value`

```html
<a href="https://google.com">Google</a>
```

```go
type Link struct {
    Text string `html:"a"`
}
```

You will get `Text = Google`. However, how should we do if we want to get `href`?

```go
type Link struct {
    Href string `html:"a" attr:"href"`
    Text string `html:"a"`
}
```

You will get `link.Href == "https://google.com"`

##### converter

Sometimes, you want to process the original data

```html
<p>2018-10-01 00:00:01</p>
```

You may unmarshal it like this

```go
type Birthday struct {
	Time time.Time `html:"p"`
}

func TestConverter(t *testing.T) {
	birthday := Birthday{}
	assert.Nil(t, Unmarshal([]byte(BirthdayHTML), &birthday))
	assert.Equal(t, 2018, birthday.Time.Year())
	assert.Equal(t, time.October, birthday.Time.Month())
	assert.Equal(t, 1, birthday.Time.Day())
}
```

Absolutely, you will fail, because you don't define the way converts string to time.Time. `unhtml` will regard it as a struct.

However, you can use `converter`

```go
type Birthday struct {
    Time time.Time `html:"p" converter:"StringToTime"`
}

const TimeStandard = `2006-01-02 15:04:05`

func (Birthday) StringToTime(str string) (time.Time, error) {
	return time.Parse(TimeStandard, str)
}

func TestConverter(t *testing.T) {
	birthday := Birthday{}
	assert.Nil(t, Unmarshal([]byte(BirthdayHTML), &birthday))
	assert.Equal(t, 2018, birthday.Time.Year())
	assert.Equal(t, time.October, birthday.Time.Month())
	assert.Equal(t, 1, birthday.Time.Day())
}
```

Make it.

The type of converter **MUST** be 

```go
func (inputType) (resultType, error)
```

`resultType` **MUST** be the same with the field type, and they can be any type.

`inputType` **MUST NOT** violate the requirements in [Types](#types).



