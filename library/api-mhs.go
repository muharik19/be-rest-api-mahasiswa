package library

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbport = "3306"
	dbaddr = "127.0.0.1"
)

func queryRow(nim string) checkNim {
	var nim_check = checkNim{}
	db, err := sql.Open("mysql", "root:@tcp("+dbaddr+":"+dbport+")/go_db")

	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	err = db.QueryRow(`SELECT nim FROM mahasiswa WHERE nim = ?`, nim).Scan(&nim_check.nim)
	return nim_check
}

func MahasiswaAdd(c *gin.Context) {
	var status bool
	var pesan string
	db, err := sql.Open("mysql", "root:@tcp("+dbaddr+":"+dbport+")/go_db")

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"Error ": err.Error()})
		return
	}
	defer db.Close()

	t := time.Now()
	datetime := t.Format("20060102150405")
	nim := c.PostForm("nim")
	nama := c.PostForm("nama")
	jurusan := c.PostForm("jurusan")
	no_hp := c.PostForm("no_hp")
	file, err := c.FormFile("image")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	imageName := "http://192.168.43.237:83/mima/v1/images/" + datetime + file.Filename
	path := "image/" + datetime + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
	}

	nim_check := queryRow(nim)

	if (checkNim{}) == nim_check {
		stmt, err := db.Prepare(`INSERT INTO mahasiswa SET nama = ?, photo = ?, nim = ?, jurusan = ?, no_hp = ?`)
		if err == nil {
			_, err := stmt.Exec(&nama, &imageName, &nim, &jurusan, &no_hp)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
		defer stmt.Close()
		status = true
		pesan = "Created Successfully."
	} else {
		status = false
		pesan = "Already Exists NIM " + nim + ""
	}
	RespVer := RespUpload{
		Status:  status,
		Message: pesan,
	}
	Respjson, _ := json.Marshal(RespVer)
	c.String(200, string(Respjson))
}

func MahasiswaAll(c *gin.Context) {
	var nim, nama, jurusan, hp, photo string
	var id int
	var RespStruct []RespMahasiswaAll

	db, err := sql.Open("mysql", "root:@tcp("+dbaddr+":"+dbport+")/go_db")

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"Error ": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, nim, nama, jurusan, no_hp, photo FROM mahasiswa ORDER BY id DESC`)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"Error ": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&id, &nim, &nama, &jurusan, &hp, &photo)
		RespStruct = append(RespStruct, RespMahasiswaAll{
			ID:      id,
			Nim:     nim,
			Nama:    nama,
			Jurusan: jurusan,
			Hp:      hp,
			Photo:   photo,
		})
	}

	RespMahasiswa := RespResult{
		Success: true,
		Message: "Success",
		Results: RespStruct,
	}
	jsonResp, _ := json.Marshal(RespMahasiswa)
	c.String(200, string(jsonResp))
}

func MahasiswaUpdate(c *gin.Context) {
	db, err := sql.Open("mysql", "root:@tcp("+dbaddr+":"+dbport+")/go_db")

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"Error ": err.Error()})
		return
	}
	defer db.Close()

	t := time.Now()
	datetime := t.Format("20060102150405")
	id := c.PostForm("id")
	nim := c.PostForm("nim")
	nama := c.PostForm("nama")
	jurusan := c.PostForm("jurusan")
	no_hp := c.PostForm("no_hp")
	file, err := c.FormFile("image")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	imageName := "http://192.168.43.237:83/mima/v1/images/" + datetime + file.Filename
	path := "image/" + datetime + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	stmt, err := db.Prepare(`UPDATE mahasiswa SET nama=?, photo=?, nim=?, jurusan=?, no_hp=? WHERE id=?`)
	if err == nil {
		_, err := stmt.Exec(&nama, &imageName, &nim, &jurusan, &no_hp, &id)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	defer stmt.Close()

	RespVer := RespUpload{
		Status:  true,
		Message: "Updated Successfully.",
	}
	Respjson, _ := json.Marshal(RespVer)
	c.String(200, string(Respjson))
}

func MahasiswaDelete(c *gin.Context) {
	db, err := sql.Open("mysql", "root:@tcp("+dbaddr+":"+dbport+")/go_db")

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{"Error ": err.Error()})
		return
	}
	defer db.Close()

	id := c.PostForm("id")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	stmt, err := db.Prepare(`DELETE FROM mahasiswa WHERE id = ?`)
	if err == nil {
		_, err := stmt.Exec(&id)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	defer stmt.Close()

	RespVer := RespUpload{
		Status:  true,
		Message: "Delete Successfully.",
	}
	Respjson, _ := json.Marshal(RespVer)
	c.String(200, string(Respjson))
}
