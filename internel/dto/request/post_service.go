package request

type PostRequest struct {
	Content string   `json:"content" validate:"required,min=10,max=100"`
	Title   string   `json:"title" validate:"required,max=20"`
	Tags    []string `json:"tags" validate:"required,dive,required,max=10"`
	UserID  int64    `json:"user_id" validate:"gt=1"`
}
