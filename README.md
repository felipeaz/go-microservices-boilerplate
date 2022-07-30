# microservices-boilerplate
Microservices boilerplate built using Hexagonal architecture, DDD principles &amp; SOLID.

It was built in a Monorepo structure where all the services are part of a big repository, even though they're 
still running independently.

## Dependencies Packages

### Tests
#### Ginkgo: Test Framework
Ref: https://github.com/onsi/ginkgo/v2
#### Gomega: Assertion
Ref: https://github.com/onsi/gomega

### Mocks
#### Testify: Mocks implementation
Ref: https://github.com/stretchr/testify
#### Vektra Mockery: Mocks Generator
Ref: https://github.com/vektra/mockery

### GORM
#### Entities & DB operations
Ref: https://gorm.io/

### Redis
#### Cache management
Ref: https://github.com/gomodule/redigo/redis

### UUID
#### Secure ID generator
Ref: https://github.com/satori/go.uuid

### Cors Middleware
#### Enable API to external origins
Ref: https://github.com/itsjamie/gin-cors

### API Documentation
#### Gin-Swagger: Swagger API Definitions
Ref: https://github.com/swaggo/gin-swagger
#### Swaggo Files: Doc files manager
Ref: https://github.com/swaggo/files
#### Swaggo: Doc Generator
Ref: https://github.com/swaggo/swag/cmd/swag