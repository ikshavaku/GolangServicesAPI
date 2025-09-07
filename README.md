# Kong Technical Assignment

## Problem Statement

https://docs.google.com/document/d/1K2PcGH_QTGd7PYwlb4Y5mRGovZ4vrbiu0uVCxF1Jrnw/edit?tab=t.0

## Solution

# 📡 Golang Services API  

This project implements a **Services API** for managing and discovering services within an organization. It is designed to support a dashboard widget where users can:  
- View a list of services (with filtering, sorting, pagination).  
- See details of a specific service.  
- Retrieve versions of a given service.  
- Search for services by name or description.  

The API is built in **Go** using modern tooling and structured for production-readiness.  

---

## ✨ Features  

- **Service Overview**  
  - Name, description, and versions of services.  
- **Filtering**  
  - Query parameters for flexible retrieval.  
- **Pagination**  
  - Efficient pagination for large datasets.  
- **Search**  
  - Search services by name or description.  
- **Service Details**  
  - Retrieve details of a specific service.  
- **Versions Endpoint**  
  - Fetch versions available for a given service.  
- **Persistence**  
  - PostgreSQL for reliable storage.  
- **Migrations**  
  - Version-controlled schema migrations with Goose.  
- **Dependency Injection**  
  - Wire for clean service wiring.  
- **Docker Support**  
  - Run API and DB via Docker Compose.  
- **Vertical Slicing**
  - APIs are segregated using vertical slicing to ensure modular code which is loosely coupled and independent.
- **Layered Approach**
  - API is divided into multiple layers i.e API-->Service-->Datastore(DB) each layer managing it's own scope.
- **JSON Logging Enabled**
- **Sorting**
  - Default sorting by name

---

## 📂 Project Structure  

```
.
├── cmd/api/ # Application entrypoint
├── internal/
│ ├── api/ # HTTP handlers, routing
│ ├── service/ # Business logic
│ ├── store/ # Database access (sqlc-generated)
│ └── version/ # Version-related logic
├── migrations/ # Goose migrations
├── Dockerfile # Build container
├── docker-compose.yml# Run DB + API together
├── Makefile # Common tasks (build, run, test, migrate)
└── README.md # Project documentation
```

### Stack

- Backend : Golang
    - Http Server : [Gin](https://github.com/gin-gonic/gin)
    - DB Integration : [sqlc](https://github.com/sqlc-dev/sqlc)
    - Build tool: [Wire](https://github.com/google/wire)
    - DB migration : [goose](https://github.com/pressly/goose)
- Database : Postgres

### Steps to run the project with db in container

```shell
docker-compose up
```

### Decisions

#### Why use sqlc

- Fast to spin up
- Generates go models & queries based on sql hence task focused.

### API Spec

please refer `swagger.yml`

### Task Summary

#### Completed

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
