package chat

import (
	"SoulChat/config"
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/model/openai"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

var App compose.Runnable[map[string]any, *schema.Message]
var ChatChain *compose.Chain[map[string]any, *schema.Message]
var Ctx context.Context

func InitChain() {
	Ctx = context.Background()
	_, err := openai.NewChatModel(Ctx, &openai.ChatModelConfig{
		APIKey:      config.Conf.ModelInfo.ApiKey,
		BaseURL:     config.Conf.ModelInfo.BaseUrl,
		Model:       config.Conf.ModelInfo.ModelName,
		Temperature: &config.Conf.ModelInfo.Temperature,
	})
	if err != nil {
		log.Println("模型初始化失败，err：", err)
		return
	}

	// 创建prompt模板
	_ = NewCompanionTemplate()

}
