package vessel

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/uptrace/bun"
)

type Service struct {
	db *bun.DB
}

/**
 * static private functions
 */

func toProfile(model ProfileModel) Profile {
	dto := Profile{
		ID:        model.ID,
		Name:      model.Name,
		Year:      model.Year,
		HIN:       model.HIN,
		CallSign:  model.CallSign,
		HomePort:  model.HomePort,
		ImageURI:  model.ImageURI,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		UserID:    model.UserID,
	}

	if model.Specification != nil {
		dto.Specification = toSpecification(*model.Specification)
	}

	return dto
}

func toSpecification(model SpecificationModel) Specification {
	dto := Specification{
		Builder:      model.Builder,
		Construction: model.Construction,
		CreatedAt:    model.CreatedAt,
		Designer:     model.Designer,
		Dimensions: Dimensions{
			BeamMM:         model.DimensionsBeamMM,
			DisplacementKg: model.DimensionsDisplacementKg,
			DraftMM:        model.DimensionsDraftMM,
			LOAMM:          model.DimensionsLOAMM,
			LWLMM:          model.DimensionsLWLMM,
		},
		FirstBuilt: model.FirstBuilt,
		HullType:   model.HullType,
		ID:         model.ID,
		LastBuilt:  model.LastBuilt,
		Make:       model.Make,
		Model:      model.Model,
		UpdatedAt:  model.UpdatedAt,
	}

	if model.EngineCooling != nil ||
		model.EngineDrive != nil ||
		model.EngineFuelType != nil ||
		model.EngineGearbox != nil ||
		model.EngineHorsepowerHP != nil ||
		model.EngineMake != nil ||
		model.EngineModel != nil ||
		model.EnginePropeller != nil {
		dto.Engine = &Engine{
			Cooling:      model.EngineCooling,
			Drive:        model.EngineDrive,
			FuelType:     model.EngineFuelType,
			Gearbox:      model.EngineGearbox,
			HorsepowerHP: model.EngineHorsepowerHP,
			Make:         model.EngineMake,
			Model:        model.EngineModel,
			Propeller:    model.EnginePropeller,
		}
	}

	if model.RiggingEMM != nil ||
		model.RiggingForestayLengthMM != nil ||
		model.RiggingIMM != nil ||
		model.RiggingJMM != nil ||
		model.RiggingPMM != nil ||
		model.RiggingType != nil {
		dto.Rigging = &Rigging{
			EMM:              model.RiggingEMM,
			ForestayLengthMM: model.RiggingForestayLengthMM,
			IMM:              model.RiggingIMM,
			JMM:              model.RiggingJMM,
			PMM:              model.RiggingPMM,
			Type:             model.RiggingType,
		}
	}

	if model.SailAreaForeCM2 != nil ||
		model.SailAreaMainCM2 != nil ||
		model.SailAreaReportedCM2 != nil {
		dto.SailArea = &SailArea{
			ForeCM2:     model.SailAreaForeCM2,
			MainCM2:     model.SailAreaMainCM2,
			ReportedCM2: model.SailAreaReportedCM2,
		}
	}

	if model.TanksFuelCapacityL != nil ||
		model.TanksWaterCapacityL != nil {
		dto.Tanks = &Tanks{
			FuelCapacityL:  model.TanksFuelCapacityL,
			WaterCapacityL: model.TanksWaterCapacityL,
		}
	}

	return dto
}

/**
 * constructor functions
 */

func NewService(db *bun.DB) *Service {
	return &Service{
		db: db,
	}
}

/**
 * public functions
 */

func (s *Service) FindAll(ctx context.Context, limit int64, offset int64) ([]Specification, int64, error) {
	var models []SpecificationModel

	results := make([]Specification, 0)
	total, err := s.db.NewSelect().
		Model(&models).
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		ScanAndCount(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return []Specification{}, 0, nil
	}

	if err != nil {
		slog.Error("failed to find boat specifications", "error", err)

		return []Specification{}, 0, err
	}

	for _, model := range models {
		results = append(results, toSpecification(model))
	}

	return results, total, nil
}
