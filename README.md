# inscription-relayer

Inscription Relayer is a bidirectional relaying tool serving Inscription and BSC cross-chain needs:
1. cross-chain transactions from Inscription to BSC
2. cross-chain packages from BSC to Inscription


## Build

Build binary:

```shell script
$ make build
```

## Run

Run locally:

```shell script
$ ./build/inscription-relayer --config-type [local or aws] --config-path config_file_path  --aws-region [aws region or omit] --aws-secret-key [aws secret key for config or omit]
```
