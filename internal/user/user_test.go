package user

import (
	"github.com/sillyhatxu/user-backend/grpc/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	//Initial("localhost:8080")
	Initial("192.168.1.87:8801")
}

func TestAddUser(t *testing.T) {
	err := AddUser(&user.AddRequest{
		LoginName: "sillyhat",
		Password:  "sillyhat",
		Channel:   user.Channel_REGISTER,
		Type:      user.Type_WORD,
	})
	assert.Nil(t, err)
}
