package slackAPI

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"os"
	"testing"

	"err"
	"err/errType"
	"slackAPI/client"

	. "github.com/r7kamura/gospel"
)

func TestLookupByEmail(t *testing.T) {
	Describe(t, "LookupByEmailの動作テスト", func() {
		os.Setenv("SLACK_TOKEN", "XXXX-XXXX-XXXX-XXXX-XXXX")
		client.Get = lookupByEmailMock
		api := New("users.lookupByEmail")

		Context("メールアドレスがある場合", func() {
			It("uzuki@pepabo.comを渡したら、trueを返す", func() {
				Expect(api.LookupByEmail("uzuki@pepabo.com")).To(Equal, true)
			})
			It("rin@pepabo.comを渡したら、trueを返す", func() {
				Expect(api.LookupByEmail("rin@pepabo.com")).To(Equal, true)
			})
			It("mio@pepabo.comを渡したら、trueを返す", func() {
				Expect(api.LookupByEmail("mio@pepabo.com")).To(Equal, true)
			})
		})

		Context("メールアドレスがない場合", func() {
			It("kaede@pepabo.comを渡したら、falseを返す", func() {
				Expect(api.LookupByEmail("kaede@pepabo.com")).To(Equal, false)
			})
			It("miyu@pepabo.comを渡したら、falseを返す", func() {
				Expect(api.LookupByEmail("miyu@pepabo.com")).To(Equal, false)
			})
		})
	})
}
jj
func lookupByEmailMock(uv *url.Values, ul string, res interface{}) {
	emailList := []string{"uzuki@pepabo.com", "rin@pepabo.com", "mio@pepabo.com"}
	email := checkEmail((*uv)["email"][0], emailList)

	body, e := ioutil.ReadFile("../test/lookupByEmail/" + email + ".json")
	err.Found(e, errType.ReadFile)

	e = json.Unmarshal(body, &res)
	err.Found(e, errType.ParseJSON)
}

func checkEmail(email string, emailList []string) string {
	for _, e := range emailList {
		if email == e {
			return email
		}
	}
	return "none@papabo.com"
}
