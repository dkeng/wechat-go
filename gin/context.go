package gin

import (
	wechat "github.com/dkeng/wechat-go"
	ggin "github.com/gin-gonic/gin"
)

// WechatContext include gin context and local context
type WechatContext struct {
	*ggin.Context
}

// WechatContextFunction wechat gin.Contenxt and local model.Context
func WechatContextFunction(ctlFunc func(ctx *WechatContext)) ggin.HandlerFunc {
	return func(ctx *ggin.Context) {

		wechatContext := &WechatContext{
			Context: ctx,
		}

		ctlFunc(wechatContext)
	}
}

// ReplyXML 回复xml
func (w *WechatContext) ReplyXML(xml wechat.XMLer) {
	w.String(200, xml.XML())
}
