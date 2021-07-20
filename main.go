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
			Value: fmt.Sprintf("[%s](%s)", c.GithubRepository, c.GetRepositoryUrl()),
		},
	}

	descriptionParts = append(descriptionParts, DescriptionPart{
		Name:  "Workflow",
		Value: c.GithubWorkflow,
	})

	if c.GithubRef != "" {
		descriptionParts = append(descriptionParts, DescriptionPart{
			Name:  "Ref",
			Value: c.GithubRef,
		})
	}

	if c.GithubActor != "" {
		descriptionParts = append(descriptionParts, DescriptionPart{
			Name: "Author",
			Value: c.GithubActor,
		})
	}

	descriptionParts = append(descriptionParts, DescriptionPart{
		Name: "Commit",
		Value: fmt.Sprintf("[%s](%s)", c.GithubSha, c.GetCommitUrl()),
	})

	embed := map[string]string{
		"description": GetDescription(descriptionParts),
	}

	runUrl := c.GetRunUrl()

	if runUrl != "" {
		embed["url"] = runUrl
	}

	if c.GithubJobStatus == "success" {
		embed["title"] = "Action is successful."
	} else if c.GithubJobStatus == "failure" {
		embed["title"] = "Action has failed."
	} else if c.GithubJobStatus == "cancelled" {
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
