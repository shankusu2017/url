package url

const (
	DOMAIN_MGR = "mgr.node.vpn.shanks.link" // 主机管理
)

const (
	URL_EVENT_POST      = "/v1/event/post"          /* 事件上报的 URL */
	URL_EVENT_GET       = "/v1/event/get"           /* 查看事件 */
	URL_EVENT_HELP      = "/v1/event/help"          /* 事件帮助 */
	URL_REPEATER_SERVER = "/v1/vpn/repeater/server" /* REPEATER SERVER 信息 */
)

const (
	PORT_REPEATER_SERVER        = 465
	PORT_REPEATER_SERVER_IPERF3 = 5201 /* 假装成 iperf3 的数据包 */
	PORT_NODEMGR                = 7080
)
