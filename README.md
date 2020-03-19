Gomicro
=======

This project has the objective of being a skeleton for new micro services projects

# What is this?

Gomicro sets a basic structure of a project using the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).  
With this skeleton you can create new APIs, where it will have [echo](https://echo.labstack.com/) installed already.

## Basic structure

* `appcontext/context.go`: is our `container` of services and object within the project.
* `config/`: defines all necessary configuration that your application needs to run.
* `controller/`: defines all controllers of the project.
* `domain/`: defines all entities of the project.
* `gateway/`: defines all connections with the external world, e.g. MongoDB, MySQL.
* `usecase/`: defines the use cases and business rules of the application.
* `version/`: defines embed version of the project.

# Build

```sh
go build
```

# Build with embed version

Run `build.sh` script to set embed version and it used build info using commit hash.  

```sh
bash build.sh 1.0.0
```

Example:
```
curl http://localhost:8080/gomicro/v1/health | jq .
{
  "status": "UP",
  "version": "0.0.1",
  "build_info": "4e27105efd550ffbf8e225673a5a8fb9bfceb859"
}
```
