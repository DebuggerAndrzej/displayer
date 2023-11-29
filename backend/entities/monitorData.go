package entities

type MonitorData struct {
	Name         string
	Enabled      bool
	Resolutions  []string
	RefreshRates []Resolution
	PositionX    int
	PositionY    int
	Scale        float32
}

type Resolution struct {
	Height int
	Width  int
}
