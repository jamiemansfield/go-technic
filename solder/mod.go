package solder

type Mod struct {
	Name string `json:"name"`
	PrettyName string `json:"pretty_name"`
	Author string `json:"author"`
	Description string `json:"description"`
	Link string `json:"link"`
	Donate string `json:"donate"`
	Versions []string `json:"versions"`
}

type ModVersion struct {
	Md5 string `json:"md5"`
	URL string `json:"url"`
}
