-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
CREATE TABLE tweets (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    username character varying NOT NULL,
    tweet_id character varying NOT NULL,
    content text NOT NULL,
    posted_at character varying NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now()
);
ALTER TABLE tweets ADD CONSTRAINT tweets_pkey PRIMARY KEY (id);
CREATE INDEX index_tweets_on_id ON tweets USING btree (id);
CREATE UNIQUE INDEX index_tweets_on_tweet_id ON tweets USING btree (tweet_id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE tweets;
DROP EXTENSION "uuid-ossp";