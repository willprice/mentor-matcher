CREATE TABLE user_profiles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    location TEXT,
    email TEXT NOT NULL UNIQUE
);