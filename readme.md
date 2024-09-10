# AwesomeHA

This is a fully handmade home assistant project.

Features:
- Ollama Conductivity
- Typesafe Infrastructure
- Fully local with optional firebase integration
- Ability to customize hardware interfacing with esp32/8266 


## Getting Started With Dev

#### Frontend Webapp

```bash
cd web/AwesomHA
npm install
npm run dev # start dev server
```

#### Server

```bash
cd server
go mod download
go run . # start db server
```

#### Android App

```bash
cd app/AwesomeHA
npm install
npm run start  # start app
```

## Commit Changes

#### Rules
- Create a new branch and work in there
- Once finished, commit your changes and create a PR
