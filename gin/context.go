package gin

import (
	wechat "github.com/dkeng/wechat-go"
	wxContext "github.com/dkeng/wechat-go/context"
	ggin "github.com/gin-gonic/gin"
)

var (
	// WxContext 微信上下文
	WxContext *wxContext.Context
)

// WechatContext include gin context and local context
type WechatContext struct {
	*ggin.Context
	WxContext *wxContext.Context
}

// WechatContextFunction wechat gin.Contenxt and local model.Context
func WechatContextFunction(ctlFunc func(ctx *WechatContext)) ggin.HandlerFunc {
	return func(ctx *ggin.Context) {

		wechatContext := &WechatContext{
			Context:   ctx,
			WxContext: WxContext,
		}

		ctlFunc(wechatContext)
	}
}

// ReplyXML 回复xml
func (w *WechatContext) ReplyXML(xml wechat.XMLer) {
	w.String(200, xml.XML())
}
