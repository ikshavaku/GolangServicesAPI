-- name: Ping :one
SELECT 1;

-- name: CountServices :one
SELECT COUNT(*) AS total
FROM services s
WHERE s.is_deleted = FALSE
  AND (s.name ILIKE COALESCE('%' || sqlc.narg('name_filter')::text || '%', '%'));

-- name: ListServices :many
SELECT
    s.id,
    s.name,
    s.description,
    s.created_at,
    s.updated_at,
    s.is_deleted
FROM services s
WHERE s.is_deleted = FALSE
  AND (s.name ILIKE COALESCE('%' || sqlc.narg('name_filter')::text || '%', '%'))
ORDER BY s.name ASC -- fixed sort by name
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetServiceByID :one
SELECT 
* 
FROM services s
WHERE
s.id = sqlc.arg(id)
AND s.is_deleted = FALSE;

-- name: ListServiceVersionsByServiceID :many
SELECT
*
FROM service_versions
WHERE
service_id = sqlc.arg(service_id)
AND is_deleted = FALSE;