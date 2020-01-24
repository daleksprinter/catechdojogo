## Usage
```
docker-compose up --build db

cd db
make db/reset
make migrate/up

cd ../
docker-compose up --build server
```