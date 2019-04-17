<table>
        <tr>
            <td><img width="120" src="https://cdnjs.cloudflare.com/ajax/libs/octicons/8.5.0/svg/rocket.svg" alt="onboarding" /></td>
            <td><strong>Archived Repository</strong><br />
            The code of this repository was written during a <a href="https://marmelab.com/blog/2018/09/05/agile-integration.html">Marmelab agile integration</a>. <a href="https://marmelab.com/blog/2018/10/09/go-go-quarto-ranger.html">The associated blog post</a> illustrates the efforts of the new hiree, who had to implement a board game in several languages and platforms as part of his initial learning.<br />
        <strong>This code is not intended to be used in production, and is not maintained.</strong>
        </td>
        </tr>
</table>

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
