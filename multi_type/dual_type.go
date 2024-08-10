package main

import (
	"context"
	"golang-examples/multi_type/a"
	"golang-examples/multi_type/b"
	"golang-examples/multi_type/c"
)

/*
func doPrint(ctx context.Context, x interface{}) {
	switch v := x.(type) {
	case *a.AWS:
		v.GetName()
	case *b.SRO:
		v.GetName()

	default:
		fmt.Println("Unknown type")
	}
}
*/
// func doPrint2(ctx context.Context, x intf.ClaimsServiceType) {

func main() {
	ctx := context.Background()

	aws := a.NewAWS()
	sro := b.NewSRO()
	ctx = context.WithValue(ctx, "sroPKI", sro)
	ctx = context.WithValue(ctx, "awsPKI", aws)
	//doPrint(ctx, sro)
	//doPrint2(ctx, sro)
	c.DoPrint(ctx)
}
