# Go Auth Server 🚀
---

Este é um servidor simples em Go para autenticação de usuários utilizando JWT (JSON Web Tokens) e um banco de dados PostgreSQL. Ele oferece rotas para registrar um novo usuário, fazer login e acessar uma lista de usuários autenticados.

## Funcionalidades 🛠️

- **Registro de Usuário**: Criação de um novo usuário com e-mail e senha.
- **Login**: Geração de um JWT ao fazer login com o e-mail e senha.
- **Proteção de Rotas**: Middleware para proteger rotas com autenticação JWT.
- **Listagem de Usuários**: Rota para obter uma lista de todos os usuários (somente acessível com token JWT válido).
---
## Tecnologias Utilizadas 💻

- **Go**: Linguagem de programação principal.
- **Fiber**: Framework web para Go.
- **JWT**: Para autenticação via tokens.
- **PostgreSQL**: Banco de dados relacional utilizado para armazenar os usuários.
- **Docker**: Para facilitar o uso do banco de dados PostgreSQL em container.
---
## Requisitos ⚙️

- Go 1.18 ou superior
- Docker (para rodar o PostgreSQL)
---
## Instalação 📝

### 1. Clonar o repositório

```bash
git clone https://github.com/PedroHenrique1606/auth-crud-go.git
cd auth-crud-go
```

### 2. Iniciarlizar o servidor
```bash
go run maing.go
```
---
## Endpoints🔧

### 1. POST /register - Registrar novo usuário
Registre um novo usuário com e-mail e senha.

- Request:
```json

{
  "email": "usuario@exemplo.com",
  "password": "senha123"
}
```
- Resposta:

```json
{
  "message": "Usuário registrado com sucesso"
}
```
### 2. POST /login - Login
Faça login e obtenha um token JWT para autenticação.

- Request:
```json

{
  "email": "usuario@exemplo.com",
  "password": "senha123"
}
```

- Resposta:

```json
{
  "token": "seu_token_jwt_aqui"
}
```

### 4. GET /users - Listar usuários (requer autenticação)
Obtenha a lista de usuários cadastrados. Esta rota requer que você envie um token JWT no cabeçalho Authorization como Bearer <token>.

Request:
```bash

Cabeçalho Authorization: Bearer <seu_token_jwt_aqui>
```

Resposta:
```json
[
    {
        "id": 1,
        "email": "teste@email.com",
        "created_at": "2025-01-29T21:50:45.761732Z"
    },
    {
        "id": 2,
        "email": "pedromelo.dev.contato@gmail.com",
        "created_at": "2025-01-29T21:51:07.539224Z"
    }
]
```
