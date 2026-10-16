package config

import "github.com/pelletier/go-toml/v2"

type Position string

const (
	Left   Position = "left"
	Right  Position = "right"
	Top    Position = "top"
	Bottom Position = "bottom"
)

type Dock struct {
	Position      Position `toml:"position" json:"position"`
	Size          int      `toml:"size" json:"size"`
	Magnification float64  `toml:"magnification" json:"magnification"`
	Autohide      bool     `toml:"autohide" json:"autohide"`
	Monitor       string   `toml:"monitor" json:"monitor"`
	Items         []Item   `toml:"-" json:"-"`
}

func (d *Dock) Decode(data []byte) error {
	if err := toml.Unmarshal(data, d); err != nil {
		return err
	}

	var raw struct {
		Items []rawItem `toml:"items"`
	}
	if err := toml.Unmarshal(data, &raw); err != nil {
		return err
	}

	items, err := buildItems(raw.Items)
	if err != nil {
		return err
	}
	d.Items = items
	return nil
}
