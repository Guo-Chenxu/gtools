package static

import (
	"embed"
)

//go:embed jwt.html
var StaticFS embed.FS
