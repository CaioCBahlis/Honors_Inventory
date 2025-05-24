CREATE TABLE locations(
    id SERIAL PRIMARY KEY,
    room_name TEXT NOT NULL,
    building_type TEXT NOT NULL
);

CREATE TABLE audit(
    id SERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE equipment(
    id SERIAL PRIMARY KEY,
    model TEXT NOT NULL,
    equipment_type TEXT NOT NULL,
    equipment_status TEXT NOT NULL,
    location_id INT NOT NULL REFERENCES locations(id),
    inserted_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

