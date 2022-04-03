# Usage

## Installation
```go
go mod init api
go mod tidy
```

## Quick Start
- run this command to start the server
```go
go run api
```
and visit
- http://localhost:8080/questions
```bash
[{"ID":1,"Year":2022,"Genre":"臨床生理学","Question":"問い1","Answer":"答え1","Commentary":"コメント1"},{"ID":2,"Year":2022,"Genre":"臨床血液学","Question":"問い2","Answer":"答え2","Commentary":"コメント2"}]
```
- http://localhost:8080/id/1
```bash
[{"ID":1,"Year":2022,"Genre":"臨床生理学","Question":"問い1","Answer":"答え1","Commentary":"コメント1"}]
```
- http://localhost:8080/year/2
```bash
[{"ID":1,"Year":2022,"Genre":"臨床生理学","Question":"問い1","Answer":"答え1","Commentary":"コメント1"},{"ID":2,"Year":2022,"Genre":"臨床血液学","Question":"問い2","Answer":"答え2","Commentary":"コメント2"}]
```

# Directory Structure
```bash
${ROOT}
├── data
│   └── questions.db
├── go.mod
├── go.sum
├── main.go
└── model
    └── db_model.go
```
