package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	httpGit "github.com/go-git/go-git/v5/plumbing/transport/http"
)

// var (
// 	ENDPOINT          = "127.0.0.1:9000"
// 	ACCESS_KEY_ID     = "xJ3aePX11FrUovXZ"
// 	SECRET_ACCESS_KEY = "9XiGR94PWmekGPKpmAXBJ7MyowSD055z"
// 	USE_SSL           = false
// )

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var repo *git.Repository
var errPlain error

func main() {
	//ignore error with defer func
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	//check if path already .git  and remote is remoted then continue, else create .git
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		fmt.Println("Creating .git")
		repo, errPlain = git.PlainInit(".", false)
		checkError(errPlain)
	} else {
		fmt.Println("Already have .git")
		repo, errPlain = git.PlainOpen(".")
		checkError(errPlain)
	}

	repo, errPlain = git.PlainOpen(".")
	checkError(errPlain)

	//check if remote is already exist, if not create it
	_, err := repo.Remote("origin")
	if err != nil {
		fmt.Println("Creating remote")
		_, errPlain = repo.CreateRemote(&config.RemoteConfig{
			Name: "origin",
			URLs: []string{"https://github.com/nnowaf/testingtesting.git"},
		})
		checkError(errPlain)
	} else {
		fmt.Println("Already have remote")
	}

	//git add
	w, errPlain := repo.Worktree()
	checkError(errPlain)
	errPlain = w.AddGlob(".")
	checkError(errPlain)

	//git commit
	_, errPlain = w.Commit("initial commit", &git.CommitOptions{All: true})
	checkError(errPlain)

	//git push
	errPlain = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &httpGit.BasicAuth{
			Username: "nnowaf",
			Password: "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",
		},
	})

	fmt.Println("Done")
}

func replaceSpaceUsingRegex(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	//if end of string is space and greater equal 0, remove it
	if len(str) >= 0 && str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}
	//if start of string is space and greater equal 0, remove it
	if len(str) >= 0 && str[0] == ' ' {
		str = str[1:]
	}

	processedString := reg.ReplaceAllString(str, "-")
	fmt.Println(processedString)

	//if start of string have a dash, remove it
	if len(processedString) >= 0 && processedString[0] == '-' {
		processedString = processedString[1:]
		fmt.Println(processedString)
	}
	//if end of string have a dash, remove it
	if len(processedString) >= 0 && processedString[len(processedString)-1] == '-' {
		processedString = processedString[:len(processedString)-1]
		fmt.Println(processedString)
	}

	return processedString
}
