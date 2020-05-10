-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public."Event"
(
    id                uuid                     NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    short_description CHARACTER VARYING(100)   NOT NULL,
    date              TIMESTAMP WITH TIME ZONE NOT NULL,
    duration          BIGINT                   NOT NULL,
    full_description  TEXT                     NULL,
    remind_before     BIGINT                   NULL,
    user_id           uuid                     NOT NULL
);

CREATE INDEX user_idx ON public."Event" (user_id);

ALTER TABLE public."Event"
    OWNER to postgres;

-- +goose Down
DROP TABLE public."Event";