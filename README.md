This is Authentication using gin and gorm using jwt token authentication
docker run --name api-container -p 8080:8080 --network test -e GIN_MODE=release test-api:latest