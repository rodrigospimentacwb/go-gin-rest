
# Go Gin CRUD API - POC de Produtos

Este é um projeto de estudo desenvolvido com Go e o framework Gin para explorar um CRUD de produtos, utilizando Docker e Docker Compose para facilitar o desenvolvimento e a execução local.

## Estrutura do Projeto

A estrutura do projeto segue o padrão Clean Code:

```
.
├── cmd
├── controller
├── db
├── model
├── repository
├── docs
└── usecase
```

## Roadmap de Desenvolvimento

1. **Criar um Repositório no GitHub**
    - Inicialize o projeto com um repositório público no GitHub.

2. **Clonar o Repositório e Preparar o Ambiente**
    - Clone o repositório para o ambiente local e abra-o em sua IDE de preferência.

3. **Configuração do Docker**
    - Crie um `Dockerfile` para configurar a imagem com Go e as dependências.
    - Crie um `docker-compose.yml` para definir o serviço e simplificar a execução do ambiente.

4. **Execução do Docker Compose**
    - Execute o comando:
      ```bash
      docker-compose up -d
      ```
    - Isso cria o ambiente de desenvolvimento sem a necessidade de instalar o Go localmente. Use o terminal para acessar o container e executar comandos.

5. **Iniciar o Módulo Go**
    - Dentro do container, inicialize o módulo Go com o comando:
      ```bash
      go mod init github.com/SEU-USER/SEU_PROJETO
      ```

6. **Criar o Arquivo main.go**
    - No diretório principal, crie um arquivo `main.go` com o seguinte conteúdo básico:
      ```go
      package main
 
      import "github.com/gin-gonic/gin"
 
      func main() {
          r := gin.Default()
          // Configuração das rotas
          r.Run()
      }
      ```

7. **Instalar o Gin**
    - No container, instale o framework Gin:
      ```bash
      go get -u github.com/gin-gonic/gin
      ```

## Executar/Desenvolver localmente

Siga os passos abaixo para configurar e executar o projeto localmente:

1. **Configurar o HOST_OS**:

   No arquivo `docker-compose.yml`, defina a variável `HOST_OS` de acordo com o seu sistema operacional (`mac` ou `linux`):

   ```yaml
   environment:
     - HOST_OS=mac # ou linux
   ```

2. **Iniciar os containers**:

   Execute o comando para iniciar os containers em segundo plano:

   ```bash
   docker compose up -d
   ```

3. **Acessar o container**:

   Entre no container do projeto:

   ```bash
   docker exec -it go-rest-gin /bin/bash
   ```

4. **Instalar as dependências**:

   Dentro do container, execute:

   ```bash
   go mod tidy
   ```

5. **Executar a aplicação**:

   Ainda dentro do container, inicie a aplicação:

   ```bash
   go run ./cmd/main.go
   ```

## Configuração para Produção

Para otimizar a imagem final, foi criado um `Dockerfile-prod` com dois estágios de build, que gera uma imagem enxuta para o ambiente de produção. O `docker-compose-prod.yaml` permite testar essa imagem localmente.

### Comandos de Build e Execução para Produção

1. **Build da Imagem**
   ```bash
   docker build -t go-gin-rest-prod -f Dockerfile-prod .
   ```

2. **Execução do Ambiente de Produção**
   ```bash
   docker compose -f docker-compose-prod.yaml up -d
   ```

## Scripts e Collection Postman

- Na pasta `docs`, foi incluído um script SQL para criar a tabela de produtos e inserir um produto de teste.
- Uma collection do Postman também foi adicionada para facilitar a utilização e testes dos endpoints do CRUD de produtos.
   
## Licença

Este projeto é para fins de aprendizado e não possui uma licença específica.
