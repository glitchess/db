package db

import (
	//"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB

func Auth(auth string){
	db, _ = sql.Open("mysql", auth) //"pt1305g3_user:1BpNohC7@/pt1305g3_user"
}

func Query(query string, args ...interface{}) map[int]map[string]string{
	db.Ping()
	var result = make(map[int]map[string]string)
	rows, err := db.Query(query, args...)
	defer rows.Close()
	if err!=nil{
		return nil
	}else{
		col, _ := rows.Columns()
		i := 0
		count := len(col)
		values := make([]string, count)
		valuePtrs := make([]interface{}, count)
		for rows.Next(){
			arr := make(map[string]string)
			for i := range col {
				valuePtrs[i] = &values[i]
			}
			rows.Scan(valuePtrs...)
			for index, val := range values {
					arr[col[index]] = val
			}
			result[i] = arr
			i++
		}
	return result
}
}

func Exec(query string, args ...interface{}) int64{
	db.Ping()
	rows, err := db.Exec(query, args...)
	if err==nil{
		result, _ := rows.LastInsertId()
		return result
	}else{
		return -1
	}
}