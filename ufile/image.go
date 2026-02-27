package ufile

import (
	"path"

	"github.com/samber/lo"
)

var validImageExt = []string{
	".jpg",
	".jpeg",
	".png",
	".gif",
	".bmp",
	".webp",
	".svg",
}

func IsValidImageExt(name string) bool {
	return lo.Contains(validImageExt, path.Ext(name))
}
