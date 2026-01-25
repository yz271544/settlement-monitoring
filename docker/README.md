# START POSTGRESQL FOR CPS-DEV

## init db tables
# START POSTGRESQL FOR CPS-DEV

```shell
docker run -d \
    --name gin-admin-post \
    -e POSTGRES_PASSWORD=comPostgres123 \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v /lyndon/iData/gin-admin-post/data:/var/lib/postgresql/data \
    -v /lyndon/iData/gin-admin-post/init:/docker-entrypoint-initdb.d \
    -p 35432:5432 \
    registry.cn-beijing.aliyuncs.com/dc_huzy/postgres:alpine3.20
```
