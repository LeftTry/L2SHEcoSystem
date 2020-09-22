package DB

import (
	"L2SHEcoSystem/BuisnessRules/Student"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//Opening database with name using root and password
func OpenDB(name string, root string, password string) {
	var err error
	db, err = sql.Open("mysql", root+"+"+password+"@/"+name)
	if err != nil {
		panic(err)
	}
}

//Creating student and adding him to database
func CreateStudentInDB(s Student.Student) bool {
	result, err := db.Exec("insert into studentsDB.Students (first_name, last_name, class, email)"+
		" values (?, ?, ?, ?)", s.FirstName, s.LastName, s.Class, s.Email)
	if err != nil {
		panic(err)
		return false
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return false
	}
	if err != nil {
		panic(err)
		return false
	}
	return true
}

//Gives student by name, class and email username and password
func RegisterStudentInDB(s Student.Student) bool {

	return true
}

func CheckInDB(s Student.Student) sql.Result {
	id, err := db.Exec("select * from studentsdb.Students where first_name = ? and last_name = ? and class = ? and email = ?", s.FirstName, s.LastName, s.Class, s.Email)
	if err != nil {
		panic(err)

	}
	return id
}
