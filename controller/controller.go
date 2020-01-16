package githubsync

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"net/http"
)

type ArgoCDinfo struct {
	Username string
	Password string
	IPport   string
	Token    string
}


//*********************** SyncApp ************************
func SampleSyncApp(cluster *ArgoCDinfo, app string){
	url:=fmt.Sprintf("http://%s/api/v1/applications/%s/sync", cluster.IPport, app)
	fmt.Println(url)

	//json data body

	jsonData := map[string]string{"name":"<REPONAME>","scmID":"git","forkable":"true"} //this data should be written with another config
	jsonValue,_:=json.Marshal(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)// handle err
	}

	//request header setting; authorization is required(to get in argocd cluster)
	req.Header.Set("Authorization", "Bearer " +cluster.Token)
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

func GetToken(cluster *ArgoCDinfo) {
	//인증서 없이 접근
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	//accountmap default id/pw setting: admin/Password
	accountmap := map[string] string{"Username" : cluster.Username, "Password" : cluster.Password}
	tokenmap := map[string]string{"Token":"None"}

	bodyjson, _ :=json.Marshal(accountmap)

	url:=fmt.Sprintf("http://%s/api/v1/session", cluster.IPport)
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

	cluster.Token = tokenmap["Token"]
	fmt.Printf("this is cluster.Token value -> %s\n", cluster.Token)
}

//GetApps; get all applications in argoCD cluster
func GetApps( cluster *ArgoCDinfo){
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
func CheckAPI( checker string , cluster *ArgoCDinfo) {
	//getapps, gettoken 등 sync, create, delete 외의 함수는 여기에서 다루지 않는다. 지금은 임시로 작
	if checker == "apps" {
		fmt.Println("CheckAPI == GetApps")
		GetApps(cluster)
	} else if checker == "sync"{
		fmt.Println("CheckAPI == sync")
		SyncApp(cluster, "appsync")
	} else if checker == "gettoken"{
		fmt.Println("CheckAPI == Token")
		GetToken(cluster)
	} else {
		fmt.Println("Invalid API")
	}
}