// Package arg 联系人
package arg

// ContactArg 联系人参数
type ContactArg struct {
	PageArg
	Userid int64 `json:"userid" form:"userid"`
	Dstid  int64 `json:"dstid" form:"dstid"`
}
