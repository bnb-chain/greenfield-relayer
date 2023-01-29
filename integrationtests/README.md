# How to run test in local

1. Set up BSC and Greenfield and modify the `config_test.json` file
2. A Mysql instance is required to be running for onchain data storage. A docker compose file is provided to quickly set it up.

Start the Mysql container in detach mode
```shell script
$ docker-compose -f integrationtests/docker/docker-compose.yml up --detach
```

Shut down the container after test
```shell script
$ docker-compose -f integrationtests/docker/docker-compose.yml down
```