package dao

import (
	"go_papers/paper6/lv2+lv3/model"
	"log"
)

var DB = Drive()

func AddUser(username, password string) bool {
	stmt, err := DB.Prepare("insert into users (username,password) values (?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	res, err := stmt.Exec(username, password)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	row, _ := res.RowsAffected()
	if row == 0 {
		return false
	} else {
		return true
	}
}

func SearchUser(username string) bool {
	stmt, err := DB.Prepare("select username from users")
	if err != nil {
		log.Println(err)
		return true
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	for rows.Next() {
		var U model.User
		err := rows.Scan(&U.Username)
		if err != nil {
			log.Println(err)
			return true
		}
		if username == U.Username {
			return true
		}
	}
	return false
}

func SearchPwdFromUsername(username, password string) bool {
	stmt, err := DB.Prepare("select password from users where username=? ")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()
	row := stmt.QueryRow(username)
	var U model.User
	err = row.Scan(&U.Password)
	if err != nil {
		log.Println(err)
		return false
	}
	if U.Password == password {
		return true
	} else {
		return false
	}
}

func SetSecurityQu(username, password, security_question, security_answer string) bool {
	if SearchPwdFromUsername(username, password) {
		stmt, err := DB.Prepare("update users set security_question=?,security_answer=? where username=?")
		if err != nil {
			log.Println(err)
			return false
		}
		defer stmt.Close()
		row, _ := stmt.Exec(security_question, security_answer, username)
		n, _ := row.RowsAffected()
		if n == 0 {
			return false
		}
		return true
	}
	return false
}

// todo 在变量名的使用和错误的处理上还需要修改

func ResetPassword(username, newpassword, security_answer string) bool {
	// 感觉这里在实现效果时候应该要加载一个自动抛出问题的界面
	stmt1, err := DB.Prepare("select security_answer from users where username=?")
	if err != nil {
		log.Println(err)
		return false
	}
	row := stmt1.QueryRow(username)
	var ans string
	_ = row.Scan(&ans)
	if ans == security_answer {
		stmt2, err := DB.Prepare("update users set password=? where username=?")
		if err != nil {
			log.Println(err)
			return false
		}
		row, _ := stmt2.Exec(newpassword, username)
		n, _ := row.RowsAffected()
		if n == 0 {
			return false
		}
		return true
	}
	return false
}
