package ports

type Validator interface {
	Validate(interface{}) error
	Messages() []string
}
