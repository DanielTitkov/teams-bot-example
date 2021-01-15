package mesga

func NewRouter() *States {

}

type RouterSetup struct {
	States []StateSetup
}

type StateSetup struct {
	Title            string
	Actions          []ActionSetup
	Default          ActionSetup
	Data             map[string]interface{} // default data
	OnEnterCallbacks []func() error
	OnExitCallbacks  []func() error
}

type ActionSetup struct {
	TriggerText          string // or
	TriggerTextRgxp      string // or
	TriggerPayloadAction string
	OnSuccessTransition  string
	OnFailTransition     string
	Function             func(*Turn, map[string]interface{}) (*Turn, error)
}
