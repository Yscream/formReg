db:
	docker-compose up -d
		
migrate-up: 
	migrate -source file://./migrations -database postgres://postgres:2201@localhost:6080/users?sslmode=disable up 
	
migrate-down: 
	migrate -source file://./migrations -database postgres://postgres:2201@localhost:6080/users?sslmode=disable down
	
run:
	docker-compose up --build

down:
	docker-compose down
