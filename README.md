# IRISHUB Chain Msg Go SDK

Irishub Chain Msg GO SDK makes a simple package of API provided by Irishub, which provides parser for users to parse tx msg on block chain.

## install

### Requirement

Go version above 1.15

### Use Go Mod

```text
require (
    github.com/weichang-bianjie/msg-sdk latest
)
```

## Usage

### Init Client

The initialization SDK code is as follows:

```go

client := msg_sdk.NewMsgClient()
```

parse Bank Msg of Tx 
```go
bankDoc, ok := client.Bank.HandleTxMsg(&msg)
if ok {
		//db save bank msg
	}
```

use in sync 
```go

var (
	docTx models.Tx
	docTxMsgs []msg_sdk.DocTxMsg
    	)
	authTx := Tx.(signing.Tx)
	......
	msgs := authTx.GetMsgs()
	for _, msg := range msgs {
        if bankDoc, ok := client.Bank.HandleTxMsg(&msg);ok {
            docTxMsgs = append(docTxMsgs, bankDoc.DocTxMsg)
            continue
        }
        .....
    }
 docTx.DocTxMsgs = docTxMsgs

```