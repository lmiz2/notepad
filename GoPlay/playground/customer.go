package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func attachCustomers() {
	// a, _ := strconv.Atoi(os.Args[1])
	// rs := doSignUp(a)
	// token := doLoginAndGetLoginToken(rs)
	// /Users/dokku/Desktop/etc/profiles.json 파일을 읽어온다. 이 파일은 프로필 정보들이 들어있다.

	// 파일을 읽어온다.
	file, err := os.Open("/Users/dokku/Desktop/etc/profiles.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 파일을 읽어서 json으로 Unmarshal한다.
	body, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var profileInfos []ProfileInfo
	json.Unmarshal(body, &profileInfos)
	// 프로필 정보들을 하나씩 보내면서 프로필 정보를 업데이트한다.
	for i, _ := range profileInfos {
		// for i, profileInfo := range profileInfos {
		rs := doSignUp(i)
		token := doLoginAndGetLoginToken(rs)
		println(token)
		// doPatchProfileInfo(token, profileInfo)

	}

}

func doSignUp(n int) LoginInfo {
	host := Host + "/api/v1/signup/id_and_pw"
	// n을 dummy 뒤에 붙이는 stirng 작성
	nStr := strconv.Itoa(n)
	dummy := "dummy" + nStr
	// Request Body 작성.
	reqBody := bytes.NewBufferString(`
	{ "loginId": "` + dummy + `", "loginPassword": "` + dummy + `"}
	`)
	// Request 보내기. application/json 타입으로 보낸다.
	resp, err := http.Post(host, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Response Body를 json으로 Unmarshal
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var rs LoginInfo
	json.Unmarshal(body, &rs)
	if rs.Id == "" || rs.LoginId == "" {
		return LoginInfo{
			Id:      "dummy" + nStr,
			LoginId: "dummy" + nStr,
		}
	}
	return rs
}

func doLoginAndGetLoginToken(loginInfo LoginInfo) string {
	host := Host + "/api/v1/login/id_and_pw"
	// Request Body 작성.
	reqBody := bytes.NewBufferString(`
	{ "loginId": "` + loginInfo.LoginId + `", "loginPassword": "` + loginInfo.LoginId + `"}
	`)
	// Request 보내기. application/json 타입으로 보낸다.
	resp, err := http.Post(host, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic("login fail")
	}
	defer resp.Body.Close()
	// Response Body를 json으로 Unmarshal
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	token := string(body)
	if token != "" {
		return token
	}
	panic(string(body))
}

func doPatchProfileInfo(token string, profileInfo ProfileInfo) {
	host := Host + "/api/v1/profile/info"
	// Request Body 작성.
	bodyStr := `{ "nickName": "` + profileInfo.NickName + `", "gender": "` + profileInfo.Gender + `", "age": "` + profileInfo.Age + `", "ofBelong": "` + profileInfo.OfBelong + `"}`
	println(bodyStr)
	reqBody := bytes.NewBufferString(bodyStr)
	// Request 보내기. application/json 타입으로 보낸다.
	req, err := http.NewRequest("PATCH", host, reqBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "*/*")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Response Body를 json으로 Unmarshal
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if string(body) != "" {
		println(string(body))
	}
	if resp.StatusCode != 200 {
		panic("patch fail")
	}
}

type LoginInfo struct {
	Id      string `json:"id"`
	LoginId string `json:"loginId"`
}

type ProfileInfo struct {
	NickName string `json:"nickName"`
	Gender   string `json:"gender"`
	Age      string `json:"age"`
	OfBelong string `json:"ofBelong"`
}
