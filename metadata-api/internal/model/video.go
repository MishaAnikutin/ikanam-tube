package model

type Video struct {
	ID          string `json:"id"`
	VideoPath   string `json:"video_path"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	ChannelID   string `json:"channel_id"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
}
