### Table of Contents
1. [About](#About)
2. [Getting Started](#GettingStarted)
    1. [Installation](#Installation)
        1. [Windows](#WindowsInstallation)
        2. [Linux/BSD/MacOSX/POSIX](#PosixInstallation)
    2. [Configuration](#Configuration)
    3. [Controlling and Querying btcd via btcctl](#BtcctlConfig)
    4. [Mining](#Mining)
3. [Help](#Help)
    1. [Startup](#Startup)
        1. [Using bootstrap.dat](#BootstrapDat)
    2. [Network Configuration](#NetworkConfig)
    3. [Wallet](#Wallet)
4. [Contact](#Contact)
    1. [IRC](#ContactIRC)
    2. [Mailing Lists](#MailingLists)
5. [Developer Resources](#DeveloperResources)
    1. [Code Contribution Guidelines](#ContributionGuidelines)
    2. [JSON-RPC Reference](#JSONRPCReference)
    3. [Conformal Bitcoin-related Go Packages](#GoPackages)

<a name="About" />
### 1. About
btcd is a full node bitcoin implementation written in [Go](http://golang.org),
licensed under the [copyfree](http://www.copyfree.org) ISC License.

This project is currently under active development and is in a Beta state. It is
extremely stable and has been in production use for over 6 months as of May
2014, however there are still a couple of major features we want to add before
we come out of beta.

It currently properly downloads, validates, and serves the block chain using the
exact rules (including bugs) for block acceptance as the reference
implementation, [bitcoind](https://github.com/bitcoin/bitcoin).  We have taken
great care to avoid btcd causing a fork to the block chain. It passes all of
the '[official](https://github.com/TheBlueMatt/test-scripts/)' block acceptance
tests.

It also properly relays newly mined blocks, maintains a transaction pool, and
relays individual transactions that have not yet made it into a block. It
ensures all individual transactions admitted to the pool follow the rules
required into the block chain and also includes the vast majority of the more
strict checks which filter transactions based on miner requirements ("standard"
transactions).

One key difference between btcd and bitcoind is that btcd does NOT include
wallet functionality and this was a very intentional design decision. See the
[blog entry](https://blog.conformal.com/btcd-not-your-moms-bitcoin-daemon) for
more details. This means you can't actually make or receive payments directly
with btcd. That functionality is provided by the
[btcwallet](https://github.com/btcsuite/btcwallet) and
[btcgui](https://github.com/btcsuite/btcgui) projects which are both under
active development.

<a name="GettingStarted" />
### 2. Getting Started

<a name="Installation" />
**2.1 Installation**<br />

The first step is to install btcd.  See one of the following sections for
details on how to install on the supported operating systems.

<a name="WindowsInstallation" />
**2.1.1 Windows Installation**<br />

* Install the MSI available at: https://github.com/btcsuite/btcd/releases
* Launch btcd from the Start Menu

<a name="PosixInstallation" />
**2.1.2 Linux/BSD/MacOSX/POSIX Installation**<br />

* Install Go according to the installation instructions here: http://golang.org/doc/install
* Run the following command to ensure your Go version is at least version 1.2: `$ go version`
* Run the following command to obtain btcd, its dependencies, and install it: `$ go get github.com/btcsuite/btcd/...`<br />
  * To upgrade, run the following command: `$ go get -u github.com/btcsuite/btcd/...`
* Run btcd: `$ btcd`

<a name="Configuration" />
**2.2 Configuration**<br />

btcd has a number of [configuration](http://godoc.org/github.com/btcsuite/btcd)
options, which can be viewed by running: `$ btcd --help`.

<a name="BtcctlConfig" />
**2.3 Controlling and Querying btcd via btcctl**<br />

btcctl is a command line utility that can be used to both control and query btcd
via [RPC](http://www.wikipedia.org/wiki/Remote_procedure_call).  btcd does
**not** enable its RPC server by default;  You must configure at minimum both an
RPC username and password:

* btcd.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
```
* btcctl.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
```
For a list of available options, run: `$ btcctl --help`

<a name="Mining" />
**2.4 Mining**<br />
btcd supports both the `getwork` and `getblocktemplate` RPCs although the
`getwork` RPC is deprecated and will likely be removed in a future release.<br />

**1. Add the payment addresses with the `miningaddr` option.**<br />

```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
miningaddr=12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX
miningaddr=1M83ju3EChKYyysmM2FXtLNftbacagd8FR
```

**2. Add btcd's RPC TLS certificate to system Certificate Authority list.**<br />

`cgminer` uses [curl](http://curl.haxx.se/) to fetch data from the RPC server.
Since curl validates the certificate by default, we must install the `btcd` RPC
certificate into the default system Certificate Authority list.

**Ubuntu**<br />

1. Copy rpc.cert to /usr/share/ca-certificates: `# cp /home/user/.btcd/rpc.cert /usr/share/ca-certificates/btcd.crt`<br />
2. Add btcd.crt to /etc/ca-certificates.conf: `# echo btcd.crt >> /etc/ca-certificates.conf`<br />
3. Update the CA certificate list: `# update-ca-certificates`<br />

**3. Set your mining software url to use https.**<br />

`$ cgminer -o https://127.0.0.1:8334 -u rpcuser -p rpcpassword`

<a name="Help" />
### 3. Help

<a name="Startup" />
**3.1 Startup**<br />

Typically btcd will run and start downloading the block chain with no extra
configuration necessary, however, there is an optional method to use a
`bootstrap.dat` file that may speed up the initial block chain download process.

<a name="BootstrapDat" />
**3.1.1 bootstrap.dat**<br />
* [Using bootstrap.dat](https://github.com/btcsuite/btcd/tree/master/docs/using_bootstrap_dat.md)

<a name="NetworkConfig" />
**3.1.2 Network Configuration**<br />
* [What Ports Are Used by Default?](https://github.com/btcsuite/btcd/tree/master/docs/default_ports.md)
* [How To Listen on Specific Interfaces](https://github.com/btcsuite/btcd/tree/master/docs/configure_peer_server_listen_interfaces.md)
* [How To Configure RPC Server to Listen on Specific Interfaces](https://github.com/btcsuite/btcd/tree/master/docs/configure_rpc_server_listen_interfaces.md)
* [Configuring btcd with Tor](https://github.com/btcsuite/btcd/tree/master/docs/configuring_tor.md)

<a name="Wallet" />
**3.1 Wallet**<br />

btcd was intentionally developed without an integrated wallet for security
reasons.  Please see [btcwallet](https://github.com/btcsuite/btcwallet) for more
information.

<a name="Contact" />
### 4. Contact

<a name="ContactIRC" />
**4.1 IRC**<br />
* [irc.conformal.com](irc://irc.conformal.com:+6697), port 6697, tls, channel #btcd

<a name="MailingLists" />
**4.2 Mailing Lists**<br />
* <a href="mailto:btcd+subscribe@opensource.conformal.com">btcd</a>: discussion
  of btcd and its packages.
* <a href="mailto:btcd-commits+subscribe@opensource.conformal.com">btcd-commits</a>:
  readonly mail-out of source code changes.

<a name="DeveloperResources" />
### 5. Developer Resources

<a name="ContributionGuidelines" />
* [Code Contribution Guidelines](https://github.com/btcsuite/btcd/tree/master/docs/code_contribution_guidelines.md)
<a name="JSONRPCReference" />
* [JSON-RPC Reference](https://github.com/btcsuite/btcd/tree/master/docs/json_rpc_api.md)
    * [RPC Examples](https://github.com/btcsuite/btcd/tree/master/docs/json_rpc_api.md#ExampleCode)
<a name="GoPackages" />
* Conformal Bitcoin-related Go Packages:
    * [btcrpcclient](https://github.com/btcsuite/btcrpcclient) - Implements a
	  robust and easy to use Websocket-enabled Bitcoin JSON-RPC client
    * [btcjson](https://github.com/btcsuite/btcjson) - Provides an extensive API
	  for the underlying JSON-RPC command and return values
    * [btcws](https://github.com/btcsuite/btcws) - Custom types for btcd
	  websocket extension commands (registers the extension commands with
	  [btcjson](https://github.com/btcsuite/btcjson))
    * [wire](https://github.com/btcsuite/btcd/tree/master/wire) - Implements the
	  Bitcoin wire protocol
    * [blockchain](https://github.com/btcsuite/btcd/tree/master/blockchain) -
	  Implements Bitcoin block handling and chain selection rules
    * [txscript](https://github.com/btcsuite/btcd/tree/master/txscript) -
	  Implements the Bitcoin transaction scripting language
    * [btcec](https://github.com/btcsuite/btcd/tree/master/btcec) - Implements
	  support for the elliptic curve cryptographic functions needed for the
	  Bitcoin scripts
    * [database](https://github.com/btcsuite/btcd/tree/master/database) -
	  Provides a database interface for the Bitcoin block chain
    * [btcutil](https://github.com/btcsuite/btcutil) - Provides Bitcoin-specific
	  convenience functions and types