## Setup Local Relayers


You might want to also bootstrap multiple relayers when running multiple `Greenfield` Nodes in local

1. Build binary
```bash
make build
```

2. Fill in config files under `./config/local`, input private keys got from `greenfield`


3. start n instance of relayer 
```bash
bash ./deployment/localup/localup.sh start ${SIZE}
```

4. stop relayer
```bash
bash ./deployment/localup/localup.sh stop
```


