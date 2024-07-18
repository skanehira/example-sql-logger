-- name: GetT1 :many
SELECT * FROM t1;

-- name: CreateT1 :execresult
INSERT INTO t1 (id, name) VALUES (?, ?);

-- name: UpdateT1 :execresult
UPDATE t1 SET name = sqlc.arg(name) WHERE id = sqlc.arg(id);

-- name: DeleteT1 :execresult
DELETE FROM t1 WHERE id = sqlc.arg(id);
