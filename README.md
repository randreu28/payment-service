## Payment service

A backend API for economic transactions, using Go and PostgreSQL. Meant for learning purposes only. It solves concurrency problems in distributed systems by the aid of SQL transactions in a relational database. It also uses JWT for authentication, a logging system and tests for debugging and a docker-compose file for easy development experience.

## Setting up a development environment

You'll need to have Go installed. If you haven't check the get started guide [here](https://golang.org/doc/install).

You'll also need to have Docker and Docker Compose installed. If you haven't check the get started guide [here](https://docs.docker.com/get-docker/).

And local enivorment variables, that are part of the .gitignore for security reason. You may create a `.env.local` file like this:

```.env
POSTGRES_PORT='5432'
POSTGRES_PASSWORD='admin'
DB_URI='postgres://admin:admin@localhost:5432/test?sslmode=disable'
POSTGRES_USER='admin'
POSTGRES_HOST='localhost'
POSTGRES_DB='test'
PORT="3001"
JWT_SECRET="my-very-long-256-bit-secret"
```

To run your PostgreSQL database locally, you can use the following command:

```shell
docker-compose up
```

To populate it with dummy data, you can use the db_populate script. The script takes a single optional argument, the number of accounts to create. If no argument is passed, it will create 100 accounts.

```shell
cd src
go run cmd/db_populate/main.go
```

To get the server up and running, you can use the following command:

```shell
go run cmd/server/main.go
```

## Routes

### Public routes

GET /health ✅

> Checks the health status of the server.

POST /auth ✅

> Authorizes an account, creating a JWT

GET /transactions/{id} ✅

> Retrieves details of a specific transaction by its ID.

### Private routes

POST /accounts { "account_owner": "string" } ✅

> Creates a new account with an initial balance of 0. 

GET /account ✅

> Retrieves details of the current account

DELETE /account ✅

> Deletes the current account

GET /account/transactions ✅

> Retrieves a list of all transactions involving a specific account, either as the sender or receiver.

POST /transfer { "account_from": "integer", "account_to": "integer", "amount": "money" } ✅

> A specialized endpoint to facilitate money transfer between accounts, wrapping the transaction creation process with additional validations.  

## Disclaimers

As the project was meant for learning purposes only, there are a couple PF security and portability considerations that weren't prioritized. 

For example, it's important to note that the communication between the database and the application is transmitted over unencrypted TCP/IP. This setup may pose a security risk, especially in a production environment where sensitive data is involved.

To address this issue and enhance security, consider the following steps:

1. **Enable SSL Encryption**: Configure PostgreSQL to use SSL encryption for connections. This ensures that data transmitted between the database and the application is encrypted, reducing the risk of interception.

2. **Use TLS Certificates**: Generate and configure TLS certificates for both the PostgreSQL server and the client applications. This helps in establishing secure connections and verifying the identity of the communicating parties.

3. **Implement Network Segmentation**: Ensure that the database server is not directly exposed to the public internet. Use network segmentation techniques to restrict access to the database only from trusted sources.

4. **Regular Security Audits**: Conduct regular security audits and vulnerability assessments to identify and address any potential security weaknesses in the setup.

The application itself isn't dockerized either. This means that the application is not portable and can't be easily deployed to different environments. Although, due to Go being a compiled language, an executable could be compiled as so:

```shell
go build cmd/server/main.go
```

The only consideration is the server's architecture, as the executable would need to be compiled for the specific hardware specs (amd64 or arm64).

By implementing these security and portability measures, you can mitigate the security risks and ensure protability to prevent vendor locking when choosing your cloud provider. ✨


## Roadmap

1. Learn the basics of Go and set up a basic Go server. ✅
2. Create a postgreSQL database. ✅
3. Write the schema for the PostgreSQL database. ✅
4. Connect the Go server to a PostgreSQL database . ✅
5. Test the schema with Postman or other Go-specific testing tools. ✅
6. Implement a simple transaction feature. ✅
7. Test the transaction feature. ✅
8. Add authentication to the server. ✅
9. Write a cool readme explaining everything! ✅
10. Make the reader star the repo because it's cool and you're cool too. 👀
