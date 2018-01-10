##  CMBToken - Blockchain ## 

12.000.000 Coins
6.000.000 On Pre-sale (ICO)

Official golang implementation of the CMBToken protocol.



## Building the source

For prerequisites and detailed build instructions please read the
[Installation Instructions](https://github.com/CoinMarketBrasil/cmbtoken/wiki/Building-CMBToken)
on the wiki.

Building geth requires both a Go (version 1.7 or later) and a C compiler.
You can install them using your favourite package manager.
Once the dependencies are installed, run

    make geth

or, to build the full suite of utilities:

    make all

## Executables

The go-CMBToken project comes with several wrappers/executables found in the `cmd` directory.

| Command    | Description |
|:----------:|-------------|
| **`geth`** | Our main CMBToken CLI client. It is the entry point into the CMBToken network (main-, test- or private net), capable of running as a full node (default) archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the CMBToken network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `geth --help` and the [CLI Wiki page](https://github.com/CoinMarketBrasil/cmbtoken/wiki/Command-Line-Options) for command line options. |
| `abigen` | Source code generator to convert CMBToken contract definitions into easy to use, compile-time type-safe Go packages. It operates on plain [CMBToken contract ABIs](https://github.com/CMBToken/wiki/wiki/CMBToken-Contract-ABI) with expanded functionality if the contract bytecode is also available. However it also accepts Solidity source files, making development much more streamlined. Please see our [Native DApps](https://github.com/CoinMarketBrasil/cmbtoken/wiki/Native-DApps:-Go-bindings-to-CMBToken-contracts) wiki page for details. |
| `bootnode` | Stripped down version of our CMBToken client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks. |
| `evm` | Developer utility version of the EVM (CMBToken Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow isolated, fine-grained debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug`). |
| `gethrpctest` | Developer utility tool to support our [CMBToken/rpc-test](https://github.com/CMBToken/rpc-tests) test suite which validates baseline conformity to the [CMBToken JSON RPC](https://github.com/CMBToken/wiki/wiki/JSON-RPC) specs. Please see the [test suite's readme](https://github.com/CMBToken/rpc-tests/blob/master/README.md) for details. |
| `rlpdump` | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://github.com/CMBToken/wiki/wiki/RLP)) dumps (data encoding used by the CMBToken protocol both network as well as consensus wise) to user friendlier hierarchical representation (e.g. `rlpdump --hex CE0183FFFFFFC4C304050583616263`). |
| `swarm`    | swarm daemon and tools. This is the entrypoint for the swarm network. `swarm --help` for command line options and subcommands. See https://swarm-guide.readthedocs.io for swarm documentation. |
| `puppeth`    | a CLI wizard that aids in creating a new CMBToken network. |


