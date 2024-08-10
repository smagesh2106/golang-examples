package c

import (
	"context"
	"golang-examples/multi_type/a"
	"golang-examples/multi_type/b"
	"os"
)

var PKI_TYPE string
var PKI_OBJ any = nil

func init() {

	if dat, ok := os.LookupEnv("PKI_PROVIDER"); ok {
		if dat == "aws" {
			PKI_TYPE = "awsPKI"
			PKI_OBJ = a.NewAWS()
		} else {
			PKI_TYPE = "sroPKI"
			PKI_OBJ = b.NewSRO()
		}
	}
}

func DoPrint(ctx context.Context) {
	//x.GetName()
	ctx.Value("sroPKI").(*b.SRO).GetName()
	ctx.Value("awsPKI").(*a.AWS).GetName()

}
