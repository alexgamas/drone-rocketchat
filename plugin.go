package main

import "./rocketchat"

type (
	Repo struct {
		Owner   string
		Name    string
		Link    string
		Avatar  string
		Branch  string
		Private bool
		Trusted bool
	}

	Build struct {
		Number   int
		Event    string
		Status   string
		Deploy   string
		Created  int64
		Started  int64
		Finished int64
		Link     string
	}

	Commit struct {
		Remote  string
		Sha     string
		Ref     string
		Link    string
		Branch  string
		Message string
		Author  Author
	}

	Author struct {
		Name   string
		Email  string
		Avatar string
	}

	Config struct {
		// plugin-specific parameters and secrets
		Channel   string
		Text      string
		Username  string
		Password  string
		Url       string
		UserId    string
		AuthToken string
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Commit Commit
		Config Config
	}
)

func (p Plugin) Exec() error {

	client := rocketchat.New(p.Config.Url, p.Config.UserId, p.Config.AuthToken)

	if p.Config.Username != "" {
		req := &rocketchat.LoginRequest{p.Config.Username, p.Config.Password}
		err := client.Login(req)
		if err != nil {
			return err
		}
	}
	return client.ChatPostMessage(p.Config.Text, p.Config.Channel)
}
