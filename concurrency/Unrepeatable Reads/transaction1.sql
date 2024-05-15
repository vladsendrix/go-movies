BEGIN TRANSACTION;

SELECT title, director
FROM movies
WHERE id = 1;

SELECT SLEEP(5);

SELECT title, director
FROM movies
WHERE id = 1;
COMMIT;
