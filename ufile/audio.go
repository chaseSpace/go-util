package ufile

import (
	"bytes"
	"os/exec"
	"path"
	"strconv"
	"strings"

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

// ParseAudioDurationByFFprobe 使用 ffprobe 命令解析音频文件时长
// file: 音频文件的完整路径
// 返回：时长（秒，float64 类型）和错误信息
func ParseAudioDurationByFFprobe(filePath string) (float64, error) {
	cmd := exec.Command("ffprobe",
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		filePath,
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return 0, err
	}

	durationStr := strings.TrimSpace(out.String())
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, err
	}

	return duration, nil
}
