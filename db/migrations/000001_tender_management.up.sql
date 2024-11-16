CREATE TYPE status_type AS ENUM ('open', 'close', 'awarded');
CREATE TYPE role_type AS ENUM ('client', 'contractor');
CREATE TYPE bid_status_type AS ENUM ('fail', 'award', 'process');

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
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),          -- Tenderning unikal identifikatori
    client_id UUID REFERENCES users(id), -- Users jadvalidagi id ustuniga chet el kaliti
    title VARCHAR(255) NOT NULL,         -- Tenderning sarlavhasi
    description TEXT,                    -- Tenderning tavsifi
    deadline TIMESTAMP NOT NULL,         -- Tenderning oxirgi muddati
    budget DOUBLE PRECISION NOT NULL,    -- Tender uchun byudjet
    status status_type DEFAULT 'open',              -- Tenderning holati (statusi)
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
    status bid_status_type DEFAULT 'process',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);


CREATE TABLE notifications (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    relation_id UUID NOT NULL,
    type VARCHAR(50) NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);
