CREATE TYPE status_type AS ENUM ('open', 'close', 'awarded');
CREATE TYPE role_type AS ENUM ('client', 'contractor');

CREATE TABLE Tenders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),          -- Tenderning unikal identifikatori
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


CREATE TABLE user (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL UNIQUE,
    full_name VARCHAR(255) ,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE bids (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tender_id INT NOT NULL REFERENCES tenders(id) ON DELETE CASCADE,
    contractor_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    price DECIMAL(10, 2) NOT NULL,
    delivery_time TIMESTAMP NOT NULL,
    comments TEXT,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);


CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    relation_id INT,
    type VARCHAR(50) NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);
