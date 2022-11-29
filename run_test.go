//Test
//go test
package main

import (
	"context"
	"testing"
	"golang.org/x/sync/errgroup" 
)

//contextにValue(RequestのParameter)を持たせてテストする。
func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, valueKey{}, RequestEvent{
		Text: "すもももももももものうち"})
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error{
		return run(ctx)
	})

	cancel()

	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}

}