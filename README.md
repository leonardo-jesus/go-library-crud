
<p align="center">
	<h1 align="center">Go Library :green_heart:</h1>
</p>

<p align="center">
  <a href="#-Must-Have">Must-Have</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-Improvement-Points">Improvement Points</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-Resources">Resources</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-Technologies">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-Project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-Requirements">Requirements</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-How-to-Run">How to Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-Work-Environment">Work Environment</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-How-to-contribute">How to contribute</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#memo-license">License</a>
</p>

<p align="center">
 <img src="https://img.shields.io/static/v1?label=PRs&message=welcome&color=7159c1&labelColor=000000" alt="PRs welcome!" />

  <img alt="License" src="https://img.shields.io/static/v1?label=license&message=MIT&color=7159c1&labelColor=000000">
</p>

<br>

## ✅ Must-Have

✔️ You should implement an application for a library to store book and authors data(**This application must provide an HTTP REST API to attend the requirements**).\
✔️ Receive a CSV with authors and import to database. Given a CSV file with many authors (_authors.csv_), you need to build a command to import the data into the database.\
✔️ You need to store the authors' data to complement the book data that will be stored\
✔️ Expose authors' data in an endpoint. This endpoint needs to return a paginated list with the authors' data. Optionally the authors can be searched by name.\
✔️ CRUD (Create, Read, Update and Delete) of books.\
✔️ To retrieve a book we can filter by 4 fields (name, publicationYear, edition and author). But these 4 filters are optional. It must be possible to navigate all the books' data without any filter.\
✔️ To create a book, the user should be able to send a field author, with the book's authors' ids.\
✔️ Provide a working environment with your project (eg. Heroku).\
✔️ Application must be written in Go.\
✔️ Every text or code must be in English.\
✔️ Write the project documentation containing:\
✔️ Provide API documentation (in English);

## 💪 Improvement Points

Due to the tight deadline, it was not possible to accomplish everything I had planned initially. I originally intended to follow the Pub/Sub Pattern, where a REST API would handle incoming requests, and for heavier tasks such as creation and updating, it would send them to a RabbitMQ queue to be consumed by a Go consumer. This is why the API is located inside the go-rest-api folder. However, due to time constraints, I ended up simplifying the implementation significantly and having only one API handle everything.
- Pub/Sub pattern implementation to store and retrieve data (eg. RabbitMQ)
- Microservices architecture (Gateway -> Message Boker -> Consumer)
- Redis caching to frequent GET requests
- Goroutine implementation to insert bulk data into database
- Unit and integration tests

## 📦 Resources

- [API's Deploy on Fly.io](https://go-library.fly.dev)
- [API's Documentation at Postman](https://documenter.getpostman.com/view/11958037/2s93zB62j2)

## 🚀 Technologies

This project was developed with the following technologies:

- [Go](https://go.dev/)
- [Postgres](https://www.postgresql.org/)
- [Fiber](https://docs.gofiber.io/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [pgx](https://pkg.go.dev/github.com/jackc/pgx)
- [Fly.io](http://fly.io/)
- [Neon.tech](https://neon.tech/)
- [migrate/migrate](https://github.com/golang-migrate/migrate)
- [mergo](https://github.com/imdario/mergo)

## 💻 Project
The ***Go Library*** is an application that uses **Go**, ***with the purpose to achieve my goal which is to work at Dialog***💚. The API saves *6.8M authors* in *Postgres DB*, and makes available a *Books' CRUD*. Made with **Go**, **Postgres**, **Fiber**, **Docker** and **pgx**.

## 🔧 Requirements
To run this Go application, you will need to have only ***Docker** and **Docker Compose** on your computer.

[🐳 **Docker**](https://www.docker.com/get-started) allows you to create and manage containers that encapsulate your application and its dependencies.

[**Docker Compose**](https://docs.docker.com/compose/install/) is a tool that allows you to define and manage multi-container Docker applications. Docker Compose usually comes bundled with Docker, so if you have installed Docker using the official installer, Docker Compose should be included. You can verify its installation by running the command `docker compose version` or `docker-compose --version` in your terminal.

## 🏃 How to Run

1. Make a clone;
2. Open the project on your terminal;
3. Run `cd ./go-rest-api`
4. Run `cp .env.example .env` to copy environment variables
5. Run `docker compose run --rm api go mod tidy` to install Dependencies;
6. Run `docker compose --profile tools run migration` to create the tables on database
7. Run `docker compose up` to run the server;
- By default the server will run at `localhost:3333` and the Postgres DB at `localhost:5432`.
#### More informations at [API's Documentation](https://documenter.getpostman.com/view/11958037/TVYKZvnE).

## ⚙️ Work Environment

- AMD Ryzen 5 3600, 16GB RAM, GTX 1660
- WSL 2 (Ubuntu 22.04.2 LTS)
- Visual Studio Code
- Docker version 23.0.5
- Docker Compose version v2.17.3
- Go 1.20.4 linux/amd64

## 🤔 How to contribute

- Make a fork;
- Create a branch with your feature: `git checkout -b my-feature`;
- Do commit with your changes: `git commit -m 'feat: My new feature'`;
- Do a push for your branch: `git push origin my-feature`.

After the merge of your pull request was made, you can delete your branch.

## 📝 License

This project is under License MIT. See the documentation [LICENSE](LICENSE) for more details.

---

<p align="center">Developed by <a href="https://www.linkedin.com/in/leonardojesus02/">Leonardo Jesus </a>:copyright:

