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
- `go test -v` results: link to asciinema
- API check: link to asciinema

or you can tesh it yourself at: http://ip:8000

**Project tasks:**
- [x] Make tests passed
- [x] Test REST API working
- [x] Write comments for code 
- [ ] Setup CI
- [x] To issue repository
- [ ] Host application on remote server
- [ ] Make screencasts for application
