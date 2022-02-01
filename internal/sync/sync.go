package sync

import (
	"errors"

	"github.com/alnvdl/gomockcmp/external"
)

// Syncer performs multiple calls to sync things to external services.
type Syncer struct {
	external.ServiceClient
}

// SyncWithExternalService triggers the sync by calling methods in an external
// service.
func (s *Syncer) SyncWithExternalService(id string) error {
	obj, err := s.ServiceClient.DoSomething(id)
	if err != nil {
		return errors.New("error in DoSomething")
	}

	obj, err = s.ServiceClient.DoSomethingSlightlyDifferent(obj.ID)
	if err != nil {
		return errors.New("error in DoSomethingSlightlyDifferent")
	}

	return nil
}
