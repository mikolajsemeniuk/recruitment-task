# Recruitment Task

- Explanation
- Motivation
- Run web API
- Running tests
- Running linter

## Explanation

- **`README.md`**: Contains project documentation, including this structure overview.
- **`cmd/`**: Holds executable commands:
  - `web/`: Command for starting the web server.
- **`go.mod`** & **`go.sum`**: Dependency management files.
- **`pkg/`**: Core business logic and package modules:
  - `docs/`: Serves API documentation and related templates.
  - `index/`: Manages index finder service with it's logic, including in-memory storage implementation.
- **`recruitment_task.md`**: Contains project-related tasks or requirements.

## Motivation

Because this is a recruitment task I decided to put extra comments in code to make this clearly intend purpose of implemented things.

I decided to implement this service as easy as possible because most of programmers work is to read somebody's code and I
don't want to keep my code over engineered and hard to read to other developers. Let's keep things simple.

I've chosen to put `main.go` which is entry file for web API in `cmd/web/main.go` and here it's why.
From my experiences of deploying applications in k8s it's very likely that applications usually have
containers, sidecar containers or initContainer for making applications related stuff like which uses the same codebase from pkg folder like:

- database migrators (`cmd/migrator/main.go`)
- worker retrieving messages from queue (`cmd/worker/main.go`)
- cronjobs generating files to analytics (`cmd/cronjob/main.go`)
- sidecar loggers (`cmd/logger/main.go`)

so `cmd` folder gives us more flexibility for dockerizing and deploying stuff instead of keeping single main.go in root of the project.

I implemented benchmark to Find function embedded in memory struct to search for best performant solution and decided to use BinarySearch.

I also decided to use as little as possible libraries which are not the part of std library and here is why:

Usually I try to name packages in `pkg` folder names like: `auth`, `cart`, `wishlist` but because this is recruitment process name `index` suits best for me.

- less binary and docker image size
- less external dependencies are very often less image vulnerabilities to fix
- from version 1.6 golang net/http is much more faster and has a better performance than frameworks like gin, echo, fiber etc because it incorporates http/2

## Run web API

```sh
make tidy
make run
make watch # use it for hot reload, make sure you have go install github.com/cosmtrek/air@latest installed on your local machine

# then visit http://localhost:8080/#/
# or curl "http://localhost:8080/index/find/300"
```

## Running tests

```sh
make test
```

## Running linter

```sh
make lint
```
