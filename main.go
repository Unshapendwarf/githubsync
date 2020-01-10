package main

import (
	"fmt"
	""
)

func main() {
	// 0. set variable
	var cluster ArgoCDinfo

	// 1. argocd cluster,
	cluster.iport = "192.168.48.12:31410" //argocd cluster
	cluster.username = "admin"
	cluster.password = "1222"
	//cluster.token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NzgyOTEyOTEsImlzcyI6ImFyZ29jZCIsIm5iZiI6MTU3ODI5MTI5MSwic3ViIjoiYWRtaW4ifQ.O_WQAZ5R6Jdca3uZji6LVrmYY461feHGwRmhvDo0uUI"
	CheckAPI("token", &cluster)
	CheckAPI("apps", &cluster)
}
