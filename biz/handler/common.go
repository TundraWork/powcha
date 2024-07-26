package handler

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tundrawork/powcha/biz/altcha"
)

func Challenge(ctx context.Context, c *app.RequestContext) {
	challenge := altcha.NewChallenge()
	c.JSON(consts.StatusOK, challenge)
}

func Validate(ctx context.Context, c *app.RequestContext) {
	msg := altcha.Message{}
	if err := c.BindAndValidate(&msg); err != nil {
		c.JSON(consts.StatusBadRequest, err)
		return
	}
	logger.CtxInfof(ctx, "validate request: %v", msg)
	if altcha.ValidateResponse(msg.EncodeWithBase64(), false) {
		c.JSON(consts.StatusOK, "valid")
	} else {
		c.JSON(consts.StatusBadRequest, "invalid")
	}
}
