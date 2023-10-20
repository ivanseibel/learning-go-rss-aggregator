# RSS Aggregator API

## How to run migrations

Up
```bash
$ goose postgres postgres://<USER>:<PASSWORD>@localhost:5432/rssagg up
```
Down
```bash
$ goose postgres postgres://<USER>:<PASSWORD>@localhost:5432/rssagg down
```