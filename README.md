# AutoERP

AutoERP is a cloud-based business management system designed for automobile accessories, customization, and service businesses.

## Features

- Customer management
- Vehicle management
- Service management
- Quotation management
- Work order management
- Invoice management
- Payment tracking
- Dashboard and reporting
- PostgreSQL database
- REST APIs

## Technology Stack

- Go
- Gin
- PostgreSQL
- SQLC
- Docker
- Git & GitHub

## Backend Architecture

Client / React Frontend
        ↓
     Gin API
        ↓
      Go
        ↓
      SQLC
        ↓
   PostgreSQL

## Business Workflow

Customer
   ↓
Quotation
   ↓
Customer Approval
   ↓
Work Order
   ↓
Invoice
   ↓
Payment

## API Modules

- Customers
- Vehicles
- Services
- Quotations
- Work Orders
- Invoices
- Payments
- Dashboard

## Running the Project

### 1. Start PostgreSQL

docker compose up -d

### 2. Set the Database URL

Create a .env file:

DATABASE_URL=postgres://autoerp:autoerp@localhost:5432/autoerp

### 3. Run the Application

go run .

The API runs on:

http://localhost:8080

## Project Structure

AutoERP/
├── db/sqlc/          # SQLC generated code
├── migrations/       # Database migrations
├── queries/          # SQL queries
├── *_handler.go      # HTTP handlers
├── *.go              # Application logic
├── docker-compose.yml
├── sqlc.yaml
├── go.mod
└── go.sum

## Status

Backend MVP completed.

React frontend is planned as the next phase.