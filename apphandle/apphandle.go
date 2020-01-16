package apphandle

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"githubsync/argocdinfo" //local
	"io/ioutil"
	"net/http"
)

//*********************** SyncApp ************************
func SyncApp(cluster *argocdinfo.ArgoCDinfo , app string){
	url:=fmt.Sprintf("http://%s/api/v1/applications/%s/sync", cluster.IPport, app)
	fmt.Println(url)
	//json data body

	jsonData := map[string] string{"name": app} //this data should be written with another config
	jsonValue,_:=json.Marshal(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	//req, err := http.NewRequest("POST", url, jsonValue)
	if err != nil {
		panic(err)// handle err
	}

	//req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NzgyOTEyOTEsImlzcyI6ImFyZ29jZCIsIm5iZiI6MTU3ODI5MTI5MSwic3ViIjoiYWRtaW4ifQ.O_WQAZ5R6Jdca3uZji6LVrmYY461feHGwRmhvDo0uUI")
	//request header setting; authorization is required(to get in argocd cluster)
	req.Header.Add("Authorization", "Bearer " +cluster.Token)
	//req.SetBasicAuth("<USERNAME>", "<PASSWORD>")

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
	err = ioutil.WriteFile("/Users/mf839-027/Documents/appsync/githubsync/apphandle/syncresp.yaml", yamlbytes, 0)
	if err != nil {
		panic(err)
	}
}

/*func Update(cluster *argocdinfo.ArgoCDinfo , app string){
	url:=fmt.Sprintf("http://%s/api/v1/applications/%s", cluster.IPport, app)
	//json data body
	jsonData := map[string] string{"name": app} //this data should be written with another config
	jsonValue,_:=json.Marshal(jsonData)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	//req, err := http.NewRequest("POST", url, jsonValue)
	if err != nil {
		panic(err)// handle err
	}

	//req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NzgyOTEyOTEsImlzcyI6ImFyZ29jZCIsIm5iZiI6MTU3ODI5MTI5MSwic3ViIjoiYWRtaW4ifQ.O_WQAZ5R6Jdca3uZji6LVrmYY461feHGwRmhvDo0uUI")
	//request header setting; authorization is required(to get in argocd cluster)
	req.Header.Add("Authorization", "Bearer " +cluster.Token)
	//req.SetBasicAuth("<USERNAME>", "<PASSWORD>")

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
	err = ioutil.WriteFile("/Users/mf839-027/Documents/appsync/githubsync/apphandle/syncresp.yaml", yamlbytes, 0)
	if err != nil {
		panic(err)
	}
}
 */

func Create(cluster *argocdinfo.ArgoCDinfo){
	url:=fmt.Sprintf("http://%s/api/v1/applications", cluster.IPport)
	//json data body

	readbytes, err := ioutil.ReadFile("/Users/mf839-027/Documents/appsync/githubsync/apphandle/basicreq.yml")
	if err != nil {
		panic(err)
	}

	jsonbytes, err := yaml.YAMLToJSON(readbytes)
	if err != nil {
		panic(err)
		return
	}

	//jsonData := map[string] string{"name": app} //this data should be written with another config
	//jsonValue,_:=json.Marshal(jsonData)

	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonbytes))
	if err != nil {
		panic(err)// handle err
	}

	//req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NzgyOTEyOTEsImlzcyI6ImFyZ29jZCIsIm5iZiI6MTU3ODI5MTI5MSwic3ViIjoiYWRtaW4ifQ.O_WQAZ5R6Jdca3uZji6LVrmYY461feHGwRmhvDo0uUI")
	//request header setting; authorization is required(to get in argocd cluster)
	req.Header.Add("Authorization", "Bearer " +cluster.Token)
	//req.SetBasicAuth("<USERNAME>", "<PASSWORD>")
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
	respbytes, _ := ioutil.ReadAll(resp.Body)

	//convert bytes from json to yaml
	yamlbytes, err := yaml.JSONToYAML(respbytes)
	if err != nil {
		panic(err)
		return
	}
	//file write
	err = ioutil.WriteFile("/Users/mf839-027/Documents/appsync/githubsync/apphandle/createcresp.yaml", yamlbytes, 0)
	if err != nil {
		panic(err)
	}
}

func Delete(cluster *argocdinfo.ArgoCDinfo, app string){
	url:=fmt.Sprintf("http://%s/api/v1/applications/%s", cluster.IPport, app)
	//json data body
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)// handle err
	}

	//request header setting; authorization is required(to get in argocd cluster)
	req.Header.Add("Authorization", "Bearer " +cluster.Token)

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
	//deletion
}