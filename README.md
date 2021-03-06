# Golang Authentication

Membuat Authentication Golang dengan menggunakan Gin Framework dan JWT

## Usage

```bash
git clone https://github.com/wildanicardi/go-gin-auth
go build
go run main.go
```

## API Endpoints

### Login

- Path : ` /api/auth/login`
- Method : `POST`
- Response : ` 200`
- Field : `email, password `

### Register

- Path : ` /api/auth/register`
- Method : `POST`
- Field : `name, email, password,password_confirm `
- Response : ` 200`

### GetOne User

- Path : ` /api/auth/users`
- Method : `POST`
- Response : ` 200`
- Authorization : `Bearer Token`

## Dokumentasi Postman
https://documenter.getpostman.com/view/6225373/TVYF7J3p

## License

[MIT](https://choosealicense.com/licenses/mit/) wildanicardi
