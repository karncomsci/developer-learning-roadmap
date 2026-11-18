-- Initial schema for Developer Roadmap v1 (see docs/SPEC.md §4).

create type learning_status as enum ('not_started', 'learning', 'done');
create type resource_type   as enum ('article', 'video', 'official_docs', 'course', 'other');

create table roles (
  id          uuid primary key default gen_random_uuid(),
  slug        text not null unique check (slug ~ '^[a-z0-9]+(-[a-z0-9]+)*$'),
  name        text not null,
  description text not null default '',
  icon        text,
  color       text,
  created_at  timestamptz not null default now(),
  updated_at  timestamptz not null default now()
);

create table topics (
  id                uuid primary key default gen_random_uuid(),
  role_id           uuid not null references roles(id) on delete cascade,
  title             text not null,
  description_md    text not null default '',
  status            learning_status not null default 'not_started',
  status_changed_at timestamptz not null default now(),
  created_at        timestamptz not null default now(),
  updated_at        timestamptz not null default now()
);

create index topics_role_id_idx on topics (role_id);

create table prerequisites (
  topic_id              uuid not null references topics(id) on delete cascade,
  prerequisite_topic_id uuid not null references topics(id) on delete cascade,
  primary key (topic_id, prerequisite_topic_id),
  check (topic_id <> prerequisite_topic_id)
);

create index prerequisites_prerequisite_topic_id_idx on prerequisites (prerequisite_topic_id);

create table resources (
  id         uuid primary key default gen_random_uuid(),
  topic_id   uuid not null references topics(id) on delete cascade,
  title      text not null,
  url        text not null,
  type       resource_type not null default 'other',
  created_at timestamptz not null default now()
);

create index resources_topic_id_idx on resources (topic_id);

create table topic_images (
  id           uuid primary key default gen_random_uuid(),
  topic_id     uuid not null references topics(id) on delete cascade,
  storage_path text not null unique,
  content_type text not null,
  size_bytes   integer not null,
  created_at   timestamptz not null default now()
);

create index topic_images_topic_id_idx on topic_images (topic_id);

-- The Go API is the only client and connects as the postgres role.
-- Lock the tables away from Supabase's public PostgREST roles.
alter table roles         enable row level security;
alter table topics        enable row level security;
alter table prerequisites enable row level security;
alter table resources     enable row level security;
alter table topic_images  enable row level security;
