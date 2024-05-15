BEGIN TRANSACTION;

UPDATE movies
SET director = 'Quentin Tarantino'
WHERE id = 1;

COMMIT;
