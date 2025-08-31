package postgres

import (
	"strconv"

	"github.com/ikshavaku/catalogue/store"
)

func (r *Service) TODAO() store.ServiceDAO {
	return store.ServiceDAO{
		ID:          strconv.Itoa(int(r.ID)),
		Name:        r.Name,
		Description: r.Description.String,
		CreatedAt:   r.CreatedAt.Time.String(),
		UpdatedAt:   r.UpdatedAt.Time.String(),
		IsDeleted:   r.IsDeleted,
	}
}

func (r *ServiceVersion) TODAO() store.ServiceVersionDAO {
	return store.ServiceVersionDAO{
		ServiceID:     strconv.Itoa(int(r.ServiceID)),
		VersionNumber: r.VersionNumber,
		CreatedAt:     r.CreatedAt.Time.String(),
		UpdatedAt:     r.UpdatedAt.Time.String(),
		IsDeleted:     r.IsDeleted,
	}
}
