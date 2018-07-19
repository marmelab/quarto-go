# quarto-go
Quarto game server using golang

## Install

Install docker image

```
make install
```

## Run

Run the game

```
make run
```

## Test

Runs all the tests of the project

```
make test
```

## Lint

Runs code linter on the project

```
make lint
```

## Deploy

Deploy project on server
  -Add ssh parameter to specify distant connection name (like make deploy sshname=quartoServer)

```
make deploy
```
