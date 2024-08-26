package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var currentStatus = &CustomStatus{}

func SetStatus(ctx context.Context, status string, emoji string) error {
	return setStatus(ctx, NewStatusRequest(status, emoji))
}

func ClearStatus(ctx context.Context) error {
	return setStatus(ctx, &StatusRequest{CustomStatus: nil})
}

func setStatus(ctx context.Context, statusReq *StatusRequest) error {
	if !shouldUpdate(statusReq) {
		return nil
	}

	url := discordAPI + userSettings

	body, err := json.Marshal(statusReq)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrorStatus, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrorStatus, err)
	}

	token, ok := os.LookupEnv(discordTokenEnv)
	if !ok {
		return ErrTokenNotFound
	}

	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrorStatus, err)
	}
	defer req.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d: %w", res.StatusCode, ErrorStatus)
	}

	currentStatus = statusReq.CustomStatus

	if statusReq.CustomStatus == nil {
		log.Print("Status cleared")
	} else {
		log.Print("Status updated to: ", statusReq.CustomStatus.Text)
	}
	return nil
}

func shouldUpdate(statusReq *StatusRequest) bool {
	custom := statusReq.CustomStatus
	if custom == nil && currentStatus == nil {
		return false
	} else if (custom == nil && currentStatus != nil) || (custom != nil && currentStatus == nil) {
		return true
	} else if custom.Text != currentStatus.Text {
		return true
	} else if custom.EmojiName != currentStatus.EmojiName {
		return true
	}
	return false
}
