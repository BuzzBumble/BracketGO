
# BracketGO
A lightweight tournament bracket creator and visualizer
## Tech Stack

**Client:** NextJS, React, TypeScript

**Server:** Go

**Database:** PostgreSQL


## Run Locally

Clone the project

```bash
  git clone https://github.com/BuzzBumble/BracketGO/
```

Go to the project directory

```bash
  cd BracketGO
```

Build & Run Docker containers

```bash
  docker compose build
  docker compose up -d
```

Access/View the database with:

```bash
  docker exec -it db psql -U postgres
```

Access the Web App at `localhost:3000`
