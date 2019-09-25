# Goal

Have a tool (CLI) to easily request data from mongo database.

## Technical

1. Support database connection check and ping
2. Support configuration via a configuration file

```
    # connect & ping via default configuration file mongocli.yml
    mongocli ping

    # connect & ping via specified configuration file custom.yml
    mongocli -f custom.yml
```

3. Support optional configuration arguments

```
    # use provided connection arguments instead of the configuration file
    mongocli -h localhost -p 27017 -db test
```

4. Support a list request

```
    # list top data in collection coll1
    mongocli list coll1
```

5. Support a create request

```
    # create a new document in collection col1
    mongocli create coll1 "{ title: 'title1', desc: 'desc1'}"
```

6. Support a create request by using input data from a file

```
    # create a document by using data from a json file
    mongocli create coll1 -d col1-doc1.json
```
