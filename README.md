![Gopher with MongoDB](https://cdn.cp.adobe.io/content/2/dcx/8182b7fd-7661-4b81-8a2e-276c203ecfa3/rendition/preview.jpg/version/0/format/jpg/dimension/width/size/1200)

# MongoCLI

**C**ommand **L**ine **I**nterface for **Mongo**DB. **MongoCLI**.

Status: development in progress

## Features

- Check database connection, ping: `mongocli ping`
- List existing database names: `mongocli dbs`
- List collection names: `mongocli colls` or `mongocli colls -d <db name>`
- Use configuration file for default connection settings and database name: `~/mongocli.yml` (example configuration file included)
- Use command line args for connection settings: `mongocli ping -s localhost -p 27017 -d config` or `mongocli ping --server localhost --port 27017 --database config`

## Usage

For now MongoCLI distributed only via GitHub.

1. Clone repository

```
git clone https://github.com/KenanBek/mongocli.git
```

2. Test & Build

```
make test
make test/e2e
go build mongocli.go
```
Note: test coverage is not full.

3. Run Mongo with Docker

```
docker run --name mongodb -p 27017:27017 mongo
```

4. Use

```
mongocli ping
```
