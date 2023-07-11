package helpers
import(
	"net/http"
	"time"
	log "github.com/sirupsen/logrus"
)
func SendRequestWithRetries(req *http.Request) (*http.Response, error) {
    var resp *http.Response
    var err error
    for i := 0; i < 3; i++ { // Retry up to 3 times
        if i > 0 {
            time.Sleep(3 * time.Second) // Wait 3 second before retrying
        }
        resp, err = http.DefaultClient.Do(req)
        if err == nil && resp.StatusCode == http.StatusOK {
            return resp, nil // Request succeeded
        }
        log.WithFields(log.Fields{
            "retry": i+1,
            "error": err,
            "status_code": resp.StatusCode,
        }).Warn("API request failed, retrying...")
    }
    return resp, err // All retries failed
}