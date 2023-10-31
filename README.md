# Simple go grpc server

Toy project to experiment with some golang technologies. The project includes a [GRCP](https://grpc.io/) server generated from .proto, a multi-module structure with dependecy injection handled at compile time by [wire](https://github.com/google/wire) a build pipeline handled using [make](https://www.gnu.org/software/make/) (Evaluate bazel for the future) and a built in gui for manually invoking service via [grpcui](https://github.com/fullstorydev/grpcui) and some unit test.

## Useful commands

Build the project:
``` bash
make
```

Clean the build directory
``` bash
make clean
```

Generate grpc and wire src
``` bash
make code_gen
```

Run unit test
``` bash
make test
```

Strat the web ui to test invokations
``` bash
make grpc_ui
```
