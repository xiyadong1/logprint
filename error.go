// 名称最好是小写字母  最好不要带下划线   见名知意
package logprint

import (
	"fmt"
	"time"
)

func Error(err error) {
	t := time.Now()
	fmt.Printf("[error] %s : %s \n", t.Format("2006-01-05 15:04:05.000"), err.Error())
}
