package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

var userOrgsQuery = `query($login:String!, $count:Int!, $cursor:String) {
  user(login: $login) {
    organizations(first: $count, after: $cursor) {
      pageInfo {
       endCursor
       hasNextPage
      }
      nodes {
        name
      }
    }
  }
}`

type UserOrgsResponseHandler struct {
  client.BasicGraphQLResponseHandler
}

func (this *UserOrgsResponseHandler) TableHeader() []string {
  return []string{"name"}
}

func (this *UserOrgsResponseHandler) TableRows(jsonObj map[string]interface{}) [][]string {
  table := [][]string{}
  nodes := this.GetNodes(jsonObj, []string{"data", "user", "organizations"})
  for _, org := range nodes {
    row := []string{fmt.Sprintf("%v", org.(map[string]interface{})["name"])}
    table = append(table, row)
  }
  return table
}

func (this *UserOrgsResponseHandler) PageInfoPath() []string {
  return []string{"data", "user", "organizations"}
}

func GetUserOrgs(server, token, user string) {
  params := map[string]interface{}{"login" : user}
  orgHandler := UserOrgsResponseHandler{}
  client.DoGraphQLApiCall(server, token, userOrgsQuery, params, &orgHandler)
}
