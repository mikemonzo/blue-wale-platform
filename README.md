# Introducción

Este repositorio contiene un proyecto personal para afianzar conceptos de Arquitectura de software y desarrollar servicios con Go. También, se aprovechará para comenzar en el desarrollo de frontend web y mobile(ios, android). Todo ello basado en un desarrollo BDD.

## Proyecto

La idea principal es llevar a cabo la construcción de un producto SaaS que nos permita gestionar ficheros (almacenamiento, envío, control de fraude, integridad, etc).
Algunos requerimientos que debe cumplir el producto:
- Para mercados B2B y B2C
- Despliegue multiproveedor, on-cloud, on-premise o híbrido
- Separación backend y los diferentes clientes web y mobile
- Cumplimiento ENS, ISO 27001 y GDPR
- Arquitecturas límpias y código límpio.
- Framework gofr
- Comunicación entre microservicios con gRPC
- Servicios publicos y comunicación con clientes a través de api Restful

## Arquitectura inicial

```mermaid
graph TD
    Client[Client Application] --> Gateway[API Gateway]
    Gateway --> Auth[IdP]
    Gateway --> Tenant[Tenant Services]
    Gateway --> Other[Other Services]
    Auth --> Cache[(Redis)]
    Auth --> DB[(Database)]
    Tenant --> DB[(Database)]
    Other --> DB[(Database)]

```

## Estructura del proyecto

```
saas-file-platform/
├── .github/
│   ├── workflows/
│   │   ├── ci.yml
│   │   └── release.yml
│   └── PULL_REQUEST_TEMPLATE.md
├── deployments/
│   ├── docker/
│   │   ├── idp/
│   │   │   └── Dockerfile
│   │   ├── tenant-service/
│   │   │   └── Dockerfile
│   │   └── api-gateway/
│   │       └── Dockerfile
│   ├── kubernetes/
│   │   ├── base/
│   │   └── overlays/
│   └── terraform/
│       ├── modules/
│       └── environments/
├── docs/
│   ├── architecture/
│   │   ├── diagrams/
│   │   └── decisions/
│   ├── api/
│   │   └── openapi/
│   └── development/
│       └── getting-started.md
├── shared/
│   ├── common/
│   │   ├── auth/
│   │   │   ├── jwt/
│   │   │   └── middleware/
│   │   ├── database/
│   │   │   ├── postgres/
│   │   │   └── migrations/
│   │   ├── logging/
│   │   ├── monitoring/
│   │   ├── testing/
│   │   ├── go.mod
│   │   └── go.sum
│   └── pkg/
│       ├── tenant/
│       ├── errors/
│       ├── validation/
│       ├── go.mod
│       └── go.sum
├── services/
│   ├── idp/
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── application/
│   │   │   │   ├── commands/
│   │   │   │   ├── queries/
│   │   │   │   └── services/
│   │   │   ├── domain/
│   │   │   │   ├── model/
│   │   │   │   ├── repository/
│   │   │   │   └── service/
│   │   │   └── infrastructure/
│   │   │       ├── persistence/
│   │   │       ├── auth/
│   │   │       └── api/
│   │   ├── config/
│   │   │   ├── config.go
│   │   │   └── config.yaml
│   │   ├── migrations/
│   │   ├── test/
│   │   ├── go.mod
│   │   └── go.sum
│   ├── tenant-service/
│   │   └── [estructura similar a idp]
│   └── api-gateway/
│       └── [estructura similar a idp]
├── scripts/
│   ├── setup.sh
│   └── dev.sh
├── tools/
│   └── go.mod
├── web/
│   └── admin-panel/
├── mobile/
│   ├── android/
│   └── ios/
├── go.work
├── go.work.sum
├── Makefile
├── docker-compose.yml
└── README.md
```