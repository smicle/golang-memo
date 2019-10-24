package slackAPI

import (
	"err"
	"err/errType"
	"slackAPI/client"

	"github.com/pkg/errors"
)

func (sl *Slack) UsersList() (users []User) {
	var cursor string
	for {
		uv := sl.createURLValues()
		uv.Add("limit", sl.limit)
		uv.Add("cursor", cursor)
		ul := sl.baseUrl + sl.endPoint
		res := new(APIResponse)

		client.Get(uv, ul, res)
		if !res.Ok {
			e := errors.New(errType.APIExecution)
			err.Found(e, errType.HTTPResponse)
		}

		members := res.getMembers()
		users = append(users, members...)
		cursor = res.Metadata.NextCursor
		if cursor == "" {
			break
		}
	}
	return
}
