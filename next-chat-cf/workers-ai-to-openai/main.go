package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

func main() {
	// 初始化 CF
	if err := cfClient.init(); err != nil {
		fmt.Printf("Init cloudflare client: %s\n", err)
		os.Exit(-1)
	}
	// 设置路由处理函数
	http.HandleFunc("/v1/chat/completions", handleCompletions)

	// 启动服务器
	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func handleCompletions(w http.ResponseWriter, r *http.Request) {
	sessionID := uuid.New().String()
	// 检查请求方法是否为 POST
	if r.Method != http.MethodPost {
		httpError(w, http.StatusMethodNotAllowed, fmt.Sprintf("SessionID: %s: Method Not Allowed", sessionID))
		return
	}

	requestData, err := io.ReadAll(r.Body)
	if err != nil {
		httpError(w, http.StatusBadRequest, fmt.Sprintf("SessionID: %s: Bad Request: read body: %s", sessionID, err))
		return
	}
	request := &openai.ChatCompletionRequest{}
	if err := json.Unmarshal(requestData, request); err != nil {
		httpError(w, http.StatusBadRequest, fmt.Sprintf("SessionID: %s: Bad Request: unmarshal request: %s", sessionID, err))
		return
	}
	fmt.Printf("SessionID: %s: Request Received: %s\n", sessionID, requestData)
	if request.Stream {
		streamer, err := cfClient.connect(r.Context(), request)
		if err != nil {
			httpError(w, http.StatusInternalServerError, fmt.Sprintf("SessionID: %s: Internal Server Error: connect: %s", sessionID, err))
			return
		}
		defer streamer.Close()
		streamer.id = fmt.Sprintf("chatcmpl-%s", sessionID)
		streamer.created = time.Now().Unix()
		fmt.Printf("SessionID: %s: Streaming Start\n", sessionID)
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Transfer-Encoding", "chunked")
		if err := cfClient.stream(r.Context(), w, streamer); err != nil {
			fmt.Printf("SessionID: %s: Streaming Error: %s\n", sessionID, err)
			return
		}
		fmt.Printf("SessionID: %s: Streaming Stop\n", sessionID)
	} else {
		response := &openai.ChatCompletionResponse{
			ID:      fmt.Sprintf("chatcmpl-%s", sessionID),
			Created: time.Now().Unix(),
			Model:   request.Model,
		}
		if err := cfClient.complete(r.Context(), request, response); err != nil {
			httpError(w, http.StatusInternalServerError, fmt.Sprintf("SessionID: %s: Internal Server Error: generate completion: %s", sessionID, err))
			return
		}
		responseData, err := json.Marshal(response)
		if err != nil {
			httpError(w, http.StatusInternalServerError, fmt.Sprintf("SessionID: %s: Internal Server Error: marshal response: %s", sessionID, err))
			return
		}
		fmt.Printf("SessionID: %s: Response Sent: %s\n", sessionID, responseData)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseData)
	}
}

func httpError(w http.ResponseWriter, code int, msg string) {
	fmt.Println(msg)
	http.Error(w, msg, code)
}

var cfClient = &client{
	Client: http.Client{
		Transport: &http.Transport{
			// 忽略证书验证，优先使用服务器偏好的加密套件
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify:       true,
				PreferServerCipherSuites: true,
			},
			// 默认启用HTTP2
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			MaxConnsPerHost:       100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		// 配置请求超时
		Timeout: 60 * time.Second,
	},
	retryCount: 3,
	retryWait:  time.Millisecond * 200,
}

type client struct {
	http.Client
	retryCount int
	retryWait  time.Duration
	accountID  string
	apiToken   string
}

func (c *client) init() error {
	accountID := os.Getenv("CF_ACCOUNT_ID")
	if len(accountID) == 0 {
		return errors.New("empty CF_ACCOUNT_ID")
	}
	apiToken := os.Getenv("CF_API_TOKEN")
	if len(apiToken) == 0 {
		return errors.New("empty CF_API_TOKEN")
	}
	c.accountID = accountID
	c.apiToken = apiToken
	return nil
}

func (c *client) connect(ctx context.Context, request *openai.ChatCompletionRequest) (*CFStreamer, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)
	cfRequest := &CFRequest{Stream: true}
	for i := range request.Messages {
		cfRequest.Messages = append(cfRequest.Messages, Message{
			Role:    request.Messages[i].Role,
			Content: request.Messages[i].Content,
		})
	}
	cfReqData, err := json.Marshal(cfRequest)
	if err != nil {
		return nil, fmt.Errorf("marshal cf request: %s", err)
	}
	req, err = http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/ai/run/%s", c.accountID, request.Model),
		bytes.NewBuffer(cfReqData),
	)
	if err != nil {
		return nil, fmt.Errorf("build request: %s", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Set("Content-Type", "application/json")
	// 获取 HTTP 响应，出错时自动重试
	for i := 0; i < c.retryCount; i++ {
		resp, err = c.Do(req)
		if err == nil {
			break
		}
		time.Sleep(c.retryWait)
	}
	if err != nil {
		return nil, fmt.Errorf("send http request: %s", err)
	}
	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})
	s := &CFStreamer{
		readCloser: resp.Body,
		scanner:    scanner,
		model:      request.Model,
	}
	return s, nil
}

func (c *client) stream(ctx context.Context, w http.ResponseWriter, s *CFStreamer) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			for s.scanner.Scan() {
				data := s.scanner.Text()
				if len(data) < len("data: ") {
					continue
				}
				data = strings.TrimPrefix(data, "data: ")
				data = strings.TrimSuffix(data, "\r")
				//fmt.Printf("SessionID: %s: Streaming: data: %s\n", s.id, data)
				if strings.HasPrefix(data, "[DONE]") {
					break
				}
				// 反序列化应用层响应
				cfStreamResponse := &CFStreamResponse{}
				if err := json.Unmarshal([]byte(data), cfStreamResponse); err != nil {
					return fmt.Errorf("unmarshal cf response: %s: %s", err, data)
				}
				// 构建流式传输结构体
				response := &openai.ChatCompletionStreamResponse{}
				response.ID = s.id
				response.Created = s.created
				response.Object = "chat.completion.chunk"
				response.Model = s.model
				response.Choices = []openai.ChatCompletionStreamChoice{
					{
						Index: 0,
						Delta: openai.ChatCompletionStreamChoiceDelta{
							Role:    openai.ChatMessageRoleAssistant,
							Content: cfStreamResponse.Response,
						},
						FinishReason: openai.FinishReasonNull,
					},
				}
				// 无内容时提示停止传输
				if len(cfStreamResponse.Response) == 0 {
					response.Choices[0].FinishReason = openai.FinishReasonStop
				}
				jsonData, err := json.Marshal(response)
				if err != nil {
					return fmt.Errorf("marshal response: %s: %s", err, data)
				}
				m := append([]byte("data: "), jsonData...)
				m = append(m, []byte("\n\n")...)
				w.Write(m)
			}
			// 结束传输
			w.Write([]byte("data: [DONE]\n\n"))
			return nil
		}
	}
}

func (c *client) complete(ctx context.Context, request *openai.ChatCompletionRequest, response *openai.ChatCompletionResponse) error {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)
	cfRequest := &CFRequest{Stream: false}
	for i := range request.Messages {
		cfRequest.Messages = append(cfRequest.Messages, Message{
			Role:    request.Messages[i].Role,
			Content: request.Messages[i].Content,
		})
	}
	cfReqData, err := json.Marshal(cfRequest)
	if err != nil {
		return fmt.Errorf("marshal cf request: %s", err)
	}
	req, err = http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/ai/run/%s", c.accountID, request.Model),
		bytes.NewBuffer(cfReqData),
	)
	if err != nil {
		return fmt.Errorf("build request: %s", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Set("Content-Type", "application/json")
	// 获取 HTTP 响应，出错时自动重试
	for i := 0; i < c.retryCount; i++ {
		resp, err = c.Do(req)
		if err == nil {
			break
		}
		time.Sleep(c.retryWait)
	}
	if err != nil {
		return fmt.Errorf("send http request: %s", err)
	}
	// 读取 body
	defer resp.Body.Close()
	cfRespData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read http response: %s", err)
	}
	// 检查 HTTP 层响应码
	if resp.StatusCode != http.StatusOK {
		errStr := fmt.Sprintf("%s: status: %s:", req.URL, resp.Status)
		if len(cfRespData) > 0 {
			errStr += fmt.Sprintf(" msg: %s", string(cfRespData))
		}
		return errors.New(errStr)
	}
	// 反序列化应用层响应
	cfResponse := &CFResponse{}
	if err := json.Unmarshal(cfRespData, cfResponse); err != nil {
		return fmt.Errorf("unmarshal response: %s: %s", err, string(cfRespData))
	}
	response.Object = "chat.completion"
	response.Choices = []openai.ChatCompletionChoice{
		{
			Index: 0,
			Message: openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: cfResponse.Result.Response,
			},
			FinishReason: openai.FinishReasonStop,
		},
	}
	return nil
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CFRequest struct {
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type CFStreamer struct {
	readCloser io.ReadCloser
	scanner    *bufio.Scanner
	id         string
	created    int64
	model      string
}

func (s *CFStreamer) Close() error {
	return s.readCloser.Close()
}

type CFResponse struct {
	Result struct {
		Response string `json:"response"`
	} `json:"result"`
	Success bool `json:"success"`
}

type CFStreamResponse struct {
	Response string `json:"response"`
}
