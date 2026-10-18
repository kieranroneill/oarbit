-- +goose Up

/* creates "boat_specifications" table */
CREATE TABLE boat_specifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    builder TEXT NULL,
    construction TEXT NULL,
    designer TEXT NULL,
    first_built INTEGER NULL,
    hull_type TEXT NULL,
    last_built INTEGER NULL,
    make TEXT NOT NULL,
    model TEXT NOT NULL,

    dimensions_ballast_kg INTEGER NULL,
    dimensions_beam_mm INTEGER NOT NULL,
    dimensions_displacement_kg INTEGER NULL,
    dimensions_draft_mm INTEGER NOT NULL,
    dimensions_loa_mm INTEGER NOT NULL,
    dimensions_lwl_mm INTEGER NOT NULL,

    engine_cooling TEXT NULL,
    engine_drive TEXT NULL,
    engine_fuel_type TEXT NULL,
    engine_gearbox TEXT NULL,
    engine_horsepower_hp INTEGER NULL,
    engine_make TEXT NULL,
    engine_model TEXT NULL,
    engine_propeller TEXT NULL,

    rigging_e_mm INTEGER NULL,
    rigging_forestay_length_mm INTEGER NULL,
    rigging_i_mm INTEGER NULL,
    rigging_j_mm INTEGER NULL,
    rigging_p_mm INTEGER NULL,
    rigging_type TEXT NULL,

    sail_area_fore_cm2 INTEGER NULL,
    sail_area_main_cm2 INTEGER NULL,
    sail_area_reported_cm2 INTEGER NULL,

    tanks_fuel_capacity_l INTEGER NULL,
    tanks_water_capacity_l INTEGER NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX ix_boat_specifications_make_model
    ON boat_specifications (make, model);

/* creates "boat_profiles" table */
CREATE TABLE boat_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    specification_id UUID NOT NULL,

    call_sign TEXT NULL,
    home_port TEXT NULL,
    hin TEXT NOT NULL,
    image_uri TEXT NULL,
    name TEXT NOT NULL,
    year INTEGER NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_boat_profiles_specification
        FOREIGN KEY (specification_id)
        REFERENCES boat_specifications (id)
        ON DELETE RESTRICT
);
CREATE INDEX ix_boat_profiles_specification_id
    ON boat_profiles (specification_id);
CREATE UNIQUE INDEX uq_boat_profiles_hin
    ON boat_profiles (hin);

-- +goose Down

DROP TABLE IF EXISTS boat_profiles;
DROP TABLE IF EXISTS boat_specifications;
