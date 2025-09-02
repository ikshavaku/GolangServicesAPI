# Kong Technical Assignment

## Problem Statement

https://docs.google.com/document/d/1K2PcGH_QTGd7PYwlb4Y5mRGovZ4vrbiu0uVCxF1Jrnw/edit?tab=t.0

## Solution

### Stack

- Backend : Golang
    - Http Server : Gin
    - DB Integration : sqlc
    - Build tool: Wire
    - DB migration : goose
- Database : Postgres

### Features

- **Vertical Slicing** : APIs are segregated using vertical slicing to ensure modular code which is loosely coupled and independent.

- **Layered Approach** : API is divided into multiple layers i.e API-->Service-->Datastore(DB) each layer managing it's own scope.

- **Dependency Injection** : Extensively used to ensure contract based development where every layer expects a contract to be followed by the downstream layer and the implementation is abstracted out. Wire is used to tie the dependencies togather.

### Steps to run the project

```shell
source local.envrc (export config via env)
make goose-up (runs migrations)
make run-server (build and run project)
```

### Decisions

#### Why use sqlc

- Fast to spin up
- Generates go models & queries based on sql hence task focused.

### API Spec

- Base Path : `/v1/service`

- Endpoints
    - `GET /` : List all services
    - `GET /:id` : List service by id
    - `GET /:id/versions` : List all versions for service

### Task Summary

#### Compelted

- Returning a list of services
    - support filtering (By name)
    - sorting (by name)
    - pagination
- Fetching a particular service
    - including a method for retrieving its versions
- Include tests (unit, integration) or a test plan
    - Integration
    - Unit test for Service Layer
- Provide authentication/authorization on the API : Not Done
- Add CRUD operations to the API
    - Read : Done
    - Create,Update,Delete : Not Done