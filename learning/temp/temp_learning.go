package main

// This is a temp file for testing purpose

import (
	"github.com/eoscanada/eos-go/ecc"
	"fmt"
	"crypto/sha256"
)

func main()  {
	sign(getHash())
}

func getHash() []byte {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d%s%s%s", 1234, "zhangxinyue1", "333fb6718abe32cb7e19edcd17e15133", "test")))
	return h.Sum(nil)
}

func sign(hashContent []byte) (){
	privateKey, err := ecc.NewPrivateKey("5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3")
	if err != nil {
		fmt.Println(err)
	}
	signature, err := privateKey.Sign(hashContent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(signature.String())
}

func verify(hashContent []byte)() {
	publickKey, err := ecc.NewPublicKey("EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV")
	if err != nil {
		fmt.Println(err)
	}
	signature, err := ecc.NewSignature("SIG_K1_KA8Jq3fdozdQSQQFp7CLc19p1PkyfrodNypNwGyzi79gkFs5GBcdKAW4fRJX5w66Vv4AzwMgnV9CG8VumhWKcdwNKfgRxp")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(signature.Verify(hashContent, publickKey))
}
