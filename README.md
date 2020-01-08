# huaweicloud-sdk

华为云简单go sdk。

### 安装

```
go get -u github.com/yangyimincn/huaweicloud-sdk
```

### 使用

```
package main

import (
	"fmt"
	vpc "github.com/yangyimincn/huaweicloud-sdk/services/vpc/v1"
)

func main() {
	hwVPC := vpc.NewClient(
		"ak",
		"sk",
		"cn-east-2",
	)
	res, err := hwVPC.DescribePublicIPList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
```
