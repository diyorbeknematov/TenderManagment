CREATE TYPE status_type AS ENUM ('open', 'close', 'awarded');

CREATE TABLE Tenders (
    id UUID PRIMARY KEY UNIQUE,          -- Tenderning unikal identifikatori
    client_id UUID REFERENCES users(id), -- Users jadvalidagi id ustuniga chet el kaliti
    title VARCHAR(255) NOT NULL,         -- Tenderning sarlavhasi
    description TEXT,                    -- Tenderning tavsifi
    deadline TIMESTAMP NOT NULL,         -- Tenderning oxirgi muddati
    budget DOUBLE PRECISION NOT NULL,    -- Tender uchun byudjet
    status INTEGER NOT NULL              -- Tenderning holati (statusi)
);
