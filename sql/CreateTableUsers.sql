CREATE TABLE if not exists public.users
(
    id          bigserial NOT NULL,
    telegram_id bigint    NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (telegram_id)
);

ALTER TABLE IF EXISTS public.users
    OWNER to test;