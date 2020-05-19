package platform

type Modpack struct {
	ID int `json:"id"`
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	User string `json:"user"`
	PlatformURL string `json:"platformUrl"`
	Ratings int `json:"ratings"`
	Downloads int `json:"downloads"`
	Runs int `json:"runs"`
	Description string `json:"description"`

	// Direct Download
	URL string `json:"url"`
	Minecraft string `json:"minecraft"`
	Version string `json:"version"`

	// Solder
	Solder string `json:"solder"`

	// Images
	Icon *Image `json:"icon"`
	Logo *Image `json:"logo"`
	Background *Image `json:"background"`
}

type Image struct {
	URL string `json:"url"`
	Md5 string `json:"md5"`
}
