package telemetry

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/valist-io/valist/log"
)

var logger = log.New()

func RecordDownload(project string) {
	telemetry_api := fmt.Sprintf("https://stats.valist.io/api/download/%s", project)
	req, err := http.NewRequest(http.MethodPut, telemetry_api, &bytes.Buffer{})

	if err != nil {
		logger.Error("%v", err)
		return
	}

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		logger.Error("%v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error("%v", err)
		return
	}

	if res.StatusCode > 299 {
		logger.Error("failed to connect to telemetry service: status=%s body=%s", res.Status, body)
	}
}
