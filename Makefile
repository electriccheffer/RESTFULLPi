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
	cd backend && go build -o ../bin/restfulPi ./cmd/main.go

build: build-frontend build-backend

test: test-backend

ship: 

deploy: build ship 

