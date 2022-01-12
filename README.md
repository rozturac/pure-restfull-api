# REST API With Mediator

# REST API

With this project, it was tried to develop a restful service using mediator pattern in golang.

## Get Records

### Request

`POST /records`

```json
{
  "startDate": "2016-01-21",
  "endDate": "2016-03-02",
  "minCount": 2900,
  "maxCount": 3000
}
```

### Response

```
HTTP/1.1 200 OK
Date: Wed, 12 Jan 2022 17:12:49 GMT
Status: 200 OK
Connection: close
Content-Type: application/json
Content-Length: 204
```

```json
{
  "code": 0,
  "msg": "Success",
  "records": [
    {
      "key": "ibfRLaFT",
      "createdAt": "2016-12-25T16:43:27.909Z",
      "totalCount": 2892
    },
    {
      "key": "pxClAvll",
      "createdAt": "2016-12-19T10:00:40.05Z",
      "totalCount": 2772
    }
  ]
}
```

## Create Config

### Request

`POST /configs`

```json
{
  "key": "Hello",
  "value": "World"
}
```
### Response
```
HTTP/1.1 201 Created
Date: Wed, 12 Jan 2022 17:12:49 GMT
Status: 201 Created
Connection: close
Content-Type: application/json
Content-Length: 0
```

## Get Config

### Request

`GET /configs?key=Hello`

    curl -i -H 'Accept: application/json' http://localhost:8080/configs?key=Hello

### Response

HTTP/1.1 200 OK
Date: Wed, 12 Jan 2022 17:13:08 GMT
Status: 200 OK
Connection: close
Content-Type: application/json
Content-Length: 32

```json
{
    "key": "Hello",
    "value": "World"
}
```
