package boat

import (
	"time"

	"github.com/google/uuid"
)

type Dimensions struct {
	BeamMM         int  `json:"beamMm"`
	DisplacementKg *int `json:"displacementKg"`
	DraftMM        int  `json:"draftMm"`
	LOAMM          int  `json:"loaMm"`
	LWLMM          int  `json:"lwlMn"`
}

type Engine struct {
	Cooling      *string `json:"cooling"`
	Drive        *string `json:"drive"`
	FuelType     *string `json:"fuelType"`
	Gearbox      *string `json:"gearbox"`
	HorsepowerHP *int    `json:"horsepowerHp"`
	Make         *string `json:"make"`
	Model        *string `json:"model"`
	Propeller    *string `json:"propeller"`
}

type Profile struct {
	CreatedAt     time.Time     `json:"createdAt"`
	Specification Specification `json:"specification"`
	CallSign      *string       `json:"callSign,omitempty"`
	HIN           string        `json:"hin"`
	HomePort      *string       `json:"homePort,omitempty"`
	ID            uuid.UUID     `json:"id"`
	ImageURI      *string       `json:"imageUri,omitempty"`
	Name          string        `json:"name"`
	UpdatedAt     time.Time     `json:"updatedAt"`
	Year          int           `json:"year"`
}

type Rigging struct {
	EMM              *int    `json:"eMm"`
	ForestayLengthMM *int    `json:"forestayLengthMm"`
	IMM              *int    `json:"iMm"`
	JMM              *int    `json:"jMm"`
	PMM              *int    `json:"pMm"`
	Type             *string `json:"type"`
}

type SailArea struct {
	ForeCM2     *int `json:"foreCm2"`
	MainCM2     *int `json:"mainCm2"`
	ReportedCM2 *int `json:"reportedCm2"`
}

type Specification struct {
	Builder      *string    `json:"builder"`
	Construction *string    `json:"construction"`
	CreatedAt    time.Time  `json:"-"`
	Designer     *string    `bun:"designer"`
	Dimensions   Dimensions `json:"dimensions"`
	Engine       *Engine    `json:"engine"`
	FirstBuilt   *int       `json:"firstBuilt"`
	HullType     *string    `json:"hullType"`
	ID           uuid.UUID  `json:"id"`
	LastBuilt    *int       `json:"lastBuilt"`
	Make         string     `json:"make"`
	Model        string     `json:"model"`
	Rigging      *Rigging   `json:"rigging"`
	SailArea     *SailArea  `json:"sailArea"`
	Tanks        *Tanks     `json:"tanks"`
	UpdatedAt    time.Time  `json:"-"`
}

type Tanks struct {
	FuelCapacityL  *int `json:"fuelCapacityL"`
	WaterCapacityL *int `json:"waterCapacityL"`
}
