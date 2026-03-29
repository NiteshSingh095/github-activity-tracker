package cmd

import (
	"fmt"
	"github-activity/api"
	"github-activity/utils"
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

	if len(events) == 0 {
		fmt.Println("No recent activity found")
		return
	}

	for _, event := range events {
		fmt.Println("--------------------------------------------")
		msg := utils.FormatEvents(event)
		fmt.Println(msg)
		fmt.Println("--------------------------------------------")
	}
}
