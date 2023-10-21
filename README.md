# RSS Aggregator API

## How to run migrations

Up
```bash
$ cd sql/schema
$ goose postgres postgres://<USER>:<PASSWORD>@localhost:5432/rssagg up
```
Down
```bash
$ cd sql/schema
$ goose postgres postgres://<USER>:<PASSWORD>@localhost:5432/rssagg down
```

## How to create go code for models

```bash
$ sqlc generate
```