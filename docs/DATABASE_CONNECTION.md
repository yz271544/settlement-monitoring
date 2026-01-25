# Database Connection Configuration

## PostgreSQL

```
Host: localhost
Port: 35432
Database: settlement
Username: postgres
Password: STmt0125
SSL Mode: disable
```

### Connection String

```
postgres://postgres:STmt0125@localhost:35432/settlement?sslmode=disable
```

### Go GORM DSN

```go
host=localhost
port=35432
user=postgres
password=STmt0125
dbname=settlement
sslmode=disable
TimeZone=Asia/Shanghai
```

### JDBC URL (Java)

```
jdbc:postgresql://localhost:35432/settlement?user=postgres&password=STmt0125
```

### psql Command

```bash
docker exec -it settlement-monitoring-pg psql -U postgres -d settlement
```

### Environment Variables

```bash
export DB_HOST=localhost
export DB_PORT=35432
export DB_NAME=settlement
export DB_USER=postgres
export DB_PASSWORD=STmt0125
```

---

## Redis

```
Host: localhost
Port: 36379
Password: (empty)
Database: 0
```

### Connection String

```
redis://localhost:36379/0
```

### Go Redis DSN

```go
地址: localhost:36379
密码: (空)
DB: 0
```

### redis-cli Command

```bash
redis-cli -p 36379
```

### Environment Variables

```bash
export REDIS_HOST=localhost
export REDIS_PORT=36379
export REDIS_PASSWORD=
export REDIS_DB=0
```

---

## Quick Test Commands

### PostgreSQL Test

```bash
# Test connection
docker exec settlement-monitoring-pg psql -U postgres -d settlement -c "SELECT version();"

# List all tables
docker exec settlement-monitoring-pg psql -U postgres -d settlement -c "\dt"

# Count rows
docker exec settlement-monitoring-pg psql -U postgres -d settlement -c "SELECT COUNT(*) FROM t_staff;"

# Insert test data
docker exec settlement-monitoring-pg psql -U postgres -d settlement -c "INSERT INTO t_staff (id, name, type, sex, tel, personnel_id) VALUES (gen_random_uuid()::varchar(32), '张三', 'staff', '1', '13800138000', '110101199001011234');"

# Query test data
docker exec settlement-monitoring-pg psql -U postgres -d settlement -c "SELECT * FROM t_staff LIMIT 5;"
```

### Redis Test

```bash
# Test connection
redis-cli -p 36379 ping

# Set value
redis-cli -p 36379 SET test_key "hello"

# Get value
redis-cli -p 36379 GET test_key

# Delete value
redis-cli -p 36379 DEL test_key

# List all keys
redis-cli -p 36379 KEYS "*"
```

---

## Example Application Configurations

### Go (Gin-Admin) config.yaml

```yaml
database:
  default:
    driver: postgres
    host: localhost
    port: 35432
    database: settlement
    username: postgres
    password: STmt0125
    options:
      sslmode: disable
      TimeZone: Asia/Shanghai
      
redis:
  default:
    address: localhost:36379
    password: ""
    db: 0
```

### Spring Boot application.yml

```yaml
spring:
  datasource:
    url: jdbc:postgresql://localhost:35432/settlement?useSSL=false&serverTimezone=Asia/Shanghai
    username: postgres
    password: STmt0125
    driver-class-name: org.postgresql.Driver
    
  redis:
    host: localhost
    port: 36379
    password:
    database: 0
```

### Docker Compose

```yaml
services:
  app:
    environment:
      - DB_HOST=host.docker.internal
      - DB_PORT=35432
      - DB_NAME=settlement
      - DB_USER=postgres
      - DB_PASSWORD=STmt0125
      - REDIS_HOST=host.docker.internal
      - REDIS_PORT=36379
```