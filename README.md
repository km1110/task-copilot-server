# task-copilot-server

### docker compose 環境のビルド

```
docker compose build --no-cache
```

### docker compose 環境の立ち上げ

```
docker compose up
```

### docker compose 環境への接続

```
docker compose exec -it api /bin/sh
```

### db への接続

```
docker compose exec -it postgresql bash
psql -d task-copilot-db -U user
```
