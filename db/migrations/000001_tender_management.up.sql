CREATE TYPE status_type AS ENUM ('open', 'close', 'awarded');
CREATE TYPE role_type AS ENUM ('client', 'contractor');

CREATE TABLE users (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE Tenders (
    id UUID PRIMARY KEY UNIQUE,          -- Tenderning unikal identifikatori
    client_id UUID REFERENCES users(id), -- Users jadvalidagi id ustuniga chet el kaliti
    title VARCHAR(255) NOT NULL,         -- Tenderning sarlavhasi
    description TEXT,                    -- Tenderning tavsifi
    deadline TIMESTAMP NOT NULL,         -- Tenderning oxirgi muddati
    budget DOUBLE PRECISION NOT NULL,    -- Tender uchun byudjet
    status INTEGER NOT NULL,              -- Tenderning holati (statusi)
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE bids (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    tender_id UUID NOT NULL REFERENCES tenders(id) ON DELETE CASCADE,
    contractor_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    price DECIMAL(10, 2) NOT NULL,
    delivery_time TIMESTAMP NOT NULL,
    comments TEXT,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);


CREATE TABLE notifications (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    relation_id INT,
    type VARCHAR(50) NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);
