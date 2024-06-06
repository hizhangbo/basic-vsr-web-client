package data

import (
	"basic-vsr-web-client/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type basicVSRRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.BasicVSRRepo {
	return &basicVSRRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/basicVSRRepo")),
	}
}

func (r *basicVSRRepo) Save(ctx context.Context, vsr *biz.BasicVSR) (*biz.BasicVSR, error) {
	return vsr, nil
}
