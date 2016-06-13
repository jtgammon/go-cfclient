package cfclient

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type User struct {
	admin        		bool	`json:"admin"`
	active        		bool	`json:"active"`
	DefaultSpaceGuid 	string	`json:"default_space_guid"`
	Username    		string	`json:"username"`
	OrganizationRoles   	[]string	`json:"organization_roles"`
	c          		*Client
}

type UserResponse struct {
	Count     int           `json:"total_results"`
	Pages     int           `json:"total_pages"`
	NextUrl   string        `json:"next_url"`
	Resources []UserResource `json:"resources"`
}

type UserResource struct {
	Meta   Meta `json:"metadata"`
	Entity User  `json:"entity"`
}

func (c *Client) UsersBySpace(guid string) ([]User, error) {
	var users []User
	var userResponse UserResponse
	r := c.NewRequest("GET", "/v2/organizations/"+guid+"/user_roles?order-direction=asc")

	resp, err := c.DoRequest(r)
	if err != nil {
		return nil, fmt.Errorf("Error requesting users %v", err)
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading users request: %v", err)
	}

	err = json.Unmarshal(resBody, &userResponse)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling users %v", err)
	}
	for _, user := range userResponse.Resources {
		users = append(users, user.Entity)
	}
	return users, nil
}
