#! /bin/bash

docker container rm -f db

docker volume rm db_vol
docker volume create db_vol

docker image build . -t rfs_db
docker container run \
    -d \
    --name db \
    -v db_vol:/var/lib/postgresql/data \
    --env-file DB.env \
    -p 2345:5432 \
    rfs_db

echo $(docker logs db)
