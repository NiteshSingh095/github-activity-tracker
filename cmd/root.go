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

	response, err := api.FetchUserActivity(username)
	if err != nil {
		fmt.Println("Error fetching user activity: ", err)
		return
	}

	fmt.Println("User activity: ", string(response))
}
