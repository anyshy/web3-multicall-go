[![CircleCI](https://circleci.com/gh/Alethio/multicall-go.svg?style=svg)](https://circleci.com/gh/Alethio/multicall-go)

### Multicall

Wrapper for [Multicall](https://github.com/bowd/multicall) which batches calls to contract
view functions into a single call and reads all the state in one EVM round-trip.

### Usage

The library is used in conjunction with [web3-go](https://github.com/Alethio/web3-go), and the first parameter to `multicall.New` is an `ethrpc.ETHInterface` as defined in the package.

#### Initialization

The library requires the [Multicall](https://github.com/bowd/multicall) contract to pe deployed on the target chain.
We have deployed two variants on Mainnet and Ropsten so far which can be used by using the provided configs.


```go
// Mainnet
mc, err := multicall.New(eth, multicall.MainnetConfig)
// Ropsten
mc, err := multicall.New(eth, multicall.RopstenConfig)
```

Otherwise you can define your own config passing in a custom address:

```go
mc, err := multicall.New(eth, multicall.Config{
    MulticallAddress: "0x0",
    Gas: "0x400000000",
})
```

In this case the contract deployed has to maintain the same function signature as the original one.

#### Calling

```go
vcs := ViewCalls{
    {
        Key:       "key-1",
        Target:    "0x5eb3fa2dfecdde21c950813c665e9364fa609bd2",
        Method:    "getLastBlockHash()(bytes32)",
        Arguments: []interface{}{},
    },
    {
        Key:        "key-2",
        Target:     "0x6b175474e89094c44da98b954eedeac495271d0f",
        Method:     "balanceOf(address)(uint256)",
        Arguments:  []interface{}{"0x8134d518e0cef5388136c0de43d7e12278701ac5"}
    },
}
block := "latest" // default block parameter
res, err := mc.Call(vcs, block)
if err != nil {
    panic(err)
}

lastBlockHashSuccess = res.Calls["key-1"].Success;
lastBlockHash := res.Calls["key-1"].ReturnValues[0].([32]byte);

someBalanceSuccess := res.Calls["key-2"].Success;
someBalance := res.Calls["key-2"].ReturnValues[0].(*big.Int);
```

In the example above we batch two calls to two different contracts and get back a map of `CallResults` which contain the exit value an array of returned values (`[]interface{}`) which are decoded by the `go-ethereum` package.
