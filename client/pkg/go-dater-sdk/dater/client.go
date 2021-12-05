package dater

import (
	"github.com/1makarov/go-dater/client/pkg/go-dater-sdk/proto"
	"google.golang.org/grpc"
)

type Client struct {
	conn        *grpc.ClientConn
	daterClient proto.DaterClient
}

func New(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := proto.NewDaterClient(conn)

	return &Client{conn: conn, daterClient: c}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
