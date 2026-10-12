package main

type App struct {
	Name    string `json:"name"`
	AppIcon []byte `json:"appIcon"`
}

type harbourService struct {
	docks map[string][]App
}

type HarbourService interface {
	Apps(dock string) []App
}

func NewHarbourService() HarbourService {
	return &harbourService{}
}

func (h *harbourService) Apps(dock string) []App {
	return h.docks[dock]
}
