package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"
	"github.com/Ann-Geo/Latency_Controller/api/latencypb"
	"github.com/Ann-Geo/Latency_Controller/helper"
)

func (p *enbWOCntlr) doENBPeerStreaming() {

	errMsg, fileList := helper.WalkAllFilesInDir(p.imageFilesPath)
	if errMsg != "file read success" {
		log.Fatalf("File read failed - %v\n", errMsg)
	}

	imBuf, err := ioutil.ReadFile(fileList[0])
	if err != nil {
		log.Fatalf("Cannot read image file %v\n", err)

	}

	stream, err := p.cl.SendImage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling SendAndMeasure %v", err)
	}

	var i uint64
	for i = 0; i < p.numImagesInsert; i++ {

		time.Sleep(time.Duration(p.frameRate) * time.Millisecond)
		ts := time.Now().Format(customTimeformat)
		err = stream.Send(&latencypb.ImageData{
			Image:     imBuf,
			Timestamp: ts,
		})

		p.numPublished++
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Failed to send image %v", err)
		}

		fmt.Println(p.numPublished)

	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response %v", err)
	}

	fmt.Printf("Response %v\n", res)

}
