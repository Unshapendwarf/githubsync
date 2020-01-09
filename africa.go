package hongfile

import (
	"fmt"
	"io/ioutil"
)

func filetest() {
	//파일 읽기
	bytes, err := ioutil.ReadFile("/Users/mf839-027/Documents/hong/README.md")
	if err != nil {
		panic(err)
	}

	//파일 출력
	fmt.Print(bytes)

	//파일 쓰기
	err = ioutil.WriteFile("Users/mf839-027/Documents/appsync/githubsync/cicd/TEST.md", bytes, 0)
	if err != nil {
		panic(err)
	}
}