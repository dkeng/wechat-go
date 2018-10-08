package material

type NewsRequest struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     bool   `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
}

// NewsReply 图文素材回复
type NewsReply struct {
	MediaID string `json:"media_id"`
}

// UploadImgReply 上传图片回复
type UploadImgReply struct {
	URL string `json:"url"`
}

// UploadFileReply 上传文件回复
type UploadFileReply struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"` // 新增的图片素材的图片URL（仅新增图片素材时会返回该字段）
}
