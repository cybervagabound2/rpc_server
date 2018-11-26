# rpc_server

This is my test exercise to get offer from awesome tech company.

**What this application can do?**

- Create a new user in response to a valid POST request at /user
- Update a user in response to a valid PUT request at /user/{id}
- Delete a user in response to a valid DELETE request at /user/{id}
- Fetch a user in response to a valid GET request at /user/{id}
- Fetch a list of users in response to a valid GET request at /users

**Installation instructions:**
- install [PostgreSQL](https://www.postgresql.org/download/)
- clone repository: `git clone https://github.com/cybervagabound2/rpc_server`
- install all necessary dependencies:
`go get github.com/gorilla/mux github.com/lib/pq`
- export following environment variables:
`TEST_DB_USERNAME`
`TEST_DB_PASSWORD`
`TEST_DB_NAME`
`APP_DB_USERNAME`
`APP_DB_PASSWORD`
`APP_DB_NAME`
- in project directory start tests:
`go test -v`
- build application:
`go build`
- run application:
`./rpc_server`

**Application info:**
- `go test -v` results: https://asciinema.org/a/ibMzb02oVDVTKV4gzXuuHSq1m
- API check: https://asciinema.org/a/OlzoWwRpqKP7swu6UROClY5eT

**Project tasks:**
- [x] Make tests passed
- [x] Test REST API working
- [x] Write comments for code 
- [x] To issue repository
- [x] Make screencasts for application
- [ ] Setup AMQP
- [ ] Try any cryptoexchange API
- [ ] Setup CI
- [ ] Host application on remote server
- [ ] Enable Docker support
