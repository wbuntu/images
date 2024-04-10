package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type RequestBody struct {
	Model string `json:"model"`
	// 其他字段根据实际情况添加
}

func main() {
	// 添加 hosts 映射
	if err := addHostsMapping(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	// 根据 model 字段返回对应的 url
	determineTargetURL := func(body []byte) (string, error) {
		// 解析请求体
		var requestBody RequestBody
		err := json.Unmarshal(body, &requestBody)
		if err != nil {
			return "", fmt.Errorf("failed to parse JSON body: %s", err)
		}

		// 根据 model 字段选择后端
		var backendPort string
		switch requestBody.Model {
		case "gpt-3.5-turbo":
			backendPort = "3040"
		case "gpt-3.5-search":
			backendPort = "8000"
		default:
			return "", fmt.Errorf("unknown model %s", requestBody.Model)
		}

		// 构造目标 URL
		return fmt.Sprintf("http://localhost:%s", backendPort), nil
	}
	// 创建反向代理器
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// 从请求体中读取字段
			body, err := io.ReadAll(req.Body)
			if err != nil {
				log.Println("Error reading request body:", err)
				return
			}
			// 将数据重写回 body
			req.Body = io.NopCloser(bytes.NewBuffer(body))
			// 根据请求体中的字段来确定目标服务器地址
			targetURL, err := determineTargetURL(body)
			if err != nil {
				log.Println("Error determining target URL:", err)
				return
			}

			// 解析目标服务器地址
			target, err := url.Parse(targetURL)
			if err != nil {
				log.Println("Error parsing target URL:", err)
				return
			}

			// 设置请求的目标地址
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.Host = target.Host
		},
	}
	// 启动代理服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 通过自定义的处理函数来处理请求
		proxy.ServeHTTP(w, r)
	})
	fmt.Println("Proxy server listening on :6789")
	err := http.ListenAndServe(":6789", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func addHostsMapping() error {
	// 获取当前主机的 hostname
	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("failed to get hostname: %w", err)
	}

	// 构造要添加的 host 字符串
	hosts := fmt.Sprintf("127.0.0.1\t%s\n127.0.0.1\tbackend\n127.0.0.1\tfreegpt35\n127.0.0.1\tsearxng\n", hostname)

	// 读取当前 /etc/hosts 文件内容
	currentHosts, err := os.ReadFile("/etc/hosts")
	if err != nil {
		return fmt.Errorf("failed to read /etc/hosts: %w", err)
	}

	// 检查是否已经写入 hosts
	if !strings.Contains(string(currentHosts), "searxng") {
		// 追加到 /etc/hosts 文件中
		err = os.WriteFile("/etc/hosts", append(currentHosts, []byte(hosts)...), 0644)
		if err != nil {
			return fmt.Errorf("failed to append to /etc/hosts: %w", err)
		}
		fmt.Println("Hosts mapping added to /etc/hosts file.")
	} else {
		fmt.Println("Hosts mapping already exists in /etc/hosts file.")
	}
	return nil
}
