package action

type Trigger interface {
	Evaluate(context interface{}) bool
}
