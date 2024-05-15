BEGIN TRANSACTION;

INSERT INTO movies
    (title, director, year)
VALUES
    ('Inception', 'Christopher Nolan', 2010),
    ('Avatar', 'James Cameron', 2009);

COMMIT;
