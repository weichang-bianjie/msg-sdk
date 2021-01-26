package types

type (
	DocTxMsg struct {
		Type string `bson:"type"`
		Msg  Msg    `bson:"msg"`
	}

	Msg interface {
		GetType() string
		BuildMsg(msg interface{})
	}
)
