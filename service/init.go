package service

import "github.com/baidubce/bce-qianfan-sdk/go/qianfan"

var chatClient *qianfan.ChatCompletion

func QianfanInit() {
	// 使用安全认证AK/SK鉴权，替换下列示例中参数，安全认证Access Key替换your_iam_ak，Secret Key替换your_iam_sk
	qianfan.GetConfig().AccessKey = "40214bab73c24ffc85ee156d800bfa40"
	qianfan.GetConfig().SecretKey = "28b97973987f46b989e11b7c29eca438"

	// 调用对话Chat，可以通过 WithModel 指定模型，例如指定ERNIE-3.5-8K，参数对应ERNIE-Bot
	chatClient = qianfan.NewChatCompletion(
		qianfan.WithModel("ERNIE-Speed-128K"),
	)
}
