package platform

import "net/http"

// ModpackService provides access to the modpack routes of the
// Technic Platform API.
type ModpackService service

// GetModpack gets information about the given modpack from the Technic
// Platform API.
func (s *ModpackService) GetModpack(slug string) (*Modpack, error) {
	request, err := s.client.NewRequest(http.MethodGet, "modpack/" + slug + "/", nil)
	if err != nil {
		return nil, err
	}

	var response Modpack
	_, err = s.client.Do(request, &response)
	return &response, err
}
