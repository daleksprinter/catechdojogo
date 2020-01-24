DBNAME:=CATechDojo
goose:= $(which goose)

migrate/create:
	$(goose) create $(NAME) sql

migrate/up:
	$(goose) up

db/reset:
	docker-compose exec db mysql -h 0.0.0.0 -u root -p -e "drop database if exists $(DBNAME); create database $(DBNAME);"
