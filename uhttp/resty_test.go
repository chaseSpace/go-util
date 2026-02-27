package uhttp

import (
	"testing"
)

func TestDownloadFile(t *testing.T) {
	// 测试音频文件下载
	url := "https://cdn.pixabay.com/audio/2025/09/17/audio_32aeb1ec12.mp3"

	t.Run("下载MP3音频文件", func(t *testing.T) {
		file, err := DownloadFile(url, "audio/mpeg")

		if err != nil {
			t.Fatalf("下载失败: %v", err)
		}

		if len(file) == 0 {
			t.Error("下载的文件内容为空")
		}

		// 验证是否为有效的MP3文件（检查MP3文件头）
		if len(file) >= 3 {
			// MP3文件通常以ID3标签或帧同步字开始
			// 简单检查文件头是否合理
			if file[0] == 0xFF && (file[1]&0xE0) == 0xE0 {
				t.Log("检测到MP3帧同步字")
			} else if string(file[:3]) == "ID3" {
				t.Log("检测到ID3标签")
			} else {
				t.Logf("文件大小: %d bytes，前几字节: %x", len(file), file[:min(10, len(file))])
			}
		}

		t.Logf("成功下载文件，大小: %d bytes", len(file))
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
