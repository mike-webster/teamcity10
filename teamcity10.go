package teamcity10

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetToken retrieves a short-use token for the user.
// Life is 60 seconds.
// The benefit here is it front-loads the authentication
// and hopefully speeds up the following calls.
func GetToken(ctx context.Context) (string, error) {
	resp, err := makeRequest(ctx, "GET", "httpAuth/app/rest/server", nil)
	if err != nil {
		return "", err
	}

	tok := ""
	for _, h := range resp.Cookies() {
		if h.Name == "TCSESSIONID" {
			tok = h.Value
		}
	}
	if len(tok) < 1 {
		fmt.Println(resp.Header)
		return "", errors.New("no token returned")
	}

	return tok, nil
}

// TriggerBuild will add a build to the build queue
// It will return true if it was able to add the build to the queue and false if not.
func TriggerBuild(ctx context.Context, buildID, branchName, projectID string) (bool, error) {
	path := "app/rest/buildQueue"
	body := struct {
		Name string `json:"branchName"`
		Type struct {
			ID  string `json:"id"`
			PID string `json:"projectId"`
		} `json:"buildType"`
	}{
		Name: branchName,
		Type: struct {
			ID  string `json:"id"`
			PID string `json:"projectId"`
		}{
			ID:  buildID,
			PID: projectID,
		},
	}
	bytes, err := json.Marshal(body)
	if err != nil {
		return false, err
	}

	resp, err := makeRequest(ctx, "POST", path, bytes)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, errors.New(fmt.Sprint("error reading response body; ", err.Error()))
	}

	if resp.StatusCode > 299 {
		return false, errors.New(fmt.Sprint("non-200: ", resp.StatusCode, ";\n", string(b)))
	}

	respStruct := struct {
		ID int32 `json:"id"`
	}{}

	err = json.Unmarshal(b, &respStruct)
	if err != nil {
		return false, err
	}

	if respStruct.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetBuilds retrieves some information about the recent builds for a given project id.
// This method expects either a teamcity token or teamcity credentials to be in the context
func GetBuilds(ctx context.Context, id string) (*[]Build, error) {
	path := fmt.Sprint("app/rest/buildTypes/id:", id, "/builds?fields=build(id,buildTypeId,number,status,state,branchName,href,finishDate,triggered(type,user(username,name)))")
	resp, err := makeRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	builds := struct {
		Builds []Build `json:"build"`
	}{}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprint("error reading response body; ", err.Error()))
	}

	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprint("non-200: ", string(b)))
	}

	err = json.Unmarshal(b, &builds)
	if err != nil {
		return nil, err
	}

	return &builds.Builds, nil
}

func buildRequest(method, url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "Application/json")
	req.Header.Add("Content-Type", "Application/json")

	return req, nil
}

func addShortLiveToken(ctx context.Context, req *http.Request) {
	// idk if this will work or if I'll need to return the request  with changes
	tok := ctx.Value(TeamCityToken)
	if tok != nil {
		c := http.Cookie{
			Name:    "TCSESSIONID",
			Value:   fmt.Sprint(tok),
			Domain:  "/",
			Path:    "/",
			Expires: time.Now().Add(1 * time.Minute),
		}

		req.AddCookie(&c)
		req.Header.Add("tcsessionid", fmt.Sprint(tok))
	}
}

func addBasicCreds(ctx context.Context, req *http.Request) {
	creds := ctx.Value(TeamCityCreds)
	if creds != nil {
		req.Header.Add("Authorization", fmt.Sprint("Basic ", creds))
	}
}

func makeRequest(ctx context.Context, method string, path string, body []byte) (*http.Response, error) {
	cfg, err := getConfig(ctx)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprint(cfg.BaseURL, "/", path)
	req, err := buildRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	addShortLiveToken(ctx, req)
	addBasicCreds(ctx, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
