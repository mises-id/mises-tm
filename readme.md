# Table of Contents <!-- omit in toc -->

- [What is Mises?](#what-is-mises)
- [Build](#build)
- [Production Environment](#production)
- [Full Node](#full-node)
  - [Join the mainnet](#join-the-mainnet)
  - [Run a single node testnet](#run-a-single-node-testnet)
- [Run as service](#service)
- [Contributing](#contributing)

# The Mises Chain
Mises is a social network protocol based on blockchain technology, it helps developer build decentralized social media applications on blockchain.


# Build
**Step 1. Install Golang**

Go v1.17+ or higher is required for Mises.

If you haven't already, install Golang by following the [official docs](https://golang.org/doc/install). Make sure that your `GOPATH` and `GOBIN` environment variables are properly set up.

**Step 2: Get Mises source code**

Use `git` to retrieve Mises from the [official repo](https://github.com/mises-id/mises-tm) and checkout the `main` branch. This branch contains the latest stable release, which will install the `misestmd` binary.
```bash
git clone https://github.com/mises-id/mises-tm/
cd mises-tm/
git checkout main
```

**Step 3: Build Mises**

Run the following command to install the executable `misestmd` to your `GOPATH` and build Mises. `misestmd` is the node daemon and CLI for interacting with a Mises node.

```bash
make install
```

**Step 4: Verify your installation**

Verify that you've installed misestmd successfully by running the following command:

```bash
misestmd version --long
```

If misestmd is installed correctly, the following information is returned:

```bash
name: mises-tm
server_name: misestmd
version: 0.0.2-28-gafca505
commit: afca505c863146d0b09733daf23965f066bbcd79
build_tags: netgo
go: go version go1.17.7 darwin/amd64
build_deps:
...
```


# `misestmd` CLI

`misestmd` is the all-in-one command for operating and interacting with a running Mises node. For comprehensive coverage on each of the available functions. To view various subcommands and their expected arguments, use the `$ misestmd --help` command:

<pre>
    <div align="left">
        <b>$ misestmd --help</b>

        Stargate Mises App

        Usage:
          misestmd [command]

        Available Commands:
        add-genesis-account Add a genesis account to genesis.json
        collect-gentxs      Collect genesis txs and output a genesis.json file
        config              Create or query an application CLI configuration file
        debug               Tool for helping with debugging your application
        eth_keys            Manage your application's ethereum keys
        export              Export state to JSON
        gentx               Generate a genesis tx carrying a self delegation
        help                Help about any command
        init                Initialize private validator, p2p, genesis, and application configuration files
        keys                Manage your application's keys
        light               Run a light client proxy server, verifying Tendermint rpc
        migrate             Migrate genesis to a specified target version
        query               Querying subcommands
        rest                Run a rest server
        rosetta             spin up a rosetta server
        start               Run the full node
        status              Query remote node for status
        tendermint          Tendermint subcommands
        testnet             Initialize files for a simapp testnet
        tx                  Transactions subcommands
        unsafe-reset-all    Resets the blockchain database, removes address book files, and resets data/priv_validator_state.json to the genesis state
        validate-genesis    validates the genesis file at the default location or at the location passed as an arg
        version             Print the application binary version information

        Flags:
        -h, --help                help for misestmd
            --home string         directory for config and data (default "/Users/baoge/.misestm")
            --log_format string   The logging format (json|plain) (default "plain")
            --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
            --trace               print out full stack trace on errors

        <b>Use "misestmd [command] --help" for more information about a command.</b>
    </div>
</pre>

# Set up a production environment

Use the following information to set up and manage your production-level full Mises node.  

For information about running a validator node, visit the [validator guide](docs/VALIDATOR.md).



## Create a dedicated user

Although `misestmd` does not require a super user account, during the setup process you'll need super user permission to create and modify some files. It is strongly recommended to use a normal user when running `misestmd`.  

## Increase the maximum files `misestmd` can open

`misestmd` is set to open 1024 files by default. It is recommended that you increase this amount.

Modify `/etc/security/limits.conf`[*](https://linux.die.net/man/5/limits.conf) to increase the amount, where `nofile` is the number of files `misestmd` can open.

```bash
# If you have never changed this system config or your system is fresh, most of this file will be commented
# ...
*                soft    nofile          65535   # Uncomment the following two lines at the bottom
*                hard    nofile          65535   # Change the default values to ~65535
# ...
```


# Full Node

Once you have `misestm` installed, you will need to set up your node to be part of the network.

## Join the mainnet

The following requirements are recommended for running a `mises` mainnet node:

- **4 or more** CPU cores
- At least **256GB** of disk storage
- At least **100mbps** network bandwidth
- An Linux distribution


**Mises Node Quick Start**
```
misestmd init mises --chain-id mainnet

sudo snap install jq
curl https://e1.mises.site:443/genesis | jq .result.genesis > ~/.misestm/config/genesis.json

PERSISTENT_PEERS="40a8318fa18fa9d900f4b0d967df7b1020689fa0@e1.mises.site:26656"

sed -i.bak -E "s|^(persistent_peers[[:space:]]+=[[:space:]]+).*$|\1\"$PERSISTENT_PEERS\"|"  ~/.misestm/config/config.toml

misestmd start

```

**Enable State Sync**
You can enable the state sync for faster block replay.
first query the lastet block height and hash from mises block explorer https://gw.mises.site/blocks
then edit the config.toml to enable it.

```
read BLOCK_HASH BLOCK_HEIGHT < <(echo $(curl https://e1.mises.site:443/block -s | jq -r '.result.block_id.hash,.result.block.header.height')) 

sed -i.bak -E "s|^(enable[[:space:]]+=[[:space:]]+).*$|\1true| ; \
s|^(rpc_servers[[:space:]]+=[[:space:]]+).*$|\1\"http://e1.mises.site:26657,http://e2.mises.site:26657,http://w1.mises.site:26657,http://w2.mises.site:26657\"| ; \
s|^(trust_height[[:space:]]+=[[:space:]]+).*$|\1$BLOCK_HEIGHT| ; \
s|^(trust_hash[[:space:]]+=[[:space:]]+).*$|\1\"$BLOCK_HASH\"|"  ~/.misestm/config/config.toml

vi ~/.misestm/config/config.toml

//you should see something like this in config.toml
[statesync]
enable = true
rpc_servers = "https://e1.mises.site:443,https://e2.mises.site:443,https://w1.mises.site:443,https://w2.mises.site:443"
trust_height = 38188  //should be the latest height
trust_hash = "8AF6C7C7607A5C49ECCEB355DD82E8479922A1CDCD6D9F4F0E7C620A2259587F" //should be the latest hash

```

**Enable mongodb backend for mises social data**
You can also store the mises social data to your local mongodb, that will help you integrate your own mises dapp with the on-chain social data.
```
misestmd start --mises-use-mongodb mongodb://127.0.0.1:27017
```


## Run a single node testnet

You can also run a local testnet using a single node. On a local testnet, you will be the sole validator signing blocks.


**Step 1. Create network and account**

First, initialize your genesis file to bootstrap your network. Create a name for your local testnet and provide a moniker to refer to your node:

```bash
misestmd init <node_moniker> --chain-id=<testnet_name> 
```

Next, create a Mises account by running the following command:

```bash
misestmd keys add <account_name>
```

**Step 2. Add account to genesis**

Next, add your account to genesis and set an initial balance to start. Run the following commands to add your account and set the initial balance:

```bash
misestmd add-genesis-account $(misestmd keys show <account_name> -a) 100000000umis
misestmd gentx <account_name> 100000000umis --chain-id=<testnet_name>
misestmd collect-gentxs
```

**Step 3. Run Mises daemon**

Now you can start your private Mises network:

```bash
misestmd start
```

Your `misestmd` node will be running a node on `tcp://localhost:26656`, listening for incoming transactions and signing blocks.

Congratulations, you've successfully set up your local Mises network!



# Run the server as a daemon

`misestmd` must be running at all times. It is recommended that you register `misestmd` as a `systemd` service so that it will be started automatically when the system reboots.

## Register `misestmd` as a service

1. Create a service definition file in `/etc/systemd/system/misestmd.service`.

     **Example**:

     ```bash
     [Unit]
     Description=Mises Daemon
     After=network.target

     [Service]
     Type=simple
     User=<Mises_USER>
     ExecStart=<PATH_TO_MISESTMD>/misestmd start  
     Restart=on-abort

     [Install]
     WantedBy=multi-user.target

     [Service]
     LimitNOFILE=65535  
     ```

2. Modify the `Service` section according to your environment:

   - Enter the user (likely your username, unless you created a user specifically for `misestmd`)
   - Enter the path to the `misestmd` executable. `<PATH_TO_MISESTMD>` is likely `/home/<YOUR_USER>/go/bin/misestmd` or `/usr/go/bin`. Confirm this with `whereis misestmd`
   - Make sure you made the correct edits to /etc/security/limits.conf


3. Run `systemctl daemon-reload` followed by `systemctl enable misestmd`. This will register `misestmd` as a system service and turn it on upon startup.

4. Now start the serivce with `systemctl start misestmd`.

### Controlling the service

Use `systemctl` to start, stop, and restart the service:

```bash
# Check health
systemctl status misestmd
# Start
systemctl start misestmd
# Stop
systemctl stop misestmd
# Restart
systemctl restart misestmd
```

### Access logs

Use `journalctl -t` to access entire logs, entire logs in reverse, and the latest and continuous log.

```bash
# Entire log reversed
journalctl -t misestmd -r
# Entire log
journalctl -t misestmd
# Latest and continuous
journalctl -t misestmd -f
```

# Contributing

If you are interested in contributing to Mises source, please review our [code of conduct](./coc.md).
