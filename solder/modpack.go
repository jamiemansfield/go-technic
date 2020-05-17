package solder

// Modpacks provides a list of modpacks available on the Solder
// instance, and the mirror (files) URL.
type Modpacks struct {
	// Modpacks is a map of modpack slugs to names. Only public
	// modpacks will be listed here.
	Modpacks map[string]string `json:"modpacks"`
	
	// MirrorURL is the URL to the mirror (files)
	MirrorURL string `json:"mirror_url"`
}

type Modpack struct {
	// The name (really slug) of the pack.
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	URL string `json:"url"`

	Icon string `json:"icon"`
	IconMd5 string `json:"icon_md5"`

	Background string `json:"background"`
	BackgroundMd5 string `json:"background_md5"`

	Recommended string `json:"recommended"`
	Latest string `json:"latest"`
	Builds []string `json:"builds"`
}

type Build struct {
	Minecraft string `json:"minecraft"`
	MinecraftMd5 string `json:"minecraft_md5"`
	Forge string `json:"forge"`
	Mods []*BuildMod `json:"mods"`
}

type BuildMod struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Md5 string `json:"md5"`
	URL string `json:"url"`
}
