package article

import "blog_backend_go/api/v1/dto/common/request"

// 文章列表请求
type ListRequest struct {
	request.PageInfo
	request.SearchInfo
	Sort []request.SortInfo `json:"sort"`
}

// 文章添加请求
type AddRequest struct {
	Title    string `json:"title"`
	Cover    string `json:"cover"`
	Abstract string `json:"abstract"`
	Content  string `json:"content"`
	Tags     []int  `json:"tags"`
}

// 文章更新请求
type UpdateRequest struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Cover    string `json:"cover"`
	Abstract string `json:"abstract"`
	Content  string `json:"content"`
	Tags     []int  `json:"tags"`
}

// 文章删除请求
type DeleteRequest struct {
	ID string `json:"id"`
}
