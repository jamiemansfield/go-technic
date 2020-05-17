package solder

import (
	"errors"
	"net/http"
)

var (
	// TODO: support these
	ErrModpackDoesNotExist = errors.New("modpack does not exist")
	ErrBuildDoesNotExist   = errors.New("build does not exist")
)

// ModpackService provides access to the modpack routes of a
// Solder API.
//
// See https://docs.solder.io/docs/apimodpack
// See https://docs.solder.io/docs/apimodpackslugbuild
type ModpackService service

// GetModpacks gets the list of public modpacks available from the
// Solder instance.
func (s *ModpackService) GetModpacks() (*Modpacks, error) {
	request, err := s.client.NewRequest(http.MethodGet, "modpack/", nil)
	if err != nil {
		return nil, err
	}

	var response Modpacks
	_, err = s.client.Do(request, &response)
	return &response, err
}

// GetModpack gets information about the given modpack from the Solder
// instance.
func (s *ModpackService) GetModpack(slug string) (*Modpack, error) {
	request, err := s.client.NewRequest(http.MethodGet, "modpack/" + slug + "/", nil)
	if err != nil {
		return nil, err
	}

	var response Modpack
	_, err = s.client.Do(request, &response)
	return &response, err
}

// GetBuild gets information about the given modpack's build from the
// Solder instance
func (s *ModpackService) GetBuild(slug string, version string) (*Build, error) {
	request, err := s.client.NewRequest(http.MethodGet, "modpack/" + slug + "/" + version + "/", nil)
	if err != nil {
		return nil, err
	}

	var response Build
	_, err = s.client.Do(request, &response)
	return &response, err
}
