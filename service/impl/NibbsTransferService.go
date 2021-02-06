package impl

import (
	"ajocard/status-indicator/dto"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type NibbsTransferService struct {}


func (nibbs NibbsTransferService) CheckTransferServiceStatus() (dto.IndicatorDto, error) {
	url := os.Getenv("NIP_LIVE_STATUS_URL")
	maximumAllowedDefaultRate, _ := strconv.ParseFloat(os.Getenv("MAXIMUM_ALLOWED_DEFAULT_RATE"), 64)
	httpClient := http.Client{Timeout: 5 * time.Second}
	res, err := httpClient.Get(url)
	if err != nil {
		log.Panic("An error occurred. ", err.Error())
		return dto.IndicatorDto{}, err
	}
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		var data map[string]map[string] string
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &data)
		failureRateStr := data["msg"]["todayFailureRateOutward"]
		failureRate, _ := strconv.ParseFloat(strings.Replace(failureRateStr, "%", "", 1), 64)
		if failureRate < maximumAllowedDefaultRate {
			return dto.IndicatorDto{true, failureRateStr}, nil
		} else {
			return dto.IndicatorDto{false, failureRateStr}, nil
		}
	} else {
		return dto.IndicatorDto{}, errors.New("Something went wrong")
	}

}

