package example

import (
	"SpiderFrame/models"
	"fmt"
)

//search sql
func GetRow(){
	dbw:=models.InitMysql()
	rows,err := dbw.Query("select bm_id from fb_ad_bm limit 5")
	if err != nil {
		fmt.Printf("get data error: %v\n", err)
	}
	defer rows.Close()
	var bmId string
	for rows.Next() {
		rows.Scan(&bmId)
		fmt.Println(bmId)
	}
	dbw.Close()
}