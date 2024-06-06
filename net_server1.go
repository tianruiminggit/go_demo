package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//type Role struct {
//	RoleUid int
//	Name    string
//}

type Role struct {
	RoleUid int    `json:"role_uid"`
	Name    string `json:"name"`
}

var db, dbErr = sql.Open("mysql", "root:database1121@tcp(192.168.2.189:3306)/go_db")

func IndexResponse(w http.ResponseWriter, req *http.Request) {
	fmt.Println("IndexResponse")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func LoginResponse(w http.ResponseWriter, req *http.Request) {
	fmt.Println("LoginResponse" + req.FormValue("name"))
	roleUid, err := strconv.Atoi(req.FormValue("role_uid"))
	name := req.FormValue("name")
	if err != nil {
		fmt.Println("参数异常", err)
	}
	var role Role
	rows, _ := db.Query("select * from role where role_uid = ?", roleUid)
	rows.Next()
	rows.Scan(&role.RoleUid, &role.Name)
	if role.Name != name {
		fmt.Fprintf(w, "login err:%s", role.Name)
	} else {
		fmt.Fprintf(w, "welcome....%s", name)
	}

	defer rows.Close()
}

func GetRoleResponse(w http.ResponseWriter, req *http.Request) {
	fmt.Println("GetRoleResponse" + req.FormValue("name"))
	fmt.Fprintf(w, "Login  loading....,")
	if dbErr != nil {
		fmt.Println("数据库链接失败:", dbErr)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("数据库链接失败:", pingErr)
	}
	rows, _ := db.Query("select * from role")
	defer rows.Close()
	for rows.Next() {
		var role Role
		rows.Scan(&role.RoleUid, &role.Name)
		dataB, _ := json.Marshal(role)
		fmt.Fprintf(w, "role:", string(dataB))
	}
}

func main() {
	http.HandleFunc("/index", IndexResponse)
	http.HandleFunc("/login", LoginResponse)
	http.HandleFunc("/get_role", GetRoleResponse)
	err := http.ListenAndServe("192.168.2.189:7662", nil)
	defer db.Close()
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}

}
