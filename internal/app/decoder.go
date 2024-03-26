package app

import "github.com/gorilla/schema"

func newDecoder() *schema.Decoder {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	return decoder
}
