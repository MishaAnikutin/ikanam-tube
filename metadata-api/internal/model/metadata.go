package model

type Metadata struct {
	ID        string
	VideoPath string

	Title       string
	Subtitle    string
	Description string
	Tag         string
	ChannelID   string

	Likes    int
	Dislikes int
}
