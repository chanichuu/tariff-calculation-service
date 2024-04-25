# tariff-calculation-service

A microservice implemented with Go using Lambda, DynamoDb, Api Gateway &amp; SQS to calculate Tariffs for Energy Consumption
Andgreen Service

# Table of Content

1. [ Description ](#description)
2. [ REST API ](#rest-api)
3. [ OpenAPI ](#openapi)
4. [ Local Development ](#local-development)

# Description

Service to manage tariffs, contracts and their providers to calculate costs of consumption.

Entities:

- Tariff
- Contract
- Provider

# REST API

## Base

- /partitions/{partitionId}

## Tariff

- GET /tariffs
- POST /tariffs
- GET /tariffs/{tariffId}
- PUT /tariffs/{tariffId}
- DELETE /tariffs/{tariffId}

## Contract

- GET /contracts
- POST /contracts
- GET /contracts/{contractId}
- PUT /contracts/{contractId}
- DELETE /contracts/{contractId}

## Provider

- GET /providers
- POST /providers
- GET /providers/{providerId}
- PUT /providers/{providerId}
- DELETE /providers/{providerId}

## Service

- GET /health
- GET /version
- GET /restversion

# OpenAPI

OpenAPI documentation can be found here:

- LOCAL: [ TariffCalculation Service OpenAPI LOCAL ](http://localhost:8000/api/schema/swagger-ui/#/)
- DEV: n/a
- TEST: n/a
- PROD: n/a

# Local Development

## Backend

0. Prerequisites: Golang Version >= 1.20, Gin
1. todo

## Frontend

todo
