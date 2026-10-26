-- +goose Up

INSERT INTO vessel_specifications (
    id,
    builder,
    construction,
    designer,
    first_built,
    hull_type,
    last_built,
    make,
    model,

    dimensions_ballast_kg,
    dimensions_beam_mm,
    dimensions_displacement_kg,
    dimensions_draft_mm,
    dimensions_loa_mm,
    dimensions_lwl_mm,

    engine_cooling,
    engine_drive,
    engine_fuel_type,
    engine_gearbox,
    engine_horsepower_hp,
    engine_make,
    engine_model,
    engine_propeller,

    rigging_e_mm,
    rigging_forestay_length_mm,
    rigging_i_mm,
    rigging_j_mm,
    rigging_p_mm,
    rigging_type,

    sail_area_fore_cm2,
    sail_area_main_cm2,
    sail_area_reported_cm2,

    tanks_fuel_capacity_l,
    tanks_water_capacity_l
) VALUES (
    '4b11f254-b2a5-4f6d-b442-452a2f6f6d12',
    'Marine Projects Ltd.',
    'GRP',
    'Bill Dixon',
    1985,
    'Fin with rudder on skeg',
    1991,
    'Moody',
    '376',

    2948,
    3810,
    7371,
    1680,
    11530,
    9500,

    'freshwater',
    'shaft',
    'diesel',
    'ZF 10 M',
    35,
    'Thornycroft',
    'T80',
    'fixed 3-blade',

    4320,
    15040,
    14330,
    4570,
    12600,
    'Masthead Sloop',

    327500,
    271900,
    602000,

    204,
    250
)
ON CONFLICT (id) DO NOTHING;

-- +goose Down

DELETE FROM vessel_specifications
WHERE id = '4b11f254-b2a5-4f6d-b442-452a2f6f6d12';
