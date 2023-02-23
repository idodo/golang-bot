package helper

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

const (
	APIMethodGet = "GET"
)

type ApiHelper struct {
	Token      string
	Type       string
	Language   string
	BaseUrl    string
	QueryParam string
}

func NewApiHelper(token, baseUrl, apiType, language string) *ApiHelper {
	apiHelper := &ApiHelper{Token: token, Type: "Bot", BaseUrl: "https://www.kaiheila.cn", Language: "zh-CN"}

	if baseUrl != "" {
		apiHelper.BaseUrl = baseUrl
	}
	if apiType != "" {
		apiHelper.Type = apiType
	}
	if language != "" {
		apiHelper.Language = language
	}
	return apiHelper
}
func (h *ApiHelper) SetQuery(values map[string]string) {

	for k, v := range values {
		if h.QueryParam == "" {
			h.QueryParam = fmt.Sprintf("&%s=%s", k, v)
		}
		h.QueryParam += fmt.Sprintf("&%s=%s", k, v)
	}

}
func (h *ApiHelper) Get(path string) ([]byte, error) {
	client := &http.Client{}
	reqPath := ""
	if strings.HasPrefix(path, "/") || strings.HasSuffix(h.BaseUrl, "/") {
		reqPath = h.BaseUrl + path
	} else {
		reqPath = h.BaseUrl + "/" + path
	}
	if h.QueryParam != "" {
		reqPath += "?" + h.QueryParam
	}
	req, _ := http.NewRequest("GET", reqPath, nil)
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", h.Type, h.Token))
	req.Header.Set("Accept-Language", h.Language)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.WithField("statusCode", resp.StatusCode).Error("http error", reqPath)
		return nil, errors.New("http error")
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
