package main

func main() {
	rs := doLoginAndGetLoginToken(LoginInfo{
		Id:      "dummy14",
		LoginId: "dummy14",
	})
	println(rs)
}
