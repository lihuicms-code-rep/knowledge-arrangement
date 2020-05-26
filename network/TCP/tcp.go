package TCP

//TCP首部格式
//注意这里只是描述TCP首部用到了哪些字段,不代表实际占用字节
type bit byte
type TCPHeader struct {
	srcPort    [16]bit //源端口(16位)
	dstPort    [16]bit //目的端口(16位)
	seqNum     [32]bit //序号(32位)用于对字节流进行编号
	ackNum     [32]bit //确认号(32位)期望收到下一个报文段的序号
	dataOffset [4]bit  //数据偏移,数据部分距离报文段的起始位置的偏移,就是首部的长度
	remain     [6]bit  //保留位(6位)
	controlU   bit     //控制位urgent,紧急指针有效标志
	controlA   bit     //控制位acknowledgement,确认序号的有效标记
	controlP   bit     //控制位push,接收方应该尽快将报文段提交到应用层
	controlR   bit     //控制位Reset,重置连接标志
	controlS   bit     //控制位Syn,同步需要标志
	controlF   bit     //控制位Fin,传输数据结束标志
	window     [16]bit //窗口大小(16位),作为接收方让发送方设置其发送窗口的依据
	checkSum   [16]bit //校验和(16位)
	urgentP    [16]bit //紧急指针(16位),到此为止，一共20字节
	options    []bit   //选项部分(0~40字节)
}
