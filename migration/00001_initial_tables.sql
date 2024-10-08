-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(100) NOT NULL COMMENT "borrower, investor, employee",
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS loans (
    id BIGINT PRIMARY KEY,
    borrower_id BIGINT NOT NULL,
    principal_amount DECIMAL(10, 2) NOT NULL,
    invested_amount DECIMAL(10, 2) NOT NULL,
    interest_rate DECIMAL(5, 2) NOT NULL COMMENT "Per annum",
    status VARCHAR(100) NOT NULL COMMENT "proposed, approved, invested, disbursed",
    approval_date TIMESTAMP,
    approval_employee_id BIGINT,
    disbursement_date TIMESTAMP,
    agreement_letter_document_url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS loan_investments (
    id BIGINT PRIMARY KEY,
    loan_id BIGINT NOT NULL,
    investor_id BIGINT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (id, name, type) VALUES 
(1, 'A', 'borrower'),
(2, 'B', 'investor'),
(3, 'C', 'investor'),
(4, 'D', 'employee');

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS loans;
DROP TABLE IF EXISTS loan_investments;