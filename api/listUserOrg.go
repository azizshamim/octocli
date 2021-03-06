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
        email
        location
        description
      }
    }
  }
}`

type UserOrgsResponseHandler struct {
  client.BasicGraphQLResponseHandler
}

func (this *UserOrgsResponseHandler) TableHeader() []string {
  return []string{"name", "email", "location", "description"}
}

func (this *UserOrgsResponseHandler) TableRows(jsonObj map[string]interface{}) [][]string {
  table := [][]string{}
  nodes := this.GetNodes(jsonObj, this.ResultPath())

  if nodes != nil {
    for _, org := range nodes {
      row := []string{
        fmt.Sprintf("%v", org.(map[string]interface{})["name"]),
        fmt.Sprintf("%v", org.(map[string]interface{})["email"]),
        fmt.Sprintf("%v", org.(map[string]interface{})["location"]),
        fmt.Sprintf("%v", org.(map[string]interface{})["description"])}
      table = append(table, row)
    }
    return table
  } else {
    return nil
  }
}

func (this *UserOrgsResponseHandler) ResultPath() []string {
  return []string{"data", "user", "organizations"}
}

func PrintUserOrgs(user string) {
  params := map[string]interface{}{"login" : user}
  orgHandler := UserOrgsResponseHandler{}
  client.GraphQLQueryAndPrintTable(userOrgsQuery, params, &orgHandler)
}
