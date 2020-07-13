package meta

import "time"

const baseFormat = "2001-01-01 14:01:01"

type ByUpLoadTime []FileMeta

// Len : 长度
func (a ByUpLoadTime) Len() int {
	return len(a)
}

// Swap ：交换函数
func (a ByUpLoadTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByUpLoadTime) Less(i, j int) bool {
	// 通过注册时间并按照指定的时间格式进行分析
	iTime, _ := time.Parse(baseFormat, a[i].UploadAt)
	jTime, _ := time.Parse(baseFormat, a[j].UploadAt)
	return iTime.UnixNano() > jTime.UnixNano()
}
