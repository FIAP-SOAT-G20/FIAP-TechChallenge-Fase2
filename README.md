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

## 💬 About

Repository for the [FIAP](https://postech.fiap.com.br/) Tech Challenge 2, focused on developing a monolith backend system for managing orders in a fast-food restaurant.

Tech Challenge 2 specifications can be found [here](docs/tc2-spec.pdf).

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
│   ├── adapters
│   │   ├── controller
│   │   ├── gateway
│   │   └── presenter
│   ├── core
│   │   ├── domain
│   │   │   ├── entity
│   │   ├── port
│   │   └── usecase
│   └── infrastructure
│       ├── config
│       ├── database
│       ├── datasources
│       ├── handler
│       ├── logger
│       ├── middleware
│       ├── routes
│       └── server
├── k8s
```

<details>
<summary>Project Structure Explanation</summary>

### **1️⃣ Core (Innermost layer)**
- `domain/`: Central business entities and rules.
- `usecases/`: Application use cases.
- `ports/`: Interfaces that define contracts between layers, ensuring independence.

### **2️⃣ Adapters (Middle layer)**
- `controller/`: Coordinates the flow of data between use cases and infrastructure.
- `presenter/`: Formats data for presentation.
- `gateway/`: Implements access to data from external sources (databases, APIs, etc.).

### **3️⃣ Infrastructure (External layer)**
- `config/`: Application configuration management.
- `database/`: Configuration and connection to the database. - `server/`: Initialization of the HTTP server.
- `routes/`: Definition of API routes.
- `middleware/`: HTTP middlewares for handling requests.
- `logger/`: Structured logger for detailed logs.
- `handler/`: Handling of HTTP requests.
- `datasources/`: Concrete implementations of data sources.

</details>



### ✨ Features
- [x] Dockerfile: small image with multi-stage docker build, and independent of the host environment
- [x] Makefile: to simplify the build and run commands
- [x] Clean architecture
- [x] PostgreSQL database
- [x] Conventional commits
- [x] Unit tests
- [x] Code coverage
- [x] Swagger documentation
- [x] Postman collection
- [x] Feature branch workflow
- [x] Air to run go
- [x] Pagination
- [x] Health Check
- [x] Lint
- [x] Vulnerability check

## :computer: Technologies

- [Go 1.23+](https://golang.org/)
- [Gin Web Framework](https://gin-gonic.com/)
- [golangci-lint](https://golangci-lint.run/)
- [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
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


## :scroll: Requirements

### Build/Run with Docker

- [Docker 1.23+](https://www.docker.com/)

### Build/Run Locally

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)

> [!NOTE]
> You need to have Go (> 1.23) installed in your machine to build, run and test the application locally

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :cd: Installation

```sh
git clone https://github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1.git
```

```sh
cd FIAP-TechChallenge-Fase1
```

Set the environment variables
```sh
cp .env.example .env
```

### :whale: Docker

```sh
make compose-build
```
> The binary will be created in the `bin` folder

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :runner: Running

### :whale: Docker

```sh
make compose-up
```

> [!NOTE]
> To stop the application, run `compose-stop`
> To remove the application, run `compose-clean`

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### :gear: Kubernetes

```bash
kubectl apply -f k8s/
```

> [!NOTE]
> You can check the application swagger documentation at http://localhost:8080/docs/index.html  
> Alternatively, a postman collection is available at [here](docs/10soat-g20-tech-challenge-1.postman_collection.json)  

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
13. Commit your changes
14. Push to the branch and Open a new PR by running `make pull-request`

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
- [DBML](https://www.dbml.org/)
- [Health Check Response Format for HTTP APIs](https://datatracker.ietf.org/doc/html/draft-inadarei-api-health-check-06)
- [Event Storming](https://www.eventstorming.com/)
- [Swagger](https://swagger.io/)
- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
- [Shields.io](https://shields.io/)
- [Alistair in the "Hexagone" 1/3](https://www.youtube.com/watch?v=th4AgBcrEHA&list=PLGl1Jc8ErU1w27y8-7Gdcloy1tHO7NriL&ab_channel=DDDFR)

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
