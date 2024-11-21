-- Create Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    country VARCHAR(50),
    city VARCHAR(50),
    address VARCHAR(255),
    postal_code VARCHAR(20),
    wanted_job_title VARCHAR(100),
    driving_license VARCHAR(20),
    nationality VARCHAR(50),
    place_of_birth VARCHAR(50),
    date_of_birth DATE,
    photo_url VARCHAR(255),
    working_experience TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Employments Table
CREATE TABLE employments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    job_title VARCHAR(100) NOT NULL,
    employer VARCHAR(100),
    start_date DATE NOT NULL,
    end_date DATE,
    city VARCHAR(50),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create Educations Table
CREATE TABLE educations (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    school VARCHAR(100) NOT NULL,
    degree VARCHAR(100),
    start_date DATE NOT NULL,
    end_date DATE,
    city VARCHAR(50),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create Skills Table
CREATE TABLE skills (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    skill VARCHAR(100) NOT NULL,
    level VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);