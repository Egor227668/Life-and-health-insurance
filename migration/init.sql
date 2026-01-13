-- Создание таблицы клиентов
CREATE TABLE clients (
    client_id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    passport_number VARCHAR(50) UNIQUE NOT NULL,
    birth_date DATE NOT NULL,
    gender VARCHAR(10) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    health_info TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы агентов
CREATE TABLE agents (
    agent_id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    position VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    work_shift VARCHAR(50) NOT NULL,
    employee_id VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы страховых компаний
CREATE TABLE insurance_companies (
    company_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    inn VARCHAR(20) UNIQUE NOT NULL,
    legal_address TEXT NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы типов страхования
CREATE TABLE insurance_types (
    type_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    risks TEXT,
    base_cost DECIMAL(15,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы страховых полисов
CREATE TABLE policies (
    policy_id SERIAL PRIMARY KEY,
    policy_number VARCHAR(50) UNIQUE NOT NULL,
    client_id INT NOT NULL,
    agent_id INT NOT NULL,
    type_id INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('active', 'suspended', 'terminated')),
    coverage_amount DECIMAL(15,2) NOT NULL,
    premium DECIMAL(15,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (client_id) REFERENCES clients(client_id) ON DELETE CASCADE,
    FOREIGN KEY (agent_id) REFERENCES agents(agent_id) ON DELETE CASCADE,
    FOREIGN KEY (type_id) REFERENCES insurance_types(type_id) ON DELETE CASCADE
);

-- Создание таблицы страховых случаев
CREATE TABLE insurance_claims (
    claim_id SERIAL PRIMARY KEY,
    claim_number VARCHAR(50) UNIQUE NOT NULL,
    policy_id INT NOT NULL,
    submission_date DATE NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('submitted', 'reviewing', 'approved', 'rejected')),
    claim_amount DECIMAL(15,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (policy_id) REFERENCES policies(policy_id) ON DELETE CASCADE
);

-- Создание таблицы бенефициаров
CREATE TABLE beneficiaries (
    beneficiary_id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    relationship VARCHAR(100) NOT NULL,
    passport_number VARCHAR(50) NOT NULL,
    policy_id INT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (policy_id) REFERENCES policies(policy_id) ON DELETE CASCADE
);

-- Создание таблицы повреждений (ущербов)
CREATE TABLE damages (
    damage_id SERIAL PRIMARY KEY,
    claim_id INT NOT NULL,
    damage_type VARCHAR(255) NOT NULL,
    severity VARCHAR(50) NOT NULL CHECK (severity IN ('low', 'medium', 'high', 'critical')),
    icd_code VARCHAR(20),
    diagnosis_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (claim_id) REFERENCES insurance_claims(claim_id) ON DELETE CASCADE
);

-- Создание таблицы коэффициентов
CREATE TABLE coefficients (
    coefficient_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    value DECIMAL(10,4) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы для связи полисов и коэффициентов (многие-ко-многим)
CREATE TABLE policy_coefficients (
    policy_id INT NOT NULL,
    coefficient_id INT NOT NULL,
    applied_value DECIMAL(10,4) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (policy_id, coefficient_id),
    FOREIGN KEY (policy_id) REFERENCES policies(policy_id) ON DELETE CASCADE,
    FOREIGN KEY (coefficient_id) REFERENCES coefficients(coefficient_id) ON DELETE CASCADE
);