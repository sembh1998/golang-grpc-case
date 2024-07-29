protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative testpb/*.proto

postgres10:
	docker build database -t golang-grpc-db && docker run -d -p 54321:5432 -e POSTGRES_PASSWORD=postgres golang-grpc-db
