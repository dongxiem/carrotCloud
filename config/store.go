package config

import (
	cmn "carrotCloud/common"
)

const (
	// TempLocalRootDir : 本地临时存储地址的路径
	TempLocalRootDir = "D:/Garmen/tmp/fileserver/"
	// TempPartRootDir : 分块文件在本地临时存储地址的路径
	TempPartRootDir = "D:/Garmen/tmp/fileserver_part/"
	// CephRootDir : Ceph的存储路径prefix
	CephRootDir = "/ceph"
	// OSSRootDir : OSS的存储路径prefix
	OSSRootDir = "oss/"
	// CurrentStoreType : 设置当前文件的存储类型
	CurrentStoreType = cmn.StoreLocal
)
