package db

import (
	mydb "carrotCloud/db/mysql"
	"fmt"
	"time"
)

// UserFile : 用户文件表结构体
type UserFile struct {
	UserName    string // 比meta原信息多了一个UserName，用了保存是哪个用户
	FileHash    string
	FileName    string
	FileSize    int64
	UploadAt    string
	LastUpdated string
}

// OnUserFileUploadFinished ： 更新文件用户表
func OnUserFileUploadFinished(userName, fileHash, fileName string, fileSize int64) bool {

	// 语句预执行将数据插入表中
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user_file (`user_name`,`file_sha1`,`file_name`," +
			"`file_size`,`upload_at`) values (?,?,?,?,?)")
	if err != nil {
		return false
	}

	defer stmt.Close()

	// 将传入的参数进行上面准备语句的执行
	_, err = stmt.Exec(userName, fileHash, fileName, fileSize, time.Now())
	if err != nil {
		return false
	}
	return true
}

// QueryUserFileMetas : 批量获取用户的文件信息
func QueryUserFileMetas(userName string, limit int) ([]UserFile, error) {

	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1,file_name,file_size,upload_at," +
			"last_update from tbl_user_file where user_name=? limit ?")

	if err != nil {
		// 如果发生错误，则返回该错误
		return nil, err
	}
	// 进行语句查询并返回多行数据
	rows, err := stmt.Query(userName, limit)
	if err != nil {
		return nil, err
	}

	// 定义一个UserFile的用以接受返回的用户信息
	var userFiles []UserFile

	for rows.Next() {
		// 定义一个零时UserFile结构体
		uFile := UserFile{}
		// 将每行的数据进行赋值
		err = rows.Scan(&uFile.FileHash, &uFile.FileName, &uFile.FileSize,
			&uFile.UploadAt, &uFile.LastUpdated)
		// 发生了错误则打印并且break
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// 将每行数据插入到userFiles切片当中
		userFiles = append(userFiles, uFile)
	}
	// 最后返回该切片，并且error为nil
	return userFiles, nil
}

// QueryUserFileMeta : 获取用户单个文件信息
func QueryUserFileMeta(username string, filehash string) (*UserFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1,file_name,file_size,upload_at," +
			"last_update from tbl_user_file where user_name=? and file_sha1=?  limit 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, filehash)
	if err != nil {
		return nil, err
	}

	ufile := UserFile{}
	if rows.Next() {
		err = rows.Scan(&ufile.FileHash, &ufile.FileName, &ufile.FileSize,
			&ufile.UploadAt, &ufile.LastUpdated)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	return &ufile, nil
}
