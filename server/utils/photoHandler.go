package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/fogleman/gg"
	"github.com/rwcarlsen/goexif/exif"
)

type PhotoHandler interface {
	IsCorrectResolution(photo io.Reader) (bool, error)
	RegisterMetadata(photo io.Reader) error
}

type InsurancePhotoHandler struct {
	apiKey     string
	pathToFont string
	pathToSave string
	fontSize   float64
}

func NewInsurancePhotoHandler(apiKey, pathToFont, pathToSave string, fontsize float64) *InsurancePhotoHandler {
	return &InsurancePhotoHandler{
		apiKey:     apiKey,
		pathToFont: pathToFont,
		pathToSave: pathToSave,
		fontSize:   fontsize,
	}
}

var (
	url = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address"
)

func (h *InsurancePhotoHandler) IsCorrectResolution(photo io.Reader) (bool, error) {
	img, _, err := image.Decode(photo)
	if err != nil {
		return false, err
	}

	// fmt.Println(img.Bounds().Dx(), "x", img.Bounds().Dy())
	if img.Bounds().Dx() < 1600 || img.Bounds().Dy() < 1200 {
		return false, nil
	}

	return true, nil
}

func (h *InsurancePhotoHandler) RegisterMetadata(photo io.Reader) error {
	x, err := exif.Decode(photo)
	if err != nil {
		log.Printf("InsurancePhotoHandler.RegisterMetadata: %v", err)
		return err
	}

	textToPrint := ""

	datetime, err := x.DateTime()
	if err != nil {
		log.Printf("Datetime Tag not found")
	} else {
		textToPrint += datetime.Format("01-02-2006 03:04:05 MST")
	}

	lat, long, err := x.LatLong()
	if err != nil {
		log.Printf("GeoTag not found")
	} else {
		output, err := h.getAddresByLatLong(lat, long)
		if err != nil {
			log.Printf("InsurancePhotoHandler.RegisterMetadata: %v", err)
			return err
		}
		textToPrint += "\n" + output.Suggestions[0].Unrestricted_value
	}

	if err := h.printTextToImage(photo, textToPrint); err != nil {
		log.Printf("InsurancePhotoHandler.RegisterMetadata: %v", err)
		return err
	}
	return nil
}

type requestData struct {
	Lat   float64 `json:"lat"`
	Long  float64 `json:"lon"`
	Count int     `json:"count"`
}
type responseOutput struct {
	Suggestions []suggest `json:"suggestions"`
}
type suggest struct {
	Value              string `json:"value"`
	Unrestricted_value string `json:"unrestricted_value"`
}

func (h *InsurancePhotoHandler) printTextToImage(img io.Reader, text string) error {
	bgImg, imgName, err := image.Decode(img)
	if err != nil {
		return err
	}
	imgWidth := bgImg.Bounds().Dx()
	imgHeight := bgImg.Bounds().Dy()

	canvas := gg.NewContextForImage(bgImg)
	canvas.SetRGB255(255, 0, 0)
	canvas.LoadFontFace(h.pathToFont, h.fontSize)
	canvas.DrawStringWrapped(text, float64(imgWidth/2), float64(imgHeight-100), 0.5, 0.5, float64(imgWidth-150), 1.0, gg.AlignRight)

	filename := filepath.Join(filepath.Dir(h.pathToSave), "modified-"+imgName)
	if err = gg.SaveJPG(filename, canvas.Image(), 60); err != nil {
		return err
	}

	return nil
}

func (h *InsurancePhotoHandler) getAddresByLatLong(lat, long float64) (responseOutput, error) {
	client := http.Client{
		Timeout: time.Second,
	}
	data := requestData{
		Lat:   lat,
		Long:  long,
		Count: 1,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return responseOutput{}, err
	}

	token := fmt.Sprintf("Token %s", h.apiKey)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return responseOutput{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return responseOutput{}, nil
	}
	defer resp.Body.Close()

	var output responseOutput
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return responseOutput{}, err
	}

	return output, nil
}
