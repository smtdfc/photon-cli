package domain

type Config struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	EntryPoint  string `json:"entryPoint"`
	CoreVersion string `json:"coreVer"`
}
