package main

import "fmt"

// паттерн «адаптер» на примере плееров

// Адаптируемый интерфейс плеера видеофайлов
type VideoPlayer interface {
	PlayVideo(filename string)
}

// Адаптируемый интерфейс плеера аудиофайлов
type AudioPlayer interface {
	PlayAudio(filename string)
}

// Реализация плеера видеофайлов
type VLCPlayer struct{}

func (v *VLCPlayer) PlayVideo(filename string) {
	fmt.Printf("Playing video: %s using VLC Player\n", filename)
}

// Реализация плеера аудиофайлов
type WinampPlayer struct{}

func (w *WinampPlayer) PlayAudio(filename string) {
	fmt.Printf("Playing audio: %s using Winamp Player\n", filename)
}

// Адаптер для плеера аудиофайлов, преобразующий его интерфейс в интерфейс плеера видеофайлов
type AudioToVideoPlayerAdapter struct {
	audioPlayer AudioPlayer
}

func (a *AudioToVideoPlayerAdapter) PlayVideo(filename string) {
	// Необходимо проигрывать только аудио часть видеофайла
	a.audioPlayer.PlayAudio(filename)
}

// Главная функция программы
func main() {
	vlcPlayer := &VLCPlayer{}
	winampPlayer := &WinampPlayer{}

	videoFile := "movie.mp4"
	audioFile := "song.mp3"

	// Создаем адаптер, присваивая ему экземпляр плеера аудиофайлов
	audioAdapter := &AudioToVideoPlayerAdapter{
		audioPlayer: winampPlayer,
	}

	// Проигрываем видеофайл через VLCPlayer
	vlcPlayer.PlayVideo(videoFile)

	// Проигрываем аудиофайл через адаптер, который внутри использует WinampPlayer
	audioAdapter.PlayVideo(audioFile)
}
