run:
	go run cmd/HomeAssist/main.go

tidy:
	go mod tidy

up:
	docker-compose up

front:
	cd HomeAssist && npm run dev

down:
	docker-compose down

build:
	go build -o homeassist backend/main.go

clean:
	rm -f homeassist
