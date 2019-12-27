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

-- insert default feeds
INSERT INTO public.feed (created_at,modified_at,category,provider,url) VALUES 
('2019-12-26 22:16:24.239','2019-12-26 22:16:24.239','UK','BBC','http://feeds.bbci.co.uk/news/uk/rss.xml')
,('2019-12-26 22:43:08.976','2019-12-26 22:43:08.976','Technology','BBC','http://feeds.bbci.co.uk/news/technology/rss.xml')
,('2019-12-26 22:44:03.088','2019-12-26 22:44:03.088','UK','Reuters','http://feeds.reuters.com/reuters/UKdomesticNews?format=xml')
,('2019-12-26 00:04:53.336','2019-12-26 00:04:53.336','Techology','Reuters','http://feeds.reuters.com/reuters/UKdomesticNews?format=xml')
;