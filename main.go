package main

import (
	"fmt"
	"tools/tool"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 14:26
 * @Desc:
 */

func main()  {
	fmt.Println(tool.GenRsaKey(1024))
}