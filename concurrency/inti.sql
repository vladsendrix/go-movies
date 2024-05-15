CREATE TABLE movies
(
    id INT NOT NULL,
    title VARCHAR(255),
    director VARCHAR(255),
    year INT,
    PRIMARY KEY (id)
);

INSERT INTO movies
    (id, title, director, year)
VALUES
    (1, 'The Shawshank Redemption', 'Frank Darabont', 1994),
    (2, 'The Godfather', 'Francis Ford Coppola', 1972),
    (3, 'The Dark Knight', 'Christopher Nolan', 2008);
