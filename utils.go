package sensemicroservice

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//Checks an error and then logs and prints accordingly
func checkErr(err error) {
	if err != nil {
		pc, file, line, _ := runtime.Caller(1)
		function := strings.TrimPrefix(filepath.Ext(runtime.FuncForPC(pc).Name()), ".")
		fmt.Println("[" + time.Now().Format("Jan-02-06 3:04pm") + "] Error Warning:" + file + " " + function + "() line:" + strconv.Itoa(line) + " " + err.Error())
	}
}
