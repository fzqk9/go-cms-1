package request

import (
	"github.com/88act/go-cms/server/model/common/request"
	"github.com/88act/go-cms/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}