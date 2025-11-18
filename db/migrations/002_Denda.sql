CREATE TABLE IF NOT EXISTS cek_denda (
    id VARCHAR(50) PRIMARY KEY NOT NULL,
    plat_id VARCHAR(255) NOT NULL,
    plat_nomor VARCHAR(255) NOT NULL,
    total_denda INTEGER,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    constraint fk_plat_kendaraan
        FOREIGN KEY (plat_id) 
        REFERENCES plat_model(id)
);