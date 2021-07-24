package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	v, err := NewValidator()
	if err != nil {
		log.Panic(err)
	}

	c, err := NewConfig(v)
	if err != nil {
		log.Panicln(err)
	}

	descriptionParts := []DescriptionPart{
		{
			Name:  "Repository",
			Value: fmt.Sprintf("[%s](%s)", c.GitHubRepository, c.GetRepositoryUrl()),
		},
	}

	descriptionParts = append(descriptionParts, DescriptionPart{
		Name:  "Workflow",
		Value: c.GitHubWorkflow,
	})

	if c.GitHubJobName != "" {
		descriptionParts = append(descriptionParts, DescriptionPart{
			Name:  "Job name",
			Value: c.GitHubJobName,
		})
	}

	if c.GitHubRef != "" {
		descriptionParts = append(descriptionParts, DescriptionPart{
			Name:  "Ref",
			Value: c.GitHubRef,
		})
	}

	if c.GitHubActor != "" {
		descriptionParts = append(descriptionParts, DescriptionPart{
			Name:  "Author",
			Value: c.GitHubActor,
		})
	}

	descriptionParts = append(descriptionParts, DescriptionPart{
		Name:  "Commit",
		Value: fmt.Sprintf("[%s](%s)", c.GitHubSha, c.GetCommitUrl()),
	})

	embed := map[string]string{
		"description": GetDescription(descriptionParts),
	}

	runUrl := c.GetRunUrl()

	if runUrl != "" {
		embed["url"] = runUrl
	}

	if c.GitHubJobStatus == "success" {
		embed["title"] = "Action is successful."
	} else if c.GitHubJobStatus == "failure" {
		embed["title"] = "Action has failed."
	} else if c.GitHubJobStatus == "cancelled" {
		embed["title"] = "Action is cancelled."
	}

	payload := map[string]interface{}{
		"embeds": []map[string]string{embed},
	}

	payloadByte, err := json.Marshal(payload)

	req, err := http.NewRequest(http.MethodPost, c.DiscordWebhookUrl, bytes.NewBuffer(payloadByte))
	if err != nil {
		log.Panicln(err)
	}

	req.Header.Set("content-type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Panicln(err)
	}
}
