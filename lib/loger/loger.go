package loger

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"runtime"
	"time"
)

var IsLogFile = false
var LogfilePath = "mange.log"

func OpenLogFile() {
	IsLogFile = true
}

func SetLogFile(logfilePath string) {
	IsLogFile = true
	LogfilePath = logfilePath
}

//\033[1;31;1m \033[0m  红色字
//\033[1;30;44m  \033[0m  蓝色底
//\033[1;30;42m  \033[0m  黄色底
//\033[1;37;41m  \033[0m  红色底
//\033[1;34;1m  \033[0m   蓝色字
//\033[1;35;1m  \033[0m   紫色字
//\033[1;32;1m  \033[0m   黄色字
//\033[1;33;1m  \033[0m   橙色字
//\033[1;36;1m  \033[0m   绿色字
//\033[1;37;1m  \033[0m   白色字
//\033[1;38;1m  \033[0m   绿色字 亮
//\033[1;39;1m  \033[0m   绿色字
func Debug(v ...interface{}) {
	t := time.Now()
	pc, file, line, _ := runtime.Caller(1)
	fun := runtime.FuncForPC(pc)
	funName := fun.Name()

	s := fmt.Sprintf("%s | [Unix:%d] | \033[1;34;1m[DEBUG]\033[0m | \033[1;37;1m[F=%s:%d] | [N=%s] \033[0m| ", t.Format("2006-01-02 03:04:05"),
		t.Unix(), file, line, funName)
	for _, i := range v {

		s += fmt.Sprintf("<%s>:%+v", reflect.TypeOf(i), i)
	}
	fmt.Println(s)
}

func Error(v ...interface{}) {

	if IsLogFile {
		f, err := os.OpenFile(LogfilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		mw := io.MultiWriter(os.Stdout, f)
		log.SetOutput(mw)
	}

	t := time.Now()
	pc, file, line, _ := runtime.Caller(1)
	fun := runtime.FuncForPC(pc)
	funName := fun.Name()

	s := fmt.Sprintf(" | [Unix:%d] | \033[1;37;41m[ERROR]\033[0m | \033[1;37;1m[F=%s:%d] | [N=%s] \033[0m| \033[1;31;1m",
		t.Unix(), file, line, funName)
	for _, i := range v {

		s += fmt.Sprintf("<%s>:%+v", reflect.TypeOf(i), i)
	}
	s += "\033[0m"
	log.Println(s)
}
