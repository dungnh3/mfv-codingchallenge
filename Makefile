init-db:
	docker exec -it saladin_mysql mysql -u root -psecret -e "CREATE DATABASE mfv;"
.PHONY: init-db