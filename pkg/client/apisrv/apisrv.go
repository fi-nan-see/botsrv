// Code generated from jsonrpc schema by rpcgen v2.4.3; DO NOT EDIT.

package apisrv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/vmkteam/zenrpc/v2"
)

var (
	// Always import time package. Generated models can contain time.Time fields.
	_ time.Time
)

type Client struct {
	rpcClient *rpcClient

	Plans *svcPlans
}

func NewDefaultClient(endpoint string) *Client {
	return NewClient(endpoint, http.Header{}, &http.Client{})
}

func NewClient(endpoint string, header http.Header, httpClient *http.Client) *Client {
	c := &Client{
		rpcClient: newRPCClient(endpoint, header, httpClient),
	}

	c.Plans = newClientPlans(c.rpcClient)

	return c
}

type HandlerPlanPreviewDto struct {
	// Текущий баланс
	Current_balance int `json:"current_balance"`
	// Дата окончания периода
	End_date string `json:"end_date"`
	// ID плана
	ID string `json:"id"`
	// Активен ли данный план сейчас
	Is_actual bool `json:"is_actual"`
	// Название плана
	Name string `json:"name"`
	// Дата начала периода
	Start_date string `json:"start_date"`
}

type svcPlans struct {
	client *rpcClient
}

func newClientPlans(client *rpcClient) *svcPlans {
	return &svcPlans{
		client: client,
	}
}

var (
	ErrPlansAddIncome404 = zenrpc.NewError(404, fmt.Errorf("plan or user were not found"))
)

// AddIncome adds income
func (c *svcPlans) AddIncome(ctx context.Context, tgId int, planId string, amount float64) (err error) {
	_req := struct {
		TgID   int
		PlanID string
		Amount float64
	}{
		TgID: tgId, PlanID: planId, Amount: amount,
	}

	err = c.client.call(ctx, "plans.AddIncome", _req, nil)

	switch v := err.(type) {
	case *zenrpc.Error:
		if v.Code == 404 {
			err = ErrPlansAddIncome404
		}
	}

	return
}

var (
	ErrPlansAddOutcome404 = zenrpc.NewError(404, fmt.Errorf("plan or user were not found"))
)

// AddOutcome adds outcome
func (c *svcPlans) AddOutcome(ctx context.Context, tgId int, planId string, amount float64) (err error) {
	_req := struct {
		TgID   int
		PlanID string
		Amount float64
	}{
		TgID: tgId, PlanID: planId, Amount: amount,
	}

	err = c.client.call(ctx, "plans.AddOutcome", _req, nil)

	switch v := err.(type) {
	case *zenrpc.Error:
		if v.Code == 404 {
			err = ErrPlansAddOutcome404
		}
	}

	return
}

var (
	ErrPlansAddSavings404 = zenrpc.NewError(404, fmt.Errorf("plan or user were not found"))
)

// AddSavings adds savings
func (c *svcPlans) AddSavings(ctx context.Context, tgId int, planId string, amount float64) (err error) {
	_req := struct {
		TgID   int
		PlanID string
		Amount float64
	}{
		TgID: tgId, PlanID: planId, Amount: amount,
	}

	err = c.client.call(ctx, "plans.AddSavings", _req, nil)

	switch v := err.(type) {
	case *zenrpc.Error:
		if v.Code == 404 {
			err = ErrPlansAddSavings404
		}
	}

	return
}

func (c *svcPlans) GetPlansHandler(ctx context.Context, tgId int) (res []HandlerPlanPreviewDto, err error) {
	_req := struct {
		TgID int
	}{
		TgID: tgId,
	}

	err = c.client.call(ctx, "plans.GetPlansHandler", _req, &res)

	return
}

type rpcClient struct {
	endpoint string
	cl       *http.Client

	requestID uint64
	header    http.Header
}

func newRPCClient(endpoint string, header http.Header, httpClient *http.Client) *rpcClient {
	return &rpcClient{
		endpoint: endpoint,
		header:   header,
		cl:       httpClient,
	}
}

func (rc *rpcClient) call(ctx context.Context, methodName string, request, result interface{}) error {
	// encode params
	bts, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("encode params: %w", err)
	}

	requestID := atomic.AddUint64(&rc.requestID, 1)
	requestIDBts := json.RawMessage(strconv.Itoa(int(requestID)))

	req := zenrpc.Request{
		Version: zenrpc.Version,
		ID:      &requestIDBts,
		Method:  methodName,
		Params:  bts,
	}

	res, err := rc.Exec(ctx, req)
	if err != nil {
		return err
	}

	if res == nil {
		return nil
	}

	if res.Error != nil {
		return res.Error
	}

	if res.Result == nil {
		return nil
	}

	if result == nil {
		return nil
	}

	return json.Unmarshal(*res.Result, result)
}

// Exec makes http request to jsonrpc endpoint and returns json rpc response.
func (rc *rpcClient) Exec(ctx context.Context, rpcReq zenrpc.Request) (*zenrpc.Response, error) {
	if n, ok := ctx.Value("JSONRPC2-Notification").(bool); ok && n {
		rpcReq.ID = nil
	}

	c, err := json.Marshal(rpcReq)
	if err != nil {
		return nil, fmt.Errorf("json marshal call failed: %w", err)
	}

	buf := bytes.NewReader(c)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rc.endpoint, buf)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header = rc.header.Clone()
	req.Header.Add("Content-Type", "application/json")

	if xRequestID, ok := ctx.Value("X-Request-Id").(string); ok && req.Header.Get("X-Request-Id") == "" && xRequestID != "" {
		req.Header.Add("X-Request-Id", xRequestID)
	}

	// Do request
	resp, err := rc.cl.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, fmt.Errorf("make request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad response (%d)", resp.StatusCode)
	}

	var zresp zenrpc.Response
	if rpcReq.ID == nil {
		return &zresp, nil
	}

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response body (%s) read failed: %w", bb, err)
	}

	if err = json.Unmarshal(bb, &zresp); err != nil {
		return nil, fmt.Errorf("json decode failed (%s): %w", bb, err)
	}

	return &zresp, nil
}
