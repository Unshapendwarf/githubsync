package argocdinfo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	//"github.com/ghodss/yaml"
	"io/ioutil"
	"net/http"
)

type ArgoCDinfo struct {
	Username string
	Password string
	IPport   string
	Token    string
}

func GetToken(cluster *ArgoCDinfo) string {
	//cluster IPport, Username, Password info validation
	if len(cluster.IPport) <= 3 {
		panic("cluster.IPport is invalid!")
	}
	if len(cluster.Username) <=0 {
		panic("cluster.Username is invalid!")
	}
	if len(cluster.Password) <=0 {
		panic("cluster.Password is invalid1")
	}
	//인증서 없이 접근
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	//accountmap default id/pw setting: admin/Password
	accountmap := map[string] string{"username" : cluster.Username, "password" : cluster.Password}
	tokenmap := map[string]string{"token":"None"}

	url:=fmt.Sprintf("http://%s/api/v1/session", cluster.IPport)
	bodyjson, _ :=json.Marshal(accountmap)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyjson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	tokenbytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(tokenbytes, &tokenmap)
	if err != nil {
		panic(err)
	}

	token := tokenmap["token"]
	fmt.Printf("token -> %s\n", token)

	return tokenmap["token"]
}

func SetToken(cluster *ArgoCDinfo, token string){
	//token validation
	if token == ""{
		panic("cluster.Token is empty!!")
	}
	cluster.Token = token
}