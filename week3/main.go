package main

import (
	"fmt"
	"encoding/json"
)

type Setting struct {
	Host string
	IP string
	Port int
}

type List struct {
	Google Setting
	Phpadmin Setting
	Jsonviewer Setting
	Github Setting
}

func main()  {
	Q1()
	Q2()
}

func Q1()  {
	type t struct {
		Ret []struct{
			UserID int `json:"id"`
		} `json:"ret"`
	}

	var userInfo t

	json.Unmarshal(getApiResult(), &userInfo)

	fmt.Println(userInfo)
}

func Q2()  {
	//a := "IP:192.17.55.3 Host:docs.google.com Port:80"
	//b := "IP:192.17.33.17 Host:ts-phpadmin.vir999.com Port:80"
	//c := "IP:192.17.99.52 Host:jsonviewer.stack.hu Port:7711"

	//d := "IP:177.18.2.33 Host:github.com Port:80"


	var list List
	list.Google.IP = "192.17.55.3"
	list.Google.Host = "docs.google.com"
	list.Google.Port = 80

	list.Phpadmin.IP = "192.17.33.17"
	list.Phpadmin.Host = "ts-phpadmin.vir999.com"
	list.Phpadmin.Port = 80


	list.Jsonviewer.IP = "192.17.99.52"
	list.Jsonviewer.Host = "jsonviewer.stack.hu"
	list.Jsonviewer.Port = 7711

	addSetting(&list)
	updateSetting(&list)

	fmt.Printf("%+v", list)
}

func addSetting(list *List)  {
	list.Github.IP = "177.18.2.33"
	list.Github.Host = "github.com"
	list.Github.Port = 80
}

func updateSetting(list *List)  {
	list.Jsonviewer.Port = 80
}

func getApiResult() []byte {
	userData := `{"ret":[{"id":148500286,"parent_id":1414663,"parent":1414663,"all_parents":[1414663,59735,58971,24367,6],"domain":6,"last_login":"2018-06-12T09:17:51+0800","currency":"CNY","password_expire_at":"2020-09-04T18:19:02+0800","password_reset":false,"last_bank":54,"username":"wesley01","block":false,"bankrupt":false,"cash":{"id":71814701,"user_id":148500286,"balance":4923642.0545,"pre_sub":0,"pre_add":0,"currency":"CNY","last_entry_at":"20180612155900"},"cash_fake":null,"card":null,"enabled_card":null}],"result":"ok","profile":{"execution_time":14,"server_name":"ipl-web01.rd5.qa"}}`

	return []byte(userData)
}