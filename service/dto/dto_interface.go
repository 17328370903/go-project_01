package dto

type Dto interface {
	Validator() error
}
