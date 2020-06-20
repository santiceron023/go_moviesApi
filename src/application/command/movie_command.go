package command

type MovieCommand struct {
	UserId   int64  `json:"user_id"`
	Title    string `json:"title"`
	Length   int    `json:"length"`
	Synopsis string `json:"synopsis"`
	ImageUrl string `json:"image_url"`
}
