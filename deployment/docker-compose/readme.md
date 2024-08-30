# Run with Docker Compose
* fix .env file

```bash
   cp .env.example .env
```

* deploy

```bash
   # deploy infrastructure services
   docker compose -f ./infrastructure.yaml  up -d
   # deploy self services
   docker compose up -d
```


