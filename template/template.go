package template

import (
	"github.com/dkeng/wechat-go"
	"github.com/dkeng/wechat-go/context"
)

// Template 消息
type Template struct {
	context *context.Context
}

// NewTemplate ...
func NewTemplate(c *context.Context) *Template {
	return &Template{
		context: c,
	}
}

// Send 发送模板消息
// https://mp.weixin.qq.com/debug/cgi-bin/readtmpl?t=tmplmsg/faq_tmpl
func (t *Template) Send(data *ReplyInfo) error {
	_, err := wechat.PostJSON("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+t.context.GetAccessToken(), data)
	if err != nil {
		return err
	}
	return nil
}
