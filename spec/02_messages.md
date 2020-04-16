<!--
order: 2
-->

# Messages

In this section we describe the processing of the token messages and the corresponding updates to the state.

## MsgCreateRecord

A record is created using the `MsgCreateRecord` message.

```go
type MsgCreateRecord struct {
	Contents []Content
	Creator  sdk.AccAddress // the creator of the record
}
```

This message is expected to fail if:
- the length of contents is 0
- the creator is empty
- each content parameters are faulty, namely:
    - the initial `Digest` is empty
    - the initial `DigestAlgo` is empty


