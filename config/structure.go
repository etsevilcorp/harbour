package config

type info struct {
	Name  string `json:"name" toml:"app"`
	Icon  string `json:"icon" toml:"icon"`
	Items []Item `json:"items"`
	Type  string `toml:"type"`
}

type Config struct {
	info
	Path     string   `toml:"path"`
	BundleID string   `toml:"bundle_id"` // prefered
	Name     string   `toml:"name"`
	Args     []string `toml:"args"`
	Command  string   `toml:"run"` // command-only
	Items    []Config `toml:"items"`
}

type Item interface {
	Run() error
	Info() info
}

func NewItem(config Config) Item {
	switch config.Type {
	case "separator":
		return &separator{}
	case "folder":
		items := make([]Item, len(config.Items))
		for i := range items {
			items[i] = NewItem(config)
		}
		return &folder{
			Name:  config.Name,
			Icon:  config.Icon,
			Type:  config.Type,
			Items: items,
		}
	case "command":
		return &command{
			info:    config.info,
			Command: config.Command,
			Args:    config.Args,
		}
	default:
		return &app{
			info:     config.info,
			BundleID: config.BundleID,
			Path:     config.Path,
			Args:     config.Args,
		}
	}
}
