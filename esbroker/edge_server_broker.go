package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Ann-Geo/Latency_Controller/api/latencypb"

	"github.com/Ann-Geo/Latency_Controller/api/testcpb"

	"google.golang.org/grpc"
)

var customTimeformat string = "Monday, 02-Jan-06 15:04:05.00000 MST"

type edgeServerBroker struct {
	ipaddrES   string
	resWOCntlr string
}

func newEdgeServerBroker(ipaddrES, resWOCntlr string) *edgeServerBroker {
	return &edgeServerBroker{
		ipaddrES:   ipaddrES,
		resWOCntlr: resWOCntlr,
	}

}

func (s *edgeServerBroker) startLatencyMeasESB() {
	//Server for latency measurement only
	fmt.Println("Starting Edge server")

	lis, err := net.Listen("tcp", s.ipaddrES)
	if err != nil {
		log.Fatalf("Failed to listen %v\n", err)
	}

	grpcServer := grpc.NewServer()
	latencypb.RegisterMeasureServiceServer(grpcServer, s)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}

func (s *edgeServerBroker) SendAndMeasure(stream latencypb.MeasureService_SendAndMeasureServer) error {
	fmt.Println("SendAndMeasure RPC was invoked")

	resultFile, err := os.Create(s.resWOCntlr)
	if err != nil {
		log.Fatalf("Cannot create result file %v\n", err)
	}

	defer resultFile.Close()

	for {
		im, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&latencypb.Response{
				Status: true,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		tsRcvd := time.Now()
		ts, _ := time.Parse(customTimeformat, im.GetTimestamp())
		imLatency := tsRcvd.Sub(ts)
		fmt.Println(len(im.GetImage()))
		fmt.Fprintf(resultFile, "current time: %s, image_size: %d, latency: %s\n", tsRcvd, len(im.GetImage()), imLatency)

	}
}

func (s *edgeServerBroker) SendImage(stream latencypb.MeasureService_SendImageServer) error {
	fmt.Println("SendImage RPC was invoked")

	for {
		im, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&latencypb.Response{
				Status: true,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		fmt.Printf("received %v bytes at %v\n", len(im.GetImage()), im.GetTimestamp())

	}
}

type EdgeServerClient struct {
	testENIp     string
	peerENIps    []string
	resWithCntlr string
	conn         *grpc.ClientConn
	cl           testcpb.TestControllerClient
}

func newEdgeServerClient(testENIp string, peerENIps []string, resWithCntlr string) *EdgeServerClient {

	conn, err := grpc.Dial(testENIp, grpc.WithInsecure())
	log.Println("connected with test Edge node")

	if err != nil {
		log.Fatalf("Could not connect with node-%s, %v\n", testENIp, err)
	}

	cl := testcpb.NewTestControllerClient(conn)
	fmt.Printf("Created client %v\n", cl)

	return &EdgeServerClient{
		testENIp:     testENIp,
		peerENIps:    peerENIps,
		resWithCntlr: resWithCntlr,
		conn:         conn,
		cl:           cl,
	}
}

func (s *EdgeServerClient) startSubscription() {
	resultFile, err := os.Create(s.resWithCntlr)
	if err != nil {
		log.Fatalf("Cannot create result file %v\n", err)
	}
	defer resultFile.Close()

	target := &testcpb.Targets{
		TargetLat: "24",
		TargetAcc: "0.47 jaad simple",
	}

	//resImages, err := s.cl.Subscribe(context.Background(), target)
	resImages, err := s.cl.SubscribeWithControl(context.Background(), target)
	if err != nil {
		log.Fatalf("error while calling Subscribe RPC: %v\n", err)
	}
	for {
		msg, err := resImages.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v\n", err)
		}

		tsRcvd := time.Now()
		ts, _ := time.Parse(customTimeformat, msg.GetTimestamp())
		imLatency := tsRcvd.Sub(ts)

		fmt.Println("latency=", imLatency)

		go s.sendMeasuredLatency(imLatency)

		imLen := len(msg.GetImage())
		achAcc := msg.GetAcheivedAcc()
		fmt.Fprintf(resultFile, "current time: %s, image_size: %d, latency: %s, accuracy: %s\n", tsRcvd, imLen, imLatency, achAcc)

		//log.Printf("Response from Subscribe: msg with len %v received at %v with accuracy %v\n",imLen, tsRcvd, achAcc)

	}

}

func (s *EdgeServerClient) sendMeasuredLatency(imLatency time.Duration) {

	imLatstr := imLatency.String()

	re := regexp.MustCompile("[0-9]+")
	integerSlice := re.FindAllString(imLatstr, -1)
	imLatstrFloatstr := integerSlice[0] + "." + integerSlice[1]
	var imLatstrFloat float64
	if s, err := strconv.ParseFloat(imLatstrFloatstr, 64); err == nil {
		imLatstrFloat = s
	}

	if strings.Contains(imLatstr, "us") {
		imLatstrFloat = imLatstrFloat / 1000
	} else if strings.Contains(imLatstr, "ms") {
		imLatstrFloat = imLatstrFloat * 1
	} else {
		imLatstrFloat = imLatstrFloat * 1000
	}

	imLatstrFloatstrCorrectFmt := fmt.Sprintf("%f", imLatstrFloat)

	lat := &testcpb.LatencyMeasured{
		CurrentLat: imLatstrFloatstrCorrectFmt,
	}

	status, err := s.cl.LatencyCalc(context.Background(), lat)
	if err != nil {
		log.Fatalf("error while calling LatencyCalc RPC: %v", err)
	}
	log.Printf("Response from LatencyCalc: %v\n", status.GetStatus())
}

func main() {

	//time.Sleep(58 * time.Second)

	//ip address is address of Edge server
	esb := newEdgeServerBroker("192.168.1.2:50051", "/home/research/goworkspace/src/vsc_workspace/cont_test/esbroker/resWOCntlr.txt")
	go esb.startLatencyMeasESB()

	/*
		time.Sleep(110 * time.Second)

		//ip address is address of test Edge node
		testENIp := "192.168.1.5:9000"
		numPeerEN := 0
		peerENIps := make([]string, numPeerEN)
		esC := newEdgeServerClient(testENIp, peerENIps, "/home/research/goworkspace/src/vsc_workspace/cont_test/esbroker/resWithCntlr.txt")
		defer esC.conn.Close()

		esC.startSubscription()

	*/
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

}
