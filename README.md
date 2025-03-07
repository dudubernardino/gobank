<h3 align="center">
  🏦 GOBANK API Project
</h3>

<p align="center">A simple rest api for bank accounts.</p>

<p align="center">
  <a href="#-about-the-project">About the project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-technologies">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-getting-started">Getting started</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-docs">Docs</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-contribute">How to contribute</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-license">License</a>
</p>

## 📊 About the project

Back-end app built with Golang and Go-Chi.

## 🚀 Technologies

Technologies that I used to develop this api

- [go](https://go.dev/)
- [go-chi](https://go-chi.io/#/)
- [tern](https://github.com/jackc/tern)
- [sqlc](https://sqlc.dev/)
- [pgx](https://github.com/jackc/pgx)
- [air](https://github.com/air-verse/air)
- [docker](https://www.docker.com/)

## 💻 Getting started

### Requirements

- [Go](https://go.dev/)
- One instance of [PostgreSQL](https://www.postgresql.org/)

> Obs.: I recommend to use docker

**Clone the project and access the folder**

```bash
$ git clone https://github.com/dudubernardino/gobank && cd gobank
```

**Follow the steps below**

```bash
# Install the dependencies
$ go mod tidy

# Make a copy of '.env.example' to '.env'
# and set with YOUR environment variables.
# The aws variables do not need to be filled for dev environment
$ cp .env.example .env

# Create the instance of postgreSQL using docker compose
$ docker compose up -d

# To finish, run the api service
$ go run ./cmd/api

# (Optional) Run the Go Server Using Air-Verse for hot reload
$ air --build.cmd "go build -o ./bin/api ./cmd/api" --build.bin "./bin/api"

# Well done, project is started!
```

## 📚 Docs

I used [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) to document the API. You can access `api.http` file.

## 🤔 How to contribute

**Make a fork of this repository**

```bash
# Fork using GitHub official command line
# If you don't have the GitHub CLI, use the web site to do that.

$ gh repo fork dudubernardino/gobank
```

**Follow the steps below**

```bash
# Clone your fork
$ git clone your-fork-url && cd gobank

# Create a branch with your feature
$ git checkout -b my-feature

# Make the commit with your changes
$ git commit -m 'feat: My new feature'

# Send the code to your remote branch
$ git push origin my-feature
```

After your pull request is merged, you can delete your branch

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Made with ❤️ &nbsp;by Eduardo Bernardino 👋 &nbsp;[See my linkedin](https://www.linkedin.com/in/dudubernardino/)
