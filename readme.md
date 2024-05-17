redis的使用更简单
---

### 说明
基于redigo的封装，写入好配置文件后，在任何地方随时调用，使用起来更简单，代码更优雅
支持和redigo同时使用

### 安装
`go get -u -v github.com/chensanle/redis-pack`

### 使用
```go
package main

import (
	"fmt"

	"github.com/chensanle/redis-pack"
)

func init() {
    redis_pack.NewConnectionWithFile("", "")
}

func main() {
	err := redis_pack.RedigoConn.String.Set("1", 2, 10).Error()
	if err != nil {
		fmt.Println(err)
		return
	}
	v, err := redis_pack.RedigoConn.String.Get("1").Int64()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}
```


### License
The MIT License.

source https://github.com/newgoo/redigo-pack