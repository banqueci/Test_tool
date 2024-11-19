package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/quic-go/quic-go"
	"io"
	"log"
)

func main() {
	// 创建一个 QUIC 监听器
	listener, err := quic.ListenAddr(":4242", generateTLSConfig(), nil)
	if err != nil {
		log.Fatalf("Failed to create listener: %v\n", err)
	}
	fmt.Println("Server is listening on port 4242...")

	for {
		// 接受一个新会话
		session, err := listener.Accept(context.Background())
		if err != nil {
			log.Printf("Failed to accept session: %v\n", err)
			continue
		}
		go handleSession(session) // 处理每个会话
	}
}

func handleSession(session quic.Connection) {
	fmt.Printf("New connection from: %s\n", session.RemoteAddr())

	for {
		// 接受一个新流
		stream, err := session.AcceptStream(context.Background())
		if err != nil {
			log.Printf("Failed to accept stream: %v\n", err)
			break
		}
		go handleStream(stream) // 处理每个流
	}
}

func handleStream(stream quic.Stream) {
	defer stream.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := stream.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Stream read error: %v\n", err)
			return
		}
		message := string(buffer[:n])
		fmt.Printf("Received message: %s\n", message)
	}
}

// 创建 TLS 配置
func generateTLSConfig() *tls.Config {
	cert, err := tls.X509KeyPair([]byte(serverCert), []byte(serverKey))
	if err != nil {
		log.Fatalf("Failed to load TLS keys: %v\n", err)
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}}
}

const serverCert = `-----BEGIN CERTIFICATE-----
MIIDTTCCAjWgAwIBAgIUDyJngCoyPRSMnSDcUGMmn7cqXJcwDQYJKoZIhvcNAQEL
BQAwNjELMAkGA1UEBhMCY24xCzAJBgNVBAgMAmpzMQswCQYDVQQHDAJuajENMAsG
A1UECgwEbWlndTAeFw0yNDExMTgwODQ1MjlaFw0yNTExMTgwODQ1MjlaMDYxCzAJ
BgNVBAYTAmNuMQswCQYDVQQIDAJqczELMAkGA1UEBwwCbmoxDTALBgNVBAoMBG1p
Z3UwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCwmfCB3Jp6sOFCnuwM
ugqJ9t5jbIX8U0GOuDPjtQLiKaL9hHMPIQTvpD8vnefP4RfOgFRjCpXbFx6zwfE+
VnR6EMNVmLeNPlumP0JvpF5vVTFKfaS6zcq+eBMJLE2MNZdqUz42rbJoUT5nqmrb
EJrDcryWQ4qpyKxB5tKHYlcbJWNAj6bHnILQnRf8+OaYQGM7x65qI/sXrmZpyCWN
P75NRvfmP50oMcROk9ckn6GeJLhWN8LITTNxTuuiij9tBmFrWHdEL6fyVZbMuVDI
qrnZsqoPp9eb4J6VwqkGSKR0+rbrjRg3PlckiPMYvUrPK8qqvOHLRBxV/6YTlAM7
mrA1AgMBAAGjUzBRMB0GA1UdDgQWBBTJ9D8OR8LnGGUMNI6ACgM2cucjrDAfBgNV
HSMEGDAWgBTJ9D8OR8LnGGUMNI6ACgM2cucjrDAPBgNVHRMBAf8EBTADAQH/MA0G
CSqGSIb3DQEBCwUAA4IBAQCjyrQI8lpQnSO+2Z0s3+EmMlLDiy0hcVsDYWvNK7hk
YyVw2EQh2pYtV0cUfaDCXocjdFCljvbgruawWUUgpnapzxe+FeTJTVRN2GDFnVPX
CBrFO3TtdmG5xgUTk4D8sYus/mPl6VAi9HK640qdRJSjYPlWgs2i8CDmHwOSzWBX
VyQzVCItztM2YuZykBe5Cpo2hk5o4WlBkFxgbFg/nhqxelxX4ADOYWSOTVvnK8nA
rspJ7mPqmpdFDrRqlhV4E/Xpc6YcoM9xREgrhZTV1B6VXekv73J+IoSK45WA6q/T
jgOi7wkM644xkJHrFVjhsz08FDEopM5pCVxqZqPnTcv7
-----END CERTIFICATE-----`

const serverKey = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCwmfCB3Jp6sOFC
nuwMugqJ9t5jbIX8U0GOuDPjtQLiKaL9hHMPIQTvpD8vnefP4RfOgFRjCpXbFx6z
wfE+VnR6EMNVmLeNPlumP0JvpF5vVTFKfaS6zcq+eBMJLE2MNZdqUz42rbJoUT5n
qmrbEJrDcryWQ4qpyKxB5tKHYlcbJWNAj6bHnILQnRf8+OaYQGM7x65qI/sXrmZp
yCWNP75NRvfmP50oMcROk9ckn6GeJLhWN8LITTNxTuuiij9tBmFrWHdEL6fyVZbM
uVDIqrnZsqoPp9eb4J6VwqkGSKR0+rbrjRg3PlckiPMYvUrPK8qqvOHLRBxV/6YT
lAM7mrA1AgMBAAECggEBAKQWVHaE3/ARuytoeFEW50XrYBSnOBML1ITkOCph/Izh
PPUrhLDQ8EItrRv0Lvhi3/jaGg5wPm70G23QTPOr5suoGabFn/6XqxZb/sG2ypvN
X2AkP9Pl9HRnIuDKDj+xZ7++Grc5SuRMYJx3ppNm6TWivQEtPoCk2RDehEyQnfi0
hxV/fzHLUk9iTQfBoH24GFuwxOzNMirLz8NR9PibgZgU9Y+fZeZqtXrUkowcI5KJ
QVctvE4GcDYkfTMVy4akH9pxCvuQpBt8KCbUlxWw78iQEWUVeNerV2MAIATAAx+7
P43SRxJIlTdx4egTvJcSlmDm9GObMTcPlmiPdGc39D0CgYEA5jWo4ICJGkwv1xkT
FCnwzIkwZytTPsYVdwabCkUaVyl+F6XGVXhLPjKvLnVvrZga6tNPWwAAPUcymmjW
FtjgSiAKKNoMP8JrTkZCsWXX6bJCBgYVqhFqaY/htVmxiLZ0UBrrIXBKTwDL/zgV
mEmM3oZR2Mqd4eCTu2jECbFd+AsCgYEAxGLPlPIQCC1u3KXEUHbK08NUw0FhHfRm
rmMoBtwzIfNUhDrrmYzPf5tuwON95gNAb35LrctTD4vfuOM4qAuIsiuxB4BmjDp1
Ou3jsJebFwuuVPQY2AMcMFRq4CNl26DvXTRA7gA3/BUHDXSYMJlxPjd9nbUbkdKr
UiRbiuB/4L8CgYEAoc3PX9QYCUrJSUcPeNmtrSUzxx/Vh8aUKa+T41kElYTNYnOa
/lHpmNLo/B+AmiPRr4FMQmqywF89evf6md5fbtosKeBwQZQ19bM+hw9M/a3T6AX4
zislfwKpItzjAnMzN2ZkI4GYSQUHXOAflYUEpRcFifmHlM5TJ6MQPrvSj1cCgYBX
nRWrB4sresl42uOIWlcGvqA0NBjVulGM/2O+G8McJGjSTU8KxA1WisuQdm2WjmDS
3O96a7l0uBxpacW/AtZLVr618AzQBsyCK9tmz1w7ndR6xiPHSyvqS98ae/BXWacw
V72X8LUJW4A3+opjNDGXZj4+e6v/FJOmI95LSPkEFQKBgHCTDE+Z4guReqKWy7xX
GfXQfg4//bgWru30E2/B3hDoJkHimBK4r3s4IXt2ogXpyK+h0ROxSukX/S7kZeX6
gyxoYmIsxG0lc3brJfKkYZv+b64NczTJrrZ9rkO4JlyZhTrytXsvgR92a3X0Zfu4
Jx3K17E822TDKEhQkD0umQPz
-----END PRIVATE KEY-----`
