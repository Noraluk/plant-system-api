package pumprepo

import (
	"context"
	"fmt"
	pumpmodel "plant-system-api/api/models/pump"
	pumprepoitf "plant-system-api/api/repositories/pump/interface"

	"firebase.google.com/go/v4/db"
)

type pumpRepository struct {
	ctx context.Context
	ref *db.Ref
}

func NewPumpRepository(firebaseClient *db.Client) pumprepoitf.PumpRepository {
	return &pumpRepository{
		ctx: context.Background(),
		ref: firebaseClient.NewRef("pump"),
	}
}

func (repo *pumpRepository) ActivePump(pump *pumpmodel.Pump) error {
	specificPump := repo.ref.Child(fmt.Sprintf("%d", pump.PumpID))
	if err := specificPump.Set(repo.ctx, pump.IsActive); err != nil {
		return fmt.Errorf("error activating pump : %w", err)
	}
	return nil
}
