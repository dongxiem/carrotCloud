package mq

import (
	cmn "carrotCloud/common"
)

// TransferData : 将要写到rabbitmq的数据的结构体
type TransferData struct {
	FileHash      string        // 被转移文件的Hash值
	CurLocation   string        // 临时存储的具体地址
	DestLocation  string        // 要转移的目标地址
	DestStoreType cmn.StoreType // 文件将要被转移到的存储类型
}
