/*
https://github.com/tencentcloud/tencentcloud-sdk-go
https://github.com/TencentCloud/tencentcloud-sdk-go/tree/master/examples/sms/v20190711
https://console.cloud.tencent.com/smsv2
https://cloud.tencent.com/document/product/382/3773 

*/


package main


import (
    "encoding/json"
    "fmt"
	"os"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
    sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)


// *v20190711.Client
func conn()(*sms.Client){
    credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
    )
    cpf := profile.NewClientProfile()
    cpf.HttpProfile.ReqMethod = "POST"
    cpf.HttpProfile.ReqTimeout = 5
    /* SDK 会自动指定域名，通常无需指定域名，但访问金融区的服务时必须手动指定域名
     * 例如 SMS 的上海金融区域名为 sms.ap-shanghai-fsi.tencentcloudapi.com */
    cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
    cpf.SignMethod = "HmacSHA1"
	client, _ :=sms.NewClient(credential, "ap-guangzhou", cpf)
	return  client

}



func send(){
	client:=conn()
    request := sms.NewSendSmsRequest()

    /* 短信应用 ID: 在 [短信控制台] 添加应用后生成的实际 SDKAppID，例如1400006666 */
	// https://console.cloud.tencent.com/smsv2/app-manage
	//id:="1001484497"
	// "1400787878"
    SmsSdkAppid:="1400391772"
	//k:="7b53ada9738c0837005dd134078f88d5"
    request.SmsSdkAppid = common.StringPtr(SmsSdkAppid)
    /* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名，可登录 [短信控制台] 查看签名信息 */
	// https://console.cloud.tencent.com/smsv2/csms-sign/create
    request.Sign = common.StringPtr("ccc")
    request.SenderId = common.StringPtr("xxx")
    request.SessionContext = common.StringPtr("xxx")
    request.ExtendCode = common.StringPtr("0")
    request.TemplateParamSet = common.StringPtrs([]string{"0"})
    /* 模板 ID: 必须填写已审核通过的模板 ID，可登录 [短信控制台] 查看模板 ID */
    request.TemplateID = common.StringPtr("449739")
    /* 下发手机号码，采用 e.164 标准，+[国家或地区码][手机号]
     * 例如+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
	phone:=[]string{"+8618576690127", }
    request.PhoneNumberSet = common.StringPtrs(phone)
    response, err := client.SendSms(request)
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    if err != nil {
        panic(err)
    }
    b, _ := json.Marshal(response.Response)
    fmt.Printf("%s", b)
}


func add_temp(){
	client:=conn()
    request := sms.NewAddSmsTemplateRequest()
    /* 基本类型的设置:
     * 短信控制台：https://console.cloud.tencent.com/smsv2
     * sms helper：https://cloud.tencent.com/document/product/382/3773 */
    request.TemplateName = common.StringPtr("腾讯云")
    request.TemplateContent = common.StringPtr("{1}为您的登录验证码，请于{2}分钟内填写，如非本人操作，请忽略本短信。")
    request.SmsType = common.Uint64Ptr(0)
    /* 是否国际/港澳台短信：
       0：表示国内短信
       1：表示国际/港澳台短信 */
    request.International = common.Uint64Ptr(0)
    request.Remark = common.StringPtr("xxx")

    response, err := client.AddSmsTemplate(request)
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    if err != nil {
        panic(err)
    }
    b, _ := json.Marshal(response.Response)
    fmt.Printf("%s", b)
}

func pull_s(){
	client:=conn()
    request := sms.NewPullSmsSendStatusRequest()

    /* 基本类型的设置:
     * SDK 采用的是指针风格指定参数，即使对于基本类型也需要用指针来对参数赋值。
     * SDK 提供对基本类型的指针引用封装函数
     * 帮助链接：
     * 短信控制台：https://console.cloud.tencent.com/smsv2
     * sms helper：https://cloud.tencent.com/document/product/382/3773 */

    /* 短信应用 ID: 在 [短信控制台] 添加应用后生成的实际 SDKAppID，例如1400006666 */
    request.SmsSdkAppid = common.StringPtr("1400787878")
    /* 拉取最大条数，最多100条 */
    request.Limit = common.Uint64Ptr(10)

    // 通过 client 对象调用想要访问的接口，需要传入请求对象
    response, err := client.PullSmsSendStatus(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非 SDK 异常，直接失败。实际代码中可以加入其他的处理
    if err != nil {
        panic(err)
    }
    b, _ := json.Marshal(response.Response)
    // 打印返回的 JSON 字符串
    fmt.Printf("%s", b)
}


func get_h(){
	client:=conn()
    request := sms.NewSendStatusStatisticsRequest()

    /* 基本类型的设置:
     * SDK 采用的是指针风格指定参数，即使对于基本类型也需要用指针来对参数赋值。
     * SDK 提供对基本类型的指针引用封装函数
     * 帮助链接：
     * 短信控制台：https://console.cloud.tencent.com/smsv2
     * sms helper：https://cloud.tencent.com/document/product/382/3773 */

    /* 短信应用 ID: 在 [短信控制台] 添加应用后生成的实际 SDKAppID，例如1400006666 */
    request.SmsSdkAppid = common.StringPtr("1400787878")
    /* 拉取最大条数，最多100条 */
    request.Limit = common.Uint64Ptr(0)
    /* 偏移量，目前固定设置为0 */
    request.Offset = common.Uint64Ptr(0)
    /* 开始时间，yyyymmddhh 需要拉取的起始时间，精确到小时 */
    request.StartDateTime = common.Uint64Ptr(2019122400)
    /* 结束时间，yyyymmddhh 需要拉取的截止时间，精确到小时
     * 注：EndDataTime 必须大于 StartDateTime */
    request.EndDataTime = common.Uint64Ptr(2019122523)

    // 通过 client 对象调用想要访问的接口，需要传入请求对象
    response, err := client.SendStatusStatistics(request)
    // 处理异常
    if _, ok := err.(*errors.TencentCloudSDKError); ok {
        fmt.Printf("An API error has returned: %s", err)
        return
    }
    // 非 SDK 异常，直接失败。实际代码中可以加入其他的处理
    if err != nil {
        panic(err)
    }
    b, _ := json.Marshal(response.Response)
    // 打印返回的 JSON 字符串
    fmt.Printf("%s", b)
}




func main() {
	send()
}


