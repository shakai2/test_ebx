# API de Gerenciamento de Contas Bancárias

Esta API, desenvolvida em Go com o framework Gin, foi criada para gerenciar operações bancárias básicas de forma simples e eficiente. Ela permite realizar ações como depósitos, saques e transferências entre contas, utilizando um armazenamento em memória para manter os dados das contas e transações.

## O que ela faz?

A API oferece funcionalidades essenciais para simular um sistema bancário básico:

- **Reset**: Limpa todas as contas e transações, restaurando o sistema ao estado inicial.
- **Consulta de saldo**: Retorna o saldo de uma conta específica com base no ID fornecido.
- **Operações bancárias**:
  - **Depósito**: Adiciona um valor a uma conta.
  - **Saque**: Remove um valor de uma conta, se houver saldo suficiente.
  - **Transferência**: Move um valor de uma conta de origem para uma conta de destino, verificando o saldo disponível.

## Como funciona?

A API é baseada em requisições HTTP e utiliza endpoints específicos:

1. **POST /reset**: Zera todas as contas e transações.
2. **GET /balance?account_id=<ID_DA_CONTA>**: Consulta o saldo de uma conta.
3. **POST /event**: Realiza operações bancárias (depósito, saque ou transferência) enviando um JSON com os detalhes da transação.

### Exemplo de uso

- **Depósito**:
  ```json
  {
    "type": "deposit",
    "destination": "100",
    "amount": 10
  }
  ```
  Resposta:
  ```json
  {
    "destination": {
      "id": "100",
      "balance": 10
    }
  }
  ```

- **Saque**:
  ```json
  {
    "type": "withdraw",
    "origin": "100",
    "amount": 5
  }
  ```
  Resposta:
  ```json
  {
    "origin": {
      "id": "100",
      "balance": 5
    }
  }
  ```

- **Transferência**:
  ```json
  {
    "type": "transfer",
    "origin": "100",
    "destination": "200",
    "amount": 5
  }
  ```
  Resposta:
  ```json
  {
    "origin": {
      "id": "100",
      "balance": 0
    },
    "destination": {
      "id": "200",
      "balance": 5
    }
  }
  ```

## Como executar

Este projeto é escrito em Go, que é uma linguagem compilada. Para rodar a API, você precisa compilar o código antes de executá-lo. Siga os passos abaixo:

1. Certifique-se de ter o Go instalado (versão 1.24.2 ou superior recomendada).
2. Clone este repositório.
3. No diretório do projeto, execute o comando para instalar as dependências:
   ```bash
   go mod tidy
   ```
4. Compile e execute o projeto com:
   ```bash
   go run main.go
   ```
   Ou, para criar um binário executável:
   ```bash
   go build -o api-bancaria
   ./api-bancaria
   ```
5. A API estará disponível em `http://localhost:8080`.

## Observações

- Os dados são armazenados em memória, ou seja, não persistem após o servidor ser reiniciado.
- Para saques e transferências, a API verifica se há saldo suficiente e retorna erro caso a conta não exista ou o saldo seja insuficiente.
- Como o Go requer compilação, certifique-se de recompilar o projeto após qualquer alteração no código.
