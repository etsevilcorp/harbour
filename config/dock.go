package config

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
	Items         []Item   `toml:"items"`
}
