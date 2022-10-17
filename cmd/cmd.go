package cmd

type Command interface {
	Execute() error
}
