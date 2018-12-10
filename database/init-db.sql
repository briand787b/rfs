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

-- 'creator' is synonymous with 'artist' for certain works
CREATE TABLE creators (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO creators
(
    name
)
VALUES
(
    'anonymous'
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
    is_port_80_blocked BOOLEAN NOT NULL, -- probably shouldnt use 80 at all anyway
    is_port_443_blocked BOOLEAN NOT NULL,
    is_private BOOLEAN NOT NULL, -- can other computers on this network be trusted
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
    -- add a file to a server
    'add' -- 1
),
-- ( MODIFY WILL BE A V2.0 FEATURE
--     -- modify a file on a server
--     'modify' -- 2
-- ),
(
    -- remove a file on a server
    'remove' -- 3
),
(
    -- receive a file from a server (typically for viewing/downloading)
    'get' -- 4
);

-- an action is a verb associated with a certain server
CREATE TABLE actions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) NOT NULL UNIQUE,
    action_type INTEGER REFERENCES action_types(id) NOT NULL,
    is_global BOOLEAN NOT NULL DEFAULT FALSE,
    server_id INTEGER REFERENCES servers(id), -- should this be NOT NULL? (some actions might not relate to servers)
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO actions
(
    name,
    action_type,
    is_global
)
VALUES
( -- Files
    'Global Add File',
    1,
    TRUE,
),
(
    'Global Modify File',
    2,
    TRUE
),
(
    'Global Remove File',
    3,
    TRUE
),
(
    'Get File',
    4,
    TRUE
), -- Media
(
    'Add Media',
    1,
    TRUE,
),
(
    'Modify Media',
    2,
    TRUE
),
(
    'Remove Media',
    3,
    TRUE
),
(
    'Get Media',
    4,
    TRUE
), -- Servers
(
    'Add Server',
    1,
    TRUE,
),
(
    'Modify Server',
    2,
    TRUE
),
(
    'Remove Server',
    3,
    TRUE
), -- Networks
(
    'Add Network',
    1,
    TRUE,
),
(
    'Modify Network',
    2,
    TRUE
),
(
    'Remove Network',
    3,
    TRUE
);

-- attempt 1
-- users can own (digital) property e.g. networks, servers, media, files, etc...
-- each property can be either public or private
-- public means anyone can view/retrieve the property
-- private means that only owners and admins can view/retrieve property
-- property can be native or non-native
-- property can be modified and deleted only by its owner, whether public or private
-- whenever an action on a property is requested, a roll up of the property hierarchy will be executed
--      and if any of those permissions grant the action, it will be performed

-- attempt 2
-- users can own (digital) property e.g. networks, servers, media, files, etc...
-- all properties are visible to all users
-- if an action is requested on a certain property, permission to perform the action
--      is granted only to users who either own that property or the hierarchical parents
--      of that property.
--      Example:
--          a) User A wants to delete a file from his own server that is owned by User B
--          b) Although User A does not own the media, he does own the server, which is
--              hierarchically superior to the media.  He is therefore permitted to delete
--              the property (in this case media).

-- attempt 3
-- users own property
-- there are two types of property - digital and physical
-- physical property can be modified by its owner
-- digital property cannot be modified within rfs (i.e. copy it, modify it, add it back to rfs, delte old version)
-- users have absolute control over their physical property
-- users can own (digital) property e.g. networks, servers, media, files, etc...
-- only the owner of a server can add/modify/delete content on that server
-- media and/or files (TBD) can be listed as 'restricted', meaning they can only
--      live on servers owned by the user
-- any non-restricted media can be added to servers not owned by the user by
--      both admins and the system
-- to ensure file durability, the system user can automatically add content to servers
--      even if the user who owns the server did not request the transfer, but only
--      if the files are not restricted
-- 

-- GENERAL RULES AND THEIR IMPLICATIONS
-- Three user types (UserType): Super, Standard, Public
-- UserType hierarchy: Super > Standard > Public
-- A user is considered to own the property of users who are of types that are lower in the hierarchy than them
--      e.g. a Super User technically owns all content of Standard and Public Users
--      e.g. a Standard User technically owns all content of Public Users
-- At startup, two users are created, 'Admin' and 'Community'
-- 'Admin' user is of type Super User
-- 'Community' user is of type Public User
-- New Users can only be added to rfs by an admin (rfs is invitation only)
-- New Users will always be signed up with a UserType of Standard unless they are intended to be an admin (rare)
-- rfs has a concept of 'property' being anything that can be owned by an entity
-- There are two types of property: Physical and Digital
--      e.g. Physical property being Servers and Networks
--      e.g. Digital property being Media and Files
-- Physical property can be mutated
--      e.g. a Server can have its IP address changed (an unavoidable real world scenario)
-- Digital property cannot be mutated (makes building rfs simpler)
-- Because of this, media owned by the Community user cannot be adulterated by other users
-- If you want to add or remove a property, you must own the parent property
--      e.g. to add media to a server, you must own the server
--      e.g. to add a child media to a parent media, you must own the parent media
--          (top level media are created by Community User at startup)
--      e.g. to add a server to a network, you must own the network
--      e.g to add a file to a media, you must own the media
-- All basal media (media without parents) will be owned by the Community user.
-- Thus, users can add media that are children to the basal media without owning the basal media
-- The application should NOT allow the creation of basal media after initialization
-- As a general rule, content can only be added or removed on a server by its owner (e.g. admin and actual owner)
--      e.g. rfs detects there are files missing from a server and asks
--      the user for permission to download those files to regain sync


-- All infrastructure that the master process runs on MUST be owned by the Admin user 
-- /\ May not be necessary since the app server should not be storing files

-- 

-- For syncing, out of sync status is only detected for media when the server already holds 
--      at least one child media
--      e.g. server A has multiple episodes of season 6 of a show and detects that one of
--          the episodes in that season is not on the server, it will request to add it
--      e.g. server A has multiple episodes of season 6 of a show but no episodes of season
--          7, it will not ask to add season 7 episodes to the server


-- a role is a collection of actions
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO roles
(
    name
)
VALUES
(
    'admin'
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

-- 'System' is the user through which the application performs automatic admin actions
INSERT INTO users 
(
    name,
    email
)
VALUES
(
    'system',
    'system_doesnt_read_emails@acme.com'
);

CREATE TABLE user_roles (
    user_id INTEGER REFERENCES users(id) NOT NULL,
    role_id INTEGER REFERENCES roles(id) NOT NULL,
    PRIMARY KEY(user_id, role_id)
);

CREATE TABLE transfers {
    id SERIAL PRIMARY KEY,
    -- \/ NOTE: requester may not necessarily be recipient
    -- example: system request for transfer between two other users to maintain file durability)
    request_user_id INTEGER REFERENCES users(id) NOT NULL, 
    requested_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    executed_at TIMESTAMP,
    completed_at TIMESTAMP,

}