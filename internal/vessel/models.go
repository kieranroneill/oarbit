package vessel

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProfileModel struct {
	bun.BaseModel `bun:"table:vessel_profiles,alias:bp"`

	// primary key
	ID uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`

	// foreign keys
	SpecificationID uuid.UUID `bun:"specification_id,type:uuid,notnull"`

	// Data
	CallSign  *string   `bun:"call_sign"`
	CreatedAt time.Time `bun:"created_at,type:timestamptz,nullzero,notnull,default:current_timestamp"`
	HomePort  *string   `bun:"home_port"`
	HIN       string    `bun:"hin,notnull"`
	ImageURI  *string   `bun:"image_uri"`
	Name      string    `bun:"name,notnull"`
	UpdatedAt time.Time `bun:"updated_at,type:timestamptz,nullzero,notnull,default:current_timestamp"`
	Year      int       `bun:"year,notnull"`

	// relationships
	Specification *SpecificationModel `bun:"rel:belongs-to,join:specification_id=id"`
}

type SpecificationModel struct {
	bun.BaseModel `bun:"table:vessel_specifications,alias:bs"`

	// primary key
	ID uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`

	// data
	Builder                  *string   `bun:"builder"`
	Construction             *string   `bun:"construction"`
	CreatedAt                time.Time `bun:"created_at,type:timestamptz,nullzero,notnull,default:current_timestamp"`
	Designer                 *string   `bun:"designer"`
	DimensionsBallastKg      *int      `bun:"dimensions_ballast_kg"`
	DimensionsBeamMM         int       `bun:"dimensions_beam_mm,notnull"`
	DimensionsDisplacementKg *int      `bun:"dimensions_displacement_kg"`
	DimensionsDraftMM        int       `bun:"dimensions_draft_mm,notnull"`
	DimensionsLOAMM          int       `bun:"dimensions_loa_mm,notnull"`
	DimensionsLWLMM          int       `bun:"dimensions_lwl_mm,notnull"`
	EngineCooling            *string   `bun:"engine_cooling"`
	EngineDrive              *string   `bun:"engine_drive"`
	EngineFuelType           *string   `bun:"engine_fuel_type"`
	EngineGearbox            *string   `bun:"engine_gearbox"`
	EngineHorsepowerHP       *int      `bun:"engine_horsepower_hp"`
	EngineMake               *string   `bun:"engine_make"`
	EngineModel              *string   `bun:"engine_model"`
	EnginePropeller          *string   `bun:"engine_propeller"`
	FirstBuilt               *int      `bun:"first_built"`
	HullType                 *string   `bun:"hull_type"`
	LastBuilt                *int      `bun:"last_built"`
	Make                     string    `bun:"make,notnull"`
	Model                    string    `bun:"model,notnull"`
	RiggingEMM               *int      `bun:"rigging_e_mm"`
	RiggingForestayLengthMM  *int      `bun:"rigging_forestay_length_mm"`
	RiggingIMM               *int      `bun:"rigging_i_mm"`
	RiggingJMM               *int      `bun:"rigging_j_mm"`
	RiggingPMM               *int      `bun:"rigging_p_mm"`
	RiggingType              *string   `bun:"rigging_type"`
	SailAreaForeCM2          *int      `bun:"sail_area_fore_cm2"`
	SailAreaMainCM2          *int      `bun:"sail_area_main_cm2"`
	SailAreaReportedCM2      *int      `bun:"sail_area_reported_cm2"`
	TanksFuelCapacityL       *int      `bun:"tanks_fuel_capacity_l"`
	TanksWaterCapacityL      *int      `bun:"tanks_water_capacity_l"`
	UpdatedAt                time.Time `bun:"updated_at,type:timestamptz,nullzero,notnull,default:current_timestamp"`
}
