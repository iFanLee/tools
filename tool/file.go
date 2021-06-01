package tool

import (
	"bufio"
	"os"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 14:48
 * @Desc:
 */
//按行读取文件内容
func ReadFileByLine(filePath string) ([]string,error) {
	s := make([]string,0)
	file, err := os.Open(filePath)
	if err != nil {
		return s,err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s,line)
	}
	if err := scanner.Err(); err != nil {
		return s,err
	}
	return s,nil
}