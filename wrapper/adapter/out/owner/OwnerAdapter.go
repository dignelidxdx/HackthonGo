package owner

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
	"github.com/mercadolibre/fury_go-core/pkg/breaker"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
)

var (
	// Circuit breaker
	CircuitBreakerSettings = gobreaker.Settings{
		MaxRequests: uint32(1),
		Timeout:     time.Minute,
		Interval:    time.Hour,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.4 // At least 40% of requests failed
		},
	}

	rg  = rand.New(rand.NewSource(time.Now().Unix()))
	r   = rate.Every(2 * time.Second)
	lim = rate.NewLimiter(r, 3)
)

const (
	BaseURLV1 = "http://localhost:9090/api/v1/employees"
)

type Client struct {
	BaseURL        string
	apiKey         string
	HTTPClient     *http.Client
	circuitBreaker *CircuitBreaker
}

type Employee struct {
	ID                int    `json:"id,omitempty"`
	CardNumberID      string `json:"card_number_id,omitempty" validate:"required"`
	FirstName         string `json:"first_name" validate:"required"`
	LastName          string `json:"last_name" validate:"required"`
	WarehouseID       int    `json:"warehouse_id" validate:"required"`
	InboundOrderCount int    `json:"inbound_orders_count,omitempty"`
}

type EmployeeData struct {
	Data []Employee `json:"data,omitempty"`
}

func NewClient(apiKey string, circuitBreaker *CircuitBreaker) *Client {
	return &Client{
		BaseURL: BaseURLV1,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
		circuitBreaker: circuitBreaker,
	}
}

func (c *Client) GetEmployees(ctx context.Context, options *domain.Employee) (*domain.Employee, error) {

	req, err := http.NewRequest("GET", BaseURLV1, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := EmployeeData{}

	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}
	res2 := domain.Employee{}
	if len(res.Data) > 0 {
		res2.ID = res.Data[0].ID
		res2.FirstName = res.Data[0].FirstName
		res2.LastName = res.Data[0].LastName
	}

	return &res2, nil
}

func (c *Client) SendRequest(req *http.Request, v *EmployeeData) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	// Circuit Breaker fury core

	failureRatio := .9 // allow 30% of failures before opening the circuit.
	cb := breaker.NewBreaker(failureRatio)

	// Assert circuit breaker status before doing operation.
	var res *http.Response

	//limiter := golimiter.New(3000, 100*time.Millisecond)
	//result, err := limiter.Action(10, )

	err := cb.Do(func() (retErr error) {
		if lim.Allow() {
			fmt.Println("was allowed")
			res, retErr = c.HTTPClient.Do(req)
			return
		}
		return
	})

	if err != nil {
		fmt.Println("was not allowed")
		if errors.Is(err, breaker.ErrCircuitOpen) {
			fmt.Println("Operation did not execute because breaker was open.")
			// Operation did not execute because breaker was open.
		}
		fmt.Println("Operation was executed, the error is from the operation.")
		// Operation was executed, the error is from the operation.
	}
	fmt.Println("was successful")

	/*executed, respX, err := c.circuitBreaker.Execute("circuitEmployee", func() (interface{}, error) {

		return c.HTTPClient.Do(req)
	})
	fmt.Println("executed:", executed)
	if !executed {
		// Not executed because of circuit breaker
		//metrics.ProcessPusherCircuitBreaker("action:limited", fmt.Sprintf("subscriber:%s", distributionMessage.Subscriber))
		fmt.Println("hubo interrupcion por circuit breaker")
		fmt.Println(err)
		return err
	}

	res := respX.(*http.Response)

	if err != nil {
		fmt.Println("hubo un error despues de circuit breaker")
		return err
	}

	fmt.Println("sucessful")*/
	if res != nil {
		defer res.Body.Close()
	}

	if res == nil || res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		//	var errRes ErrorResponse

		// if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {

		// 	return errors.New(errRes.Message)
		// }

		if res != nil {
			return fmt.Errorf("unknown error, status code: %d", res.StatusCode)

		}
		return fmt.Errorf("unknown error, status code: 500")
	}
	fmt.Println("paso 3")
	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type GoCircuitBreaker interface {
	State() gobreaker.State
	Execute(f func() (interface{}, error)) (interface{}, error)
	Name() string
}

type CircuitBreaker struct {
	cBreakers map[string]GoCircuitBreaker
}

func NewCircuitBreaker() *CircuitBreaker {
	return &CircuitBreaker{
		cBreakers: make(map[string]GoCircuitBreaker),
	}
}

// Execute returns a boolean indicating if the function was executed and it's return values
func (c *CircuitBreaker) Execute(key string, f func() (interface{}, error)) (bool, interface{}, error) {
	breaker, exists := c.cBreakers[key]
	if !exists {
		breaker = newCircuitBreaker(key)
		c.cBreakers[key] = breaker
	}
	res, err := breaker.Execute(f)
	if err == gobreaker.ErrOpenState || err == gobreaker.ErrTooManyRequests {
		return false, nil, nil
	}
	return true, res, err
}

func newCircuitBreaker(subscriber string) GoCircuitBreaker {
	settings := CircuitBreakerSettings
	settings.Name = "cb_" + subscriber

	return gobreaker.NewCircuitBreaker(settings)
}
