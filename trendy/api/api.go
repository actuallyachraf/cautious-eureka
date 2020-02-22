package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/actuallyachraf/backend-challenge/trendy/models"
	"github.com/gorilla/mux"
)

const (
	// TrendingRoute is the main API route for fetching the top trending projects.
	TrendingRoute = "/trending"
	// LanguageRoute is the subroute to fetch the top trending projects by language.
	LanguageRoute = "/trending/{language}"
	// Port is the configuration port to expose the HTTP service.
	Port = ":3000"
	// dateFormat is a layout specifier for dates
	dateFormat = "2006-01-02"
)

var (
	// substract 1 month from today's date to get the 30 days result
	aMonthAgo = time.Now().AddDate(0, -1, 0).Format(dateFormat)
	// Github API URL
	url = fmt.Sprintf("https://api.github.com/search/repositories?q=created:>%s&sort=stars&order=desc&per_page=100", aMonthAgo)
)

// fetchData executes the Get request to the remote API and returns the
// json result as binary data.
func fetchData(startDate string, url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("failed to make GET request with error :", err)
		return nil, err
	}
	defer resp.Body.Close()

	jsonBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return jsonBody, nil

}

// TrendingHandler handles request to the trending route
func TrendingHandler(w http.ResponseWriter, r *http.Request) {

	jsonBody, err := fetchData(aMonthAgo, url)
	if err != nil {
		log.Fatal("failed to read json body with error :", err)
		w.Write(jsonErrResponse(err, "failed to fetch data from Github API"))
		return
	}

	req := &models.RequestedData{}
	err = json.Unmarshal(jsonBody, req)
	if err != nil {
		log.Fatal("failed to unmarshal json response with error:", err)
		w.Write(jsonErrResponse(err, "failed to fetch data from Github API"))
		return
	}
	response, err := json.MarshalIndent(req, "", " ")
	if err != nil {
		log.Fatal("failed to marshal json response with error:", err)
		w.Write(jsonErrResponse(err, "failed to build proper response"))
		return
	}
	w.Write(response)
}

// LanguageHandler handles request to the language route
func LanguageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	language, ok := vars["language"]
	if !ok {
		log.Fatal("no language in path")
		w.Write(jsonErrResponse(errors.New("bad API request"), "language specifier can't be null"))
		return
	}

	jsonBody, err := fetchData(aMonthAgo, url)
	req := &models.RequestedData{
		Count: 100,
	}
	err = json.Unmarshal(jsonBody, req)
	if err != nil {
		log.Fatal("failed to parse json response with error:", err)
		w.Write(jsonErrResponse(err, "failed to unmarshal json response"))
		return
	}

	resp := &models.LanguageInfo{
		Language: language,
	}
	var count int
	for _, item := range req.Items {
		if item.Language == language {
			count++
			resp.Links = append(resp.Links, item.URL)
		}
	}
	resp.Count = count

	response, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		log.Fatal("failed to marshal json response with error:", err)
		w.Write(jsonErrResponse(err, "failed to marshal response"))
		return
	}
	w.Write(response)

}
func jsonErrResponse(err error, msg string) []byte {

	errResponse := struct {
		Error   error
		Message string
	}{
		Error:   err,
		Message: msg,
	}
	// we are ommiting serialization errors because we're already failing
	errResponseBytes, _ := json.Marshal(errResponse)

	return errResponseBytes
}
