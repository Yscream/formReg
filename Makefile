run:
	docker-compose build
	docker-compose up postgresql
		
mig: 
	docker run -v /mnt/c/Users/Spark/Desktop/clone/go-form-reg/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://postgres:2201@localhost:6080/users?sslmode=disable" up
	docker-compose up

stop:
	docker-compose down
