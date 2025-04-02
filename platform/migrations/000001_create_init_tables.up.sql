-- Set timezone
SET TIMEZONE="Europe/Moscow";

-- Create users table
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  email VARCHAR(255) NOT NULL,
  pass_hash VARCHAR(255) NOT NULL
);