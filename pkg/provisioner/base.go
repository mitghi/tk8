package provisioner

var IOnly bool = false

type Provisioner interface {
	Init(args []string)
	Setup(args []string)
	Scale(args []string)
	Remove(args []string)
	Reset(args []string)
	Upgrade(args []string)
	Destroy(args []string)
}
