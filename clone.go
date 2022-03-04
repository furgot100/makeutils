package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	var repoList []string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many repos?:")
	repos, _ := reader.ReadString('\n')
	repos = strings.TrimSuffix(repos, "\n")

	reposNum, err := strconv.Atoi(repos)

	i := 0

	for i < reposNum {
		fmt.Print("Link of repo: ")
		repo, _ := reader.ReadString('\n')

		repo = strings.TrimSuffix(repo, "\n")

		repoList = append(repoList, repo)

		if err != nil {
			fmt.Println(err)
		}

		i++
	}

	if err != nil {
		fmt.Println(err)
	}

	testrepo := "https://github.com/furgot100/testrepo.git"

	cmd := exec.Command("git", "clone", testrepo)

	errors := cmd.Run()

	fmt.Printf("command finshed with err %v", errors)
}
