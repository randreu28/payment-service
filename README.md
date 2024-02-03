## Payment service

A backend API for economic transactions, using Go and PostgreSQL.

## Roadmap

1. Learn the basics of Go and set up a basic Go server. ✅
2. Create a postgreSQL database.✅
3. Write the schema for the PostgreSQL database ✅.
4. Connect the Go server to a PostgreSQL database  ✅.
5. Test the schema with Postman or other Go-specific testing tools.
6. Implement a simple transaction feature.
7. Test the transaction feature.
8. Add authentication to the server.
9. Dockerize the app for easy deployment and scaling.
10. Write a cool readme explaining everything!

## Todo's:

- Talk about the local docker compose not using SSL, but unencripted TCP/IP, which may pose a risk in a production enviroment. Talk about how it could be addressed.

- Explain the populate script 

```shell
go run src/db_populate/main.go 123
```


## Schema

### Accounts

POST /accounts { "account_owner": "string" }

> Creates a new account with an initial balance of 0.

GET /accounts

> Retrieves a list of all accounts.

GET /accounts/{id}

> Retrieves details of a specific account by its ID.

PUT /accounts/{id} { "account_owner": "string", "balance": "money" }

> Updates the specified fields of an account. (Note: Directly updating the balance might not be a good practice; typically, balance updates would be handled through transactions.)

DELETE /accounts/{id}

> Deletes a specific account by its ID.


### Transactions

POST /transactions { "account_from": "integer", "account_to": "integer", "amount": "money", "description": "string" }

> Creates a new transaction, transferring money from one account to another.

GET /transactions

> Retrieves a list of all transactions.

GET /transactions/{id}

> Retrieves details of a specific transaction by its ID.
Additional Functionalities

GET /accounts/{id}/balance

> Retrieves the current balance of a specific account.

POST /transfer { "account_from": "integer", "account_to": "integer", "amount": "money" }

> A specialized endpoint to facilitate money transfer between accounts, wrapping the transaction creation process with additional validations.

GET /accounts/{id}/transactions

> Retrieves a list of all transactions involving a specific account, either as the sender or receiver.