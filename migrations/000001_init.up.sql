CREATE SCHEMA pvz_management;

-- Таблица пользователей
CREATE TABLE pvz_management.users (
                                      id UUID PRIMARY KEY,
                                      email VARCHAR(255) UNIQUE NOT NULL,
                                      password_hash VARCHAR(255) NOT NULL,
                                      role VARCHAR(20) NOT NULL CHECK (role IN ('employee', 'moderator')),
                                      created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Таблица ПВЗ
CREATE TABLE pvz_management.pvz (
                                    id UUID PRIMARY KEY,
                                    city VARCHAR(50) NOT NULL CHECK (city IN ('Москва', 'Санкт-Петербург', 'Казань')),
                                    registration_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Таблица приемок
CREATE TABLE pvz_management.receptions (
                                           id UUID PRIMARY KEY,
                                           pvz_id UUID NOT NULL REFERENCES pvz_management.pvz(id),
                                           status VARCHAR(20) NOT NULL CHECK (status IN ('in_progress', 'close')),
                                           date_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                           created_by UUID REFERENCES pvz_management.users(id)
);

-- Таблица продуктов
CREATE TABLE pvz_management.products (
                                         id UUID PRIMARY KEY,
                                         reception_id UUID NOT NULL REFERENCES pvz_management.receptions(id),
                                         type VARCHAR(50) NOT NULL CHECK (type IN ('одежда', 'обувь', 'электроника')),
                                         date_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для ускорения поиска надо потом добавить
