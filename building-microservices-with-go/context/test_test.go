package context

import (
	"testing"
	"net/http"
	"time"
	"fmt"
	"context"
)

// Samples and documentation
// https://golang.org/pkg/context/#WithTimeout
// https://talks.golang.org/2014/gotham-context.slide#11

func TestSiteRequestCanceledByTimeout(t *testing.T){

	r, _ := http.NewRequest("GET", "https://google.com", nil)

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)
	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func TestAnyMethodCallCanceledByTimeout(t *testing.T) {

	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
