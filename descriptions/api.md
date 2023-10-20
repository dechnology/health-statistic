This is the API documentation for 「健康資料公鏈」

## Flow Diagram

![flow_diagram](https://cdn.discordapp.com/attachments/874556062815100940/1132920083174408222/App-.drawio.png)

---

## Auth0 Setup

This API uses Auth0 as its authorization provider, and all endpoints are protected using Auth0 token.
Therefore, all your requests must have **Authorization** header with a bearer token.

Use the following parameters to login to our Auth0 services:

- `domain`: `itri-dechnology.jp.auth0.com`

- `clientId`: `holxN6SuSQtRV5oOSwOIXWYwJnvioObh`

- `audience`: `https://health-statistic.dechnology.com.tw/api/v1/`

**_Reference_**: [**Auth0 | Quickstart**](https://auth0.com/docs/quickstarts)

## Schema Overview

### Quesionnaires

For the sake of scalibility, we make the questionnaire system a general one. That is, we can later create more questionnaires in our application. Upon initialization of our database, the registration questionnaire is created and assigned a fixed UUID.

---

## Client Endpoints

Full descriptions of all endpoints are described below; however, more common endpoints for a client application are described here. For detailed types of params, body, or response, please refer to the full descriptiions below.

### Registraion

Endpoints required to register an user.

- `/questionnaires/registration` [**GET**]

  Before the user register, we need to have them fill in the registration questionnaire. This questionniare can be fetch using this endpoint. Note that this questionnaire correspond to question 16 to 30 since the answers to question 1 to 15 are part of the user data.

- `/register` [**POST**]

  The only endpoint to create an user, note that we will use the ID from Auth0 as the primary key in our database, for example: `auth0|888888888888888888888888`.

### Utilities

Some useful endpoints.

- `/health_check` [**GET**]

  This endpoint checks if the server is alive and the database is connected. It returns a code of 200 if ture.

### Client's User Data

A normal user has access to all resources owned by him/her-self. This is available via the `/user` endpoint.
This is not meant to be confused with the plural counterpart, `/users`, whose endpoints are only accessible for users with `read:users` or `write:users` scope. The server will look into your Auth0 token and get the corresponding user data.

- `/user` [**GET**]

- `/user/healthkit` [**POST**]

- `/user/notifications` [**GET**]

  **Not yet implemented**

### Public Resources

The following endpoints are for reading public resources. Currently, they are **"quesionnaires"** and **"prices"**.

- `/quesionnaires` [**GET**]

- `/quesionnaires/{id}` [**GET**]

- `/prices` [**GET**]

- `/prices/{id}` [**GET**]

### Deegoo

The following APIs are for Deegoo server to hit.

- `/deegoo` [**POST**]

  This endpoint is for Deegoo server to post deegoo scores data to our server. The server will then process the data and store it in our database.
