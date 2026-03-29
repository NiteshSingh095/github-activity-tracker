package utils

import (
	"fmt"
	"github-activity/models"
)

func FormatEvents(event models.Event) string {

	switch event.Type {
		case "PushEvent":
			commits := len(event.Payload.Commits)
			if commits == 0 {
				return fmt.Sprintf("Pushed commits to %s", event.Repo.Name)
			}
			return fmt.Sprintf("Pushed %d commits to %s", commits, event.Repo.Name)

		case "IssuesEvent":
			return fmt.Sprintf("Opened a new issue in %s", event.Repo.Name)

		case "PullRequestEvent":
			return fmt.Sprintf("Created a pull request in %s", event.Repo.Name)
			
		case "WatchEvent":
			return fmt.Sprintf("Starred %s", event.Repo.Name)

	    case "ForkEvent":
			return fmt.Sprintf("Forked %s", event.Repo.Name)

		default:
			return fmt.Sprintf("Performed %s on %s", event.Type, event.Repo.Name)
	}
}