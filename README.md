# blockchsin_api


sudo docker run --name=postgres -e POSTGRES_PASSWORD='postgres' -p 7432:5432 -d postgres

migrate -path "./schema" -database "postgres://postgres:postgres@localhost:7432/postgres?sslmode=disable" up
 
sudo cat ./postgres/db/seeds/seeds.sql | psql postgres://postgres:postgres@localhost:7323

migrate create -ext sql -dir ./schema/migrations -seq init

sudo docker run -v schema:/migrations --network host migrate/migrate
    -path=schema/migrations/000001_init.up.sql -database postgres://postgres:postgres@localhost:7323/postgres?sslmode=disable up



# connect to DB:
- sudo docker exec -it e8e7b24f4f08 /bin/bash
- psql -U postgres
