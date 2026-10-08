package factory

import (
	"solopg/app/shared/types/primitive"
)

type Mill struct {
	primitive.Fallible
}

/*
	NOTE Snippet - Mill

	type Mill = mill

	type mill struct {
		factory.Mill

	}

	type Template struct {

	}

	func New(params Template) *mill {
		return &mill{

		}
	}

	// Getters & Setters

	// Handlers

	// Helpers

*/
