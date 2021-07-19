package logic

import (
	"context"
	"time"

	"glass/internal/svc"
	"glass/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) PingLogic {
	return PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req types.PingReq) (*types.PongResp, error) {
	return &types.PongResp{T: time.Now().Unix()}, nil
}
