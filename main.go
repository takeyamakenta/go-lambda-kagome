package main

import (
	"fmt"
	"context"
	"errors"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type RequestEvent struct {
	Text string `json:"text"`
}

type Response struct {
	Tokens []Token `json:"tokens"`
}

type Token struct {
	POS []string `json:"pos"`
	Surface  string `json:"surface"`
	Reading string `json:"reading"`
}

var T *tokenizer.Tokenizer

func main() {
	err := run(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func HandleRequest(ctx context.Context, ev RequestEvent) (Response, error) {
	var err error
	T, err = tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return Response{}, fmt.Errorf("tokeninzer.New :%w", err)
	}

	res := &Response{Tokens: []Token{}}
	tokens := T.Tokenize(ev.Text)
	for _, token := range tokens {
		reading, _ := token.Reading()
		res.Tokens = append(res.Tokens, Token{
			POS: token.POS(),
			Surface: token.Surface,
			Reading: reading})
	}
	return *res, nil
}

type valueKey struct{}

func run(ctx context.Context) error{
	if v, err := GetParameters(ctx); err == nil {
		if res, err := HandleRequest(ctx, v); err != nil {
			return fmt.Errorf("HandleRequest :%w", err)
		}else{
			log.Printf("%v", res)
			return nil
		}
	}else{
		lambda.Start(HandleRequest)
	}

	<-ctx.Done()
	return errors.New("terminated")
}

func GetParameters(ctx context.Context) (RequestEvent, error) {
	if v, ok := ctx.Value(valueKey{}).(RequestEvent); ok{
		return v, nil
	}
	return RequestEvent{}, errors.New("空") //空
}