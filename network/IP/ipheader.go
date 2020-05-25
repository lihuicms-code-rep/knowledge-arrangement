package IP

//IP数据报首部格式
//注意这里只是描述,不是字段真正的字节数
//比如4位这种不足一字节的部分
type bit byte
type IPHeader struct {
	version [4]bit  //版本信息(4位),目前是IPV4,IPV6
	headLen [4]bit  //首部长度(4位),如果只为[]bit{1,1,1,1}表示IP首部长度为15*4=60字节
	dSrv    [8]bit  //区分服务(8位),一般不用
	allLen  [16]bit //总长度(16位)
	sign    [16]bit //标识(16位),标志,片偏移都是关于数据分片的
	flag    [3]bit  //标志(3位)
	offset  [13]bit //片偏移(13位)
	sTime   [8]bit  //生存时间(8位),数据报在网络中的寿命
	proto   [8]bit  //协议(8位),比如,6:TCP,17:UDP
	headCrc [16]bit //首部校验和(16位),只校验首部
	src     [32]bit //源地址
	dst     [32]bit //目的地址
}
