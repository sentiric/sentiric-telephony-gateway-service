# Sentiric Telephony Gateway Service

**Description:** Integrates the Sentiric platform with traditional telephone networks (PSTN), TDM, or other legacy telecommunication systems.

**Core Responsibilities:**
*   Translating traditional telephony protocols (e.g., ISDN, SS7, TDM) to Sentiric's internal communication protocols (SIP/RTP) and vice-versa.
*   Interacting with specialized telecommunication hardware drivers if necessary.
*   Handling call setup and teardown across different network types.

**Technologies:**
*   Go (or C++)
*   Specialized telephony libraries or direct hardware interfaces.

**API Interactions (As a Protocol Bridge):**
*   Communicates directly with `sentiric-sip-gateway-service` (via SIP protocol for converted calls) or `sentiric-sip-signaling-service`.
*   Interacts with external telecommunication networks/hardware.

**Local Development:**
1.  Clone this repository: `git clone https://github.com/sentiric/sentiric-telephony-gateway-service.git`
2.  Navigate into the directory: `cd sentiric-telephony-gateway-service`
3.  Install dependencies (Go: `go mod tidy` / C++: Build system config).
4.  Create a `.env` file from `.env.example` to configure hardware interfaces or legacy network connections.
5.  Start the service: (Go: `go run main.go` / C++: Build and run executable).

**Configuration:**
Refer to `config/` directory and `.env.example` for service-specific configurations, including hardware parameters and protocol settings.

**Deployment:**
Designed for containerized deployment (e.g., Docker, Kubernetes), potentially requiring privileged access or specific hardware drivers. Refer to `sentiric-infrastructure`.

**Contributing:**
We welcome contributions! Please refer to the [Sentiric Governance](https://github.com/sentiric/sentiric-governance) repository for coding standards and contribution guidelines.

**License:**
This project is licensed under the [License](LICENSE).
