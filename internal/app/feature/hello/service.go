package hello

type (
	//GateKeeper is a interface to implement in repo layer, we can use this for many database form NoSQL and traditional SQL
	GateKeeper interface {
		SayHello(mgs string) string
	}
	//Service the interface for expose to handler of RestAPI Or MQ, gRPC
	Service interface {
		Say(mgs string) string
	}
	//MakeFun is a implementation oi hello service is defined above
	MakeFun struct {
		Mgs string
		R   GateKeeper
	}
)

// NewHelloService create a instand of service
func NewHelloService(r GateKeeper) *MakeFun {
	return &MakeFun{
		Mgs: "Hello",
		R:   r,
	}
}

//Say return a phase
func (h MakeFun) Say(mgs string) string {
	return h.Mgs + h.R.SayHello(mgs)
}
