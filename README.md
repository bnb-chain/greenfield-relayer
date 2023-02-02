# greenfield-relayer

Greenfield Relayer is a bidirectional relaying tool serving communication
requirement between Greenfield and BSC. It is a standalone process that can only be run by Greenfield validators. 
The relayer watches cross chain events happen on BSC and Greenfield independently and persist them into the database, 
after a few blocks confirmation to reach finality, the relayer will sign a message by the BLS private key to confirm the event, 
and broadcast the signed event which called "the vote" through the p2p network on Greenfield network. Once enough votes from 
the Greenfield relayer are collected, the relayer will assemble a cross chain package transaction and submit it to the 
BSC or Greenfield network


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
$ ./build/greenfield-relayer --config-type [local or aws] --config-path config_file_path  --aws-region [aws region or omit] --aws-secret-key [aws secret key for config or omit]
```

Run docker:
```shell script
$ docker run -it -v /your/data/path:/greenfield-relayer -e CONFIG_TYPE="local" -e CONFIG_FILE_PATH=/your/config/file/path/in/container -d greenfield-relayer
```