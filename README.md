# Küfa-Karben

## Development

```bash
cd frontend
yarn
yarn dev
```

### API-Client Generation
You need the latest version of [go-swagger](https://github.com/go-swagger/go-swagger/releases/tag/v0.29.0).
After downloading, you can run following snippet. 
```bash
go generate ./src/rest
cd frontend
yarn generate:api
```

## Default User Creation
Set environment variables `KUEFA_DEFAULT_USER` and `KUEFA_DEFAULT_PASSWORD` to create a new user with given credentials. 