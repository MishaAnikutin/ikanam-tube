package repository

import "os"

type Video struct {
	path string
}

func New(video_name string) (Video, error) {
	path := "../../static/" + video_name + ".m8e3"

	_, err := os.Stat(path) // Проверка на существование видео

	video := Video{path: path} // Создаем экземпляр видео

	return video, err
}

func (v *Video) GetChunk(number int) {

}
