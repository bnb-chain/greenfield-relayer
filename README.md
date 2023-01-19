# inscription-relayer

Inscription Relayer is a bidirectional relaying tool serving communication
requirement between Inscription and BSC. It is a standalone process that can only be run by Inscription validators. The relayer watches cross chain events happen on BSC and Inscription independently and persist them into the database, after a few blocks confirmation to reach finality, the relayer will sign a message by the BLS key to confirm the event, and broadcast the signed event which called vote through the p2p network on Inscription network. Once enough votes from the Inscription relayer are collected, the relayer will assemble a cross chain package transaction and submit it to the BSC or Greenfield network


## Build

Build binary:

```shell script
$ make build
```

Build docker image:

```shell script
$ make build_docker
```


## Run

Run locally:

```shell script
$ ./build/inscription-relayer --config-type [local or aws] --config-path config_file_path  --aws-region [aws region or omit] --aws-secret-key [aws secret key for config or omit]
```

Run docker:
```shell script
$ docker run -it -v /your/data/path:/inscription-relayer -e CONFIG_TYPE="local" -e CONFIG_FILE_PATH=/your/config/file/path/in/container -d inscription-relayer
```