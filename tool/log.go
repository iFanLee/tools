package tool

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 16:00
 * @Desc:
 */

const (
	format = "2006/0102"
	ext = ".log"
)
var _pwd,_path string
func init()  {
	_pwd,_ = os.Getwd()
	_pwd = strings.ReplaceAll(_pwd,"\\","/")
	_pwd += "/"
	_path = fmt.Sprintf("%vlog/",_pwd)
}

func LogCustom(path string,args ...interface{})  {
	logData(fmt.Sprintf("%v/custom/%v/",_path,path),args)
}
func LogInfo(args ...interface{})  {
	logData(fmt.Sprintf("%v/info/",_path),args)
}
func LogWarn(args ...interface{})  {
	logData(fmt.Sprintf("%v/warn/",_path),args)
}
func LogError(args ...interface{})  {
	logData(fmt.Sprintf("%v/error/",_path),args)
}
func LogFatal(args ...interface{})  {
	logData(fmt.Sprintf("%v/fatal/",_path),args)

}

//获取文件和行号
func getLine() (string) {
	if _,file,line,ok := runtime.Caller(4);ok{
		if strings.Contains(file,"/runtime/"){
			return "[goroutine]"
		}
		return fmt.Sprintf("[%v][%v]",strings.ReplaceAll(strings.ReplaceAll(file,_pwd,""),".go",""),line)
	}
	return ""
}
func logData(path string,args []interface{})  {
	defer func() {
		if err := recover();err!=nil{
			fmt.Println(err)
		}
	}()

	logStr := ""
	lens := len(args)
	if lens==0{
		return
	}
	if lens==1&&reflect.TypeOf(args[0]).Name()=="string"{
		logStr = fmt.Sprintf("%v",args[0])
	}else {
		logData := make([]interface{},0)
		for _,v := range args{
			logData = append(logData,v)
		}
		jsonStr,_ := json.Marshal(logData)
		logStr = string(jsonStr)
	}
	write(path,logStr)
}
func write(path,content string)  {
	//文件名和行号
	line := getLine()
	if line==""{
		content = fmt.Sprintf("[%v] %v\r\n",time.Now().Format("15:04:05"),content)
	}else {
		content = fmt.Sprintf("[%v]%v %v\r\n",time.Now().Format("15:04:05"),line,content)
	}
	//日志目录为自定义目录+方法+日期
	path += fmt.Sprintf("/%v/",time.Now().Format(format))
	//文件夹不存在则创建
	if isE,_ := PathExists(path);!isE{
		if err := os.MkdirAll(path,os.ModePerm);err!=nil{
			return
		}
	}
	logName := time.Now().Format("15")+ext
	fileObj,err := os.OpenFile(path+logName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		return
	}
	defer fileObj.Close()
	io.WriteString(fileObj,content)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}