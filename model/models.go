package model

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/solenovex/web-tuborial/common"
)

func Get_all() (all []common.App, err error) {
	Sql := "SELECT * FROM dbo.userRole"
	rows, err := common.Db.Query(Sql)
	if err != nil {
		log.Fatalln(err)
	}
	L := common.App{}
	for rows.Next() {
		err = rows.Scan(&L.Id, &L.Key, &L.Role)
		all = append(all, L)
	}
	return
}

//get one
func Get_one(id string) (record common.App) {
	Sql := "SELECT * FROM dbo.staff WHERE number = @ID"
	err := common.Db.QueryRow(Sql, sql.Named("ID", id)).Scan(&record.Id, &record.Key)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Add_one(number, name, password string) (err error) {
	Sql := "INSERT INTO dbo.User_1 (u_id, userName, password, Name) VALUES (@nu, @na, @pas, @nm)"
	stmt, err := common.Db.Prepare(Sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(sql.Named("nu", number), sql.Named("na", name), sql.Named("pas", password), sql.Named("nm", "默认"))
	if err != nil {
		return
	}
	return
}

func Edit(record common.App) (err error) {
	Sql := "UPDATE dbo.staff SET name=@Name WHERE number=@Id"
	_, err = common.Db.Exec(Sql, sql.Named("Name", record.Key), sql.Named("Id", record.Id))
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Delete(id string) (err error) {
	Sql := "DELETE FROM dbo.User_1 WHERE u_id=@Id"
	_, err = common.Db.Exec(Sql, sql.Named("Id", id))
	return
}

//get purview
func Get_purview(uid string) (all []string, err error) {
	Sql := `SELECT p.descrip FROM User_1 u, U_Role ur, R_Purview rp, Purview p 
	WHERE u.u_id = ur.u_id AND ur.r_id = rp.r_id AND rp.qx_id = p.qx_id And u.u_id = @uid`
	rows, err := common.Db.Query(Sql, sql.Named("uid", uid))
	if err != nil {
		log.Fatalln(err)
	}
	var L string
	for rows.Next() {
		err = rows.Scan(&L)
		all = append(all, L)
	}
	return
}

func Test(uid string) (all []common.User_p, err error) {
	Sql := `SELECT u.u_id, u.Name, p.descrip FROM User_1 u, U_Role ur, R_Purview rp, Purview p 
	WHERE u.u_id = ur.u_id AND ur.r_id = rp.r_id AND rp.qx_id = p.qx_id And u.u_id = @uid`
	rows, err := common.Db.Query(Sql, sql.Named("uid", uid))
	if err != nil {
		return
	}
	L := common.User_p{}
	for rows.Next() {
		rows.Scan(&L.Id, &L.Name, L.Purview)
		all = append(all, L)
	}
	return
}
