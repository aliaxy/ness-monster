// Package model 信息模型
package model

// Message 信息
type Message struct {
	Id int64 `json:"id,omitempty" form:"id"` // 消息ID
	// 谁发的
	Userid int64 `json:"userid,omitempty" form:"userid"` // 谁发的
	// 什么业务
	Cmd int `json:"cmd,omitempty" form:"cmd"` // 群聊还是私聊
	// 发给谁
	Dstid int64 `json:"dstid,omitempty" form:"dstid"` // 对端用户ID/群ID
	// 怎么展示
	Media int `json:"media,omitempty" form:"media"` // 消息按照什么样式展示
	// 内容是什么
	Content string `json:"content,omitempty" form:"content"` // 消息的内容
	// 图片是什么
	Pic string `json:"pic,omitempty" form:"pic"` // 预览图片
	// 连接是什么
	Url string `json:"url,omitempty" form:"url"` // 服务的URL
	// 简单描述
	Memo string `json:"memo,omitempty" form:"memo"` // 简单描述
	// 其他的附加数据，语音长度/红包金额
	Amount int `json:"amount,omitempty" form:"amount"` // 其他和数字相关的
}

const (
	CmdSingleMsg = 10 // CmdSingleMsg 点对点单聊 dstid 是用户 ID
	CmdRoomMsg   = 11 // CmdRoomMsg 群聊消息 dstid 是群 id
	CmdHeart     = 0  // CmdHeart 心跳消息,不处理
)

const (
	MediaTypeText       = 1   // MediaTypeText 文本样式
	MediaTypeNews       = 2   // MediaTypeNews 新闻样式,类比图文消息
	MediaTypeVoice      = 3   // MediaTypeVoice 语音样式
	MediaTypeImg        = 4   // MediaTypeImg 图片样式
	MediaTypeRedPackage = 5   // MediaTypeRedPackage 红包样式
	MediaTypeEmoji      = 6   // MediaTypeEmoji emoji 表情样式
	MediaTypeLink       = 7   // MediaTypeLink 超链接样式
	MediaTypeVideo      = 8   // MediaTypeVideo 视频样式
	MediaTypeConcat     = 9   // MediaTypeConcat 名片样式
	MediaTypeUndefined  = 100 // MediaTypeUndefined 其他自己定义,前端做相应解析即可
)
