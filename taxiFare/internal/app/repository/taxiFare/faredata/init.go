package faredata

import "taxiFare/internal/app/repository/taxiFare"

type taxiFareDataRepository struct{}

func New() taxiFare.Repository {
	return &taxiFareDataRepository{}
}
