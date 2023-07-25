-- Create the cakes table
CREATE TABLE cakes (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    rating FLOAT NOT NULL,
    image TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME,
    deleted_at DATETIME
);

-- Insert data into the cakes table
INSERT INTO cakes (title, description, rating, image, created_at, updated_at, deleted_at)
VALUES ('Lemon cheesecake', 'A cheesecake made of lemon', 7.2, 'https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/', '2020-02-01 10:56:31', '2020-02-13 09:30:23', '2020-02-13 09:30:23');
