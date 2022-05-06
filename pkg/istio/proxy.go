package istio

import (
	"fmt"
	"net/http"
	"time"
)

const WAIT_ATTEMPTS = 10

func WaitForProxyAvailability() error {
	for i := 0; i < WAIT_ATTEMPTS; i++ {
		resp, _ := http.Get("http://localhost:15021/healthz/ready")
		if resp != nil && resp.StatusCode == 200 {
			return nil
		}

		time.Sleep(time.Second)
	}

	return fmt.Errorf("proxy was not available in suitable time frame")
}

func TriggerProxyShutdown() {
	_, _ = http.Post("http://localhost:15020/quitquitquit", "plain/text", nil)
}
