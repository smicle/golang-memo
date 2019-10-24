package slackAPI

import (
	"slackAPI/client"
)

func (sl *Slack) LookupByEmail(email string) bool {
	uv := sl.createURLValues()
	uv.Add("email", email)
	ul := sl.baseUrl + sl.endPoint
	res := new(APIResponse)

	client.Get(uv, ul, res)
	return res.Ok
}
