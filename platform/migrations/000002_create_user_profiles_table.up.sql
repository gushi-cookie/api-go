-- Create user_profiles table
CREATE TABLE user_profiles (
  user_id UUID PRIMARY KEY,
  nickname VARCHAR(32) NOT NULL,
  bio VARCHAR(620),
  CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);