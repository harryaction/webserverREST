CREATE TABLE IF NOT EXISTS public.api_users (
uuid varchar not null primary key,
name varchar,
lastname varchar,
birthdate timestamptz not null
);
