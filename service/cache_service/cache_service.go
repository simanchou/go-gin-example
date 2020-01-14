package cache_service

import (
	"strconv"
	"strings"

	"github.com/simanchou/go-gin-example/pkg/e"
)

//Tag tag struct
type Tag struct {
	ID    int
	Name  string
	State int

	PageNum  int
	PageSize int
}

//GetTagsKey get tags key
func (t *Tag) GetTagsKey() string {
	keys := []string{
		e.CACHE_TAG,
		"LIST",
	}

	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State >= 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if t.PageNum > 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}
	if t.PageSize > 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}

	return strings.Join(keys, "_")
}
