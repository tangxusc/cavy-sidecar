package event

func CallHandler(events []*Event) {
	go func() {
		//TODO:发起rpc,是否需要得到返回结果的event? saga模式则需要再发送command?
	}()
}
