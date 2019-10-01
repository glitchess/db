package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var db map[string]*sql.DB

func Conn(name, user string){
	db[name], _ = sql.Open("mysql", a)
}

func Query(name, query string, args ...interface{}) map[int]map[string]string{
	db[name].Ping()
	var result = make(map[int]map[string]string)
	rows, err := db[name].Query(query, args...)
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
		fmt.Println(err)
	return result
}
}

func Exec(name, query string, args ...interface{}) int64{
	db[name].Ping()
	rows, err := db[name].Exec(query, args...)
	fmt.Println(err)
	if err==nil{
		result, _ := rows.LastInsertId()
		return result
	}else{
		return -1
	}
}
