package solder

// ApiInfo provides information about the software running the
// Solder instance.
//
// See https://docs.solder.io/docs/api
type ApiInfo struct {
	Api string `json:"api"`
	Version string `json:"version"`
	Stream string `json:"stream"`
}
