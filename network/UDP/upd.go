package UDP
//UDP首部格式描述

//首部字段8字节
type UDPHeader struct {
	srcPort uint16       //源端口(2字节)
	dstProt uint16       //目的端口(2字节)
	length  uint16       //长度
	checkSum uint16      //校验和
}

//伪首部12字节



