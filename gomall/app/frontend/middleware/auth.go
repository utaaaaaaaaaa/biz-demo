package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	frontendUtils "github.com/utaaaaaaaaaa/biz-demo/gomall/app/frontend/utils"
)

type SessionUserIdKey string

const SessionUserId SessionUserIdKey = "user_id"

// 中间件返回一个函数
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//todo
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		//todo
		s := sessions.Default(ctx)
		userId := s.Get("user_id")
		if userId == nil || "" == userId {
			ctx.Redirect(302, []byte("/sign-in?next="+ctx.FullPath()))
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
