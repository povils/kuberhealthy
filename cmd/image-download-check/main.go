package main

import (
	"os"
	"strconv"
	"time"

	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/kuberhealthy/kuberhealthy/v2/pkg/checks/external/checkclient"
	"github.com/kuberhealthy/kuberhealthy/v2/pkg/checks/external/nodeCheck"
	log "github.com/sirupsen/logrus"
)

var (

	/*
		// privateRegistryURL sets the URL for a private image registry
		privateRegistryURL = os.Getenv("PRIVATE_REGISTRY_URL")

		// imageName sets the name for the image to be pulled
		imageName = os.Getenv("IMAGE_NAME")

		// imageTag sets the tag for the impage to be pulled
		imageTag = os.Getenv("IMAGE_TAG")
	*/

	// fullImageURL is the full registry + image name + tag URL for ease of testing
	fullImageURL = os.Getenv("FULL_IMAGE_URL")

	// timeoutLimit sets the maximum amount of time in seconds that an
	// an expected image pull should not breach
	timeoutLimit = os.Getenv("TIMEOUT_LIMIT")
)

func init() {

	// set debug mode for nodeCheck pkg
	nodeCheck.EnableDebugOutput()

	/*
		// check to make sure privateRegistryURL string is provided
		if privateRegistryURL == "" {
			reportErrorAndStop("No PRIVATE_REGISTRY_URL string provided in YAML")
		}

		// check to make sure imageName string is provided
		if imageName == "" {
			reportErrorAndStop("No IMAGE_NAME string provided in YAML")
		}
		// check to make sure imageTag string is provided
		if imageTag == "" {
			reportErrorAndStop("No IMAGE_TAG string provided in YAML")
		}
	*/

	// check to make sure fullImageURL string is provided
	if fullImageURL == "" {
		reportErrorAndStop("No FULL_IMAGE_URL string provided in YAML")
	}

}

func main() {

	// run check
	pass := checkPass()

	// report to kh
	if pass {
		log.Println("it passed!")
	} else {
		log.Println("it failed =(")
	}

}

// checkPass implements the logic to pull an image, track a start and end time, then
// determines if the actual pull time is greater than the expected limit threshold
func checkPass() bool {

	// initialize a start time
	startTime := time.Now()

	// download image
	img, err := downloadImage()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("successfully downloaded image: ", img)

	// calculate time it took to complete image download
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	durationSeconds := int(duration.Seconds())
	log.Println("image took this long to download: ", durationSeconds)

	// determine if duration exceeds the time limit threshold
	timeoutLimitInt, err := strconv.Atoi(timeoutLimit)
	if err != nil {
		log.Println("there was an error converting string of timeoutLimit to an int ", err)
	}

	if durationSeconds < timeoutLimitInt {
		return true
	}

	return false
}

func downloadImage() (v1.Image, error) {

	// pull image
	// i, err := crane.Pull("balenalib/rpi-alpine-node")
	// i, err := crane.Pull("docker-proto.repo.theplatform.com/kube-deploy:1.22")
	i, err := crane.Pull(fullImageURL)
	if err != nil {
		return nil, err
	}

	// save image tarball to path
	err = crane.Save(i, "saved_image", "./image")
	if err != nil {
		return nil, err
	}

	// get layer count - informative
	l, err := i.Layers()
	if err != nil {
		return nil, err
	}
	log.Println("layer count", len(l))

	// get image size - informative
	s, err := i.Size()
	if err != nil {
		return nil, err
	}
	log.Println("image size", s)

	return i, nil
}

// reportKHSuccess reports success to Kuberhealthy servers and verifies the report successfully went through
func reportKHSuccess() error {
	err := checkclient.ReportSuccess()
	if err != nil {
		log.Println("Error reporting success to Kuberhealthy servers:", err)
		return err
	}
	log.Println("Successfully reported success to Kuberhealthy servers")
	return err
}

// reportKHFailure reports failure to Kuberhealthy servers and verifies the report successfully went through
func reportKHFailure(errorMessage string) error {
	err := checkclient.ReportFailure([]string{errorMessage})
	if err != nil {
		log.Println("Error reporting failure to Kuberhealthy servers:", err)
		return err
	}
	log.Println("Successfully reported failure to Kuberhealthy servers")
	return err
}

// reportErrorAndStop reports to kuberhealthy of error and exits program when called
func reportErrorAndStop(s string) {
	log.Infoln("attempting to report error to kuberhealthy:", s)
	err := checkclient.ReportFailure([]string{s})
	if err != nil {
		log.Errorln("failed to report to kuberhealthy servers:", err)
		os.Exit(1)
	}
	log.Infoln("Successfully reported to Kuberhealthy")
	os.Exit(0)
}