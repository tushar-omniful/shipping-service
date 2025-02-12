package ajex

import (
	"context"
	"fmt"

	"github.com/omniful/shipping-service/pkg/api"

	"github.com/omniful/shipping-service/internal/domain/interfaces"
	"github.com/omniful/shipping-service/internal/domain/models"
)

type AuthTokenAdapter struct {
	psmRepo interfaces.PartnerShippingMethodRepository
	client  *api.RESTClient
}

type authResponse struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
}

func NewAuthTokenAdapter(client *api.RESTClient, psmRepo interfaces.PartnerShippingMethodRepository) *AuthTokenAdapter {
	return &AuthTokenAdapter{
		client:  client,
		psmRepo: psmRepo,
	}
}

// GetAuthToken gets the current token from PSM
func (a *AuthTokenAdapter) GetAuthToken(psm *models.PartnerShippingMethod) (string, error) {
	token, ok := psm.Credentials["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access token not found in credentials")
	}
	return token, nil
}

// HandleUnauthorized refreshes token and updates PSM
func (a *AuthTokenAdapter) HandleUnauthorized(ctx context.Context, psm *models.PartnerShippingMethod) (*models.PartnerShippingMethod, error) {
	// Get username/password from credentials
	//username, ok := psm.Credentials["username"].(string)
	//if !ok {
	//	return nil, fmt.Errorf("username not found in credentials")
	//}
	//password, ok := psm.Credentials["password"].(string)
	//if !ok {
	//	return nil, fmt.Errorf("password not found in credentials")
	//}

	// Make auth request
	//authReq := map[string]string{
	//	"username": username,
	//	"password": password,
	//}

	//respBody, err := a.client.Send(ctx, &client.Request{
	//	Method: "POST",
	//	Path:   "/authentication-service/api/auth/login",
	//	Headers: map[string][]string{
	//		"Content-Type": {"application/json"},
	//	},
	//	Body: authReq,
	//})
	//if err != nil {
	//	return nil, fmt.Errorf("auth request failed: %w", err)
	//}

	// Parse response
	//var authResp authResponse
	//if err := json.Unmarshal((&respBody).([]byte), &authResp); err != nil {
	//	return nil, fmt.Errorf("unmarshal auth response: %w", err)
	//}
	//
	//// Update PSM credentials with new token
	//psm.Credentials["access_token"] = fmt.Sprintf("%s %s", authResp.TokenType, authResp.AccessToken)
	//psm.Credentials["token_type"] = authResp.TokenType
	//
	//// Update PSM in database
	//if err := a.psmRepo.UpdatePartnerShippingMethod(ctx, map[string]interface{}{
	//	"id": psm.ID,
	//}, psm); err.Exists() {
	//	return nil, fmt.Errorf("failed to update PSM: %w", err)
	//}

	return psm, nil
}
