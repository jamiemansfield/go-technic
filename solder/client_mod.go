package solder

import "net/http"

// ModService provides access to the mod related routes on a Solder
// instance.
//
// PLEASE NOTE THAT THE TECHNIC PACK'S INSTANCE DOES NOT SUPPORT THESE
// ROUTES.
type ModService service

// GetMod gets information about the given mod, from the Solder instance.
func (s *ModService) GetMod(slug string) (*Mod, error) {
	request, err := s.client.NewRequest(http.MethodGet, "mod/" + slug + "/", nil)
	if err != nil {
		return nil, err
	}

	var response Mod
	_, err = s.client.Do(request, &response)
	return &response, err
}

// GetModVersion gets information about the given mod version, from the
// Solder instance.
func (s *ModService) GetModVersion(slug string, version string) (*ModVersion, error) {
	request, err := s.client.NewRequest(http.MethodGet, "mod/" + slug + "/" + version + "/", nil)
	if err != nil {
		return nil, err
	}

	var response ModVersion
	_, err = s.client.Do(request, &response)
	return &response, err
}
