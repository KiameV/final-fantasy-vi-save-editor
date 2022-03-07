package io

type IOer interface {
	Load(fileName string) error
	Save() error
}
