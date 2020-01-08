package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//***********************************************************
	// 0. set variable
	filetest()
	var argocdCluster string
	var appname string
	var argocdtoken string

	// 1. argocd cluster,
	argocdCluster = "192.168.48.12:31410" //argocd cluster
	appname = "handson1" //argocd application name

	url:=fmt.Sprintf("http://%s/api/v1/applications", argocdCluster) //API calling for get application list, not completed
	url2:=fmt.Sprintf("https://%s/api/v1/applications/%s/managed-resources", argocdCluster, appname) //API call for sync, not completed
	fmt.Printf("url : %s\n", url)
	fmt.Printf("url2: %s\n", url2)
	//여기는 http request를 보내는 부분인데 아마 여기에 내용을 정의해줘야한다
	//jsonData := map[string]string{"name":"<REPONAME>","scmID":"git","forkable":"true"} //this data should be written with another config
	//jsonValue,_:=json.Marshal(jsonData)
	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	//req.Header.Set("Content-Type", "application/json")
	//req.SetBasicAuth("<USERNAME>", "<PASSWORD>")

	//argoCD token; 이 부분은 따로 파일에서 이 값을 받아오도록하자 지금은 그 토큰이 로컬에 있어서 이렇게 함.
	//argocdtoken = os.ExpandEnv("Bearer $ARGOCD_TOKEN")
	argocdtoken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NzgyOTEyOTEsImlzcyI6ImFyZ29jZCIsIm5iZiI6MTU3ODI5MTI5MSwic3ViIjoiYWRtaW4ifQ.O_WQAZ5R6Jdca3uZji6LVrmYY461feHGwRmhvDo0uUI"

	//client gen; insecure connection(without certification)
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
	req.Header.Set("Authorization", argocdtoken)
	fmt.Println("++", req)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)// handle err
	}

	fmt.Println("++")
	defer resp.Body.Close()

	//data reading from response
	bytes, _ := ioutil.ReadAll(resp.Body)
	strresp := string(bytes) //바이트를 문자열로
	fmt.Println(strresp)
}