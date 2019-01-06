package core

import (
	"fmt"
	face "gRPC-Remote-Procedure-Call-/face-detection/algorithm-face"
	"log"
	"path/filepath"
)

const dataDir = "face-detection/algorithm-face/dlib-model"

func FaceDetection(image []byte) error {

	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		log.Println("Cannot initialize recognizer", err)
		return err
	}
	defer rec.Close()

	log.Println("Recognizer Initialized")

	dataImage := filepath.Join(dataDir, "sumit.jpg")

	faces, err := rec.RecognizeFile(dataImage)
	if err != nil {
		log.Println("Can't recognize: %v", err)
		return err
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

	sumit, err := rec.RecognizeSingle(image)
	if err != nil {
		log.Println("Can't recognize: %v", err)
		return err
	}
	if sumit == nil {
		log.Panicln("Not a single face on the image")
		return err
	}

	fmt.Println("Face Recorganize")

	return nil
}
