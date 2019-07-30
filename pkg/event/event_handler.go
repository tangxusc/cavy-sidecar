package event

func CallHandler(events []*Event) {
	//TODO:发起rpc,是否需要得到返回结果的event? saga模式则需要再发送command?
	//TODO:rpc完成后,需要更新事件表中的事件状态
}
