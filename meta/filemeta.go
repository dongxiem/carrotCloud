package meta

import (
	mydb "carrotCloud/db"
	"sort"
)

// 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

// init ：初始文件元信息结构
func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta ：新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// UpdateFileMetaDB : 新增/更新文件元信息到mysql当中
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return mydb.OnFileUploadFinished(
		fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

// GetFileMeta : 通过Sha1值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// GetFileMetaDB : 通过Sha1值从mysql获取文件的元信息对象
func GetFileMetaDB(fileSha1 string) (*FileMeta, error) {
	tFile, err := mydb.GetFileMeta(fileSha1)
	// 如果获得信息为空或者存在error
	if tFile == nil || err != nil {
		return nil, err
	}
	// 进行赋值
	fMeta := FileMeta{
		FileSha1: tFile.FileHash,
		FileName: tFile.FileName.String,
		FileSize: tFile.FileSize.Int64,
		Location: tFile.FileAddr.String,
	}
	return &fMeta, nil
}

// GetLastFileMetas : 批量获取文件的元信息列表（按照注册时间排序的钱count个）
func GetLastFileMetas(count int) []FileMeta {
	// 创建一个FileMeta的数组，大小为已存在的所有fileMeta的总和（即当前最大）
	fMetaArray := make([]FileMeta, len(fileMetas))
	// 全部装入该数组
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}
	// 按照更新时间进行排序
	sort.Sort(ByUpLoadTime(fMetaArray))
	// 最后返回钱count个
	return fMetaArray[0:count]
}

// GetLastFileMetasDB : 从Mysql批量获取文件的元信息列表
func GetLastFileMetasDB(limit int) ([]FileMeta, error) {
	// 调用mysql接口函数
	tFiles, err := mydb.GetFileMetaList(limit)
	// err不为空，则返回空的切片
	if err != nil {
		return make([]FileMeta, 0), err
	}
	// 根据取得的长度创建一个FileMeta切片
	tFilesm := make([]FileMeta, len(tFiles))
	// 循环获取想要的数据
	for i := 0; i < len(tFilesm); i++ {
		tFilesm[i] = FileMeta{
			FileSha1: tFiles[i].FileHash,
			FileName: tFiles[i].FileName.String,
			FileSize: tFiles[i].FileSize.Int64,
			Location: tFiles[i].FileAddr.String,
		}
	}
	return tFilesm, nil
}

// RemoveFileMeta : 删除指定元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
