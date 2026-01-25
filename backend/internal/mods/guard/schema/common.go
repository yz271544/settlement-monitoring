package schema

// PaginationReq 分页请求
type PaginationReq struct {
	Page     int `json:"page,omitempty" form:"page" binding:"min=1,default=1"`
	PageSize int `json:"page_size,omitempty" form:"page_size" binding:"min=1,max=100,default=10"`
}

// QueryOptions 查询选项
type QueryOptions struct {
	Sort  string `json:"sort,omitempty" form:"sort" binding:"omitempty"`
	Order string `json:"order,omitempty" form:"order" binding:"omitempty,oneof=asc desc"`
}

// PageResult 分页结果
type PageResult[T any] struct {
	Total int64 `json:"total"`
	Items []T   `json:"items"`
}

// CommonResponse 通用响应
type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// BatchDeleteReq 批量删除请求
type BatchDeleteReq struct {
	IDs string `json:"ids" validate:"required"`
}

// BoolResult 布尔操作结果
type BoolResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
