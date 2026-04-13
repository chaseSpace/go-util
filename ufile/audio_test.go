package ufile

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMp3AudioSeconds(t *testing.T) {
	b, _ := os.ReadFile("../testdata/14kb.oga")
	t.Logf("len: %d", len(b))
	t.Log(ParseMp3AudioSeconds(b))
}

func TestParseAudioDurationByFFprobe(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		wantDur  float64
	}{
		{
			name:     "mp3 file",
			fileName: "../testdata/173kb.mp3",
			wantDur:  5.5,
		},
		{
			name:     "oga file",
			fileName: "../testdata/14kb.oga",
			wantDur:  3.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAudioDurationByFFprobe(tt.fileName)
			if err != nil {
				t.Fatalf("ParseAudioDurationByFFprobe() error = %v", err)
			}
			assert.InDelta(t, tt.wantDur, got, 0.1, "duration should be close enough")
		})
	}
}
