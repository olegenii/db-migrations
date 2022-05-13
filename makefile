migrateup:
	migrate -path db/migration -database "postgresql://api:GQt5MTyVPuf9vsVWoWDT9YCn@oleg-web.devops.rebrain.srwx.net:5432/api_test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://api:GQt5MTyVPuf9vsVWoWDT9YCn@oleg-web.devops.rebrain.srwx.net:5432/api_test?sslmode=disable" -verbose down

.PHONY: migrateup migratedown