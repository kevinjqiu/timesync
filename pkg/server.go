package pkg

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type timeSyncServer struct {
}

func (s timeSyncServer) GetServerTime(context.Context, *GetServerTimeParams) (*ServerTime, error) {
	log := logrus.WithField("method", "GetServerTime")
	utcNowNano := time.Now().UTC().UnixNano()
	log.Infof("current server time=%v", utcNowNano)
	serverTime:= ServerTime{
		Ts: utcNowNano,
	}
	return &serverTime, nil
}

func NewServer() TimeSyncServer {
	return &timeSyncServer{}
}