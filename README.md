# PSGA METADATA READER

A Cobra based CLI †ool to parse attributes out of CSV files.

## Running in dev

```
⇒  go run . get -c ./data/config.csv -n pypgx
```

## Installing

```
go build
go install

psga-metadata-reader -c data/config.csv -n pypgx
> /app/reference/pypgx-bundle

```

## Help

```
psga-metadata-reader --help
```

-c, --config = path to config file
-n, --name = name of resource row to fetch
