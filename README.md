# MongoCLI

Command line management tool for MongoDB.

## TODO

- Func `ListDocuments` of `pkg.mongo` now directly prints to console
- Abstraction layer for mongo's `Cursor`
- Order for config file lookup: current dir, home dir

## Try Locally

1. Clone repository

    git clone ...

2. Test & Build

    make

3. Run Mongo with Docker

    docker run --name mongodb -p 27017:27017 mongo

4. Use

    mongocli ping

