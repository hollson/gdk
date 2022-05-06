package apix

import "math"

type PageResult struct {
	PageInfo
	Data interface{}
}

type PageInfo struct {
	Current int64 `json:"current"` // 当前页
	Prev    int64 `json:"prev"`    // 前一页
	Next    int64 `json:"next"`    // 后一页
	First   int64 `json:"first"`   // 第一页
	Last    int64 `json:"last"`    // 最后一页
	Pages   int64 `json:"pages"`   // 总页数
}

// MakePageInfo 生成分页信息
//  current: 当前页号
//  pageSize: 分页尺寸
//  count: 查询结果总数量
func MakePageInfo(current int64, pageSize int64, count int64) PageInfo {
	res := PageInfo{}
	res.Pages = count / pageSize
	pages := int64(math.Ceil(float64(count) / float64(pageSize)))

	if current > pages {
		current = pages
	}
	if current <= 0 {
		current = 1
	}

	prevPage := current - 1
	if prevPage <= 1 {
		prevPage = 0
	}

	nextPage := current + 1
	if nextPage >= pages {
		nextPage = 0
	}

	res.Pages = pages
	res.Current = current
	res.First = 1
	res.Last = pages
	res.Prev = prevPage
	res.Next = nextPage
	return res
}
