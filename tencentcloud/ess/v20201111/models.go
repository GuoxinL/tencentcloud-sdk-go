// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20201111

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Agent struct {
}

type ApproverInfo struct {

	// 参与者类型：
	// 0：企业
	// 1：个人
	// 3：企业静默签署
	// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。
	ApproverType *int64 `json:"ApproverType,omitempty" name:"ApproverType"`

	// 本环节需要操作人的名字
	ApproverName *string `json:"ApproverName,omitempty" name:"ApproverName"`

	// 本环节需要操作人的手机号
	ApproverMobile *string `json:"ApproverMobile,omitempty" name:"ApproverMobile"`

	// 本环节操作人签署控件配置，为企业静默签署时，只允许类型为SIGN_SEAL（印章）和SIGN_DATE（日期）控件，并且传入印章编号。
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 如果是企业,则为企业的名字
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 身份证号
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitempty" name:"ApproverIdCardNumber"`

	// 证件类型 
	// ID_CARD 身份证
	// HONGKONG_AND_MACAO 港澳居民来往内地通行证
	// HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证(格式同居民身份证)
	ApproverIdCardType *string `json:"ApproverIdCardType,omitempty" name:"ApproverIdCardType"`

	// sms--短信，none--不通知
	NotifyType *string `json:"NotifyType,omitempty" name:"NotifyType"`

	// 1--收款人、2--开具人、3--见证人
	ApproverRole *int64 `json:"ApproverRole,omitempty" name:"ApproverRole"`

	// 签署意愿确认渠道,WEIXINAPP:人脸识别
	VerifyChannel []*string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 合同的强制预览时间：3~300s，未指定则按合同页数计算
	PreReadTime *int64 `json:"PreReadTime,omitempty" name:"PreReadTime"`
}

type Caller struct {

	// 应用号
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 主机构ID
	OrganizationId *string `json:"OrganizationId,omitempty" name:"OrganizationId"`

	// 下属机构ID
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`

	// 经办人的用户ID
	OperatorId *string `json:"OperatorId,omitempty" name:"OperatorId"`
}

type CancelFlowRequest struct {
	*tchttp.BaseRequest

	// 操作用户id
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 流程id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 撤销原因
	CancelMessage *string `json:"CancelMessage,omitempty" name:"CancelMessage"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *CancelFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "CancelMessage")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcInfo struct {

	// 被抄送人手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

type Component struct {

	// 如果是 Component 控件类型，则可选类型为：
	// TEXT - 内容文本控件
	// DATE - 内容日期控件
	// SELECT - 勾选框控件
	// 如果是 SignComponent 控件类型，则可选类型为：
	// SIGN_SEAL - 签署印章控件
	// SIGN_DATE - 签署日期控件
	// SIGN_SIGNATURE - 手写签名控件
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// 参数控件宽度，单位pt
	ComponentWidth *float64 `json:"ComponentWidth,omitempty" name:"ComponentWidth"`

	// 参数控件高度，单位pt
	ComponentHeight *float64 `json:"ComponentHeight,omitempty" name:"ComponentHeight"`

	// 参数控件所在页码，取值为：1-N
	ComponentPage *int64 `json:"ComponentPage,omitempty" name:"ComponentPage"`

	// 参数控件X位置，单位pt
	ComponentPosX *float64 `json:"ComponentPosX,omitempty" name:"ComponentPosX"`

	// 参数控件Y位置，单位pt
	ComponentPosY *float64 `json:"ComponentPosY,omitempty" name:"ComponentPosY"`

	// 控件所属文件的序号（模板中的resourceId排列序号，取值为：0-N）
	FileIndex *int64 `json:"FileIndex,omitempty" name:"FileIndex"`

	// GenerateMode==KEYWORD 指定关键字
	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`

	// GenerateMode==FIELD 指定表单域名称
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 是否必选，默认为false
	ComponentRequired *bool `json:"ComponentRequired,omitempty" name:"ComponentRequired"`

	// 扩展参数：
	// ComponentType为SIGN_SIGNATURE类型可以控制签署方式
	// {“ComponentTypeLimit”: [“xxx”]}
	// xxx可以为：
	// HANDWRITE – 手写签名
	// BORDERLESS_ESIGN – 自动生成无边框腾讯体
	// OCR_ESIGN -- AI智能识别手写签名
	// ESIGN -- 个人印章类型
	// 如：{“ComponentTypeLimit”: [“BORDERLESS_ESIGN”]}
	ComponentExtra *string `json:"ComponentExtra,omitempty" name:"ComponentExtra"`

	// 控件关联的签署人ID
	ComponentRecipientId *string `json:"ComponentRecipientId,omitempty" name:"ComponentRecipientId"`

	// 控件所填写的内容
	ComponentValue *string `json:"ComponentValue,omitempty" name:"ComponentValue"`

	// 是否是表单域类型，默认不存在
	IsFormType *bool `json:"IsFormType,omitempty" name:"IsFormType"`

	// NORMAL 正常模式，使用坐标制定签署控件位置
	// FIELD 表单域，需使用ComponentName指定表单域名称
	// KEYWORD 关键字，使用ComponentId指定关键字
	GenerateMode *string `json:"GenerateMode,omitempty" name:"GenerateMode"`

	// 日期控件类型字号
	ComponentDateFontSize *int64 `json:"ComponentDateFontSize,omitempty" name:"ComponentDateFontSize"`

	// 指定关键字时横坐标偏移量
	OffsetX *float64 `json:"OffsetX,omitempty" name:"OffsetX"`

	// 指定关键字时纵坐标偏移量
	OffsetY *float64 `json:"OffsetY,omitempty" name:"OffsetY"`
}

type CreateDocumentRequest struct {
	*tchttp.BaseRequest

	// 无
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 用户上传的模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 文件名列表
	FileNames []*string `json:"FileNames,omitempty" name:"FileNames"`

	// 内容控件信息数组
	FormFields []*FormField `json:"FormFields,omitempty" name:"FormFields"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 客户端Token，保持接口幂等性
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 是否需要生成预览文件 默认不生成；
	// 预览链接有效期300秒；
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`
}

func (r *CreateDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "TemplateId")
	delete(f, "FlowId")
	delete(f, "FileNames")
	delete(f, "FormFields")
	delete(f, "Agent")
	delete(f, "ClientToken")
	delete(f, "NeedPreview")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDocumentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的电子文档ID
		DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

		// 返回合同文件的预览地址 5分钟内有效。仅当NeedPreview为true 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
		PreviewFileUrl *string `json:"PreviewFileUrl,omitempty" name:"PreviewFileUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlowByFilesRequest struct {
	*tchttp.BaseRequest

	// 调用方用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 流程名称
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署pdf文件的资源编号列表
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 签署参与者信息
	Approvers []*ApproverInfo `json:"Approvers,omitempty" name:"Approvers"`

	// 流程描述
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 发送类型：
	// true：无序签
	// false：有序签
	// 注：默认为false（有序签）
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 流程的类型
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 流程的签署截止时间
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 应用号信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 经办人内容控件配置。可选类型为：
	// TEXT - 内容文本控件
	// MULTI_LINE_TEXT - 多行文本控件
	// 注：默认字体大小为 字号12
	Components []*Component `json:"Components,omitempty" name:"Components"`

	// 被抄送人的信息列表
	CcInfos []*CcInfo `json:"CcInfos,omitempty" name:"CcInfos"`

	// 是否需要预览，true：预览模式，false：非预览（默认）
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`
}

func (r *CreateFlowByFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowByFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowName")
	delete(f, "FileIds")
	delete(f, "Approvers")
	delete(f, "FlowDescription")
	delete(f, "Unordered")
	delete(f, "FlowType")
	delete(f, "Deadline")
	delete(f, "Agent")
	delete(f, "Components")
	delete(f, "CcInfos")
	delete(f, "NeedPreview")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowByFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlowByFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流程编号
		FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

		// 合同预览链接
	// 注意：此字段可能返回 null，表示取不到有效值。
		PreviewUrl *string `json:"PreviewUrl,omitempty" name:"PreviewUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFlowByFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowByFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlowRequest struct {
	*tchttp.BaseRequest

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 流程的名字, 长度不能超过200，中文字母数字下划线
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 参与者信息
	Approvers []*FlowCreateApprover `json:"Approvers,omitempty" name:"Approvers"`

	// 流程的描述, 长度不能超过1000
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 发送类型(true为无序签,false为顺序签)
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 流程的种类(如销售合同/入职合同等)
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 过期时间戳,如果是0则为不过期
	DeadLine *int64 `json:"DeadLine,omitempty" name:"DeadLine"`

	// 执行结果的回调URL(需要以http://或者https://)开头
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 用户自定义字段(需进行base64 encode),回调的时候会进行透传, 长度需要小于20480
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 客户端Token，保持接口幂等性
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
}

func (r *CreateFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowName")
	delete(f, "Approvers")
	delete(f, "FlowDescription")
	delete(f, "Unordered")
	delete(f, "FlowType")
	delete(f, "DeadLine")
	delete(f, "CallbackUrl")
	delete(f, "UserData")
	delete(f, "Agent")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流程编号
		FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSchemeUrlRequest struct {
	*tchttp.BaseRequest

	// 调用方用户信息，参考通用结构
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 调用方渠道信息，参考通用结构
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 跳转页面 1: 小程序合同详情 2: 小程序合同列表页 0: 不传, 默认主页
	PathType *uint64 `json:"PathType,omitempty" name:"PathType"`

	// 合同详情 id (PathType=1时必传)
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 企业名称
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 链接类型
	// HTTP：跳转电子签小程序的http_url，
	// APP：第三方APP或小程序跳转电子签小程序的path。
	// 默认为HTTP类型
	EndPoint *string `json:"EndPoint,omitempty" name:"EndPoint"`

	// 是否自动回跳 true：是， false：否。该参数只针对"APP" 类型的签署链接有效
	AutoJumpBack *bool `json:"AutoJumpBack,omitempty" name:"AutoJumpBack"`
}

func (r *CreateSchemeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchemeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Agent")
	delete(f, "Name")
	delete(f, "Mobile")
	delete(f, "PathType")
	delete(f, "FlowId")
	delete(f, "OrganizationName")
	delete(f, "EndPoint")
	delete(f, "AutoJumpBack")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchemeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSchemeUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 小程序链接地址
		SchemeUrl *string `json:"SchemeUrl,omitempty" name:"SchemeUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSchemeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchemeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileUrlsRequest struct {
	*tchttp.BaseRequest

	// 文件对应的业务类型，目前支持：
	// - 模板 "TEMPLATE"
	// - 文档 "DOCUMENT"
	// - 印章  “SEAL”
	// - 流程 "FLOW"
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 操作者信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 业务编号的数组，如模板编号、文档编号、印章编号
	BusinessIds []*string `json:"BusinessIds,omitempty" name:"BusinessIds"`

	// 文件类型，"JPG", "PDF","ZIP"等
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 下载后的文件命名，只有fileType为zip的时候生效
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 指定资源起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 指定资源数量，查询全部资源则传入-1
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 下载url过期时间，0: 按默认值5分钟，允许范围：1s~24*60*60s(1天)
	UrlTtl *int64 `json:"UrlTtl,omitempty" name:"UrlTtl"`

	// 流程校验发送邮件权限
	CcToken *string `json:"CcToken,omitempty" name:"CcToken"`

	// 场景
	Scene *string `json:"Scene,omitempty" name:"Scene"`
}

func (r *DescribeFileUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUrlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessType")
	delete(f, "Operator")
	delete(f, "BusinessIds")
	delete(f, "FileType")
	delete(f, "FileName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Agent")
	delete(f, "UrlTtl")
	delete(f, "CcToken")
	delete(f, "Scene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFileUrlsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// URL信息
		FileUrls []*FileUrl `json:"FileUrls,omitempty" name:"FileUrls"`

		// URL数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUrlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowBriefsRequest struct {
	*tchttp.BaseRequest

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 需要查询的流程ID列表
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 代理商信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *DescribeFlowBriefsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowBriefsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowBriefsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowBriefsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流程列表
		FlowBriefs []*FlowBrief `json:"FlowBriefs,omitempty" name:"FlowBriefs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowBriefsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowBriefsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowTemplatesRequest struct {
	*tchttp.BaseRequest

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 查询偏移位置，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询个数，默认20，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索条件，具体参考Filter结构体。本接口取值：template-id：按照【 **模板唯一标识** 】进行过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 暂未开放
	GenerateSource *uint64 `json:"GenerateSource,omitempty" name:"GenerateSource"`

	// 查询内容：0-模板列表及详情（默认），1-仅模板列表
	ContentType *int64 `json:"ContentType,omitempty" name:"ContentType"`
}

func (r *DescribeFlowTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Agent")
	delete(f, "GenerateSource")
	delete(f, "ContentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板详情列表
		Templates []*TemplateInfo `json:"Templates,omitempty" name:"Templates"`

		// 查询到的总个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeThirdPartyAuthCodeRequest struct {
	*tchttp.BaseRequest

	// AuthCode 值
	AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`
}

func (r *DescribeThirdPartyAuthCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThirdPartyAuthCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuthCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeThirdPartyAuthCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeThirdPartyAuthCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户是否实名，VERIFIED 为实名，UNVERIFIED 未实名
		VerifyStatus *string `json:"VerifyStatus,omitempty" name:"VerifyStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeThirdPartyAuthCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThirdPartyAuthCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileInfo struct {

	// 文件Id
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件大小，单位为Byte
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件上传时间，10位时间戳（精确到秒）
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`
}

type FileUrl struct {

	// 下载文件的URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// 下载文件的附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Option *string `json:"Option,omitempty" name:"Option"`
}

type Filter struct {

	// 查询过滤条件的Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 查询过滤条件的Value列表
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FlowBrief struct {

	// 流程的编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 流程的名称
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 流程的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 流程的类型
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 流程状态
	// - `1` 未签署
	// - `2`  部分签署
	// - `3`  已退回
	// - `4`  完成签署
	// - `5`  已过期
	// - `6`  已取消
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowStatus *int64 `json:"FlowStatus,omitempty" name:"FlowStatus"`

	// 流程创建的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 拒签或者取消的原因描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMessage *string `json:"FlowMessage,omitempty" name:"FlowMessage"`
}

type FlowCreateApprover struct {

	// 参与者类型：
	// 0：企业
	// 1：个人
	// 3：企业静默签署
	// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。
	ApproverType *int64 `json:"ApproverType,omitempty" name:"ApproverType"`

	// 如果签署方为企业，需要填入企业全称
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 是否需要签署
	// - `false`: 不需要签署
	// -  `true`:  需要签署
	Required *bool `json:"Required,omitempty" name:"Required"`

	// 签署方经办人姓名
	ApproverName *string `json:"ApproverName,omitempty" name:"ApproverName"`

	// 签署方经办人手机号码
	ApproverMobile *string `json:"ApproverMobile,omitempty" name:"ApproverMobile"`

	// 签署方经办人证件号码
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitempty" name:"ApproverIdCardNumber"`

	// 签署方经办人证件类型ID_CARD 身份证
	// HONGKONG_AND_MACAO 港澳居民来往内地通行证
	// HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证(格式同居民身份证)
	ApproverIdCardType *string `json:"ApproverIdCardType,omitempty" name:"ApproverIdCardType"`

	// 签署方经办人在模板中的角色ID
	RecipientId *string `json:"RecipientId,omitempty" name:"RecipientId"`

	// 签署方经办人的用户ID,和签署方经办人姓名+手机号+证件必须有一个
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 签署前置条件：是否需要阅读全文，默认为不需要
	IsFullText *bool `json:"IsFullText,omitempty" name:"IsFullText"`

	// 签署前置条件：阅读时长限制，默认为不需要
	PreReadTime *uint64 `json:"PreReadTime,omitempty" name:"PreReadTime"`

	// 是否发送短信，sms--短信通知，none--不通知，默认为sms
	NotifyType *string `json:"NotifyType,omitempty" name:"NotifyType"`

	// 签署意愿确认渠道,WEIXINAPP:人脸识别
	VerifyChannel []*string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`
}

type FormField struct {

	// 控件填充value
	ComponentValue *string `json:"ComponentValue,omitempty" name:"ComponentValue"`

	// 控件id
	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`

	// 控件名字
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

type Recipient struct {

	// 签署参与者ID
	RecipientId *string `json:"RecipientId,omitempty" name:"RecipientId"`

	// 参与者类型（ENTERPRISE/INDIVIDUAL）
	RecipientType *string `json:"RecipientType,omitempty" name:"RecipientType"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 是否需要验证，默认为false
	RequireValidation *bool `json:"RequireValidation,omitempty" name:"RequireValidation"`

	// 是否需要签署，默认为true
	RequireSign *bool `json:"RequireSign,omitempty" name:"RequireSign"`

	// 添加序列
	RoutingOrder *int64 `json:"RoutingOrder,omitempty" name:"RoutingOrder"`

	// 是否需要发送，默认为true
	RequireDelivery *bool `json:"RequireDelivery,omitempty" name:"RequireDelivery"`

	// 邮箱地址
	Email *string `json:"Email,omitempty" name:"Email"`

	// 电话号码
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 关联的用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 发送方式（EMAIL/MOBILE）
	DeliveryMethod *string `json:"DeliveryMethod,omitempty" name:"DeliveryMethod"`

	// 附属信息
	RecipientExtra *string `json:"RecipientExtra,omitempty" name:"RecipientExtra"`
}

type StartFlowRequest struct {
	*tchttp.BaseRequest

	// 用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 客户端Token，保持接口幂等性
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`
}

func (r *StartFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "Agent")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回描述，START-发起成功， REVIEW-提交审核成功，EXECUTING-已提交发起任务
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateInfo struct {

	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名字
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模板描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 模板关联的资源IDs
	DocumentResourceIds []*string `json:"DocumentResourceIds,omitempty" name:"DocumentResourceIds"`

	// 返回的文件信息结构
	FileInfos []*FileInfo `json:"FileInfos,omitempty" name:"FileInfos"`

	// 附件关联的资源ID是
	AttachmentResourceIds []*string `json:"AttachmentResourceIds,omitempty" name:"AttachmentResourceIds"`

	// 签署顺序
	SignOrder []*int64 `json:"SignOrder,omitempty" name:"SignOrder"`

	// 签署参与者的信息
	Recipients []*Recipient `json:"Recipients,omitempty" name:"Recipients"`

	// 模板信息结构
	Components []*Component `json:"Components,omitempty" name:"Components"`

	// 签署区模板信息结构
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 模板状态(-1:不可用；0:草稿态；1:正式态)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 模板的创建人
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 模板创建的时间戳（精确到秒）
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 发起人角色信息
	Promoter *Recipient `json:"Promoter,omitempty" name:"Promoter"`
}

type UploadFile struct {

	// Base64编码后的文件内容
	FileBody *string `json:"FileBody,omitempty" name:"FileBody"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type UploadFilesRequest struct {
	*tchttp.BaseRequest

	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 文件对应业务类型，用于区分文件存储路径：
	// 1. TEMPLATE - 模板； 文件类型：.pdf/.html
	// 2. DOCUMENT - 签署过程及签署后的合同文档 文件类型：.pdf/.html
	// 3. FLOW - 签署过程 文件类型：.pdf/.html
	// 4. SEAL - 印章； 文件类型：.jpg/.jpeg/.png
	// 5. BUSINESSLICENSE - 营业执照 文件类型：.jpg/.jpeg/.png
	// 6. IDCARD - 身份证 文件类型：.jpg/.jpeg/.png
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 上传文件内容数组，最多支持20个文件
	FileInfos []*UploadFile `json:"FileInfos,omitempty" name:"FileInfos"`

	// 上传文件链接数组，最多支持20个URL
	FileUrls *string `json:"FileUrls,omitempty" name:"FileUrls"`

	// 是否将pdf灰色矩阵置白
	// true--是，处理置白
	// false--否，不处理
	CoverRect *bool `json:"CoverRect,omitempty" name:"CoverRect"`

	// 特殊文件类型需要指定文件类型：
	// HTML-- .html文件
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 用户自定义ID数组，与上传文件一一对应
	CustomIds []*string `json:"CustomIds,omitempty" name:"CustomIds"`
}

func (r *UploadFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "BusinessType")
	delete(f, "FileInfos")
	delete(f, "FileUrls")
	delete(f, "CoverRect")
	delete(f, "FileType")
	delete(f, "CustomIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UploadFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件id数组
		FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

		// 上传成功文件数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {

	// 用户在平台的编号
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户的来源渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 用户在渠道的编号
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 用户真实IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 用户代理IP
	ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
}
