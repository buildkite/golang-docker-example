# Buildkite Golang Docker Example

This repository is an example on how to test a Golang project through Docker
using Buildkite.

[Add this example to your Buildkite organization](https://buildkite.com/new)

## Using in your own build pipelines

1. Ensure `docker-compose` is installed on your build system. For details on how to do this, see: https://docs.docker.com/compose/install/

2. Use our `Dockerfile` and `docker-compose.yml` files as defaults:

   ```sh
   cd /your/golang/repo
   curl -o Dockerfile https://raw.githubusercontent.com/buildkite/golang-golang-example/master/Dockerfile
   curl -o docker-compose.yml https://raw.githubusercontent.com/buildkite/golang-golang-example/master/docker-compose.yml
   ```

3. Replace `/go/src/github.com/buildkite/golang-docker-example` in the `Dockerfile` and
   `docker-compose.yml` files with your own Golang import path.  For example,
   if your import path in Golang looks like this:

   ```go
   import (
     "github.com/keithpitt/project/sub-package"
   )
   ```

   Then you would replace `/go/src/github.com/buildkite/golang-docker-example`
   with `/go/src/github.com/keithpitt/project` (note the `sub-package` part of
   the import is not included). This path should also match the directory
   structure within the `$GOPATH` on your own development machine.

4. Add to your build pipeline and add the `BUILDKITE_DOCKER_COMPOSE_CONTAINER` env:

   ```yml
   steps:
     - command: "./scripts/test.sh"
       env:
         BUILDKITE_DOCKER_COMPOSE_CONTAINER: "app"
   ```

## License

See [Licence.md](Licence.md) (MIT)
