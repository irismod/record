package types

import (
	"fmt"
	"testing"
)

func TestPacket(t *testing.T) {
	pubKey := []byte(`{"type":"tendermint/PubKeySm2","value":"AgAApPG8mUm1MGBBpRvhb+dfIcXzeC5PmmBUocJaLL3p"}`)

	packet := Packet{
		Digest:     "9143406651ba1d0db8f9baa7bd104611a2458865f648143b1b05901d39263636",
		DigestAlgo: "sm3",
		URI:        "",
		Metadata:   `{"address":"13821324323","signoffTime":"2020-06-11 10:32:19","method":"短信","sender":"XX法院","idcard":"3456784567845678","people":"张三","sendTime":"2020-06-10 15:32:23"}`,
		PubKey:     pubKey,
	}
	fmt.Println(string(ModuleCdc.MustMarshalJSON(packet)))
}
