test-frontend:
	@echo "🧪 TESTING FRONTEND 🧪"
	cd frontend/RESTFULPiFrontend && npm run test:ci && echo "✅ TESTS PASSED ✅"

build-frontend: 
	@echo "👷 BUILDING FRONTEND 👷"
	cd frontend/RESTFULPiFrontend && npm run build:ci && echo "✅ BUILD COMPLETE ✅"

test-backend: 
	@echo "🧪 TESTING BACKEND 🧪"
	cd backend && go test ./... && echo "✅ TESTS PASSED ✅"

build-backend:
	@echo "👷 BUILDING BACKEND 👷"

build: build-frontend build-backend

ship: 

deploy: build ship 

