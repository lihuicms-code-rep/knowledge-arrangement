package TCP

//以太网(数据链路层)
/*
  6字节    6字节  2字节   46~1500字节   4字节
  |目的地址|源地址|类型|  |数据和填充|   |CRC|
*/
//以太网头,14字节
const EthLength = 6

type ethHeader struct {
	dest  [EthLength]byte //目的地址
	src   [EthLength]byte //源地址
	proto uint16          //协议类型
}

//IP协议头,
//具体查看
type ipHeader struct {
	version  byte   //版本
	headLen  byte   //首部长度
	tos      byte   //服务类型
	totalLen uint16 //总长度
	id       uint16 //16位标识符
	flagOff  uint16 //标志
	ttl      byte   //生存时间
	protocol byte   //8位协议
	check    uint16 //16位检查
	sAddr    uint32 //32位源地址
	dAddr    uint32 //32位目的地址
}

//UDP协议头
type udpHeader struct {
	sPort uint16 //源端口号
	dPort uint16 //目的端口号
	len   uint16 //长度
	check uint16 //检查位
}

//UPD的包
type updPkt struct {
	eh ethHeader     //以太网头
	ih ipHeader      //ip头
	uh udpHeader     //udp头
	body []byte      //数据域
}

//依靠netmap