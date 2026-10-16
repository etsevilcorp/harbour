package config

type folder struct {
	info
	Children []Item
}

func (f *folder) Run() error {
	return nil
}

func (f *folder) Info() info {
	return f.info
}
