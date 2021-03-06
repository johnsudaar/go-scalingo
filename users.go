package scalingo

import (
	"encoding/json"

	"gopkg.in/errgo.v1"
)

type User struct {
	ID                  string `json:"id"`
	Username            string `json:"username"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	Email               string `json:"email"`
	AuthenticationToken string `json:"authentication_token"`
}

type SelfResponse struct {
	User *User `json:"user"`
}

func (c *Client) Self() (*User, error) {
	req := &APIRequest{
		Client:   c,
		Endpoint: "/users/self",
	}
	res, err := req.Do()
	if err != nil {
		return nil, errgo.Mask(err, errgo.Any)
	}
	defer res.Body.Close()
	var u *SelfResponse
	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		return nil, errgo.Mask(err, errgo.Any)
	}
	return u.User, nil
}
