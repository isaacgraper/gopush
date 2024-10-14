package trigger

type Trigger interface {
	Evaluate(context interface{}) bool
}
