package main

func main() {

	//activate controller flag
	actController := 0

	if actController == 0 {

		enb := newEnbWOCntlr("127.0.0.1:50051")

		defer enb.conn.Close()

		enb.doENBPeerStreaming()

	} else {
		enb := newEnbWithCntlrServer("127.0.0.1:9000")

		enb.startEnbWithCntlrServer()
	}

	//on local host with controller
	/*
		enb := newEnbWithCntlrServer("127.0.0.1:9000")

		enb.startEnbWithCntlrServer()
	*/
}
