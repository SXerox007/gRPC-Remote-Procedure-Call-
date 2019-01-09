package core

import (
	"fmt"
	face "gRPC-Remote-Procedure-Call-/face-detection/algorithm-face"
	"log"
	"path/filepath"
)

const dataDir = "face-detection/algorithm-face/dlib-model"

func FaceDetection(image []byte) error {

	fmt.Println("Facial Recognition System v0.01")

	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		fmt.Println("Cannot initialize recognizer")
	}
	defer rec.Close()

	fmt.Println("Recognizer Initialized")

	avengersImage := filepath.Join(dataDir, "mix.jpg")

	faces, err := rec.RecognizeFile(avengersImage)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	fmt.Println("Number of Faces in Image: ", len(faces))

	var samples []face.Descriptor
	var avengers []int32
	for i, f := range faces {
		samples = append(samples, f.Descriptor)
		// Each face is unique on that image so goes to its own category.
		avengers = append(avengers, int32(i))
	}
	// Pass samples to the recognizer.
	rec.SetSamples(samples, avengers)

	// Now let's try to classify some not yet known image.
	testTonyStark := filepath.Join(dataDir, "eminem.jpg")
	tonyStark, err := rec.RecognizeSingleFile(testTonyStark)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if tonyStark == nil {
		log.Fatalf("Not a single face on the image")
	}
	avengerID := rec.Classify(tonyStark.Descriptor)
	if avengerID < 0 {
		log.Fatalf("Can't classify")
	}

	fmt.Println(avengerID)
	log.Fatalf("classify Success")
	return nil
}
