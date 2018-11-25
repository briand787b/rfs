/*CREATE TABLE test_models (
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
*/

CREATE TABLE media_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) UNIQUE NOT NULL
);

INSERT INTO media_types 
(
    name
)
VALUES
(
    'video' -- 1
),
(
    'text' -- 2
),
(
    'image' -- 3
),
(
    'audio' -- 4
),
(
    'interactive' -- 5
);

CREATE TABLE creators (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE media (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    media_type_id INTEGER REFERENCES media_types(id) NOT NULL,
    parent_id INTEGER REFERENCES media(id),
    creator_id INTEGER REFERENCES creators(id),
    release_year INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO media
(
    name,
    media_type_id
)
VALUES
(
    'Movies',
    1
),
(
    'TV Shows',
    1
),
(
    'Lectures',
    1
),
(
    'Books',
    2
),
(
    'Pictures',
    3
),
(
    'Music',
    4
),
(
    'Video Games',
    5
);

CREATE TABLE parent_child_media (
    parent_id INTEGER REFERENCES media(id) NOT NULL,
    child_id INTEGER REFERENCES media(id) NOT NULL,
    PRIMARY KEY(parent_id, child_id)
);

-- end-users should not need to be aware of files when managing their content,
-- other than during the initial upload.  the data structure needs to be 
-- modeled to reflect that
CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    media_id INTEGER NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    file_extension VARCHAR(16) NOT NULL,
    md5_checksum UUID NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE networks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) NOT NULL UNIQUE,
    internet_ip VARCHAR(15) NOT NULL UNIQUE,
    port_80_blocked BOOLEAN NOT NULL, -- probably shouldnt use 80 at all anyway
    port_443_blocked BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE servers (
    id SERIAL PRIMARY KEY,
    network_id INTEGER REFERENCES networks(id) NOT NULL,
    local_ip VARCHAR(15) NOT NULL,
    is_master BOOLEAN NOT NULL,
    max_storage_bytes BIGINT NOT NULL,
    working_dir VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE server_media (
    server_id INTEGER REFERENCES servers(id) NOT NULL,
    media_id INTEGER REFERENCES media(id) NOT NULL,
    PRIMARY KEY(server_id, media_id)
);

CREATE TABLE server_files (
    server_id INTEGER REFERENCES servers(id) NOT NULL,
    file_id INTEGER REFERENCES files(id) NOT NULL,
    PRIMARY KEY(server_id, file_id)
);

CREATE TABLE action_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO action_types
(
    name
)
VALUES
(
    'add'
),
(
    'remove'
);

-- an action is a verb associated with a certain server
CREATE TABLE actions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) NOT NULL UNIQUE,
    action_type INTEGER REFERENCES action_types(id) NOT NULL,
    server_id INTEGER REFERENCES servers(id) NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- a role is a collection of actions
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE role_actions (
    role_id INTEGER REFERENCES roles(id) NOT NULL,
    action_id INTEGER REFERENCES actions(id) NOT NULL,
    PRIMARY KEY(role_id, action_id)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) NOT NULL UNIQUE,
    email VARCHAR(55) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_roles (
    user_id INTEGER REFERENCES users(id) NOT NULL,
    role_id INTEGER REFERENCES roles(id) NOT NULL,
    PRIMARY KEY(user_id, role_id)
);