# Health Statistic API

## API Documentation

> Navigate to [the API page in our domain](https://health-statistic.dechnology.com.tw/swagger/index.html#/Health/get_health_check) or [our GitHub page](https://eesoymilk.github.io/health-statistic-api/) to view all API endpoints.

## Before you commit

If you have made changes to the router.go (API endpoints), please also add declarative comments according to the standard from [`swaggo/swag`](https://github.com/swaggo/swag#declarative-comments-format).
After so, you want to also update the API documentation before making a commit. You can update the documentation by following the below steps.

### Install `go` and `swag` if you haven't already

- [Go Installation](https://go.dev/doc/install)
- [`swaggo/swag`](https://github.com/swaggo/swag#declarative-comments-format)
    ```shell
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

### Generate new API documentations

```shell
swag fmt # This line is optional since it only helps formatting the comments
swag init
```

## Run the API locally

### Step 1 - Setup `.env`

Setup a `.env` file in the project directory according to `.env.example`

```
POSTGRES_PORT=<POSTGRES_PORT>
POSTGRES_USER=<POSTGRES_USER>
POSTGRES_DB=<POSTGRES_DB>
POSTGRES_PASSWORD=<POSTGRES_PASSWORD>

AUTH0_ISSUER_URL=<AUTH0_ISSUER_URL>
AUTH0_AUDIENCE=<AUTH0_AUDIENCE>
```
> Note that we don't need to provide POSTGRES_HOST since it's taken care of in `docker-compose.yml`.

### Step 2 - Run the app via `docker compose` with the `--env-file`

```shell
docker compose --env-file .env up
```
