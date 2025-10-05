package util

import (
	"os"
)

type Config struct {
	GatewayPort      string
	GatewayHost      string
	SipGatewayHost   string
	SipGatewayPort   string
	SipSignalingHost string
	SipSignalingPort string
	DevicePath       string
	MockMode         bool
}

func LoadConfig() Config {
	cfg := Config{
		GatewayPort:      getEnv("TELEPHONY_GATEWAY_SERVICE_PORT", "8080"),
		GatewayHost:      getEnv("TELEPHONY_GATEWAY_SERVICE_HOST", "telephony-gateway"),
		SipGatewayHost:   getEnv("SIP_GATEWAY_SERVICE_HOST", "sip-gateway"),
		SipGatewayPort:   getEnv("SIP_GATEWAY_SERVICE_PORT", "5060"),
		SipSignalingHost: getEnv("SIP_SIGNALING_SERVICE_HOST", "sip-signaling"),
		SipSignalingPort: getEnv("SIP_SIGNALING_SERVICE_PORT", "5060"),
		DevicePath:       getEnv("DEVICE_PATH", "/dev/dahdi"),
		MockMode:         getEnv("MOCK_MODE", "true") == "true",
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
