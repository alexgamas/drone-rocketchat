package main

import (
	"drone-rocketchat-plugin/rocketchat"
	"fmt"
	"strings"

	"github.com/drone/drone-template-lib/template"
)

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
		Pull    string
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
		Template  string
		UserId    string
		AuthToken string
		IconURL   string
		IconEmoji string
		ImageURL  string
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Commit Commit
		Config Config
	}
)

func isEmpty(s *string) bool {
	return s == nil || len(strings.TrimSpace(*s)) == 0
}

func (p Plugin) Exec() error {

	client := rocketchat.New(p.Config.Url, p.Config.UserId, p.Config.AuthToken)

	attachment := rocketchat.Attachment{
		Text:     message(p.Repo, p.Build, p.Commit),
		Color:    color(p.Build),
		ImageURL: p.Config.ImageURL,
	}

	payload := rocketchat.WebHookPostPayload{}
	payload.Username = p.Config.Username
	payload.Attachments = []*rocketchat.Attachment{&attachment}
	payload.IconUrl = p.Config.IconURL
	payload.IconEmoji = p.Config.IconEmoji

	if !isEmpty(&p.Config.Template) {

		txt, err := template.RenderTrim(p.Config.Template, p)

		if err != nil {
			return err
		}

		attachment.Text = txt
	}

	if !isEmpty(&p.Config.Username) && !isEmpty(&p.Config.Password) {
		req := &rocketchat.LoginRequest{Username: p.Config.Username, Password: p.Config.Password}
		err := client.Login(req)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("username and password not provided")
	}

	return client.PostMessage(&payload)
}

func message(repo Repo, build Build, commit Commit) string {
	return fmt.Sprintf("*%s* <%s|%s/%s#%s> (%s) by %s",
		build.Status,
		build.Link,
		repo.Owner,
		repo.Name,
		commit.Sha[:8],
		commit.Branch,
		commit.Author,
	)
}

func fallback(repo Repo, build Build, commit Commit) string {
	return fmt.Sprintf("%s %s/%s#%s (%s) by %s",
		build.Status,
		repo.Owner,
		repo.Name,
		commit.Sha[:8],
		commit.Branch,
		commit.Author,
	)
}

func color(build Build) string {
	switch build.Status {
	case "success":
		return "rgb(46, 184, 134)"
	case "failure", "error", "killed":
		return "rgb(163, 2, 0)"
	default:
		return "rgb(218, 160, 56)"
	}
}
