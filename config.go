package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	DiscordWebhookUrl string `env:"DISCORD_WEBHOOK_URL" validate:"required,url"`
	GitHubActor       string `env:"GITHUB_ACTOR"`
	GitHubJobName     string `env:"GITHUB_JOB_NAME"`
	GitHubJobStatus   string `env:"GITHUB_JOB_STATUS" validate:"required"`
	GitHubRef         string `env:"GITHUB_REF"`
	GitHubRepository  string `env:"GITHUB_REPOSITORY" validate:"required"`
	GitHubRunId       string `env:"GITHUB_RUN_ID"`
	GitHubServerUrl   string `env:"GITHUB_SERVER_URL" validate:"required"`
	GitHubSha         string `env:"GITHUB_SHA" validate:"required"`
	GitHubWorkflow    string `env:"GITHUB_WORKFLOW" validate:"required"`
}

func NewConfig(v *Validator) (*Config, error) {
	c := Config{}

	v, err := NewValidator()
	if err != nil {
		return nil, err
	}

	t := reflect.TypeOf(c)
	vp := reflect.ValueOf(&c)

	for i := 0; i < t.NumField(); i++ {
		ti := t.Field(i)
		vpi := vp.Elem().Field(i)

		envVarKey := ti.Tag.Get("env")

		if envVarKey != "" {
			envVarValue := os.Getenv(envVarKey)

			vpi.SetString(envVarValue)
		}
	}

	err = v.Validate.Struct(c)
	if err != nil {
		output := "There are errors with some environment variables:\n"

		for fieldName, fieldMessage := range v.Map(err) {
			structField, _ := t.FieldByName(fieldName)

			output = output + fmt.Sprintf("* %s: %s\n", structField.Tag.Get("env"), fieldMessage)
		}

		return nil, errors.New(output)
	}

	return &c, nil
}

func (c *Config) GetRepositoryUrl() string {
	return c.GitHubServerUrl + "/" + c.GitHubRepository
}

func (c *Config) GetRefUrl() string {
	if strings.HasPrefix(c.GitHubRef, "refs/heads") {
		s := strings.Split(c.GitHubRef, "/")
		branchName := s[len(s)-1]

		return c.GetRepositoryUrl() + "/tree/" + branchName
	}

	if strings.HasPrefix(c.GitHubRef, "refs/tags") {
		s := strings.Split(c.GitHubRef, "/")
		tagName := s[len(s)-1]

		return c.GetRepositoryUrl() + "/releases/tags/" + tagName
	}

	return ""
}

func (c *Config) GetCommitUrl() string {
	return c.GetRepositoryUrl() + "/commit/" + c.GitHubSha
}

func (c *Config) GetRunUrl() string {
	return c.GetCommitUrl() + "/checks"
}
