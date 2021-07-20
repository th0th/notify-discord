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
	GithubActor       string `env:"GITHUB_ACTOR"`
	GithubJobStatus   string `env:"GITHUB_JOB_STATUS" validate:"required"`
	GithubRef         string `env:"GITHUB_REF"`
	GithubRepository  string `env:"GITHUB_REPOSITORY" validate:"required"`
	GithubRunId       string `env:"GITHUB_RUN_ID"`
	GithubServerUrl   string `env:"GITHUB_SERVER_URL" validate:"required"`
	GithubSha         string `env:"GITHUB_SHA" validate:"required"`
	GithubWorkflow    string `env:"GITHUB_WORKFLOW" validate:"required"`
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
	return c.GithubServerUrl + "/" + c.GithubRepository
}

func (c *Config) GetRefUrl() string {
	if strings.HasPrefix(c.GithubRef, "refs/heads") {
		s := strings.Split(c.GithubRef, "/")
		branchName := s[len(s)-1]

		return c.GetRepositoryUrl() + "/tree/" + branchName
	}

	if strings.HasPrefix(c.GithubRef, "refs/tags") {
		s := strings.Split(c.GithubRef, "/")
		tagName := s[len(s)-1]

		return c.GetRepositoryUrl() + "/releases/tags/" + tagName
	}

	return ""
}

func (c *Config) GetCommitUrl() string {
	return c.GetRepositoryUrl() + "/commit/" + c.GithubSha
}

func (c *Config) GetRunUrl() string {
	return c.GetCommitUrl() + "/checks"
}
