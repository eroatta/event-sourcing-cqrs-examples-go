package model

type Specification interface {
	IsSatisfiedBy(value interface{}) bool
}
