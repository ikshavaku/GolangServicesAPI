package repositories

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/ikshavaku/catalogue/store"
	"github.com/ikshavaku/catalogue/store/postgres"
	"github.com/jackc/pgx/v5/pgtype"
)

type SerivcesRepository struct {
	Queries *postgres.Queries
}

func (s *SerivcesRepository) ListServices(ctx context.Context, params store.ListServicesParams) (store.PaginatedResponse[store.ServiceDAO], error) {
	// 1. Count total
	name := ""
	if params.Name != nil {
		name = *params.Name
	}
	nameFilter := pgtype.Text{
		String: name,
		Valid:  name != "",
	}
	total, err := s.Queries.CountServices(ctx, nameFilter)
	if err != nil {
		return store.PaginatedResponse[store.ServiceDAO]{}, err
	}
	page := params.Page
	size := params.Size
	// 2. Compute pagination
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * size

	// 3. Get data
	queryParams := postgres.ListServicesParams{
		NameFilter: nameFilter,
		Limit:      int32(size),
		Offset:     int32(offset),
	}
	log.Printf("%+v", queryParams)
	rows, err := s.Queries.ListServices(ctx, queryParams)
	if err != nil {
		return store.PaginatedResponse[store.ServiceDAO]{}, err
	}
	dao := make([]store.ServiceDAO, 0, len(rows))
	for _, r := range rows {
		dao = append(dao, r.TODAO())
	}
	log.Printf("%+v", dao)
	// 4. Wrap response
	return store.PaginatedResponse[store.ServiceDAO]{
		Total: total,
		Page:  page,
		Size:  size,
		Data:  dao,
	}, nil
}

func (s *SerivcesRepository) GetServiceByID(ctx context.Context, params store.GetServiceByIDParams) (store.ServiceDAO, error) {
	id, err := strconv.ParseInt(params.ID, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return store.ServiceDAO{}, err
	}
	res, err := s.Queries.GetServiceByID(ctx, int32(id))
	if err != nil {
		return store.ServiceDAO{}, err
	}
	return res.TODAO(), err
}

func (s *SerivcesRepository) ListServiceVersions(ctx context.Context, params store.ListServiceVersionsByServiceIDParams) ([]store.ServiceVersionDAO, error) {
	id, err := strconv.ParseInt(params.ID, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return []store.ServiceVersionDAO{}, err
	}
	res, err := s.Queries.ListServiceVersionsByServiceID(ctx, int32(id))
	if err != nil {
		return []store.ServiceVersionDAO{}, err
	}
	data := make([]store.ServiceVersionDAO, 0, len(res))
	for _, r := range res {
		data = append(data, r.TODAO())
	}
	return data, err
}
