package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"github.com/Ann-Geo/Latency_Controller/api/controller"
	"github.com/Ann-Geo/Latency_Controller/api/latencypb"
	"github.com/Ann-Geo/Latency_Controller/api/testcpb"
	"github.com/Ann-Geo/Latency_Controller/helper"

	"google.golang.org/grpc"
)

var customTimeformat string = "Monday, 02-Jan-06 15:04:05.00000 MST"

type enbWOCntlr struct {
	numPublished    uint64
	numImagesInsert uint64
	frameRate       uint64
	imageFilesPath  string
	conn            *grpc.ClientConn
	cl              latencypb.MeasureServiceClient
}

func newEnbWOCntlr(ipaddrES string) *enbWOCntlr {
	numImages := os.Args[1]
	numImagesInsert, _ := strconv.ParseUint(numImages, 10, 64)
	fps := os.Args[2]
	frameRate, _ := strconv.ParseUint(fps, 10, 64)
	imageFilesPath := os.Args[3]
	fmt.Println("Hello ths is a Peer EN broker")

	conn, err := grpc.Dial(ipaddrES, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to Edge server %v\n", err)
	}
	cl := latencypb.NewMeasureServiceClient(conn)
	fmt.Printf("Created client %v\n", cl)

	return &enbWOCntlr{
		numPublished:    0,
		numImagesInsert: numImagesInsert,
		frameRate:       frameRate,
		imageFilesPath:  imageFilesPath,
		conn:            conn,
		cl:              cl,
	}

}

type image struct {
	im []byte
}

type dataSet struct {
	filePath string
	imageLog []image
}

type enbWithCntlrServer struct {
	numPublished uint64
	dSets        []dataSet
	ipaddrEN     string
}

func newEnbWithCntlrServer(ipaddrEN string) *enbWithCntlrServer {
	return &enbWithCntlrServer{
		numPublished: 0,
		dSets:        make([]dataSet, 8),
		ipaddrEN:     ipaddrEN,
	}

}

func (s *enbWithCntlrServer) loadImageFiles(imPaths []string) {

	for i, path := range imPaths {

		errMsg, fileList := helper.WalkAllFilesInDir(path)
		if errMsg != "file read success" {
			log.Fatalf("File read failed - %v\n", errMsg)
		}

		s.dSets[i].filePath = path

		for _, file := range fileList {
			imBuf, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatalf("Cannot read image file %v\n", err)

			}
			imByte := image{
				im: imBuf,
			}
			s.dSets[i].imageLog = append(s.dSets[i].imageLog, imByte)
		}

	}

}

func (s *enbWithCntlrServer) startEnbWithCntlrServer() {

	//load jaad dataset - simple, medium, complex
	imPaths := []string{"/home/research/goworkspace/src/vsc_workspace/lat_meas/controller/jaad/simple/",
		"/home/research/goworkspace/src/vsc_workspace/lat_meas/controller/jaad/medium/",
		"/home/research/goworkspace/src/vsc_workspace/lat_meas/controller/jaad/complex/"}

	s.loadImageFiles(imPaths)
	fmt.Println("loaded image files - jaad simple, medium and complex")

	//Server for subscribe and latency calc
	fmt.Println("Starting Edge node broker server")

	lis, err := net.Listen("tcp", s.ipaddrEN)
	if err != nil {
		log.Fatalf("Failed to listen %v\n", err)
	}

	grpcServer := grpc.NewServer()
	testcpb.RegisterTestControllerServer(grpcServer, s)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}

type enbWithCntlrClient struct {
	numPublished uint64
	initialLat   string
	subResChan   chan controller.CustomImage
	conn         *grpc.ClientConn
	cl           controller.LatencyControllerClient
}

func newEnbWithCntlrClient(ipaddrCont string) *enbWithCntlrClient {
	//new client to python server
	conn, err := grpc.Dial(ipaddrCont, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v\n", err)
	}

	cl := controller.NewLatencyControllerClient(conn)
	fmt.Printf("Created client %v\n", cl)
	return &enbWithCntlrClient{
		numPublished: 0,
		initialLat:   "25",
		subResChan:   make(chan controller.CustomImage),
		conn:         conn,
		cl:           cl,
	}
}
