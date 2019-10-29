package slackAPI

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"err"
	"err/errType"
)

type Slack struct {
	baseUrl  string
	token    string
	endPoint string
	limit    string
}

type APIResponse struct {
	Ok         bool            `json:"ok"`
	RawMembers json.RawMessage `json:"members"`
	Metadata   struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

type User struct {
	Id                string `json:"id"`
	TeamId            string `json:team_id`
	Name              string `json:"name"`
	Deleted           bool   `json:"deleted"`
	Color             string `json:"color"`
	RealName          string `json:"real_name"`
	Tz                string `json:"tz"`
	TzLabel           string `json:"tz_label"`
	TzOffset          int    `json:"tz_offset"`
	Profile           *ProfileInfo
	IsAdmin           bool `json:"is_admin"`
	IsOwner           bool `json:"is_owner"`
	IsPrimaryOwner    bool `json:"is_primary_owner"`
	IsRestricted      bool `json:"is_restricted"`
	IsUltraRestricted bool `json:"is_ultra_restricted"`
	IsBot             bool `json:"is_bot"`
	IsAppUser         bool `json:"is_app_user"`
	Updated           int  `json:"updated"`
}

type ProfileInfo struct {
	Title                 string `json:"title`
	Phone                 string `json:"phone"`
	Skype                 string `json:"skype"`
	RealName              string `json:"real_name"`
	RealNameNormalized    string `json:"real_name_normalized"`
	DisplayName           string `json:"display_name"`
	DisplayNameNormalized string `json:"display_name_normalized"`
	StatusText            string `json:"status_text"`
	StatusEmoji           string `json:"status_emoji"`
	StatusExpiration      int    `json:"status_expiration"`
	AvatarHash            string `json:"avatar_hash"`
	ImageOriginal         string `json:"image_original"`
	Email                 string `json:"email"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Image24               string `json:"image_24"`
	Image32               string `json:"image_32"`
	Image48               string `json:"image_48"`
	Image72               string `json:"image_72"`
	Image192              string `json:"image_192"`
	Image512              string `json:"image_512"`
	Image1024             string `json:"image_1024"`
	StatusTextCanonical   string `json:"status_text_canonical"`
	Team                  string `json:"team"`
	IsCustomImage         bool   `json:"jis_custom_image"`
}

const (
	baseUrl = "https://slack.com/api/"
	limit   = "500"
)

func New(endPoint string) *Slack {
	return &Slack{
		baseUrl:  baseUrl,
		token:    tryGetEnv("SLACK_TOKEN"),
		endPoint: endPoint,
		limit:    limit,
	}
}

func tryGetEnv(env string) (v string) {
	v, ok := os.LookupEnv(env)
	if !ok {
		e := fmt.Errorf("Missing environment variable [%s]", env)
		err.Found(e, errType.GetEnv)
	}
	return
}

func (res *APIResponse) getMembers() (u []User) {
	e := json.Unmarshal(res.RawMembers, &u)
	err.Found(e, errType.JSONParse)
	return
}

func (sl *Slack) createURLValues() *url.Values {
	uv := url.Values{}
	uv.Add("token", sl.token)
	return &uv
}
