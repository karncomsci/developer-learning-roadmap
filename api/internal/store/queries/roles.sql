-- name: ListRolesWithProgress :many
-- Roles ordered oldest first (R3) with Role Progress counts (P1-P3).
SELECT
  r.id,
  r.slug,
  r.name,
  r.description,
  r.icon,
  r.color,
  r.created_at,
  r.updated_at,
  count(t.id) FILTER (WHERE t.status = 'done')::int     AS done_count,
  count(t.id) FILTER (WHERE t.status = 'learning')::int AS learning_count,
  count(t.id)::int                                      AS total_count
FROM roles r
LEFT JOIN topics t ON t.role_id = r.id
GROUP BY r.id
ORDER BY r.created_at, r.id;
