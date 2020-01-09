package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"net/http"
)

type argoCDinfo struct {
	username string
	password string
	iport string
	token string
}

func filetest() {
	//파일 읽기
	bytes, err := ioutil.ReadFile("/Users/mf839-027/Documents/hong/README.md")
	if err != nil {
		panic(err)
	}

	//파일 내용 출력, bytes to string
	//fmt.Print(string(bytes))

	err = ioutil.WriteFile("/Users/mf839-027/Documents/appsync/githubsync/cicd/TEST2.md", bytes, 0)
	if err != nil {
		panic(err)
	}
}

//*********************** syncapp ************************
func syncapp(cluster argoCDinfo, app string){
	url:=fmt.Sprintf("http://%s/api/v1/applications/%s/sync", cluster.iport, app)
	fmt.Println(url)

	//json data body

	jsonData := map[string]string{"name":"<REPONAME>","scmID":"git","forkable":"true"} //this data should be written with another config
	jsonValue,_:=json.Marshal(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)// handle err
	}

	//request header setting; authorization is required(to get in argocd cluster)
	req.Header.Set("Authorization", "Bearer " +cluster.token)
	req.SetBasicAuth("<USERNAME>", "<PASSWORD>")

	//client gen; insecure connection(without certification)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)// handle err
	}

	defer resp.Body.Close()

	//data reading from response
	bytes, _ := ioutil.ReadAll(resp.Body)

	//convert bytes from json to yaml
	yamlbytes, err := yaml.JSONToYAML(bytes)
	if err != nil {
		panic(err)
		return
	}

	//file write
	err = ioutil.WriteFile("/Users/mf839-027/Documents/appsync/githubsync/cicd/response.yaml", yamlbytes, 0)
	if err != nil {
		panic(err)
	}

}
//********************************************************

func gettoken(cluster *argoCDinfo) {
	//인증서 없이 접근
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	//accountmap default id/pw setting: admin/password
	accountmap := map[string] string{"username" : cluster.username, "password" : cluster.password}
	tokenmap := map[string]string{"token":"None"}

	bodyjson, _ :=json.Marshal(accountmap)

	url:=fmt.Sprintf("http://%s/api/v1/session", cluster.iport)
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

	cluster.token = tokenmap["token"]
	fmt.Printf("this is cluster.token value -> %s\n", cluster.token)
}

//getapps; get all applications in argoCD cluster
func getapps( cluster argoCDinfo ){
	//appname := "appsync"
	url:=fmt.Sprintf("http://%s/api/v1/applications", cluster.iport) //API calling for get application list, not completed

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	//request generating to get application info, http method is GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)// handle err
	}

	//request header setting; authorization is required(to get in argocd cluster)
	req.Header.Set("Authorization", "Bearer " +cluster.token)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)// handle err
	}
	defer resp.Body.Close()

	//data reading from response
	bytes, _ := ioutil.ReadAll(resp.Body)

	//convert bytes from json to yaml
	yamlbytes, err := yaml.JSONToYAML(bytes)
	if err != nil {
		panic(err)
		return
	}

	//file write
	err = ioutil.WriteFile("/Users/mf839-027/Documents/appsync/githubsync/cicd/response.yaml", yamlbytes, 0)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(y))
	//strresp := string(bytes) //바이트를 문자열로
	//fmt.Println(strresp)
}

// argoCDAPIchecker; sync, create, delete 중 어떤 것인지 확인하고 해당하는 api call을 실행한다. 지금은 일단 sync만 해본다.
func argoCDAPIchecker( checker string , cluster argoCDinfo ) {
	if checker == "apps" {
		fmt.Println("argoCDAPIchecker == getapps")
		getapps(cluster)
	} else if checker == "sync"{
		fmt.Println("argoCDAPIchecker == sync")
		syncapp(cluster, "appsync")
	}
}

func main() {
	// 0. set variable
	filetest()

	var cluster argoCDinfo

	// 1. argocd cluster,
	cluster.iport = "192.168.48.12:31410" //argocd cluster
	cluster.username = "admin"
	cluster.password = "1222"
	//cluster.token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NzgyOTEyOTEsImlzcyI6ImFyZ29jZCIsIm5iZiI6MTU3ODI5MTI5MSwic3ViIjoiYWRtaW4ifQ.O_WQAZ5R6Jdca3uZji6LVrmYY461feHGwRmhvDo0uUI"
	gettoken(&cluster)
	argoCDAPIchecker("getapps", cluster)
}