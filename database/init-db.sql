CREATE TABLE test_models (
    id SERIAL UNIQUE PRIMARY KEY,
    name VARCHAR(55) NOT NULL UNIQUE
);

INSERT INTO test_models
(
    name
)
VALUES
(
    'Ryan'
),
(
    'Brian'
),
(
    'Matt'
);