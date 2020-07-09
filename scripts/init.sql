-- SCHEMA
CREATE TABLE timezones
(
    id          serial                      NOT NULL,
    created     timestamptz DEFAULT now()   NOT NULL,
    modified    timestamptz DEFAULT now()   NOT NULL,
    name        text                        NOT NULL,
    identifier  text                        NOT NULL
);

-- SEED
INSERT INTO timezones (id, name, identifier) VALUES (1, 'eastern', 'est');