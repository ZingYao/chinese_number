# chinese_number
<big><u>**中文文档**</u></big> | <big><u>[English Document](https://github.com/ZingYao/chinese_number/blob/master/README_EN.md)</u></big>

### 中文数字和int64 类型互转的工具

#### 获取方法
项目根目录下执行命令，获取依赖
```shell
go get github.com/ZingYao/chinese_number
```

#### 使用方法

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