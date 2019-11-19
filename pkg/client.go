package pkg

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type Client struct {
	TimeSyncClient
}

func (c *Client) Sync() {
	t1 := time.Now().UTC().UnixNano()
	serverTime, err := c.GetServerTime(context.Background(), &GetServerTimeParams{})
	t2 := time.Now().UTC().UnixNano()

	if err != nil {
		logrus.Fatal(err)
	}
	syncedTime := serverTime.Ts + (t2-t1)/2
	logrus.Infof("Server time:  %v", serverTime.Ts)
	logrus.Infof("RTT:          %v", time.Duration(t2-t1))
	logrus.Infof("Sync'ed time: %v", syncedTime)
}

func NewClient(tsc TimeSyncClient) *Client {
	return &Client{
		TimeSyncClient: tsc,
	}
}
