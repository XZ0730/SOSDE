// Code generated by goctl. DO NOT EDIT.
// Source: mq.proto

package mqclient

import (
	"context"

	"mq_server/mq"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AttRequest   = mq.AttRequest
	LeaveRequest = mq.LeaveRequest
	Request      = mq.Request
	Response     = mq.Response

	Mq interface {
		Publish(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		PublishLeave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*Response, error)
		PublishPull(ctx context.Context, in *AttRequest, opts ...grpc.CallOption) (*Response, error)
	}

	defaultMq struct {
		cli zrpc.Client
	}
)

func NewMq(cli zrpc.Client) Mq {
	return &defaultMq{
		cli: cli,
	}
}

func (m *defaultMq) Publish(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := mq.NewMqClient(m.cli.Conn())
	return client.Publish(ctx, in, opts...)
}

func (m *defaultMq) PublishLeave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*Response, error) {
	client := mq.NewMqClient(m.cli.Conn())
	return client.PublishLeave(ctx, in, opts...)
}

func (m *defaultMq) PublishPull(ctx context.Context, in *AttRequest, opts ...grpc.CallOption) (*Response, error) {
	client := mq.NewMqClient(m.cli.Conn())
	return client.PublishPull(ctx, in, opts...)
}
