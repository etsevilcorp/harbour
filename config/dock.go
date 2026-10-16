package config

// position = "left"          # left | right | bottom | top
// size = 48                  # icon size in px
// magnification = 1.4
// autohide = true
// monitor = "primary"

type Position string

const (
	Left   Position = "left"
	Right  Position = "right"
	Top    Position = "top"
	Bottom Position = "bottom"
)

type Dock struct {
	Position      Position `toml:"position"`
	Size          int      `toml:"size"`
	Magnification float64  `toml:"magnification"`
	Autohide      bool     `toml:"autohide"`
	Monitor       string   `toml:"monitor"`
}
