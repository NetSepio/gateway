# NETSEPIO Engine

REST APIs for Web3 Auth and Smart Contract Functionalities

[![.github/workflows/test.yml](https://github.com/NetSepio/Engine/actions/workflows/test.yml/badge.svg)](https://github.com/NetSepio/gateway/actions/workflows/test.yml)
[![Lint](https://github.com/NetSepio/Engine/actions/workflows/lint.yml/badge.svg)](https://github.com/NetSepio/gateway/actions/workflows/lint.yml)

# Getting Started

## Postgres for development

```bash
docker run --name="netsepio" --rm -d -p 5432:5432 \
-e POSTGRES_PASSWORD=netsepio \
-e POSTGRES_USER=netsepio \
-e POSTGRES_DB=netsepio \
postgres -c log_statement=all
```

## Steps to get started

- Run `go get ./...` to install dependencies
- Set up env variables or create `.env` file as per [`.env-sample`](https://github.com/NetSepio/gateway/blob/main/.env-sample) file
- Run `go test ./...` to make sure setup is working
- Run `go run main.go` to start server

## API Reference

### Auth

For protected APIs use PASETO token which can be obtained after calling authenticate API.

Use `Authorization` key in header in order to send token

### APIs

#### Returns flow ID and EULA which should be signed and send to authenticate API in order to get the PASETO which can be used for accessing all other APIs

```
  GET /flowid?walletAddress={{wallet address}}
```

#### Request PASETO for the user who accepted the EULA by signing it with flow ID, i.e. flowid+EULA.

```
  POST /authenticate
```

| Parameter   | Type     | Description                                             |
| :---------- | :------- | :------------------------------------------------------ |
| `flowId`    | `string` | **Required**. flowId you got from flowId API            |
| `signature` | `string` | **Required**. signature obtained by signing flowId+EULA |

#### Get profile details of user present in the system.

Note - Some unset data is emitted.

```
  GET /profile
```

#### Updates profile data like name, profile picture URL, etc.

```
  PATCH /profile
```

| Parameter           | Type     |
| :------------------ | :------- |
| `name`              | `string` |
| `country`           | `string` |
| `profilePictureUrl` | `string` |

#### Returns flow ID and Eula which should be signed and passed to claim Role in order to successful verification and claim of role

```
  GET /roleId/{{roleId}}
```

#### Successfully complete role claim by sending signature which is obtained from signing eula+flowId which was returned from roleId API (aka Request role)

```
  POST /claimrole
```

| Parameter   | Type                   |
| :---------- | :--------------------- |
| `flowId`    | **Required**. `string` |
| `signature` | **Required**. `string` |

#### Delegate review creation to other voter

```
  POST /delegateReviewCreation
```

| Parameter     | Type                   |
| :------------ | :--------------------- |
| `voter`       | **Required**. `string` |
| `MetaDataUri` | **Required**. `string` |

```
  POST /feedback
```

| Parameter  | Type                                      |
| :--------- | :---------------------------------------- |
| `feedback` | **Required**. `string`                    |
| `rating`   | **Required**. `int` `ranging from 0 to 5` |
