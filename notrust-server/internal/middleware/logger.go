package middleware

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuditEvent struct {
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
	Event     string `json:"event"`
	Detail    string `json:"detail,omitempty"`
}

func AuditLog(c *fiber.Ctx, event string, detail string) {
	entry := AuditEvent{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		IP:        c.IP(),
		Event:     event,
		Detail:    detail,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Error serializando log: %v", err)
		return
	}

	log.Println(string(data))
}

