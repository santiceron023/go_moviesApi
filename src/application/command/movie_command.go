package command

type MovieCommand struct {
	User_id int64 `json:"user_id"`
	Title string `json:"title"`
	Length int `json:"length"`
	Synopsis string `json:"synopsis"`
	Image_url string `json:"image_url"`
}
