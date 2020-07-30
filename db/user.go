package db

import (
	mydb "carrotCloud/db/mysql"
	"fmt"
)

// User : 用户表model
type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

// UserSignUp : 通过用户名及密码完成user表的注册操作
func UserSignUp(username string, passwd string) bool {

	// 语句预执行
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`, `user_pwd`) values (?,?) ")

	if err != nil {
		fmt.Println("Failed to insert, err :" + err.Error())
	}
	defer stmt.Close()

	// 语句执行
	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}

	// 影响行数测试,大于0则返回true
	if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
		return true
	}

	return false
}

// UserSignIn : 判断用户密码是否一致
func UserSignIn(username string, encpwd string) bool {

	stmt, err := mydb.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	// 指定的用户名进行查询
	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if rows == nil {
		// 用户名不存在
		fmt.Println("username not found: " + username)
		return false
	}

	pRows := mydb.ParseRows(rows)
	// 密码判断一致，则返回true
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}
	return false
}

// UpdateToken : 刷新用户登录的token
func UpdateToken(username string, token string) bool {
	// 预处理，进行更新操作
	stmt, err := mydb.DBConn().Prepare(
		"replace into tbl_user_token (`user_name`, `user_token`) values (?,?)")

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

// GetUserInfo : 查询用户信息
func GetUserInfo(username string) (User, error) {

	user := User{}

	stmt, err := mydb.DBConn().Prepare(
		"select user_name,signup_at from tbl_user where user_name=? limit 1")

	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}

	defer stmt.Close()

	// 执行查询的操作
	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

// UserExist : 查询用户是否存在
func UserExist(username string) (bool, error) {

	stmt, err := mydb.DBConn().Prepare(
		"select 1 from tbl_user where user_name = ? limit 1")
	// 预执行出错
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	defer stmt.Close()
	// 执行数据库查询
	rows, err := stmt.Query(username)
	if err != nil {
		return false, err
	}
	// 返回数据
	return rows.Next(), nil
}
