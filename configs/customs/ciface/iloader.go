package ciface

type ILoader interface {
	Load(dir []byte) error
}
