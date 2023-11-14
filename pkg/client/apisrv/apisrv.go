// Code generated from jsonrpc schema by rpcgen v2.4.1; DO NOT EDIT.

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

	Application *svcApplication
	Message     *svcMessage
	Offer       *svcOffer
	Production  *svcProduction
	User        *svcUser
}

func NewDefaultClient(endpoint string) *Client {
	return NewClient(endpoint, http.Header{}, &http.Client{})
}

func NewClient(endpoint string, header http.Header, httpClient *http.Client) *Client {
	c := &Client{
		rpcClient: newRPCClient(endpoint, header, httpClient),
	}

	c.Application = newClientApplication(c.rpcClient)
	c.Message = newClientMessage(c.rpcClient)
	c.Offer = newClientOffer(c.rpcClient)
	c.Production = newClientProduction(c.rpcClient)
	c.User = newClientUser(c.rpcClient)

	return c
}

type Application struct {
	ApplicationComment         *string  `json:"applicationComment,omitempty"`
	ApplicationKey             string   `json:"applicationKey"`
	ApplicationResponseTimeout *int     `json:"applicationResponseTimeout,omitempty"`
	CreatedAt                  *string  `json:"createdAt"`
	DetailComment              *string  `json:"detailComment,omitempty"`
	DetailHeight               *int     `json:"detailHeight,omitempty"`
	DetailLength               *int     `json:"detailLength,omitempty"`
	DetailName                 string   `json:"detailName"`
	DetailQuantity             *int     `json:"detailQuantity,omitempty"`
	DetailWeight               float64  `json:"detailWeight"`
	DetailWidth                *int     `json:"detailWidth,omitempty"`
	Drawings                   []string `json:"drawings"`
	Email                      *string  `json:"email,omitempty"`
	ID                         int      `json:"id"`
	IsTollingMaterial          bool     `json:"isTollingMaterial"`
	LimitingPrice              *float64 `json:"limitingPrice,omitempty"`
	MaterialName               string   `json:"materialName"`
	OffersQuantity             *int     `json:"offersQuantity,omitempty"`
	OffersReceivedCount        *int     `json:"offersReceivedCount,omitempty"`
	OperationList              []string `json:"operationList"`
	Phone                      *string  `json:"phone,omitempty"`
	ProductionRequestedCount   *int     `json:"productionRequestedCount,omitempty"`
	Region                     *string  `json:"region,omitempty"`
	RegionSearch               string   `json:"regionSearch"`
	ShippingAt                 *string  `json:"shippingAt,omitempty"`
	StatusID                   int      `json:"statusId"`
	TelegramUserID             *int     `json:"telegramUserId,omitempty"`
	TelegramUsername           *string  `json:"telegramUsername,omitempty"`
	UnitServiceCost            *float64 `json:"unitServiceCost,omitempty"`
	UnitTotalCost              *float64 `json:"unitTotalCost,omitempty"`
}

type ApplicationSearch struct {
	ApplicationComment         *string  `json:"applicationComment,omitempty"`
	ApplicationKey             *string  `json:"applicationKey,omitempty"`
	ApplicationResponseTimeout *int     `json:"applicationResponseTimeout,omitempty"`
	CreatedAt                  *string  `json:"createdAt,omitempty"`
	DetailComment              *string  `json:"detailComment,omitempty"`
	DetailHeight               *int     `json:"detailHeight,omitempty"`
	DetailLength               *int     `json:"detailLength,omitempty"`
	DetailName                 *string  `json:"detailName,omitempty"`
	DetailQuantity             *int     `json:"detailQuantity,omitempty"`
	DetailWeight               *float64 `json:"detailWeight,omitempty"`
	DetailWidth                *int     `json:"detailWidth,omitempty"`
	Email                      *string  `json:"email,omitempty"`
	ID                         *int     `json:"id,omitempty"`
	Ids                        []int    `json:"ids"`
	IsTollingMaterial          *bool    `json:"isTollingMaterial,omitempty"`
	LimitingPrice              *float64 `json:"limitingPrice,omitempty"`
	MaterialName               *string  `json:"materialName,omitempty"`
	OffersQuantity             *int     `json:"offersQuantity,omitempty"`
	OffersReceivedCount        *int     `json:"offersReceivedCount,omitempty"`
	Phone                      *string  `json:"phone,omitempty"`
	ProductionRequestedCount   *int     `json:"productionRequestedCount,omitempty"`
	Region                     *string  `json:"region,omitempty"`
	RegionSearch               *string  `json:"regionSearch,omitempty"`
	ShippingAt                 *string  `json:"shippingAt,omitempty"`
	StatusID                   *int     `json:"statusId,omitempty"`
	TelegramUserID             *int     `json:"telegramUserId,omitempty"`
	TelegramUsername           *string  `json:"telegramUsername,omitempty"`
	UnitServiceCost            *float64 `json:"unitServiceCost,omitempty"`
	UnitTotalCost              *float64 `json:"unitTotalCost,omitempty"`
}

type ApplicationSummary struct {
	ApplicationComment         *string  `json:"applicationComment,omitempty"`
	ApplicationKey             string   `json:"applicationKey"`
	ApplicationResponseTimeout *int     `json:"applicationResponseTimeout,omitempty"`
	CreatedAt                  *string  `json:"createdAt"`
	DetailComment              *string  `json:"detailComment,omitempty"`
	DetailHeight               *int     `json:"detailHeight,omitempty"`
	DetailLength               *int     `json:"detailLength,omitempty"`
	DetailName                 string   `json:"detailName"`
	DetailQuantity             *int     `json:"detailQuantity,omitempty"`
	DetailWeight               float64  `json:"detailWeight"`
	DetailWidth                *int     `json:"detailWidth,omitempty"`
	Drawings                   []string `json:"drawings"`
	Email                      *string  `json:"email,omitempty"`
	ID                         int      `json:"id"`
	IsTollingMaterial          bool     `json:"isTollingMaterial"`
	LimitingPrice              *float64 `json:"limitingPrice,omitempty"`
	MaterialName               string   `json:"materialName"`
	OffersQuantity             *int     `json:"offersQuantity,omitempty"`
	OffersReceivedCount        *int     `json:"offersReceivedCount,omitempty"`
	OperationList              []string `json:"operationList"`
	Phone                      *string  `json:"phone,omitempty"`
	ProductionRequestedCount   *int     `json:"productionRequestedCount,omitempty"`
	Region                     *string  `json:"region,omitempty"`
	RegionSearch               string   `json:"regionSearch"`
	ShippingAt                 *string  `json:"shippingAt,omitempty"`
	StatusID                   *int     `json:"statusId,omitempty"`
	TelegramUserID             *int     `json:"telegramUserId,omitempty"`
	TelegramUsername           *string  `json:"telegramUsername,omitempty"`
	UnitServiceCost            *float64 `json:"unitServiceCost,omitempty"`
	UnitTotalCost              *float64 `json:"unitTotalCost,omitempty"`
}

type Message struct {
	Application   *ApplicationSummary `json:"application,omitempty"`
	ApplicationID *int                `json:"applicationId,omitempty"`
	ChatID        int                 `json:"chatId"`
	MessageID     int                 `json:"messageId"`
	Offer         *Offer              `json:"offer,omitempty"`
	OfferID       *int                `json:"offerId,omitempty"`
	StatusID      int                 `json:"statusId"`
	Type          string              `json:"type"`
}

type MessagePair struct {
}

type MessageSearch struct {
	ApplicationID *int    `json:"ApplicationID,omitempty"`
	ChatID        *int    `json:"ChatID,omitempty"`
	MessageID     *int    `json:"MessageID,omitempty"`
	OfferID       *int    `json:"OfferID,omitempty"`
	StatusID      *int    `json:"StatusID,omitempty"`
	Type          *string `json:"Type,omitempty"`
}

type Offer struct {
	ApplicationID            int      `json:"ApplicationID"`
	Comment                  *string  `json:"Comment,omitempty"`
	CreatedAt                *string  `json:"CreatedAt"`
	DeliveryCost             *float64 `json:"DeliveryCost,omitempty"`
	DetailQuantity           int      `json:"DetailQuantity"`
	ID                       int      `json:"ID"`
	IsTollingMaterial        bool     `json:"IsTollingMaterial"`
	ProductionBotUserID      int      `json:"ProductionBotUserID"`
	ProductionID             int      `json:"ProductionID"`
	Rating                   *int     `json:"Rating,omitempty"`
	ShippingAt               string   `json:"ShippingAt"`
	StatusID                 int      `json:"StatusID"`
	TechnologicalProcessList []string `json:"TechnologicalProcessList"`
	TotalUnitCost            *float64 `json:"TotalUnitCost,omitempty"`
	UnitServiceCost          *float64 `json:"UnitServiceCost,omitempty"`
	UnitWeight               *float64 `json:"UnitWeight,omitempty"`
}

type Production struct {
	BIK                          *string  `json:"bIK,omitempty"`
	BankAccount                  *string  `json:"bankAccount,omitempty"`
	CityLocation                 *string  `json:"cityLocation,omitempty"`
	CompanyName                  *string  `json:"companyName,omitempty"`
	Email                        *string  `json:"email,omitempty"`
	INN                          *string  `json:"iNN,omitempty"`
	ID                           int      `json:"id"`
	Logo                         *string  `json:"logo,omitempty"`
	Phone                        *string  `json:"phone,omitempty"`
	ProductionDescription        *string  `json:"productionDescription,omitempty"`
	Rating                       *float64 `json:"rating,omitempty"`
	ServiceAndProductDescription *string  `json:"serviceAndProductDescription,omitempty"`
	TechnologicalProcessList     []string `json:"technologicalProcessList"`
	Website                      *string  `json:"website,omitempty"`
}

type ProductionRequest struct {
	Application                 *ApplicationSummary `json:"application,omitempty"`
	ApplicationFilingDeadlineAt *string             `json:"applicationFilingDeadlineAt,omitempty"`
	ApplicationID               int                 `json:"applicationId"`
	CancelReason                *string             `json:"cancelReason,omitempty"`
	ID                          int                 `json:"id"`
	IsExpired                   bool                `json:"isExpired"`
	Production                  *ProductionSummary  `json:"production,omitempty"`
	ProductionID                int                 `json:"productionId"`
	ReminderSendTimeAt          *string             `json:"reminderSendTimeAt,omitempty"`
	RequestAcceptDeadlineAt     string              `json:"requestAcceptDeadlineAt"`
	StatusID                    int                 `json:"statusId"`
}

type ProductionRequestSearch struct {
	ApplicationFilingDeadlineAt *string `json:"applicationFilingDeadlineAt,omitempty"`
	ApplicationID               *int    `json:"applicationId,omitempty"`
	CancelReason                *string `json:"cancelReason,omitempty"`
	ID                          *int    `json:"id,omitempty"`
	Ids                         []int   `json:"ids"`
	IsExpired                   *bool   `json:"isExpired,omitempty"`
	ProductionID                *int    `json:"productionId,omitempty"`
	ReminderSendTimeAt          *string `json:"reminderSendTimeAt,omitempty"`
	RequestAcceptDeadlineAt     *string `json:"requestAcceptDeadlineAt,omitempty"`
	StatusID                    *int    `json:"statusId,omitempty"`
}

type ProductionSearch struct {
	BIK                          *string  `json:"bIK,omitempty"`
	BankAccount                  *string  `json:"bankAccount,omitempty"`
	CityLocation                 *string  `json:"cityLocation,omitempty"`
	CompanyName                  *string  `json:"companyName,omitempty"`
	Email                        *string  `json:"email,omitempty"`
	INN                          *string  `json:"iNN,omitempty"`
	ID                           *int     `json:"id,omitempty"`
	Ids                          []int    `json:"ids"`
	Logo                         *string  `json:"logo,omitempty"`
	Phone                        *string  `json:"phone,omitempty"`
	ProductionDescription        *string  `json:"productionDescription,omitempty"`
	Rating                       *float64 `json:"rating,omitempty"`
	ServiceAndProductDescription *string  `json:"serviceAndProductDescription,omitempty"`
	StatusID                     *int     `json:"statusId,omitempty"`
	Website                      *string  `json:"website,omitempty"`
}

type ProductionSummary struct {
	BIK                          *string   `json:"bIK,omitempty"`
	BankAccount                  *string   `json:"bankAccount,omitempty"`
	CityLocation                 string    `json:"cityLocation"`
	CompanyName                  *string   `json:"companyName,omitempty"`
	Email                        string    `json:"email"`
	INN                          string    `json:"iNN"`
	ID                           int       `json:"id"`
	Logo                         *string   `json:"logo,omitempty"`
	Phone                        *string   `json:"phone,omitempty"`
	ProductionDescription        *string   `json:"productionDescription,omitempty"`
	Rating                       *float64  `json:"rating,omitempty"`
	ServiceAndProductDescription *string   `json:"serviceAndProductDescription,omitempty"`
	Similarity                   *float64  `json:"similarity,omitempty"`
	Status                       *VtStatus `json:"status,omitempty"`
	TechnologicalProcessList     []string  `json:"technologicalProcessList"`
	Website                      *string   `json:"website,omitempty"`
	WordList                     []string  `json:"wordList"`
}

type User struct {
	IsAdmin          bool    `json:"isAdmin"`
	Name             *string `json:"name,omitempty"`
	ProductionID     int     `json:"productionId"`
	TelegramUserID   *int    `json:"telegramUserId,omitempty"`
	TelegramUsername string  `json:"telegramUsername"`
}

type VtStatus struct {
	Alias string `json:"alias"`
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type VtViewOps struct {
	// page number, default - 1
	Page int `json:"page"`
	// items count per page, max - 500
	PageSize int `json:"pageSize"`
	// sort by column name
	SortColumn string `json:"sortColumn"`
	// descending sort
	SortDesc bool `json:"sortDesc"`
}

type svcApplication struct {
	client *rpcClient
}

func newClientApplication(client *rpcClient) *svcApplication {
	return &svcApplication{
		client: client,
	}
}

// Add application
func (c *svcApplication) Add(ctx context.Context, application Application) (res *ApplicationSummary, err error) {
	_req := struct {
		Application Application
	}{
		Application: application,
	}

	err = c.client.call(ctx, "application.Add", _req, &res)

	return
}

// Get application
func (c *svcApplication) Get(ctx context.Context, search *ApplicationSearch, viewOps *VtViewOps) (res []ApplicationSummary, err error) {
	_req := struct {
		Search  *ApplicationSearch
		ViewOps *VtViewOps
	}{
		Search: search, ViewOps: viewOps,
	}

	err = c.client.call(ctx, "application.Get", _req, &res)

	return
}

func (c *svcApplication) GetByID(ctx context.Context, id int) (res *Application, err error) {
	_req := struct {
		ID int
	}{
		ID: id,
	}

	err = c.client.call(ctx, "application.GetByID", _req, &res)

	return
}

// GetByKey Get get application by unique key
func (c *svcApplication) GetByKey(ctx context.Context, hash string) (res []ApplicationSummary, err error) {
	_req := struct {
		Hash string
	}{
		Hash: hash,
	}

	err = c.client.call(ctx, "application.GetByKey", _req, &res)

	return
}

// GetHistory gets all applications history for a certain provider
func (c *svcApplication) GetHistory(ctx context.Context, id int) (res []ApplicationSummary, err error) {
	_req := struct {
		ID int
	}{
		ID: id,
	}

	err = c.client.call(ctx, "application.GetHistory", _req, &res)

	return
}

// Update application
func (c *svcApplication) Update(ctx context.Context, application Application) (res bool, err error) {
	_req := struct {
		Application Application
	}{
		Application: application,
	}

	err = c.client.call(ctx, "application.Update", _req, &res)

	return
}

type svcMessage struct {
	client *rpcClient
}

func newClientMessage(client *rpcClient) *svcMessage {
	return &svcMessage{
		client: client,
	}
}

func (c *svcMessage) Add(ctx context.Context, message Message) (res bool, err error) {
	_req := struct {
		Message Message
	}{
		Message: message,
	}

	err = c.client.call(ctx, "message.Add", _req, &res)

	return
}

func (c *svcMessage) Get(ctx context.Context, search MessageSearch) (res MessagePair, err error) {
	_req := struct {
		Search MessageSearch
	}{
		Search: search,
	}

	err = c.client.call(ctx, "message.Get", _req, &res)

	return
}

func (c *svcMessage) Update(ctx context.Context, message Message) (res bool, err error) {
	_req := struct {
		Message Message
	}{
		Message: message,
	}

	err = c.client.call(ctx, "message.Update", _req, &res)

	return
}

type svcOffer struct {
	client *rpcClient
}

func newClientOffer(client *rpcClient) *svcOffer {
	return &svcOffer{
		client: client,
	}
}

func (c *svcOffer) Add(ctx context.Context, offer Offer) (res *Offer, err error) {
	_req := struct {
		Offer Offer
	}{
		Offer: offer,
	}

	err = c.client.call(ctx, "offer.Add", _req, &res)

	return
}

func (c *svcOffer) GetByID(ctx context.Context, id int) (res *Offer, err error) {
	_req := struct {
		ID int
	}{
		ID: id,
	}

	err = c.client.call(ctx, "offer.GetByID", _req, &res)

	return
}

func (c *svcOffer) GetByIDs(ctx context.Context, productionId int, applicationId int) (res []Offer, err error) {
	_req := struct {
		ProductionID  int
		ApplicationID int
	}{
		ProductionID: productionId, ApplicationID: applicationId,
	}

	err = c.client.call(ctx, "offer.GetByIDs", _req, &res)

	return
}

type svcProduction struct {
	client *rpcClient
}

func newClientProduction(client *rpcClient) *svcProduction {
	return &svcProduction{
		client: client,
	}
}

func (c *svcProduction) Add(ctx context.Context, production Production) (res *Production, err error) {
	_req := struct {
		Production Production
	}{
		Production: production,
	}

	err = c.client.call(ctx, "production.Add", _req, &res)

	return
}

func (c *svcProduction) AddProductionRequest(ctx context.Context, productionRequest ProductionRequest) (res *ProductionRequest, err error) {
	_req := struct {
		ProductionRequest ProductionRequest
	}{
		ProductionRequest: productionRequest,
	}

	err = c.client.call(ctx, "production.AddProductionRequest", _req, &res)

	return
}

func (c *svcProduction) AddRelevantProductionsToDB(ctx context.Context, orderID int, viewOps *VtViewOps) (res bool, err error) {
	_req := struct {
		OrderID int
		ViewOps *VtViewOps
	}{
		OrderID: orderID, ViewOps: viewOps,
	}

	err = c.client.call(ctx, "production.AddRelevantProductionsToDB", _req, &res)

	return
}

func (c *svcProduction) All(ctx context.Context) (res []ProductionSummary, err error) {
	_req := struct {
	}{}

	err = c.client.call(ctx, "production.All", _req, &res)

	return
}

func (c *svcProduction) AllProductionRequests(ctx context.Context) (res []ProductionRequest, err error) {
	_req := struct {
	}{}

	err = c.client.call(ctx, "production.AllProductionRequests", _req, &res)

	return
}

func (c *svcProduction) Get(ctx context.Context, search *ProductionSearch, viewOps *VtViewOps) (res []ProductionSummary, err error) {
	_req := struct {
		Search  *ProductionSearch
		ViewOps *VtViewOps
	}{
		Search: search, ViewOps: viewOps,
	}

	err = c.client.call(ctx, "production.Get", _req, &res)

	return
}

func (c *svcProduction) GetByID(ctx context.Context, id int) (res *Production, err error) {
	_req := struct {
		ID int
	}{
		ID: id,
	}

	err = c.client.call(ctx, "production.GetByID", _req, &res)

	return
}

func (c *svcProduction) GetByKey(ctx context.Context, hash string) (res *Production, err error) {
	_req := struct {
		Hash string
	}{
		Hash: hash,
	}

	err = c.client.call(ctx, "production.GetByKey", _req, &res)

	return
}

func (c *svcProduction) GetKey(ctx context.Context, p Production) (res string, err error) {
	_req := struct {
		P Production
	}{
		P: p,
	}

	err = c.client.call(ctx, "production.GetKey", _req, &res)

	return
}

func (c *svcProduction) GetProductionRequest(ctx context.Context, orderID int, productionID int) (res *ProductionRequest, err error) {
	_req := struct {
		OrderID      int
		ProductionID int
	}{
		OrderID: orderID, ProductionID: productionID,
	}

	err = c.client.call(ctx, "production.GetProductionRequest", _req, &res)

	return
}

func (c *svcProduction) GetRelevantList(ctx context.Context, applicationId int, pageSize int, page int) (res []ProductionSummary, err error) {
	_req := struct {
		ApplicationID int
		PageSize      int
		Page          int
	}{
		ApplicationID: applicationId, PageSize: pageSize, Page: page,
	}

	err = c.client.call(ctx, "production.GetRelevantList", _req, &res)

	return
}

func (c *svcProduction) SearchProductionRequests(ctx context.Context, search *ProductionRequestSearch) (res []ProductionRequest, err error) {
	_req := struct {
		Search *ProductionRequestSearch
	}{
		Search: search,
	}

	err = c.client.call(ctx, "production.SearchProductionRequests", _req, &res)

	return
}

func (c *svcProduction) Update(ctx context.Context, production Production) (res bool, err error) {
	_req := struct {
		Production Production
	}{
		Production: production,
	}

	err = c.client.call(ctx, "production.Update", _req, &res)

	return
}

func (c *svcProduction) UpdateProductionRequest(ctx context.Context, productionRequest ProductionRequest) (res bool, err error) {
	_req := struct {
		ProductionRequest ProductionRequest
	}{
		ProductionRequest: productionRequest,
	}

	err = c.client.call(ctx, "production.UpdateProductionRequest", _req, &res)

	return
}

type svcUser struct {
	client *rpcClient
}

func newClientUser(client *rpcClient) *svcUser {
	return &svcUser{
		client: client,
	}
}

func (c *svcUser) Add(ctx context.Context, user User) (res *User, err error) {
	_req := struct {
		User User
	}{
		User: user,
	}

	err = c.client.call(ctx, "user.Add", _req, &res)

	return
}

func (c *svcUser) All(ctx context.Context) (res []User, err error) {
	_req := struct {
	}{}

	err = c.client.call(ctx, "user.All", _req, &res)

	return
}

func (c *svcUser) GetByProductionID(ctx context.Context, id int) (res *User, err error) {
	_req := struct {
		ID int
	}{
		ID: id,
	}

	err = c.client.call(ctx, "user.GetByProductionId", _req, &res)

	return
}

func (c *svcUser) GetByTgID(ctx context.Context, id int) (res *User, err error) {
	_req := struct {
		ID int
	}{
		ID: id,
	}

	err = c.client.call(ctx, "user.GetByTgID", _req, &res)

	return
}

func (c *svcUser) Update(ctx context.Context, user User) (res bool, err error) {
	_req := struct {
		User User
	}{
		User: user,
	}

	err = c.client.call(ctx, "user.Update", _req, &res)

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
