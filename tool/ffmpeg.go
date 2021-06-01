package tool

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 14:47
 * @Desc:ffmpeg 需要系统安装ffmpeg
 */
func FFGetDuration(url string) int {
	orderArguments := []string{"-i", url }
	order := exec.Command("ffmpeg", orderArguments...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	order.Stdout = &out
	order.Stderr = &stderr
	if err := order.Run(); err != nil {
		return d(stderr.String())
	}else {
		return d(out.String())
	}
}

func d(data string) int {
	if data==""{
		return 0
	}
	ix1 := strings.Index(data,"Duration:")
	ix2 := strings.Index(data,"start:")
	if ix1==-1||ix2==-1{
		return 0
	}
	d := data[ix1+9:ix2]
	d1 := strings.ReplaceAll(d," ","")
	d2 := strings.ReplaceAll(d1,",","")[:8]
	dur := 0
	a1 := strings.Split(d2,":")
	for k,v := range a1{
		if a,err := strconv.Atoi(v);err==nil {
			switch k {
			case 0:
				dur += a * 3600
			case 1:
				dur += a * 60
			case 2:
				dur += a
			}
		}
	}
	return dur
}