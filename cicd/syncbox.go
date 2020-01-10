package cicd

import (
	""
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"net/http"
)


//*********************** SyncApp ************************
func SSyncApp(cluster *AgoCDinfo, app string){
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
