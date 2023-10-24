# chinese_number
<big><u>[中文文档](https://github.com/ZingYao/chinese_number/blob/master/README.md)</u></big> | <big><u>**English Document**</u></big>

### Tool for converting Chinese numbers to int64 and vice versa Installation

#### Installation
Execute the following command in the root directory of your project to get the dependency:
```shell
go get github.com/ZingYao/chinese_number
```

#### Usage

```go
package main

import (
	"fmt"
	"github.com/ZingYao/chinese_number"
)

func main() {
	strNum := "一百九十二亿三千一百一十六万八千九百五十二亿四千零八十六万五千二百一十三"
	num, err := chinese_number.Simplified2Number(strNum)
	if err != nil {
		fmt.Printf("str:%q to number has an error:%+v\n",strNum,err)
		return
	}
	fmt.Printf("str number:%q to number is:%d\n",strNum,num)
	// output is:str number:"一百九十二亿三千一百一十六万八千九百五十二亿四千零八十六万五千二百一十三" to number is:1923116895240865213
	
	//or
	num,err = chinese_number.Chinese2Number(chinese_number.Simplified,strNum)
	if err != nil {
		fmt.Printf("str:%q to number has an error:%+v\n",strNum,err)
		return
	}
	fmt.Printf("str number:%q to number is:%d\n",strNum,num)
	// output is:str number:"一百九十二亿三千一百一十六万八千九百五十二亿四千零八十六万五千二百一十三" to number is:1923116895240865213
}
```