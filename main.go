package main

import (
	"githubsync/argocdinfo"
	"githubsync/controller"
)

func main() {
	// 0. set variable
	var cluster argocdinfo.ArgoCDinfo

	// 1. argocd cluster,
	cluster.IPport = "192.168.48.12:31410" //argocd cluster
	cluster.Username = "admin"
	cluster.Password = "1222"
	//cluster.token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NzgyOTEyOTEsImlzcyI6ImFyZ29jZCIsIm5iZiI6MTU3ODI5MTI5MSwic3ViIjoiYWRtaW4ifQ.O_WQAZ5R6Jdca3uZji6LVrmYY461feHGwRmhvDo0uUI"
	controller.CheckAPI("gettoken", &cluster) //done
	controller.CheckAPI("sync", &cluster) // basic sync done
	controller.CheckAPI("create", &cluster)  //deletion

}
