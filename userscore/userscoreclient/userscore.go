// Code generated by goctl. DO NOT EDIT.
// Source: userscore.proto

package userscoreclient

import (
	"context"

	"go-zero/mall/userscore/types/userscore"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request           = userscore.Request
	Response          = userscore.Response
	UserScoreRequest  = userscore.UserScoreRequest
	UserScoreResponse = userscore.UserScoreResponse

	Userscore interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		SaveUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error)
	}

	defaultUserscore struct {
		cli zrpc.Client
	}
)

func NewUserscore(cli zrpc.Client) Userscore {
	return &defaultUserscore{
		cli: cli,
	}
}

func (m *defaultUserscore) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := userscore.NewUserscoreClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultUserscore) SaveUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error) {
	client := userscore.NewUserscoreClient(m.cli.Conn())
	return client.SaveUserScore(ctx, in, opts...)
}
