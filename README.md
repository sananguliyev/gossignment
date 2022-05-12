# Gossignment

**Demo:** https://gossignment.guliev.info

## Setup
Just execute 
```shell
docker compose up -d
```

## Testing
The mock files are not added in order to keep the repository clean. You can use the commands for generating mock files:

```shell
cd /path/to/project
go install github.com/vektra/mockery/v2@latest
mockery --dir=domain/interactor --name=MemInteractor
mockery --dir=domain/interactor --name=RecordInteractor
mockery --dir=domain/repository --name=MemRepository
mockery --dir=domain/repository --name=RecordRepository
mockery --dir=domain/validator --name=MemValidator
mockery --dir=domain/validator --name=RecordValidator
```

## Endpoints

#### Filter records
```http request
POST http://gossignment.host/records/filter
Content-Type: application/json

{
  "startDate": "2017-01-26",
  "endDate": "2018-02-02",
  "minCount": 100,
  "maxCount": 300
}
```

#### Create in-memory record
```http request
POST http://gossignment.host/in-memory
Content-Type: application/json

{
  "key": "new-getirian",
  "value": "Sanan Guliyev"
}
```

#### Read in-memory record
```http request
GET http://gossignment.host/in-memory?key=salam1
```
