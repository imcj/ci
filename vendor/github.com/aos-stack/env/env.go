package env; 

const (
	Local = 0
	Develop = 1
	Testing = 2
	Staging = 3
	Production = 4
)

// env.Get() == env.Local
func Get() int {
	initialize()
	
	return provider.Get()
}

func GetLabel() string {
	return provider.GetLabel()
}

var provider EnvProvider

type EnvProvider interface {
	Get() int
	GetLabel() string
}

type DefaultEnvProvider struct {
}

func (d DefaultEnvProvider)Get() int {
	return Local
}

func (d DefaultEnvProvider)GetLabel() string {
	return "local"
}

func initialize() {
	if nil == provider {
		provider = DefaultEnvProvider{}
	}
}

func SetProvider(aProvider EnvProvider) {
	provider = aProvider
}
