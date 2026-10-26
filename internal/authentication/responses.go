package authentication

import "github.com/kieranroneill/oarbit/internal/user"

type StatusResponse struct {
	SetupRequired bool       `json:"setupRequired"`
	User          *user.User `json:"user"`
}
