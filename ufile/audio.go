package ufile

import (
	"bytes"
	"path"

	"github.com/hajimehoshi/go-mp3"
	"github.com/samber/lo"
)

var validAudioExt = []string{
	".mp3",
	".wav",
	".ogg",
	".flac",
	".aac",
}

func IsValidAudioExt(name string) bool {
	return lo.Contains(validAudioExt, path.Ext(name))
}

func ParseMp3AudioSeconds(file []byte) (uint32, error) {
	d, err := mp3.NewDecoder(bytes.NewReader(file))
	if err != nil {
		return 0, err
	}
	sec := float64(d.Length()) / float64(d.SampleRate()*4)
	return uint32(sec), nil
}
