
user-service:
	cd user-service && go run main.go
container-up:
	docker compose up -d
container-down:
	docker compose down
sync-gateway:
	deck gateway sync kong.yaml
update-gateway:
	deck gateway dump -o kong.yaml 
.PHONY: user-service container-up container-down sync-gateway update-gateway