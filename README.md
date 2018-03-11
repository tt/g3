# G3

G3 is a proof-of-concept implementation of a high-level, highly available API
layer for an imaginary bank.

G3 attempts to answer how to expose a coherent, user-facing API in a
service-oriented architecture.


## Usage

1. Start the services: `docker-compose up`

2. Navigate to the GraphiQL interface (available on `docker-compose port api 8080`)

3. Open an account by executing the following mutation:

    ```graphql
    mutation {
      openAccount {
        id
      }
    }
    ```

4. List the accounts with the following query:

    ```graphql
    query {
      accounts {
        id
      }
    }
    ```


## License

This project is dedicated to the public domain.

[![CC0](https://i.creativecommons.org/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)
