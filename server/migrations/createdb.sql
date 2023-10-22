-- Создание таблицы "Clients"
CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    last_name TEXT,
    first_name TEXT,
    patronymic TEXT,
    phone VARCHAR(15),
    link_to_chat TEXT,
    login TEXT UNIQUE,
    password VARCHAR(70),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы "Insurers"
CREATE TABLE insurers (
    id SERIAL PRIMARY KEY,
    last_name TEXT,
    first_name TEXT,
    patronymic TEXT,
    phone VARCHAR(15),
    login TEXT UNIQUE,
    password VARCHAR(70),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы "Insurance_applications"
CREATE TABLE insurance_applications (
    id SERIAL PRIMARY KEY,
    client_id INT REFERENCES clients(id),
    insurer_id INT REFERENCES insurers(id),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы "CarInsurance"
CREATE TABLE car_insurance (
    id SERIAL PRIMARY KEY,
    application_id INT REFERENCES insurance_applications(id),
    description TEXT
);

-- Создание таблицы "HomeInsurance"
CREATE TABLE home_insurance (
    id SERIAL PRIMARY KEY,
    application_id INT REFERENCES insurance_applications(id),
    description TEXT
);
