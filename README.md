properties
==========
这是一个golang的lib，读取 properties 文件
```shell
#it is global elements
key = value
...

#it is section
[section]
key = value
...
```

#安装
```shell
go get -u github.com/xiaoyu830411/properties
```

#快速使用
```golang
package main

import (
	"github.com/xiaoyu830411/properties"
)

func main() {
  //get a value by key
  p, err := properties.Load("file path")
  if err != nil {
     ...
  }
  
  value, ok := p.Get("key")
  if ok {
    ...
  }
  
  //get section by section id
  section, ok := p.GetSection("section id")
  
  if ok {
    //get value by key from section
    value, ok := section.Get("key")
  }
}
```
