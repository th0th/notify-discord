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

	var descriptionParts []DescriptionPart

	descriptionParts = append(descriptionParts, DescriptionPart{
		Name:  "Repository",
		Value: fmt.Sprintf("[%s](%s)", c.GithubRepository, c.GetRepositoryUrl()),
	})

	if c.GithubAction != "" {
		descriptionParts = append(descriptionParts, DescriptionPart{
			Name:  "Action",
			Value: c.GithubAction,
		})
	}

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
		"avatar_url": "https://user-images.githubusercontent.com/698079/126237897-88c5d9fb-a1d9-4421-955d-152c985726cf.png",
		"embeds":     []map[string]string{embed},
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
