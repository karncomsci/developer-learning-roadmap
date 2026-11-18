-- Local development seed only. Production starts empty (docs/SPEC.md §11).

insert into roles (id, slug, name, description, color, created_at) values
  ('00000000-0000-0000-0000-000000000001', 'backend-developer', 'Backend Developer',
   'Server-side development with Go and Postgres.', '#2563eb', now() - interval '1 day');

insert into topics (id, role_id, title, description_md, status) values
  ('10000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000001', 'Internet basics', 'How HTTP, DNS and the web work.', 'done'),
  ('10000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000001', 'Git', 'Version control fundamentals.', 'done'),
  ('10000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000001', 'Go language', 'Syntax, types, modules and tooling.', 'learning'),
  ('10000000-0000-0000-0000-000000000004', '00000000-0000-0000-0000-000000000001', 'SQL', 'Queries, joins and indexes.', 'learning'),
  ('10000000-0000-0000-0000-000000000005', '00000000-0000-0000-0000-000000000001', 'PostgreSQL', 'Transactions, locking and migrations.', 'not_started'),
  ('10000000-0000-0000-0000-000000000006', '00000000-0000-0000-0000-000000000001', 'REST API design', 'Resources, status codes and versioning.', 'not_started'),
  ('10000000-0000-0000-0000-000000000007', '00000000-0000-0000-0000-000000000001', 'Authentication', 'Sessions, JWT and OAuth.', 'not_started'),
  ('10000000-0000-0000-0000-000000000008', '00000000-0000-0000-0000-000000000001', 'Testing', 'Unit and integration tests in Go.', 'not_started'),
  ('10000000-0000-0000-0000-000000000009', '00000000-0000-0000-0000-000000000001', 'Docker', 'Images, containers and compose.', 'not_started'),
  ('10000000-0000-0000-0000-000000000010', '00000000-0000-0000-0000-000000000001', 'CI/CD', 'Automated build, test and deploy.', 'not_started');

insert into prerequisites (topic_id, prerequisite_topic_id) values
  ('10000000-0000-0000-0000-000000000003', '10000000-0000-0000-0000-000000000001'),
  ('10000000-0000-0000-0000-000000000005', '10000000-0000-0000-0000-000000000004'),
  ('10000000-0000-0000-0000-000000000006', '10000000-0000-0000-0000-000000000003'),
  ('10000000-0000-0000-0000-000000000007', '10000000-0000-0000-0000-000000000006'),
  ('10000000-0000-0000-0000-000000000008', '10000000-0000-0000-0000-000000000003'),
  ('10000000-0000-0000-0000-000000000010', '10000000-0000-0000-0000-000000000002'),
  ('10000000-0000-0000-0000-000000000010', '10000000-0000-0000-0000-000000000009'),
  ('10000000-0000-0000-0000-000000000010', '10000000-0000-0000-0000-000000000008');

insert into resources (topic_id, title, url, type) values
  ('10000000-0000-0000-0000-000000000003', 'A Tour of Go', 'https://go.dev/tour/', 'official_docs'),
  ('10000000-0000-0000-0000-000000000005', 'PostgreSQL Documentation', 'https://www.postgresql.org/docs/', 'official_docs');
