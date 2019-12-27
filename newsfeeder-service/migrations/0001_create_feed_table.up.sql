CREATE SEQUENCE IF NOT EXISTS public.feed_id_seq;

CREATE TABLE public.feed (
                feed_id BIGINT NOT NULL DEFAULT nextval('public.feed_id_seq'),
                created_at TIMESTAMP DEFAULT (now() at time zone 'utc') NOT NULL,
                modified_at TIMESTAMP DEFAULT (now() at time zone 'utc') NOT NULL,
                category VARCHAR NOT NULL,
                provider VARCHAR NOT NULL,
                url VARCHAR NOT NULL,
                CONSTRAINT feed_pk PRIMARY KEY (feed_id)
);
ALTER SEQUENCE public.feed_id_seq OWNED BY public.feed.feed_id;