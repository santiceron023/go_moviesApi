package command

type MovieCommand struct {
	user_id int64 `json:"user_id"`
	title string `json:"title"`
	duration int `json:"duration"`
	synopsis string `json:"synopsis"`
	image_url string `json:"image_url"`
}
