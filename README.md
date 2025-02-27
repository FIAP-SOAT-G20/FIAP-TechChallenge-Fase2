<a name="readme-top"></a>

# <p align="center">FIAP Tech Challenge 2 - G18 Fast Food</p>

<p align="center">
    <img src="https://img.shields.io/badge/Code-Go-informational?style=flat-square&logo=go&color=00ADD8" alt="Go" />
    <img src="https://img.shields.io/badge/Tools-Gin-informational?style=flat-square&logo=go&color=00ADD8" alt="Gin" />
    <img src="https://img.shields.io/badge/Tools-PostgreSQL-informational?style=flat-square&logo=postgresql&color=4169E1" alt="PostgreSQL" />
    <img src="https://img.shields.io/badge/Tools-Swagger-informational?style=flat-square&logo=swagger&color=85EA2D" alt="Swagger" />
    <img src="https://img.shields.io/badge/Tools-Docker-informational?style=flat-square&logo=docker&color=2496ED" alt="Docker" />
    <img src="https://img.shields.io/badge/Tools-Kubernetes-informational?style=flat-square&logo=kubernetes&color=326CE5" alt="Kubernetes" />
</p>


<p align="center">
    <img src="docs/gopher.png" alt="Fast Food" />
</p>

## 💬 About

Repository for the [FIAP](https://postech.fiap.com.br/) Tech Challenge 2, focused on developing a monolith backend system for managing orders in a fast-food restaurant.

Tech Challenge 2 specifications can be found [here](docs/tc2-spec.pdf).

> Tech Challenge 1 repository can be found [here](https://github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1)

## 📚 Dictionary - Ubiquitous Language

- Customer (actor): Actor responsible for initiating the purchasing process
- Cook (actor): Actor responsible for preparing the customer's order
- Attendant (actor): Actor responsible for interacting with the customer, providing support for the order
- Identification method: Format in which the customer is identified on the platform: via CPF or anonymous.
- Identification: Customer identification on the platform
- Authorization: Grants permission to the customer to perform operations on the platform, such as placing an order, changing registration information
- Order: Represents all items selected by the customer in the store
- Order Status: Represents the stage of order preparation after payment is confirmed.

### :open_file_folder: Project Structure

```sh
.
├── cmd
│   └── server
├── docs
├── internal
│   ├── adapter
│   │   ├── controller
│   │   ├── gateway
│   │   └── presenter
│   ├── core
│   │   ├── domain
│   │   │   ├── entity
│   │   │   ├── value_object
│   │   ├── port
│   │   └── usecase
│   └── infrastructure
│       ├── config
│       ├── database
│       ├── datasource
│       ├── handler
│       ├── logger
│       ├── middleware
│       ├── route
│       └── server
└── k8s
```

<details>
<summary>Project Structure Explanation</summary>

### **1️⃣ Core (Innermost layer)**

- `domain/`: Central business entities and rules.
- `usecase/`: Application use cases.
- `port/`: Interfaces that define contracts between layers, ensuring independence.

### **2️⃣ Adapter (Middle layer)**

- `controller/`: Coordinates the flow of data between use cases and infrastructure.
- `presenter/`: Formats data for presentation.
- `gateway/`: Implements access to data from external sources (databases, APIs, etc.).

### **3️⃣ Infrastructure (External layer)**

- `config/`: Application configuration management.
- `database/`: Configuration and connection to the database. - `server/`: Initialization of the HTTP server.
- `route/`: Definition of API routes.
- `middleware/`: HTTP middlewares for handling requests.
- `logger/`: Structured logger for detailed logs.
- `handler/`: Handling of HTTP requests.
- `datasource/`: Concrete implementations of data sources.

</details>

### :bulb: Decisions

- **Clean Architecture structure**: The project was structured using the Clean Architecture pattern, which aims to separate the application into layers, making it easier to maintain and test. The project is divided into three layers: Core, Adapter, and Infrastructure.
- **Presenter**: The presenter (from Adapter layer) was created to format the data to be returned to the client. This layer is responsible for transforming the data into the desired format, such as JSON, XML, etc. Also, it is responsible for handling errors and returning the appropriate HTTP status code.
- **Use Case**: The use case (from Core layer) was created to define the business rules of the application. This layer is responsible for orchestrating the flow of data between the entities and the data sources.
- **Middleware to handle errors**: A middleware was created to handle errors and return the appropriate HTTP status code. This middleware is responsible for catching errors and returning the appropriate response to the client.
- **Structured Logger**: A structured logger was created to provide detailed logs. This logger is responsible for logging information about the application, such as requests, responses, errors, etc.
- **Database Connection**: The database connection was created using GORM, a popular ORM library for Go. This library provides an easy way to interact with the database and perform CRUD operations.
- **Database Migrations**: Database migrations were created to manage the database schema. This allows us to version control the database schema and apply changes to the database in a structured way.
- **HTTP Server**: The HTTP server was created using the Gin framework, a lightweight web framework for Go. This framework provides a fast and easy way to create web applications in Go.

### ✨ Features

- [x] Dockerfile: small image with multi-stage docker build, and independent of the host environment
- [x] Makefile: to simplify the build and run commands
- [x] Clean architecture
- [x] PostgreSQL database
- [x] Conventional commits

<details>
<summary>more (click to expand or collapse ↕️)</summary>

- [x] Unit tests (testify)
- [x] Code coverage report (go tool cover)
- [x] Swagger documentation
- [x] Postman collection
- [x] Feature branch workflow
- [x] Live reload (air)
- [x] Pagination
- [x] Health Check
- [x] Lint (golangci-lint)
- [x] Vulnerability check (govulncheck)
- [x] Mocks (gomock)
- [x] Environment variables
- [x] Graceful shutdown
- [x] Kubernetes deployment
- [x] GitHub Actions
- [x] Structured logs (slog)
- [x] Database migrations (golang-migrate)

</details>

## :computer: Technologies

- [Go 1.23+](https://golang.org/)
- [Gin Web Framework](https://gin-gonic.com/)
- [golangci-lint](https://golangci-lint.run/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)

<details>
<summary>more (click to expand or collapse ↕️)</summary>

- [gomock](https://github.com/uber-go/mock)
- [go-playground/validator](https://github.com/go-playground/validator)
- [godotenv](https://github.com/joho/godotenv)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Air](https://github.com/air-verse/air)
- [slog](https://pkg.go.dev/log/slog)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)
- [Swagger](https://swagger.io/)
- [Make](https://www.gnu.org/software/make/)
- [Testify](https://github.com/stretchr/testify)

</details>

## :scroll: Requirements

### Build/Run with Docker

- [Docker](https://www.docker.com/)

### Build/Run Locally (development)

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)

> [!NOTE]
> You need to have Go (> 1.23) installed in your machine to build, run and test the application locally

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :cd: Installation

```sh
git clone https://github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2.git
```

```sh
cd FIAP-TechChallenge-Fase2
```

Set the environment variables

```sh
cp .env.example .env
```

### :whale: Docker

```sh
make compose-build
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :runner: Running

### :whale: Docker

```sh
make compose-up
```

> To stop the application, run `compose-down`  
> To remove the application, run `compose-clean`  

> [!NOTE]
> You can check the application swagger documentation at http://localhost:8080/docs/index.html  
> Alternatively, a postman collection is available at [here](docs/10soat-g20-tech-challenge-1.postman_collection.json)  

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### :gear: Kubernetes

```bash
kubectl apply -f k8s/
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :hammer_and_wrench: Development

1. Install Go: https://golang.org/doc/install
2. Clone this repository: `git clone https://github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2`
3. Change to the project directory: `cd FIAP-TechChallenge-Fase2`
4. Checkout to the development branch: `make new-branch`
5. Set the environment variables: `cp .env.example .env`
6. Install dependencies by running `make install`
7. Run the application by running `make run-air` or `make run`
8. Access the application at `http://localhost:8080`
9. Dont forget to run the tests by running `make test`
10. Check the coverage report by running `make coverage`
11. Check the lint by running `make lint`
12. Update the swagger documentation by running `make swagger`
13. Commit your changes following the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) standard
14. Push to the branch and Open a new PR by running `make pull-request`
15. After the PR is approved, merge it to the main branch
16. The GitHub Actions will run the tests, lint, and build the Docker image
<!-- 17. The Kubernetes deployment will be updated automatically -->

> [!NOTE]
> `make run` will run the application locally, and will build and run PostgreSQL container using Docker Compose  
> Alternatively, you can run `make run-air` to run the application using Air (live reload)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :white_check_mark: Tests

```sh
make test
```

> [!NOTE]
> It will run the unit tests and generate the coverage report as `coverage.out`  
> You can check the coverage report by running `make coverage`  

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :clap: Acknowledgments

- [Hexagonal Architecture, Ports and Adapters in Go](https://medium.com/@kyodo-tech/hexagonal-architecture-ports-and-adapters-in-go-f1af950726b)
- [Building RESTful API with Hexagonal Architecture in Go](https://dev.to/bagashiz/building-restful-api-with-hexagonal-architecture-in-go-1mij)
- [Hexagonal Architecture in Go](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)
- [The S.O.L.I.D Principles in Pictures](https://medium.com/backticks-tildes/the-s-o-l-i-d-principles-in-pictures-b34ce2f1e898)
- [Health Check Response Format for HTTP APIs](https://datatracker.ietf.org/doc/html/draft-inadarei-api-health-check-06)
- [Event Storming](https://www.eventstorming.com/)
- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
- [Alistair in the "Hexagone" 1/3](https://www.youtube.com/watch?v=th4AgBcrEHA&list=PLGl1Jc8ErU1w27y8-7Gdcloy1tHO7NriL&ab_channel=DDDFR)
- [Alistair in the "Hexagone" 2/3](https://www.youtube.com/watch?v=iALcE8BPs94&list=PLGl1Jc8ErU1w27y8-7Gdcloy1tHO7NriL&index=2&ab_channel=DDDFR)
- [Alistair in the "Hexagone" 3/3](https://www.youtube.com/watch?v=DAe0Bmcyt-4&list=PLGl1Jc8ErU1w27y8-7Gdcloy1tHO7NriL&index=3&ab_channel=DDDFR)
- [Clean Architecture - Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [How to implement Clean Architecture in Go (Golang)](https://threedots.tech/post/introducing-clean-architecture)
- [Clean Architecture — A essência direto ao ponto](https://medium.com/@rvf.vazquez/clean-architecture-a-ess%C3%AAncia-draft-6d902e10d4b2)
- [Clean Architecture with Golang](https://medium.com/nerd-for-tech/clean-architecture-with-golang-3fa1a1c2b6d6)
- [Generate Go Project with Clean Architecture - Thiago Adriano (FIAP)](https://github.com/tadrianonet/go-clean-arch-generator)
- [POSTECH_SOAT_CleanArchitecture](https://github.com/FIAP/POSTECH_SOAT_CleanArchitecture/blob/main/aula_08/src/index.ts)
- [Fiap - Clean Architecture Usando Go - Erick Muller (FIAP)](https://github.com/proferickmuller/cleanarch-golang)
- [Clean Architecture, o início - Erick Muller (FIAP)](https://medium.com/@oerickmuller/clean-architecture-o-in%C3%ADcio-4fd74578155d)
- [Building Better Go Applications with Clean Architecture: A Practical Guide for Beginners](https://medium.com/@jamal.kaksouri/building-better-go-applications-with-clean-architecture-a-practical-guide-for-beginners-98ea061bf81a)
- [Clean Architecture, 2 years later](https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/)
- [Implementando Clean Architecture com Golang](https://dev.to/booscaaa/implementando-clean-architecture-com-golang-4n0a)
- [Go Backend Clean Architecture](https://outcomeschool.com/blog/go-backend-clean-architecture)
- [golang-clean-architecture](https://pkg.go.dev/github.com/hendrorahmat/golang-clean-architecture#section-readme)
- [[Hands-on Guide] How to Implement Clean Architecture in Golang?](https://reliasoftware.com/blog/clean-architecture-golang)
- [Clean DDD lessons: presenters](https://medium.com/unil-ci-software-engineering/clean-ddd-lessons-presenters-6f092308b75e)
- [Clean Architecture: Understanding the Role of Presenters](https://medium.com/@erickzanetti/clean-architecture-understanding-the-role-of-presenters-8707ff018aa3)
- [Golang Microservices Boilerplate - Clean Architecture](https://github.com/gbrayhan/microservices-go)
- [GRACEFULL SHUTDOWN EM GOLANG - Finalizando requisições antes de desligar o projeto!](https://www.youtube.com/watch?v=V52Th2h_8FM&ab_channel=HunCoding)
- [Implement value objects with Domain-Driven Design (DDD)](https://medium.com/@nirajranasinghe/implement-value-objects-with-domain-driven-design-ddd-3aeb4e88cee5)
- [Entendendo Presenters na Clean Architecture](https://www.youtube.com/watch?v=zrYAnqA-VQs&ab_channel=FullCycle)
- [RFC 8977 Registration Data Access Protocol (RDAP) Query Parameters for Result Sorting and Paging](https://www.rfc-editor.org/rfc/rfc8977.html#name-sort-parameter)
- [PostgreSQL - 7.5. Sorting Rows (ORDER BY) #](https://www.postgresql.org/docs/current/queries-order.html#QUERIES-ORDER)
- [WordPress API Reference - Posts](https://developer.wordpress.org/rest-api/reference/posts/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :busts_in_silhouette: Contributors

<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/atomaz"><img src="https://github.com/atomaz.png" width="100px;" alt=""/><br /><sub><b>Alice Tomaz</b></sub></a><br />
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/filipe1309"><img src="https://github.com/filipe1309.png" width="100px;" alt=""/><br /><sub><b>Filipe Leuch Bonfim</b></sub></a><br />
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/hugokishi"><img src="https://github.com/hugokishi.png" width="100px;" alt=""/><br /><sub><b>Hugo Kishi</b></sub></a><br />
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/marcos-nsantos"><img src="https://github.com/marcos-nsantos.png" width="100px;" alt=""/><br /><sub><b>Marcos Santos</b></sub></a><br />
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/vitorparras"><img src="https://github.com/vitorparras.png" width="100px;" alt=""/><br /><sub><b>Vitor Parras</b></sub></a><br />
    </tr>
  </tbody>
</table>

<p align="right">(<a href="#readme-top">back to top</a>)</p>
