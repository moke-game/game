package wordsfilter

var ins = instance()

type Filter struct {
	*WordsFilter
	container map[string]*Node
}

func instance() *Filter {
	return &Filter{
		WordsFilter: New(),
		container:   make(map[string]*Node),
	}
}

func Load(conf []string) {
	container := ins.Generate(conf)
	ins.container = container
}

func Reload(conf []string) {
	_instance := instance()
	container := _instance.Generate(conf)
	_instance.container = container
	ins = _instance
}

func IsSensitive(s string) bool {
	return ins.Contains(s, ins.container)
}
