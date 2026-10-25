-- +goose Up

/* creates "vessel_specifications" table */
CREATE TABLE vessel_specifications (
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
CREATE INDEX ix_vessel_specifications_make_model
    ON vessel_specifications (make, model);

/* creates "vessel_profiles" table */
CREATE TABLE vessel_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    specification_id UUID NOT NULL,
    user_id UUID NOT NULL,

    call_sign TEXT NULL,
    home_port TEXT NULL,
    hin TEXT NOT NULL,
    image_uri TEXT NULL,
    name TEXT NOT NULL,
    year INTEGER NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_vessel_profiles_specification
        FOREIGN KEY (specification_id)
        REFERENCES vessel_specifications (id)
        ON DELETE RESTRICT

    CONSTRAINT fk_vessel_profiles_user
        FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE RESTRICT
);
CREATE INDEX ix_vessel_profiles_specification_id
    ON vessel_profiles (specification_id);
CREATE INDEX ix_vessel_profiles_user_id
    ON vessel_profiles (user_id);
CREATE UNIQUE INDEX uq_vessel_profiles_hin
    ON vessel_profiles (hin);

-- +goose Down

DROP TABLE IF EXISTS vessel_profiles;
DROP TABLE IF EXISTS vessel_specifications;
