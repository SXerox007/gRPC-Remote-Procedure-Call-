package main

import (
	"fmt"
	face "gRPC-Remote-Procedure-Call-/face-detection/algorithm-face"
	"log"
	"path/filepath"
)

const dataDir = "face-detection/algorithm-face/dlib-model"

func main() {
	fmt.Println("Facial Recognition System v0.01")

	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		log.Fatalln("Cannot initialize recognizer", err)
	}
	defer rec.Close()

	log.Println("Recognizer Initialized")

	dataImage := filepath.Join(dataDir, "sumit.jpg")

	faces, err := rec.RecognizeFile(dataImage)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	fmt.Println("Number of Faces in Image: ", len(faces))

	var samples []face.Descriptor
	var totalF []int32
	for i, f := range faces {
		samples = append(samples, f.Descriptor)
		// Each face is unique on that image so goes to its own category.
		totalF = append(totalF, int32(i))
	}

	fmt.Println("SamplesData", samples)
	fmt.Println("totalFaces", totalF)

	// Pass samples to the recognizer.
	rec.SetSamples(samples, totalF)

	// Now let's try to classify some not yet known image.
	testSumit := filepath.Join(dataDir, "sumit.jpg")
	sumit, err := rec.RecognizeSingleFile(testSumit)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if sumit == nil {
		log.Fatalf("Not a single face on the image")
	}
	id := rec.Classify(sumit.Descriptor)
	if id < 0 {
		log.Fatalf("Can't classify")
	}

	fmt.Println("Face Recorganize", id)

}
