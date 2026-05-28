# Users CRUD API 👥

Uma API RESTful completa para gerenciamento de usuários desenvolvida em Go com PostgreSQL, implementando todas as operações básicas (Create, Read, Update, Delete) com validações robustas e tratamento de erros.

---

## 📋 Sobre a Aplicação

A **Users CRUD API** é uma aplicação backend moderna que fornece endpoints para gerenciar usuários de forma segura e eficiente. Desenvolvida com as melhores práticas do Go, utiliza:

- **Framework HTTP**: Chi v5 para roteamento e middleware
- **Banco de Dados**: PostgreSQL com suporte a queries parameterizadas
- **Geração de IDs**: UUID v4 para identificadores únicos
- **Logging**: slog para registros estruturados
- **Containerização**: Docker Compose para ambiente isolado

A aplicação foi construída seguindo o padrão de separação de responsabilidades, com handlers para lógica de negócio, funções de banco de dados para persistência e structs bem definidas para transferência de dados.

---

## 🛠️ Tecnologias Utilizadas

| Tecnologia         | Versão    | Descrição                         |
| ------------------ | --------- | --------------------------------- |
| **Go**             | 1.25.6    | Linguagem de programação          |
| **Chi**            | v5.2.5    | Router e framework HTTP           |
| **PostgreSQL**     | 16-alpine | Banco de dados relacional         |
| **UUID**           | 1.6.0     | Geração de identificadores únicos |
| **lib/pq**         | 1.10.9    | Driver PostgreSQL para Go         |
| **Docker Compose** | -         | Orquestração de containers        |

---

## 📦 Requisitos

Antes de iniciar, certifique-se de ter instalado:

- **Go** >= 1.25.6
- **Docker** e **Docker Compose**
- **curl** ou **Postman** (para testar os endpoints)

---

## 🚀 Instalação e Setup

### 1. Clonar o Repositório

```bash
cd ~/www/rocketseat/golang/modulo-3/users-crud
```

### 2. Instalar Dependências

```bash
go mod download
go mod tidy
```

### 3. Iniciar o PostgreSQL com Docker Compose

```bash
docker compose up -d
```

Este comando inicia um container PostgreSQL com as configurações:

- **Usuário**: user
- **Senha**: password
- **Banco de Dados**: users_db
- **Porta**: 5432

Verificar status dos containers:

```bash
docker compose ps
```

### 4. Executar a Aplicação

```bash
go run m.go
```

A aplicação será iniciada em `http://localhost:8080`

Você deve ver a saída:

```
time=... level=INFO msg=...
```

---

## 📊 Estrutura do Banco de Dados

### Tabela: `users`

A aplicação cria automaticamente a seguinte tabela ao iniciar:

```sql
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    biography TEXT NOT NULL
);
```

**Campos:**

- `id` (TEXT): Identificador único em formato UUID
- `first_name` (TEXT): Primeiro nome do usuário (2-20 caracteres)
- `last_name` (TEXT): Sobrenome do usuário (2-20 caracteres)
- `biography` (TEXT): Biografia do usuário (20-450 caracteres)

---

## 🔌 Rotas da API

### Base URL

```
http://localhost:8080/api/users
```

### 1️⃣ Criar Usuário (POST)

**Endpoint:**

```
POST /api/users
```

**Headers:**

```
Content-Type: application/json
```

**Body (JSON):**

```json
{
  "first_name": "João",
  "last_name": "Silva",
  "biography": "Desenvolvedor full stack apaixonado por tecnologia e inovação."
}
```

**Resposta de Sucesso (201 Created):**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "first_name": "João",
    "last_name": "Silva",
    "biography": "Desenvolvedor full stack apaixonado por tecnologia e inovação."
  }
}
```

**Resposta de Erro (400 Bad Request):**

```json
{
  "error": "invalid fields"
}
```

**Validações:**

- ✓ `first_name`: mínimo 2, máximo 20 caracteres
- ✓ `last_name`: mínimo 2, máximo 20 caracteres
- ✓ `biography`: mínimo 20, máximo 450 caracteres
- ✓ Body JSON válido e obrigatório

**cURL de Exemplo:**

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Maria",
    "last_name": "Santos",
    "biography": "Engenheira de software com 5 anos de experiência em desenvolvimento web."
  }'
```

---

### 2️⃣ Listar Todos os Usuários (GET)

**Endpoint:**

```
GET /api/users
```

**Resposta de Sucesso (200 OK):**

```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "first_name": "João",
      "last_name": "Silva",
      "biography": "Desenvolvedor full stack apaixonado por tecnologia e inovação."
    },
    {
      "id": "660e8400-e29b-41d4-a716-446655440111",
      "first_name": "Maria",
      "last_name": "Santos",
      "biography": "Engenheira de software com 5 anos de experiência em desenvolvimento web."
    }
  ]
}
```

**Resposta Vazia (200 OK):**

```json
{
  "data": []
}
```

**Sem Parâmetros de Paginação**

**cURL de Exemplo:**

```bash
curl http://localhost:8080/api/users
```

---

### 3️⃣ Buscar Usuário por ID (GET)

**Endpoint:**

```
GET /api/users/{id}
```

**Resposta de Sucesso (200 OK):**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "first_name": "João",
    "last_name": "Silva",
    "biography": "Desenvolvedor full stack apaixonado por tecnologia e inovação."
  }
}
```

**Resposta de Erro (404 Not Found):**

```json
{
  "error": "user not found"
}
```

**cURL de Exemplo:**

```bash
curl http://localhost:8080/api/users/550e8400-e29b-41d4-a716-446655440000
```

---

### 4️⃣ Atualizar Usuário (PUT)

**Endpoint:**

```
PUT /api/users/{id}
```

**Headers:**

```
Content-Type: application/json
```

**Body (JSON):**

```json
{
  "first_name": "João Carlos",
  "last_name": "Silva Santos",
  "biography": "Desenvolvedor full stack com experiência em várias linguagens de programação."
}
```

**Resposta de Sucesso (200 OK):**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "first_name": "João Carlos",
    "last_name": "Silva Santos",
    "biography": "Desenvolvedor full stack com experiência em várias linguagens de programação."
  }
}
```

**Respostas de Erro:**

- 404 Not Found: Se o usuário não existir
- 400 Bad Request: Se os campos forem inválidos

**Validações:**

- ✓ Mesmas validações do POST (first_name, last_name, biography)
- ✓ Usuário deve existir no banco de dados

**cURL de Exemplo:**

```bash
curl -X PUT http://localhost:8080/api/users/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "João Carlos",
    "last_name": "Silva Santos",
    "biography": "Desenvolvedor full stack com experiência em várias linguagens de programação."
  }'
```

---

### 5️⃣ Deletar Usuário (DELETE)

**Endpoint:**

```
DELETE /api/users/{id}
```

**Resposta de Sucesso (200 OK):**

```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "first_name": "João",
    "last_name": "Silva",
    "biography": "Desenvolvedor full stack apaixonado por tecnologia e inovação."
  }
}
```

**Resposta de Erro (404 Not Found):**

```json
{
  "error": "user not found"
}
```

**Observação:** Retorna os dados do usuário deletado como confirmação

**cURL de Exemplo:**

```bash
curl -X DELETE http://localhost:8080/api/users/550e8400-e29b-41d4-a716-446655440000
```

---

## 🧪 Testes

### Teste 1: Criar um Novo Usuário

```bash
# Request
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Pedro",
    "last_name": "Oliveira",
    "biography": "Especialista em backend com foco em arquitetura de sistemas escaláveis."
  }'

# Response esperada (201 Created)
# Retorna o usuário criado com ID gerado automaticamente
```

### Teste 2: Listar Todos os Usuários

```bash
# Request
curl http://localhost:8080/api/users

# Response esperada (200 OK)
# Retorna array de todos os usuários cadastrados
```

### Teste 3: Buscar Usuário Específico

```bash
# Substitua {id} pelo ID obtido no Teste 1
curl http://localhost:8080/api/users/{id}

# Response esperada (200 OK)
# Retorna os dados do usuário
```

### Teste 4: Atualizar Usuário

```bash
# Substitua {id} pelo ID obtido no Teste 1
curl -X PUT http://localhost:8080/api/users/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Pedro Silva",
    "last_name": "Oliveira Santos",
    "biography": "Especialista em backend com foco em arquitetura de sistemas escaláveis e microserviços."
  }'

# Response esperada (200 OK)
# Retorna usuário atualizado
```

### Teste 5: Deletar Usuário

```bash
# Substitua {id} pelo ID obtido no Teste 1
curl -X DELETE http://localhost:8080/api/users/{id}

# Response esperada (200 OK)
# Retorna dados do usuário deletado como confirmação
```

### Teste 6: Validações de Campo (Deve Falhar)

```bash
# Primeiro nome muito curto (< 2 caracteres)
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "P",
    "last_name": "Oliveira",
    "biography": "Especialista em backend com foco em arquitetura de sistemas escaláveis."
  }'

# Response esperada (400 Bad Request)
# {"error":"invalid fields"}
```

```bash
# Biografia muito curta (< 20 caracteres)
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Pedro",
    "last_name": "Oliveira",
    "biography": "Dev"
  }'

# Response esperada (400 Bad Request)
# {"error":"invalid fields"}
```

```bash
# Body JSON inválido
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d 'invalid json'

# Response esperada (400 Bad Request)
# {"error":"invalid body"}
```

### Teste 7: Buscar Usuário Inexistente

```bash
curl http://localhost:8080/api/users/id-inexistente-12345

# Response esperada (404 Not Found)
# {"error":"user not found"}
```

### Teste 8: Atualizar Usuário Inexistente

```bash
curl -X PUT http://localhost:8080/api/users/id-inexistente-12345 \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Teste",
    "last_name": "Teste",
    "biography": "Uma biografia válida com mais de vinte caracteres."
  }'

# Response esperada (404 Not Found)
# {"error":"user not found"}
```

### Teste 9: Deletar Usuário Inexistente

```bash
curl -X DELETE http://localhost:8080/api/users/id-inexistente-12345

# Response esperada (404 Not Found)
# {"error":"user not found"}
```

---

## 📝 Validações Implementadas

### Campos Obrigatórios

- ✅ `first_name` - Obrigatório
- ✅ `last_name` - Obrigatório
- ✅ `biography` - Obrigatório

### Regras de Validação

| Campo          | Mínimo        | Máximo         | Tipo   | Detalhes                                    |
| -------------- | ------------- | -------------- | ------ | ------------------------------------------- |
| **first_name** | 2 caracteres  | 20 caracteres  | String | Nome não pode ser vazio ou muito longo      |
| **last_name**  | 2 caracteres  | 20 caracteres  | String | Sobrenome não pode ser vazio ou muito longo |
| **biography**  | 20 caracteres | 450 caracteres | String | Biografia necessita de conteúdo relevante   |

### Validações de Negócio

- ✅ UUID gerado automaticamente para cada novo usuário
- ✅ Usuário deve existir para ser atualizado ou deletado
- ✅ Não é possível criar usuários duplicados (ID é único)
- ✅ Body JSON deve ser válido

---

## 🏗️ Estrutura do Projeto

```
users-crud/
├── api/
│   ├── api.go         # Configuração do router e handlers
│   ├── handlers.go    # Lógica dos endpoints (CRUD)
│   ├── db.go          # Operações de banco de dados
│   └── user.go        # Struct User e modelos
├── m.go               # Arquivo main da aplicação
├── go.mod             # Dependências do projeto
├── compose.yaml       # Configuração Docker Compose
└── README.md          # Esta documentação
```

### Descrição dos Arquivos

- **`m.go`**: Ponto de entrada da aplicação. Inicializa conexão com BD e inicia servidor HTTP
- **`api/api.go`**: Configura as rotas usando Chi router e middleware
- **`api/handlers.go`**: Implementa os 5 handlers principais (Create, Read, ReadAll, Update, Delete)
- **`api/db.go`**: Funções de persistência e esquema do banco de dados
- **`api/user.go`**: Definição da struct User com tags JSON

---

## 🔒 Segurança

- ✅ Queries parameterizadas para prevenção de SQL Injection
- ✅ Validação rigorosa de entrada
- ✅ Logging de erros estruturado
- ✅ Middleware de recuperação de panics
- ✅ Headers Content-Type validados
- ✅ Timeouts configurados no servidor

---

## 📈 Métricas de Performance

- **Read Timeout**: 10 segundos
- **Write Timeout**: 10 segundos
- **Idle Timeout**: 1 minuto
- **Port**: 8080

---

## 🐛 Troubleshooting

### Erro: "connection refused"

```bash
# Verifique se o PostgreSQL está rodando
docker compose ps

# Reinicie os containers
docker compose restart
```

### Erro: "user not found" em operações

- Verifique se o ID está correto
- Confirme que o usuário foi criado anteriormente
- Use GET /api/users para listar todos

### Erro: "invalid fields"

- Revise os limites de caracteres para cada campo
- Certifique-se de que nenhum campo está vazio
- Respeite as validações na tabela acima

### Erro: "internal server error"

- Verifique os logs da aplicação
- Confirme a conexão com o banco de dados
- Valide o JSON enviado

---

## 📄 Licença

Projeto desenvolvido para fins educacionais.

---
