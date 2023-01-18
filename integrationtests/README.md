# How to run test

To run the integration tests, Database Mysql instance is required to be running for onchain data storage.
A docker compose file is provided to quickly set it up.

Start the Mysql container in detach mode
```shell script
$ docker-compose -f integrationtests/docker/docker-compose.yml up --detach
```

Shut down the container after test
```shell script
$ docker-compose -f integrationtests/docker/docker-compose.yml down
```