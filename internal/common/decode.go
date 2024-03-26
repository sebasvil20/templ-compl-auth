package common

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

func DecodeAndValidateForm(d *schema.Decoder, val *validator.Validate, target any, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	return decodeAndValidate(d, val, target, r)
}

func DecodeAndValidateMultipartForm(d *schema.Decoder, val *validator.Validate, target any, r *http.Request) error {
	err := r.ParseMultipartForm(6000)
	if err != nil {
		return err
	}

	return decodeAndValidate(d, val, target, r)
}

func decodeAndValidate(d *schema.Decoder, val *validator.Validate, target any, r *http.Request) error {
	err := d.Decode(target, r.PostForm)
	if err != nil {
		return err
	}

	err = val.Struct(target)
	if err != nil {
		return err
	}

	return nil
}
