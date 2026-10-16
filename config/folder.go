package config

type folder struct {
	info
}

func (f *folder) Run() error {
	return nil
}

func (f *folder) Info() info {
	return f.info
}
