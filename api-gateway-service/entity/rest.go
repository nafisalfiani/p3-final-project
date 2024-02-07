package entity

type HTTPResp struct {
	Message    HTTPMessage `json:"message"`
	Meta       Meta        `json:"metadata"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type HTTPMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Meta struct {
	Path        string     `json:"path"`
	StatusCode  int        `json:"status_code"`
	Status      string     `json:"status"`
	Message     string     `json:"message"`
	Timestamp   string     `json:"timestamp"`
	Error       *MetaError `json:"error,omitempty"`
	RequestID   string     `json:"request_id"`
	TimeElapsed string     `json:"time_elapsed"`
}

type MetaError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PaginationParam struct {
	Limit   int64    `form:"limit" param:"limit" db:"limit"`
	Page    int64    `form:"page" param:"page" db:"page"`
	SortBy  []string `form:"sort_by" param:"sort_by" db:"sort_by"`
	GroupBy []string `form:"-" param:"-" db:"-"`
}

type Pagination struct {
	CurrentPage     int64    `json:"current_page"`
	CurrentElements int64    `json:"current_elements"`
	TotalPages      int64    `json:"total_pages"`
	TotalElements   int64    `json:"total_elements"`
	SortBy          []string `json:"sort_by"`
	CursorStart     *string  `json:"cursor_start,omitempty"`
	CursorEnd       *string  `json:"cursor_end,omitempty"`
}

func (p *Pagination) ProcessPagination(limit int64) {
	if p.SortBy == nil {
		p.SortBy = []string{}
	}

	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	}

	if limit < 1 {
		limit = 10
	}

	totalPage := p.TotalElements / limit
	if p.TotalElements%limit > 0 || p.TotalElements == 0 {
		totalPage++
	}

	p.TotalPages = 1
	if totalPage > 1 {
		p.TotalPages = totalPage
	}
}

type Ping struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}
