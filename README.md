# Health Statistic API

## Usage

### Step 1 - Setup `.env`

Setup a `.env` file in the project directory according to `.env.example`

```
POSTGRES_HOST=<POSTGRES_HOST>
POSTGRES_PORT=<POSTGRES_PORT>
POSTGRES_USER=<POSTGRES_USER>
POSTGRES_DB=<POSTGRES_DB>
POSTGRES_PASSWORD=<POSTGRES_PASSWORD>
PGDATA=<PGDATA>

AUTH0_ISSUER_URL=<AUTH0_ISSUER_URL>
AUTH0_AUDIENCE=<AUTH0_AUDIENCE>
```

### Step 2 - Run the app via `docker-compose` with the `--env-file`

```shell
docker compose --env-file .env up
```

## API Documentation

> Navigate to [our GitHub page](https://eesoymilk.github.io/health-statistic-api/) to view all API endpoints.
