package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// Lark消息结构
type LarkMessage struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

// 初始化配置
func init() {
	viper.SetConfigName("alert.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/opt/xidp/conf")
	viper.ReadInConfig()
}

func main() {
	//如果不存在，则只打印日志
	webhookURL := viper.GetString("Lark.custom_bot")
	if webhookURL == "" {
		log.Println("Lark.custom_bot is not set")
		return
	} else {
		SendToLark("test message", webhookURL)
	}
}

func SendToLark(text string, webhookURL string) {

	log.Println(webhookURL)

	// 创建Lark消息结构
	message := LarkMessage{
		MsgType: "text",
	}
	message.Content.Text = text

	// 转换为JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Printf("JSON编码失败: %v", err)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	defer resp.Body.Close()

	log.Println(resp)

	if resp.StatusCode == http.StatusOK {
		log.Println("消息已成功发送。")
	} else {
		log.Printf("消息发送失败,HTTP 状态码: %d\n", resp.StatusCode)
	}
}

// func SendToLarkSealsuite(data string) {
// 	webhookURL := viper.GetString("alert.lark_sealsuite_bot")
// 	resp, err := http.Post(webhookURL, "application/json", strings.NewReader(data))
// 	if err != nil {
// 		fmt.Println("Error sending message:", err)
// 		return
// 	}
// 	log.Println(resp)
// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		log.Println("消息已成功发送。")
// 	} else {
// 		log.Printf("消息发送失败,HTTP 状态码: %d\n", resp.StatusCode)
// 	}
// }
