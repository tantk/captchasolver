package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/png"
	"log"
	"os"
)

func Preprocess(fileSRC string,fileDST string){

	img := gocv.IMRead(fileSRC, gocv.IMReadUnchanged)
	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", fileSRC)
		return
	}
	processed := gocv.NewMat()
	//binarize
	gocv.Threshold(img,&processed,200,255,gocv.ThresholdBinary)
	//invert
	gocv.BitwiseNot(processed,&processed)
	//enlarge
	gocv.Resize(processed,&processed,image.Point{X: 0,Y:0},2,2, gocv.InterpolationDefault)
	//Erode kernel
	//trialKernal := gocv.GetGaussianKernel(5,0.1)
	morphKernel:= gocv.GetStructuringElement(gocv.MorphRect,image.Point{X: 3,Y:3})
	//Erode twice
	for i:=0;i<2;i++{
		gocv.Erode(processed,&processed,morphKernel)
	}
	//dilate
	morphKernelDilate:= gocv.GetStructuringElement(gocv.MorphRect,image.Point{X: 2,Y:2})

	gocv.Dilate(processed,&processed,morphKernelDilate)
	gocv.Resize(processed,&processed,image.Point{X: 0,Y:0},0.5,0.5,gocv.InterpolationDefault)
	//grayscale
	gocv.CvtColor(processed,&processed,7)
	//convert back to white background
	gocv.BitwiseNot(processed,&processed)

	iProcessed,err := processed.ToImage()
	check(err)
	f, err := os.Create(fileDST)
	check(err)
	if err := png.Encode(f, iProcessed); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
