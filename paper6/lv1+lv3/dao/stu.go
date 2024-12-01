package dao

import (
	"go_papers/paper6/lv1+lv3/model"
	"log"
	"time"
)

var DB = Drive()

// 添加学生
func AddStudent(sex, name, comefrom string) bool {
	t := time.Now()
	tTos := t.Format("2006-01-02 15:04:05")
	// 写入数据库
	stmt, err := DB.Prepare("insert into stumanage (name,sex,comefrom,optime) values (?,?,?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	res, err := stmt.Exec(name, sex, comefrom, tTos)
	if err != nil {
		log.Println(err)
		return false
	}
	row, _ := res.RowsAffected()
	if row == 0 {
		return false
	}
	return true
}

// 删除学生
func Delete(name string) bool {
	stmt, err := DB.Prepare("delete from stumanage where name=?")
	if err != nil {
		log.Println(err)
	}
	res, err := stmt.Exec(name)
	if err != nil {
		log.Println(err)
	}
	row, _ := res.RowsAffected()
	if row == 0 {
		return false
	}
	return true
}

// 修改信息
func Profile(S model.Student) bool {
	t := time.Now()
	tTos := t.Format("2006-01-02 15:04:05")

	stmt, err := DB.Prepare("update stumanage set sex=?,comefrom=?,optime=? where name=?")
	if err != nil {
		log.Println(err)
		return false
	}
	row, err := stmt.Exec(S.Sex, S.Comefrom, tTos, S.Name)
	if err != nil {
		log.Println(err)
		return false
	}
	n, _ := row.RowsAffected()
	if n == 0 {
		return false
	} else {
		return true
	}
}

// 删除指定信息
func DeleteSingle(args string, name string) bool {
	t := time.Now()
	tTos := t.Format("2006-01-02 15:04:05")
	s := "update stumanage set optime=?," + args + "=?" + " " + " where name=?"
	stmt, err := DB.Prepare(s)
	if err != nil {
		log.Println(err)
		return false
	}
	row, err := stmt.Exec(tTos, "", name)
	if err != nil {
		log.Println(err)
		return false
	}
	n, _ := row.RowsAffected()
	if n == 0 {
		return false
	} else {
		return true
	}
}
