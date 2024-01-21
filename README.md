# 🍀 kafka-transaction-validator
> :warning: This repository was created to practice using Kafka and is not code to be used in production.

## Problem
Every time a financial transaction is created it must be validated by our anti-fraud microservice and then the same service sends a message back to update the transaction status. For now, we have only three transaction statuses:

- pending
- approved
- rejected

Every transaction with a value greater than 1000 should be rejected.
![image](https://github.com/Caixetadev/fraud-check-kafka-integration/assets/87894998/36d9501d-fa50-4dce-a84a-79ab46d20b2b)

## Requirements
- Go installed and configured
- Docker and Docker Compose installed

## How to Run

1. Clone repository:
    ```bash
    git clone https://github.com/Caixetadev/kafka-transaction-validator.git && cd kafka-transaction-validator
    ```

1. Start the services using Docker Compose:

    ```bash
    docker compose up
    ```

2. Run the `anti-fraud` service:

    ```bash
    cd anti-fraud && go run cmd/main.go
    ```

3. Run the `transaction` service:

    ```bash
    cd transaction && go run cmd/main.go
    ```

4. Access the application at: [http://localhost:8080/transaction](http://localhost:8080/transaction).

## API Endpoint

### Create Transaction

- **URL:** `/transaction`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "accountExternalIdDebit": "string",
    "accountExternalIdCredit": "string",
    "tranferTypeId": 1,
    "value": 1
  }
