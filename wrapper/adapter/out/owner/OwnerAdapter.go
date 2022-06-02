package owner

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
)

const (
	BaseURLV1 = "http://localhost:9090/api/v1/employees"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
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

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BaseURLV1,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
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

	fmt.Println("paso 2")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ErrorResponse

		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {

			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

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
