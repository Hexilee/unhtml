package unhtml

import (
	"encoding/json"
	"fmt"
)

const (
	AllTypesHTML = `
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
`
)

func ExampleUnmarshal() {
	allTypes := AllTypeTest{}
	_ = Unmarshal(AllTypeHTML, &allTypes)
	result, _ := json.Marshal(&allTypes)
	fmt.Println(string(result))
	// Output:
	// {"Slice":[0,1,2,3],"Struct":{"Name":"Hexilee","Age":20,"LikeLemon":true},"String":"Hello World!","Int":10,"Int8":10,"Int16":10,"Int32":10,"Int64":10,"Uint":10,"Uint8":10,"Uint16":10,"Uint32":10,"Uint64":10,"Float32":3.14,"Float64":3.14,"Bool":true}
}
