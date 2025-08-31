package repositories

import "github.com/ikshavaku/catalogue/store/postgres"

func ProvideServicesRepository(q *postgres.Queries) *SerivcesRepository {
	return &SerivcesRepository{
		Queries: q,
	}
}
