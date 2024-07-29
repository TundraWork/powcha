package handler

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tundrawork/powcha/biz/altcha"
	"github.com/tundrawork/powcha/biz/typ"
)

func Challenge(ctx context.Context, c *app.RequestContext) {
	challenge := altcha.NewChallenge()
	c.JSON(consts.StatusOK, challenge)
}

func Validate(ctx context.Context, c *app.RequestContext) {
	ok, err := altcha.ValidateResponse(string(c.Request.Body()), true)
	if err != nil {
		logger.Errorf("error decoding response: %v", err)
		c.JSON(consts.StatusBadRequest, typ.CommonResponse{
			Code:    consts.StatusBadRequest,
			Message: "invalid payload",
		})
		return
	}
	if ok {
		c.JSON(consts.StatusOK, typ.CommonResponse{
			Code:    consts.StatusOK,
			Message: "OK",
		})
	} else {
		c.JSON(consts.StatusUnauthorized, typ.CommonResponse{
			Code:    consts.StatusUnauthorized,
			Message: "invalid solution",
		})
	}
}
