CREATE TABLE if not exists public.sessions
(
    id               bytea  NOT NULL,
    user_id          bigint NOT NULL,
    private_key      bytea  NOT NULL,
    refresh_token    bytea  NOT NULL,
    creation_time    bigint NOT NULL,
    last_access_time bigint NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (refresh_token)
);

ALTER TABLE IF EXISTS public.sessions
    OWNER to test;

CREATE INDEX IF NOT EXISTS index_user_id
    ON public.sessions USING btree
        (user_id ASC NULLS LAST);