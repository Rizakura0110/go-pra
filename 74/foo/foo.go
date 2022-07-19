package foo

const (
	Max = 100 //頭文字が大文字なので外部から参照可能
	min = 1   // 小文字なので参照不可能
)

func ReturnMin() int {
	return min
}
