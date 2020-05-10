package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"
	"github.com/Ann-Geo/Latency_Controller/api/controller"
	"github.com/Ann-Geo/Latency_Controller/api/testcpb"
)

var curLatLock sync.Mutex
var currentLat string

func (s *enbWithCntlrServer) LatencyCalc(ctx context.Context, lat *testcpb.LatencyMeasured) (*testcpb.Status, error) {
	fmt.Printf("LatencyCalc RPC was invoked with %v\n", lat)

	curLatLock.Lock()
	currentLat = lat.GetCurrentLat()
	curLatLock.Unlock()
	fmt.Println(lat.GetCurrentLat())

	status := &testcpb.Status{
		Status: true,
	}
	return status, nil
}

func (s *enbWithCntlrServer) Subscribe(targets *testcpb.Targets, stream testcpb.TestController_SubscribeServer) error {
	fmt.Printf("Subscribe RPC was invoked with %v\n", targets)

	accStr := targets.GetTargetAcc()
	accSlice := strings.Fields(accStr)
	dSet := accSlice[1]
	regime := accSlice[2]
	var useDSet dataSet
	if dSet == "jaad" {
		if regime == "simple" {
			useDSet = s.dSets[0]
		} else if regime == "medium" {
			useDSet = s.dSets[1]
		} else {
			useDSet = s.dSets[2]
		}
	}

	for _, result := range useDSet.imageLog {
		modIm := &testcpb.CustomImage{
			Image:       result.im,
			Timestamp:   time.Now().Format(customTimeformat),
			AcheivedAcc: "1",
		}
		stream.Send(modIm)
	}

	return nil

}

func (s *enbWithCntlrServer) SubscribeWithControl(targets *testcpb.Targets, stream testcpb.TestController_SubscribeWithControlServer) error {
	curLatLock.Lock()
	currentLat = "1.0"
	curLatLock.Unlock()

	fmt.Printf("SubscribeWithControl RPC was invoked with %v\n", targets)

	enbC := newEnbWithCntlrClient("0.0.0.0:9002")

	conTargets := &controller.Targets{
		TargetLat: targets.GetTargetLat(),
		TargetAcc: targets.GetTargetAcc(),
	}

	_, err := enbC.cl.SetTarget(context.Background(), conTargets)
	if err != nil {
		log.Fatalf("error while calling SetTarget RPC: %v\n", err)
	}

	//format = "0.45 jaad complex"
	accStr := targets.GetTargetAcc()
	accSlice := strings.Fields(accStr)
	dSet := accSlice[1]
	regime := accSlice[2]
	var useDSet dataSet
	if dSet == "jaad" {
		if regime == "simple" {
			useDSet = s.dSets[0]
		} else if regime == "medium" {
			useDSet = s.dSets[1]
		} else {
			useDSet = s.dSets[2]
		}
	}

	//do bidi streaming
	go enbC.doControl(useDSet)

	for result := range enbC.subResChan {
		modIm := &testcpb.CustomImage{
			Image:       result.Image,
			Timestamp:   time.Now().Format(customTimeformat),
			AcheivedAcc: result.AcheivedAcc,
		}
		stream.Send(modIm)
	}

	return nil

}

func (e *enbWithCntlrClient) doControl(useDSet dataSet) {
	fmt.Println("Starting to do a Control RPC")

	stream, err := e.cl.Control(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)

	}
	waitc := make(chan struct{})

	go func() {

		for _, im := range useDSet.imageLog {

			fmt.Println("Sending original image to controller")
			time.Sleep(200 * time.Millisecond)
			curLatLock.Lock()
			req := &controller.OriginalImage{
				Image:      im.im,
				CurrentLat: currentLat,
			}
			fmt.Println(currentLat)
			curLatLock.Unlock()
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v\n", err)
				break
			}
			fmt.Println("Received from controller")
			e.subResChan <- controller.CustomImage{
				Image: res.GetImage(),

				AcheivedAcc: res.GetAcheivedAcc(),
			}
		}
		close(e.subResChan)
		close(waitc)
	}()

	// block until everything is done
	<-waitc
}
