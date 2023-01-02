# Get started

## Run the game

The game can be run by the running
```bash
$ go run main.go
```
This will run the game with basic game inputs, 2 aliens and one of the example map files.

The app takes in different arguments:

| Short flag | Description | Example |
|------------|-------------|---------|
| `-n`       | Number of invading aliens | `-n 2`|
| `-p`       | Path to map file | `-p example/map1.txt`|

## Run tests

```bash
$ make test
```

## Run linter

```bash
make linter
```
