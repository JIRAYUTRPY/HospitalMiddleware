CREATE TABLE IF NOT EXISTS patient (
    id SERIAL PRIMARY KEY,
    national_id VARCHAR(255) NOT NULL,
    passport_id VARCHAR(255) NOT NULL,
    first_name_th VARCHAR(255) NOT NULL,
    middle_name_th VARCHAR(255) NOT NULL,
    last_name_th VARCHAR(255) NOT NULL,
    first_name_en VARCHAR(255) NOT NULL,
    middle_name_en VARCHAR(255) NOT NULL,
    last_name_en VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    gender VARCHAR(2) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    patent_hn VARCHAR(255) NOT NULL DEFAULT 'A',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

INSERT INTO patient (national_id, passport_id, first_name_th, middle_name_th, last_name_th, first_name_en, middle_name_en, last_name_en, birth_date, gender, phone_number, email, patent_hn) VALUES ('1234567890123', '1234567890123', 'John', 'Doe', 'Smith', 'John', 'Doe', 'Smith', '1990-01-01', 'M', '081234567890', 'john.doe@example.com', 'A');