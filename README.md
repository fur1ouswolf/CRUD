## CRUD (Create, Read, Update, Delete) Operations with Go (Gin, GORM) and PostgreSQL


### Overview
This is a simple CRUD application written in Go using Gin and GORM. It uses PostgreSQL as the database and Docker to run the application.

### Technical Assignment
##### Tables
- Regions: id, name
- Persons: id, forename, surname, patronymic, gender, region_id

##### Endpoints
- GET /region - returns a list of regions
- POST /region - creates a region
- DELETE /region/:id - deletes a region
- GET /person - returns a list of persons
- POST /person - creates a person
- GET /person/:id - returns a person
- DELETE /person/:id - deletes a person
- POST /person/:id - updates a person
- GET /person/region/:id - returns a region of a person

### Prerequisites
- Go 1.21
- PostgreSQL
- Docker

### How to run
- Clone this repository
- Run `make run` to start the application
