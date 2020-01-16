package controller

import (
	//"bytes"
	"crypto/tls"
	//"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"net/http"
	"githubsync/argocdinfo"
	"githubsync/apphandle"
)

//GetApps; get all applications in argoCD cluster
func GetApps( cluster *argocdinfo.ArgoCDinfo){
	//appname := "appsync"
	url:=fmt.Sprintf("http://%s/api/v1/applications", cluster.IPport) //API calling for get application list, not completed

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
	req.Header.Set("Authorization", "Bearer " +cluster.Token)

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

	if yamlbytes == nil {
		panic("yamlbytes is empty")
	}
	//file write
	/*
	err = ioutil.WriteFile("./response/response.yaml", yamlbytes, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(y))
	strresp := string(bytes) //바이트를 문자열로
	fmt.Println(strresp)
	 */
}

// CheckAPI; sync, create, delete 중 어떤 것인지 확인하고 해당하는 api call을 실행한다. 지금은 일단 sync만 해본다.
func CheckAPI( checker string , cluster *argocdinfo.ArgoCDinfo) {
	//getapps, gettoken 등 sync, create, delete 외의 함수는 여기에서 다루지 않는다. 지금은 임시로 작
	if checker == "apps" {
		fmt.Println("CheckAPI == GetApps")
		GetApps(cluster)
	} else if checker == "sync"{
		fmt.Println("CheckAPI == sync")
		apphandle.SyncApp(cluster, "appsync")
	} else if checker == "gettoken"{
		fmt.Println("CheckAPI == Token")
		token := argocdinfo.GetToken(cluster)
		argocdinfo.SetToken(cluster, token)
	} else if checker == "delete"{
		fmt.Println("CheckAPT == Delete")
		apphandle.Delete(cluster, "appsync")
	} else if checker == "create" {
		fmt.Println("CheckAPT == create")
		apphandle.Create(cluster)
	} else{
		fmt.Println("Invalid API")
	}
}