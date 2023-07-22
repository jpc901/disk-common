package response

import "github.com/jpc901/disk-common/model"

type UserLogin struct {
	Token string `json:"token"`
	model.User
}