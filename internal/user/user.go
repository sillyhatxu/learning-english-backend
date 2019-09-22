package user

import (
	"context"
	"fmt"
	"github.com/sillyhatxu/user-backend/grpc/user"
	"google.golang.org/grpc"
	"time"
)

var internalAddress string

func Initial(address string) {
	internalAddress = address
}

func AddUser(req *user.AddRequest) error {
	conn, err := grpc.Dial(internalAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	userClient := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, err := userClient.Add(ctx, req)
	if err != nil {
		return err
	}
	if response.Code == user.ResponseCode_ERROR {
		return fmt.Errorf("add user error; %s", response.Message)
	}
	return nil
}
