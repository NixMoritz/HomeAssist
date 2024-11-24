run:
		go run cmd/HomeAssist/main.go

tidy:
		go mod tidy

up:
		docker-compose up

down:
		docker-compose down

build:
		go build -o homeassist backend/main.go

clean:
		rm -f homeassist
