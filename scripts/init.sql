-- SCHEMA
CREATE TABLE timezones
(
    id          serial                      NOT NULL,
    created     timestamptz DEFAULT now()   NOT NULL,
    modified    timestamptz DEFAULT now()   NOT NULL,
    name        text                        NOT NULL,
    offset      smallint                    NOT NULL,
    identifier  text                        NOT NULL
);

-- SEED
INSERT INTO timezones (id, name, identifier) VALUES (1, 'eastern', '-5', 'est');
INSERT INTO timezones (id, name, identifier) VALUES (2, 'central', '-6', 'cst');
INSERT INTO timezones (id, name, identifier) VALUES (3, 'mountain', '-7', 'mst');
INSERT INTO timezones (id, name, identifier) VALUES (4, 'pacific', '-8', 'pst');
INSERT INTO timezones (id, name, identifier) VALUES (5, 'alaska', '-9', 'ast');
