db:
	docker-compose up -d
		
migrate-up: 
	migrate -source file://./pkg/repository/postgresql/migrations/ -database postgres://postgres:2201@localhost:6080/users?sslmode=disable up 
	
migrate-down: 
	migrate -source file://./pkg/repository/postgresql/migrations -database postgres://postgres:2201@localhost:6080/users?sslmode=disable down
	
up:
	docker-compose up --build

down:
	docker-compose down
