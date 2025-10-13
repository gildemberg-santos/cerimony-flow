# Cerimony Flow

**Cerimony Flow** é um site para compartilhar com os convidados os detalhes do seu casamento: data, local, fotos, história do casal e a lista de presentes. Uma espécie de “site oficial do casamento”, com visual elegante, responsivo e fácil de navegar.

---

## 🥂 Funcionalidades principais

* Exibição da data, horário e local do casamento
* Galeria de fotos do casal e momentos especiais
* Informações de mapa / endereço com link para GPS
* Lista de presentes com integração para que os convidados vejam ideias de presentes
* Versão mobile responsiva

---

## Arquitetura & Tecnologias

Embora os detalhes exatos dependam do que já foi implementado no repositório, a estrutura ideal pode ser algo como:

* **Frontend**: React.js — escolha moderna para interface bonita
* **Backend**: API em Go para servir dados dinâmicos (lista de presentes, fotos, etc.)
* **Hospedagem / Deployment**: pode usar serviço estático (OnRender) para frontend + backend, tudo em container Docker
* **Armazenamento de fotos**: pode usar serviços de armazenamento em nuvem (S3, Google Cloud Storage) ou diretamente no servidor

---

## Pré-requisitos

Para executar localmente:

* Node.js / npm ou yarn (para frontend)
* Ambiente backend compatível (Go)
* Git
* (Opcional) Docker

---

## Como instalar & executar

### 1. Backend

1. Vá para a pasta do backend (ex: `backend/`)
2. Instale dependências (ex: `go mod tidy`)
3. Defina as variáveis de ambiente (veja seção abaixo)
4. Rode o servidor (ex: `go run main.go`)

### 2. Frontend

1. Vá para o diretório do frontend (ex: `frontend/`)
2. Instale dependências (`npm install` ou `yarn install`)
3. Configure variáveis de ambiente (URL da API, chaves etc.)
4. Inicie o servidor de desenvolvimento (`npm start` ou `yarn start`)

### 3. Docker (opcional)

Se quiser empacotar tudo em containers:

```bash
docker-compose up --build
```

Ou, se for um Dockerfile simples:

```bash
docker build -t cerimony-flow .
docker run -p 3000:3000 cerimony-flow
```

Ajuste as portas conforme o projeto.

---

## Estrutura sugerida de pastas

Aqui vai uma sugestão de organização:

```
cerimony-flow/
├── backend/
│   ├── controllers/
│   ├── models/
│   ├── routes/
│   ├── services/
│   └── main.go (ou server.js etc.)
├── frontend/
│   ├── public/
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   ├── assets/
│   │   └── services/ (para comunicação com API)
│   └── package.json
├── .env.example
├── docker-compose.yml
├── Dockerfile
└── README.md
```

---

## Contribuição

Mesmo sendo um projeto mais pessoal, você pode organizar contribuições:

1. Abra *issues* para melhorias visuais, correção de bugs ou novas funcionalidades
2. Use *branches* descritivas
3. Documente bem o que for implementado
4. Faça *pull requests* bem descritas
