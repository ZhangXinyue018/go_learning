package main

// This is a temp file for testing purpose

import (
	"crypto/sha256"
	"fmt"
	"github.com/eoscanada/eos-go/ecc"
	"regexp"
	"sort"
)

func main() {
	//sign(getHash())
	//TempTest()
	//TempTestSort()

}

func TempTest() () {
	CertainIpPattern := []string{"^172.17.[0-9]+.[0-9]+$", "^172.18.[0-9]+.[0-9]+$", "^127.0.0.1$"}
	for _, ipPattern := range CertainIpPattern {
		matched, err := regexp.Match(ipPattern, []byte("127.0.0.1"))
		if err != nil {
			panic(err)
		}
		if matched {
			fmt.Println(true)
			return
		}
	}
	fmt.Println(false)
}

func TempTestSort() () {
	test := []string{"abc", "cba", "bac"}
	fmt.Println(test)
	sort.Sort(sort.StringSlice(test))
	fmt.Println(test)
}

func getHash() []byte {
	h := sha256.New()
	h.Write([]byte("/v1/wsts=123456account_name=eosio.token"))
	return h.Sum(nil)
}

func sign(hashContent []byte) () {
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

func verify(hashContent []byte) () {
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
