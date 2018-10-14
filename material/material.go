package material

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/dkeng/wechat-go"
	"github.com/dkeng/wechat-go/context"
)

// Material 素材
type Material struct {
	context *context.Context
}

// NewMaterial ...
func NewMaterial(c *context.Context) *Material {
	return &Material{
		context: c,
	}
}

// AddNews 新增永久图文素材
func (m *Material) AddNews(reqs []*NewsRequest) (*NewsReply, error) {
	result, err := wechat.PostJSON("https://api.weixin.qq.com/cgi-bin/material/add_news?access_token="+m.context.GetAccessToken(), map[string]interface{}{
		"articles": reqs,
	})
	if err != nil {
		return nil, err
	}
	reply := new(NewsReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// UploadImg 上传图文消息内的图片获取URL
// 本接口所上传的图片不占用公众号的素材库中图片数量的5000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下。
func (m *Material) UploadImg(filename string, srcFile io.Reader) (*UploadImgReply, error) {
	result, err := wechat.Upload("https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token="+m.context.GetAccessToken(), filename, nil, srcFile)
	if err != nil {
		return nil, err
	}
	reply := new(UploadImgReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// UploadFile 上传图文消息内的图片获取URL
// 通过POST表单来调用接口，表单id为media，包含需要上传的素材内容，有filename、filelength、content-type等信息。请注意：图片素材将进入公众平台官网素材管理模块中的默认分组。
func (m *Material) UploadFile(filename, fileType string, description *wechat.VideoDescription, srcFile io.Reader) (*UploadFileReply, error) {
	if fileType == TypeVideo && description == nil {
		return nil, errors.New("请填写视频素材的描述信息")
	}

	result, err := wechat.Upload("https://api.weixin.qq.com/cgi-bin/material/add_material?access_token="+m.context.GetAccessToken()+"&type="+fileType, filename, description, srcFile)
	if err != nil {
		return nil, err
	}
	reply := new(UploadFileReply)
	json.Unmarshal(result, reply)
	return reply, nil
}
