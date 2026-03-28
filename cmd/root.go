package cmd

import (
	"fmt"
	"github-activity/api"
	"os"
)

func Execute() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a github username as an argument")
		return
	}

	username := os.Args[1]

	fmt.Println("Fetching repositories for user: ", username)

	events, err := api.FetchUserActivity(username)
	if err != nil {
		fmt.Println("Error fetching user activity: ", err)
		return
	}

	for _, event := range events {
		fmt.Println("--------------------------------------------")
		fmt.Println("Type : ", event.Type)
		fmt.Println("Repo :", event.Repo.Name)
		fmt.Println("Commits : ", event.Payload.Commits)
		fmt.Println("--------------------------------------------")
	}
}
