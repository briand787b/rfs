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

CREATE TABLE media (
    id SERIAL UNIQUE PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    parent_id INTEGER REFERENCES media(id),
    feature_file_id INTEGER REFERENCES files(id) UNIQUE,
    release_year INTEGER,

);

CREATE TABLE parent_child_media (
    parent_id INTEGER REFERENCES media(id) NOT NULL,
    child_id INTEGER REFERENCES media(id) NOT NULL,
    PRIMARY KEY(parent_id, child_id)
);

CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    media_id INT NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    md5_checksum UUID NOT NULL UNIQUE,
);

CREATE TABLE workers (
    id SERIAL PRIMARY KEY,
    network_id INTEGER REFERENCES networks(id) NOT NULL,
    local_ip INET NOT NULL
);

CREATE TABLE networks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) NOT NULL UNIQUE,
    external_ip INET NOT NULL UNIQUE,
    cidr_block CIDR NOT NULL,

);