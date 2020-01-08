package main

import (
	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	var argocdCluster string
	var appname string
	argocdCluster = "192.168.48.12:31410" //argocd private cluster
	appname = "handson1"
	//fmt.Scanf("%s", &ServerBitbucket);
	//fmt.Scanf("%s", &projectname)

	url:=fmt.Sprintf("https://%s/api/v1/applications", argocdCluster) //API calling for get application list, not completed
	//url:=fmt.Sprintf("https://%s/api/v1/applications/%s/sync", argocdCluster, appname) //API call for sync, not completed
	//jsonData := map[string]string{"name":"<REPONAME>","scmID":"git","forkable":"true"} //this data should be written with another config
	jsonValue,_:=json.Marshal(jsonData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("<USERNAME>", "<PASSWORD>")
	fmt.Println("++",req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}