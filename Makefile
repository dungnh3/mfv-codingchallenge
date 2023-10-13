HOST = 127.0.0.1
PORT = 3306
DATABASE = mfv
USER = root
PASSWORD = secret

init-db:
	docker exec -it saladin_mysql mysql -u root -psecret -e "CREATE DATABASE mfv;"
.PHONY: init-db

migrate-up:
	migrate -source "file://migrations" -database "mysql://$(USER):$(PASSWORD)@tcp($(HOST):$(PORT))/$(DATABASE)" up
.PHONY: migration-up