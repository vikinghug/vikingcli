package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("usage: vikingcli command [arguments]")
		os.Exit(2)
	}

	if args[0] == "clone" {
		resp, err := http.Get("https://api.github.com/orgs/vikinghug/repos")
		if err != nil {
			fmt.Println("failed to fetch repos from github:", err)
			os.Exit(3)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("failed to read response from github:", err)
			os.Exit(3)
		}

		resp.Body.Close()

		type Repo struct {
			CloneUrl string `json:"clone_url"`
		}

		var repos []Repo
		err = json.Unmarshal(body, &repos)
		if err != nil {
			fmt.Println("failed to read response from github:", err)
			os.Exit(3)
		}

		for _, repo := range repos {
			fmt.Printf("cloning repository %s\n", repo.CloneUrl)
			cmd := exec.Command("git", "clone", repo.CloneUrl)
			if err := cmd.Run(); err != nil {
				fmt.Printf("failed to clone repository %s: %v", repo.CloneUrl, err)
				os.Exit(3)
			}
		}
	}
}
