BEGIN TRANSACTION;

SELECT title, director
FROM movies
WHERE year >= 2000;

SELECT SLEEP(5);

SELECT title, director
FROM movies
WHERE year >= 2000;

COMMIT;
