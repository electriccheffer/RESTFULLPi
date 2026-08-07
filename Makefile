-include .env

test-frontend:
	@echo "🧪 TESTING FRONTEND 🧪"
	cd frontend/RESTFULPiFrontend && npm run test:ci && echo "✅ TESTS PASSED ✅"

build-frontend: 
	@echo "👷 BUILDING FRONTEND 👷"
	cd frontend/RESTFULPiFrontend && npm run build:ci && echo "✅ BUILD COMPLETE ✅"

test-backend: build-frontend
	@echo "🧪 TESTING BACKEND 🧪"
	cd backend && go test ./... && echo "✅ TESTS PASSED ✅"

build-backend: test-backend
	@echo "👷 BUILDING BACKEND 👷"
	cd backend && GOOS=linux GOARCH=arm64 go build -o ../bin/restfulPi ./cmd/main.go

build: build-frontend build-backend

test: test-backend

connect:
	@echo "🛜 CONNECTING TO PiGPSNet 🛜"
	networksetup -setairportnetwork en0 "$(PI_NETWORK)" "$(PI_PASSWORD)"
send:
	@echo "🚀 Sending RestfulPi binary to server 🚀"
	ssh piThree "sudo systemctl stop $(SERVICE_NAME)"
	scp ./bin/restfulPi piThree:~/

service:
	@echo "🚚 RESTARTING SERVICE 🚚"
	ssh piThree "sudo systemctl daemon-reload && sudo systemctl restart $(SERVICE_NAME)"	

browser:
	@echo "🦁 OPENING SAFARI 🦁"
	open http://$(PI_HOST):8080
	
switch:
	@echo "🛜 CONNECTING TO PalmyraCatColony 🛜"
	networksetup -setairportnetwork en0 "$(NORMAL_NETWORK)" "$(NORMAL_PASSWORD)"

ship: connect send service browser

deploy: build ship 

