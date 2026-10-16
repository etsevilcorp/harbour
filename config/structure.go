package config

import "fmt"

type info struct {
	Name string `json:"name" toml:"app"`
	Icon string `json:"icon" toml:"icon"`
	Type string `toml:"type"`
}

type rawItem struct {
	Type     string    `toml:"type"`
	Name     string    `toml:"name"`
	Icon     string    `toml:"icon"`
	BundleID string    `toml:"bundle_id"`
	Path     string    `toml:"path"`
	Args     []string  `toml:"args"`
	Run      string    `toml:"run"`
	Items    []rawItem `toml:"items"`
}

type Item interface {
	Run() error
	Info() info
}

func buildItem(r rawItem) (Item, error) {
	base := info{Name: r.Name, Icon: r.Icon, Type: r.Type}

	switch r.Type {
	case "separator":
		return &separator{}, nil
	case "folder":
		children, err := buildItems(r.Items)
		if err != nil {
			return nil, err
		}
		return &folder{info: base, Children: children}, nil
	case "command":
		if r.Run == "" {
			return nil, fmt.Errorf("command %q has no run", r.Name)
		}
		return &command{info: base, Command: r.Run, Args: r.Args}, nil
	case "", "app":
		return &app{info: base, BundleID: r.BundleID, Path: r.Path, Args: r.Args}, nil
	default:
		return nil, fmt.Errorf("unknown item type %q", r.Type)
	}
}

func buildItems(raws []rawItem) ([]Item, error) {
	items := make([]Item, 0, len(raws))
	for i, r := range raws {
		it, err := buildItem(r)
		if err != nil {
			return nil, fmt.Errorf("items[%d]: %w", i, err)
		}
		items = append(items, it)
	}
	return items, nil
}
