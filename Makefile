run:
		go run cmd/HomeAssist/main.go

build:
		go build -o homeassist backend/main.go

clean:
		rm -f homeassist
