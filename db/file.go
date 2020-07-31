package db

import (
	mydb "carrotCloud/db/mysql"
	"database/sql"
	"fmt"
)

// TableFile : 文件表结构体
type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

// GetFileMeta : 从mysql当中获取文件元信息
func GetFileMeta(fileHash string) (*TableFile, error) {

	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1,file_addr,file_name,file_size from tbl_file " +
			"where file_sha1=? and status=1 limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()

	// 新建一个空的TableFile结构体用来接收查询返回的数据
	tFile := TableFile{}
	// 进行查询，得到一行数据，并存入结构体当中
	err = stmt.QueryRow(fileHash).Scan(
		&tFile.FileHash, &tFile.FileAddr, &tFile.FileName, &tFile.FileSize)
	// 进行错误信息处理
	if err != nil {
		if err == sql.ErrNoRows {
			// 查不到对应的记录，返回参数及错误都为nil
			return nil, nil
		} else {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	return &tFile, nil
}

// GetFileMetaList : 从mysql当中根据limit的数目批量获取文件元信息
func GetFileMetaList(limit int) ([]TableFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1, file_addr, file_name, file_size, from tbl_file" +
			"where status=1 limit?")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()
	// 进行查询并返回多行数据
	rows, err := stmt.Query(limit)
	// 进行错误处理
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	// 新建一个空的TableFile用来存储并返回
	var tFiles []TableFile
	// 将每行数据进行处理
	for i := 0; i < len(values) && rows.Next(); i++ {
		tFile := TableFile{}
		err = rows.Scan(&tFile.FileHash, &tFile.FileAddr,
			&tFile.FileName, &tFile.FileSize)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		tFiles = append(tFiles, tFile)
	}
	// 打印接收了多少行数据
	fmt.Println(len(tFiles))
	// 将数据返回
	return tFiles, nil
}

// OnFileUploadFinished : 文件上传完成，保存meta
func OnFileUploadFinished(fileHash, fileName string, fileSize int64, fileAddr string) bool {
	// Mysql语句欲写入，注意status为1
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file (`file_sha1`, `file_name`, `file_size`," +
			"`file_addr`, `status`) values(?,?,?,?,1)")
	// 错误处理
	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return false
	}
	defer stmt.Close()
	// 传入参数并且执行语句
	ret, err := stmt.Exec(fileHash, fileName, fileSize, fileAddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	// 返回影响行数，如果小于0则证明插入失败
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("File with hash:%s has been uploaded before", fileHash)
		}
		return true
	}
	return false

}
