CREATE TABLE
    movies (
        id SERIAL PRIMARY KEY,
        title VARCHAR(100),
        release_year INTEGER,
        director_id INTEGER
    );

CREATE TABLE
    actors (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        birth_year INTEGER
    );

CREATE TABLE
    directors (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        birth_year INTEGER
    );

CREATE TABLE
    genres (id SERIAL PRIMARY KEY, name VARCHAR(100));

CREATE TABLE
    movie_actors (
        movie_id INTEGER REFERENCES movies (id),
        actor_id INTEGER REFERENCES actors (id),
        PRIMARY KEY (movie_id, actor_id)
    );

CREATE TABLE
    movie_genres (
        movie_id INTEGER REFERENCES movies (id),
        genre_id INTEGER REFERENCES genres (id),
        PRIMARY KEY (movie_id, genre_id)
    );

ALTER TABLE movies ADD FOREIGN KEY (director_id) REFERENCES directors (id);