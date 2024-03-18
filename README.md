<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
# go-clean-architecture

<h1 align='center'>
    ⚡ go-clean-architecture
</h1>
<h4 align='center'>
    GoFiber code structure for developing using clean architecture
</h4>

This structure, created following the development guide's for vertical slice architecture, will help to isolate the dependencies, make development easier and have a cleaner and testable code in every package.

## 👩‍💻 Authors

#### Sebastián Ituarte

- [@sebajax](https://www.github.com/sebajax)

## 🖍 Clean architecture

Clean architecture is an approach to software development where code and functionality are organized around individual features or user stories, encompassing all layers of the application from user interface to data access, promoting autonomy, reduced dependencies, and iterative development.

![alt text](./image/clean-architecture.webp)

## 📚 Code Structure

A brief description of the layout:

- `.github` has two template files for creating PR and issue. Please see the files for more details.
- `.gitignore` varies per project, but all projects need to ignore `bin` directory.
- `.golangci.yml` is the golangci-lint config file.
- `Makefile` is used to build the project. **You need to tweak the variables based on your project**.
- `CHANGELOG.md` contains auto-generated changelog information.
- `README.md` is a detailed description of the project.
- `cmd` contains the main.go file that is our starting point to execute
- `pkg` places most of project business logic.
- `migrations` contains all vendored code.
- `internal` contains all the api logic.
- `domain layer`: Contains the business logic and entities.
- `service layer`: Contains the application-specific logic and use cases.
- `infrastructure layer:` Deals with external dependencies and frameworks.
- `interface layer:` Handles communication with external systems or UI.

## 🚀 Stack

#### Programming language [Go](https://go.dev/)

#### Framework [Fiber](https://docs.gofiber.io/)

#### Dependency injection [Uber dig](https://github.com/uber-go/dig)

#### Database [Postgre SQL](https://www.postgresql.org/)

#### Container [Docker](https://www.docker.com/)

#### Live reload [Air](https://github.com/cosmtrek/air)

## 🧐 This app uses conventional commits

[Conventional commits url](https://www.conventionalcommits.org/en/v1.0.0/)

## 🤜 How to create a new use case (Example)

```http
  POST /api/product
```

| Parameter  | Type     | Description                              |
| :--------- | :------- | :--------------------------------------- |
| `name`     | `string` | **Required**. Product Name               |
| `sku`      | `string` | **Required**. Product Sku must be Unique |
| `category` | `string` | **Required**. Product Category           |
| `price`    | `float`  | **Required**. Product Price              |

### Internal folder structure for a new domain all folders and files go in internal/product/

```tree
├───internal
│   ├───product
│   │   │   port.go
│   │   │   product.go
│   │   │
│   │   ├───handler
│   │   │       createProduct.go
│   │   │       handler.go
│   │   │
│   │   ├───infrastructure
│   │   │       productRepository.go
│   │   │
│   │   ├───mock
│   │   │       mockProductRepository.go
│   │   │
│   │   └───service
│   │           createProduct.go
│   │           service.go
```

#### 1 - Create product.go (domain)

https://github.com/sebajax/go-vertical-slice-architecture/blob/d4501917930ef2263551bee3ee529de49b6d6fc5/internal/product/product.go#L1-L58

#### 2 - Create infrastructure/productRepository.go (DB query)

https://github.com/sebajax/go-vertical-slice-architecture/blob/872df7def565c7e0a95216b0b93d8166c3912ef5/internal/product/infrastructure/productRepository.go#L1-L61

```tree
├───internal
│   ├───product
│   │   ├───infrastructure
│   │   │       productRepository.go
```

#### 3 - Create port.go (dp injection for the service)

https://github.com/sebajax/go-vertical-slice-architecture/blob/872df7def565c7e0a95216b0b93d8166c3912ef5/internal/product/port.go#L1-L7

#### 4 - Create service/createProduct.go (create a new product use case implementation)

```tree
├───internal
│   ├───product
│   │   └───service
│   │           createProduct.go
│   │           service.go
```

https://github.com/sebajax/go-vertical-slice-architecture/blob/872df7def565c7e0a95216b0b93d8166c3912ef5/internal/product/service/createProduct.go#L1-L51

#### 5 - Create service/service.go (dependency injection service using uber dig)

```tree
├───internal
│   ├───product
│   │   └───service
│   │           createProduct.go
│   │           service.go
```

https://github.com/sebajax/go-vertical-slice-architecture/blob/872df7def565c7e0a95216b0b93d8166c3912ef5/internal/product/service/service.go#L1-L42

#### 6 - Create handler/createProduct.go (create a new hanlder presenter to call the use case)

```tree
├───internal
│   ├───product
│   │   ├───handler
│   │   │       createProduct.go
│   │   │       handler.go
```

https://github.com/sebajax/go-vertical-slice-architecture/blob/872df7def565c7e0a95216b0b93d8166c3912ef5/internal/product/handler/createProduct.go#L1-L60

#### 7 - Create handler/handler.go (handles all the routes for all the presenters)

```tree
├───internal
│   ├───product
│   │   ├───handler
│   │   │       createProduct.go
│   │   │       handler.go
```

https://github.com/sebajax/go-vertical-slice-architecture/blob/872df7def565c7e0a95216b0b93d8166c3912ef5/internal/product/handler/handler.go#L1-L11

#### 8 - Create mock/ProductRepository.go (mocks the user repository implementation for unit testing)

```tree
├───internal
│   ├───product
│   │   ├───mock
│   │   │       mockProductRepository.go
```

```bash
    # It uses mockery Run
        mockery
```

#### 9 - Create service/createProduct_test.go (create a unit test for the service)

```tree
├───internal
│   ├───product
│   │   └───service
│   │           createProduct.go
│   │           service.go
```

https://github.com/sebajax/go-vertical-slice-architecture/blob/9ff7ae658ee5bada18b88ec4a03994a105ccbe02/internal/product/service/createProduct_test.go#L1-L97

#### 10 - Add dependency injection service using uber dig

```tree
├───pkg
│   ├───injection
```

https://github.com/sebajax/go-vertical-slice-architecture/blob/eb79ccae805d23b6f77385a5f7ebfc81bb6174e0/pkg/injection/injection.go#L1-L73

## ⚙️ Usage without Make

### Docker usage

```bash
    # Build server
        docker-compose -p go-vertical-slice-architecture build

    # Start server
        docker-compose up -d

    # Stop server
        docker-compose down
```

### Standalone usage

```bash
    # Live reload
        air
```

### Testing

```bash
    # To run unit testing
        go test ./...

    # To run unit testing coverage
        go test -cover ./...
```

### Formatting, Linting and Vetting

```bash
    # Clean dependencies
        go mod tidy

    # Run formating
        go fmt ./...

    # Remove unused imports
        goimports -l -w .

    # Run linting
        golangci-lint run ./...

    # Run vetting
        go vet ./...

    # Run shadow to check shadowed variables
        # Install shadow
        go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
        # Run shadow
        shadow ./...
```

### Database migration script

```bash
    # Create the script
        migrate create -ext sql -dir /migrations -seq [script_name]
    # Run the script
        migrate -database ${POSTGRESQL_URL} -path /migrations up

    # It will run automatically when the database initializes
```

## ⚙️ Usage with Make

### Docker usage

```bash
    # Build server
        make build-server

    # Start server
        make start-server

    # Stop server
        make stop-server
```

### Standalone usage

```bash
    # Live reload
        make live-reload
```

### Testing

```bash
    # To run unit testing
        make test

    # To run unit testing coverage
        make test-coverage
```

### Formatting, Linting and Vetting

```bash
    # Clean dependencies
        make clean-deps

    # Run formating
        make format

    # Remove unused imports
        make clean-imports

    # Run linting
        make lint

    # Run vetting
        make vet

    # Run shadow to check shadowed variables
        # Install shadow
        go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
        # Run shadow
        make check-shadow

    # Run vetting to lint, format and vet your once
        make lint-format
```

### Database migration script

```bash
    # Create the script (replace your_script_name with the actual name)
        make migrate-create name=your_script_name
    # Run the script
        make migrate-up

    # It will run automatically when the database initializes
```

## 💻 Environment variables

To modify/add configuration via environment variables, use the `.env` file, which contains basic app configuration.
