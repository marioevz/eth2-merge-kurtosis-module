package el

type ELClientContext struct {
	enr            string
	enode          string
	ipAddr         string
	rpcPortNum     uint16
	authRpcPortNum uint16
	wsPortNum      uint16
	authWsPortNum  uint16
	miningWaiter   ELClientMiningWaiter
}

func NewELClientContext(enr string, enode string, ipAddr string, rpcPortNum uint16, authRpcPortNum uint16, wsPortNum uint16, authWsPortNum uint16, miningWaiter ELClientMiningWaiter) *ELClientContext {
	return &ELClientContext{enr: enr, enode: enode, ipAddr: ipAddr, rpcPortNum: rpcPortNum, authRpcPortNum: authRpcPortNum, wsPortNum: wsPortNum, authWsPortNum: authWsPortNum, miningWaiter: miningWaiter}
}

func (ctx *ELClientContext) GetENR() string {
	return ctx.enr
}
func (ctx *ELClientContext) GetEnode() string {
	return ctx.enode
}
func (ctx *ELClientContext) GetIPAddress() string {
	return ctx.ipAddr
}
func (ctx *ELClientContext) GetRPCPortNum() uint16 {
	return ctx.rpcPortNum
}
func (ctx *ELClientContext) GetAuthRPCPortNum() uint16 {
	return ctx.authRpcPortNum
}
func (ctx *ELClientContext) GetWSPortNum() uint16 {
	return ctx.wsPortNum
}
func (ctx *ELClientContext) GetAuthWSPortNum() uint16 {
	return ctx.authWsPortNum
}
func (ctx *ELClientContext) GetMiningWaiter() ELClientMiningWaiter {
	return ctx.miningWaiter
}
