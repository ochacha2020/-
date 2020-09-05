package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
)

type visionRequest struct {
	BaseImg string `json:"base_img"`
}

type visionResponse struct {
	Anger    int `json:"anger"`
	Joy      int `json:"joy"`
	Surprise int `json:"surprise"`
	Sorrow   int `json:"sorrow"`
}

type errResponse struct {
	Error string `json:"error"`
}

func CallVisionApi(w http.ResponseWriter, r *http.Request) {
	// CORSのためにset
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	visionRequest := visionRequest{}
	err := json.NewDecoder(r.Body).Decode(&visionRequest)
	if err != nil {
		handleError(w, err.Error())
		return
	}

	fmt.Println(len(visionRequest.BaseImg))

	visionResponse, err := detectFaces(w, visionRequest.BaseImg)
	if err != nil {
		handleError(w, err.Error())
		return
	}

	response, err := json.Marshal(visionResponse)
	if err != nil {
		handleError(w, err.Error())
		return
	}

	writeJsonFormat(w, string(response))

}

// detectFaces gets faces from the Vision API for an image at the given file path.
func detectFaces(w io.Writer, imageBase64 string) (*visionResponse, error) {
	// stringのbase64画像をreaderに変換
	data, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return nil, fmt.Errorf("DecodeString %v", err)
	}
	wb := new(bytes.Buffer)
	wb.Write(data)
	reader := bytes.NewReader(wb.Bytes())

	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewImageAnnotatorClient %v", err)
	}
	defer client.Close()

	image, err := vision.NewImageFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("NewImageFromReader %v", err)
	}

	annotations, err := client.DetectFaces(ctx, image, nil, 10)
	if err != nil {
		return nil, fmt.Errorf("DetectFaces %v", err)
	}

	visionResponse := visionResponse{}
	if len(annotations) == 0 {
		fmt.Println("No faces found.")
	} else {
		annotation := annotations[0]
		visionResponse.Anger = calcApiScore(annotation.AngerLikelihood.String())
		visionResponse.Joy = calcApiScore(annotation.JoyLikelihood.String())
		visionResponse.Surprise = calcApiScore(annotation.SurpriseLikelihood.String())
		visionResponse.Sorrow = calcApiScore(annotation.SorrowLikelihood.String())

	}
	return &visionResponse, nil
}

func calcApiScore(likelyhood string) int {
	switch likelyhood {
	case "VERY_LIKELY":
		return 100
	case "LIKELY":
		return 70
	case "POSSIBLE":
		return 50
	case "UNLIKELY":
		return 20
	case "VERY_UNLIKELY":
		return 0
	case "UNKNOWN":
		return 0
	}
	return 0
}

func writeJsonFormat(w http.ResponseWriter, contents string) {
	// jsonヘッダーを出力
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, contents)
}

func handleError(w http.ResponseWriter, errStr string) {
	errResponse := errResponse{}
	errResponse.Error = errStr
	outputJSON, err := json.Marshal(&errResponse)
	if err != nil {
		panic(err)
	}

	writeJsonFormat(w, string(outputJSON))
}
