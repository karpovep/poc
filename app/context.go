package app

type (
	IAppContext interface {
		Set(name string, resource interface{})
		Get(name string) interface{}
	}

	AppContext struct {
		resources map[string]interface{}
	}
)

func NewApplicationContext() IAppContext {
	return &AppContext{
		resources: make(map[string]interface{}),
	}
}

func (ac *AppContext) Set(name string, resource interface{}) {
	ac.resources[name] = resource
}

func (ac *AppContext) Get(name string) interface{} {
	return ac.resources[name]
}
