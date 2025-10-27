package mem0

import (
	"context"
	"testing"

	"github.com/spf13/viper"
)

// TestMem0ClientBasicFunctionality 测试 Mem0Client 的基本功能
func TestMem0ClientBasicFunctionality(t *testing.T) {
	// 设置测试配置
	viper.Set("memory.mem0.api_key", "m0-1wHmTD2cB0m00zEqOCGOMHHPIcXEZREAhrM9rPbQ")
	viper.Set("memory.mem0.host", "https://api.mem0.ai")

	// 获取客户端实例
	client, err := GetMem0Client()
	if err != nil {
		t.Fatalf("Failed to get Mem0Client instance: %v", err)
	}
	if client == nil {
		t.Fatal("Mem0Client instance is nil")
	}

	// 测试初始化
	err = client.Init()
	if err != nil {
		t.Fatalf("Failed to initialize client: %v", err)
	}

	// 测试用户ID
	testUserID := "shijingbo_news"
	ctx := context.Background()

	context_str, err := client.Search(ctx, testUserID, "总结下我今天的行程", 20, 0)
	if err != nil {
		t.Logf("Search failed (expected with test API key): %v", err)
	} else {
		t.Logf("Retrieved context: %s", context_str)
	}

	// 测试添加消息
	/*testMessage := schema.Message{
		Role:    schema.User,
		Content: "我喜欢吃意大利面",
	}

	err = client.AddMessage(ctx, testUserID, testMessage)
	if err != nil {
		t.Logf("AddMessage failed (expected with test API key): %v", err)
	}

	// 测试获取消息
	messages, err := client.GetMessages(ctx, testUserID, 10)
	if err != nil {
		t.Logf("GetMessages failed (expected with test API key): %v", err)
	} else {
		t.Logf("Retrieved %d messages", len(messages))
	}

	// 测试获取上下文
	context_str, err := client.GetContext(ctx, testUserID, 10)
	if err != nil {
		t.Logf("GetContext failed (expected with test API key): %v", err)
	} else {
		t.Logf("Retrieved context: %s", context_str)
	}

	// 测试批量添加消息
	batchMessages := []schema.Message{
		{Role: schema.User, Content: "我今天很开心"},
		{Role: schema.Assistant, Content: "很高兴听到你今天心情不错！"},
	}

	err = client.AddBatchMessages(ctx, testUserID, batchMessages)
	if err != nil {
		t.Logf("AddBatchMessages failed (expected with test API key): %v", err)
	}

	// 测试重置记忆
	err = client.ResetMemory(ctx, testUserID)
	if err != nil {
		t.Logf("ResetMemory failed (expected with test API key): %v", err)
	}

	// 测试关闭客户端
	err = client.Close()
	if err != nil {
		t.Fatalf("Failed to close client: %v", err)
	}*/

	t.Log("All basic functionality tests completed")
}

/*
// TestMem0ClientSingleton 测试单例模式
func TestMem0ClientSingleton(t *testing.T) {
	// 设置测试配置
	viper.Set("memory.mem0.api_key", "test-api-key")
	viper.Set("memory.mem0.host", "https://api.mem0.ai")

	// 获取两个客户端实例
	client1, err1 := GetMem0Client()
	if err1 != nil {
		t.Fatalf("Failed to get first client instance: %v", err1)
	}
	client2, err2 := GetMem0Client()
	if err2 != nil {
		t.Fatalf("Failed to get second client instance: %v", err2)
	}

	// 验证是同一个实例
	if client1 != client2 {
		t.Error("GetMem0Client should return the same instance (singleton pattern)")
	}

	t.Log("Singleton pattern test passed")
}

// TestMem0ClientConfiguration 测试配置读取
func TestMem0ClientConfiguration(t *testing.T) {
	// 设置测试配置
	testAPIKey := "test-api-key-12345"
	testHost := "https://custom.mem0.ai"
	testOrgName := "test-org"
	testProjectName := "test-project"

	viper.Set("memory.mem0.api_key", testAPIKey)
	viper.Set("memory.mem0.host", testHost)
	viper.Set("memory.mem0.organization_name", testOrgName)
	viper.Set("memory.mem0.project_name", testProjectName)

	// 创建新的客户端实例来测试配置读取
	client, err := newMem0Client()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// 验证配置是否正确读取
	if client.config.APIKey != testAPIKey {
		t.Errorf("Expected API key %s, got %s", testAPIKey, client.config.APIKey)
	}

	if client.config.Host != testHost {
		t.Errorf("Expected host %s, got %s", testHost, client.config.Host)
	}

	if client.config.OrganizationName != testOrgName {
		t.Errorf("Expected organization name %s, got %s", testOrgName, client.config.OrganizationName)
	}

	if client.config.ProjectName != testProjectName {
		t.Errorf("Expected project name %s, got %s", testProjectName, client.config.ProjectName)
	}

	t.Log("Configuration test passed")
}
*/
