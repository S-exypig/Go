/*编写一个程序，要求能够反转字符串，即将 “Google” 转换成 “elgooG”
（提示：使用 []byte 类型的切片）。
如果您使用两个切片来实现反转，请再尝试使用一个切片（提示：使用交换法）。
如果您想要反转 Unicode 编码的字符串，请使用 []int32 类型的切片。*/

package practise

func ReverseStr(str string) (res string) {
	sl := make([]byte, 0)
	size := len(str)
	for i := 0; i < size; i++ {
		sl = append(sl, str[size-i-1])
	}
	res = string(sl)
	return
}
