// 编写一个名字为 MySqrt 的函数，计算一个 float64 类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。
package practise

import (
	"fmt"
	"math"
)

func MySqrt(num float64) (ret float64, err error) {
	if num < 0 {
		err = fmt.Errorf("负数无法求平方根")
	}
	ret = math.Sqrt(num)
	return
}
