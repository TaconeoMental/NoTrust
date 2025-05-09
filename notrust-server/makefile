include .env
export $(shell sed 's/=.*//' .env)

# --------------------------
# Servidores
# --------------------------

run-server:
	@echo ">>> Corriendo Servidor Público en http://localhost:$(SERVER_PORT)"
	go run ./cmd/server/main.go

run-admin:
	@echo ">>> Corriendo Admin Panel en http://localhost:$(ADMIN_SERVER_PORT)"
	go run ./cmd/admin/main.go

build-server:
	@echo ">>> Compilando Servidor Público..."
	go build -o $(BINARY_PUBLIC) ./cmd/server

build-admin:
	@echo ">>> Compilando Admin Panel..."
	go build -o $(BINARY_ADMIN) ./cmd/admin

# --------------------------
# Base de Datos
# --------------------------

migrate:
	export $$(cat .env | xargs) && for file in $$(ls db/migrations/*.sql | sort); do \
		echo ">>> Ejecutando $$file"; \
		PGPASSWORD=$$DB_PASSWORD psql -h $$DB_HOST -p $$DB_PORT -U $$DB_USER -d $$DB_NAME -f $$file; \
	done
# --------------------------
# Docker
# --------------------------

docker-up:
	@echo ">>> Levantando stack Docker..."
	docker-compose up --build

docker-down:
	@echo ">>> Bajando stack Docker..."
	docker-compose down

docker-logs:
	@echo ">>> Logs Docker..."
	docker-compose logs -f

# --------------------------
# Otros
# --------------------------

clean:
	@echo ">>> Limpiando binarios..."
	rm -f $(BINARY_PUBLIC) $(BINARY_ADMIN)


