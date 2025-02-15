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

package v20210519

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateCredentialRequest struct {
	*tchttp.BaseRequest

	// 参数集合，详见示例
	FunctionArg *FunctionArg `json:"FunctionArg,omitempty" name:"FunctionArg"`

	// 参数集合，详见示例
	TransactionArg *TransactionArg `json:"TransactionArg,omitempty" name:"TransactionArg"`

	// 版本
	VersionCredential *string `json:"VersionCredential,omitempty" name:"VersionCredential"`

	// 是否未签名
	UnSigned *bool `json:"UnSigned,omitempty" name:"UnSigned"`
}

func (r *CreateCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	delete(f, "TransactionArg")
	delete(f, "VersionCredential")
	delete(f, "UnSigned")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCredentialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Credential的具体信息
		CredentialData *string `json:"CredentialData,omitempty" name:"CredentialData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSelectiveCredentialRequest struct {
	*tchttp.BaseRequest

	// 参数集合
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitempty" name:"FunctionArg"`

	// 批露策略id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *CreateSelectiveCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSelectiveCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSelectiveCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSelectiveCredentialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 凭证字符串
		CredentialData *string `json:"CredentialData,omitempty" name:"CredentialData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSelectiveCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSelectiveCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTDidByPublicKeyRequest struct {
	*tchttp.BaseRequest

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 身份公钥
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// 加密公钥
	EncryptPubKey *string `json:"EncryptPubKey,omitempty" name:"EncryptPubKey"`
}

func (r *CreateTDidByPublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPublicKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "PublicKey")
	delete(f, "EncryptPubKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByPublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTDidByPublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// did具体信息
		Did *string `json:"Did,omitempty" name:"Did"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTDidByPublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTDidRequest struct {
	*tchttp.BaseRequest

	// 群组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 网络ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署机构为1，否则为0
	Relegation *uint64 `json:"Relegation,omitempty" name:"Relegation"`
}

func (r *CreateTDidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ClusterId")
	delete(f, "Relegation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTDidResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TDID
		Did *string `json:"Did,omitempty" name:"Did"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTDidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CredentialStatus struct {

	// 凭证唯一id
	CredentialId *string `json:"CredentialId,omitempty" name:"CredentialId"`

	// 凭证状态（0：吊销；1：有效）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 凭证颁发者Did
	Issuer *string `json:"Issuer,omitempty" name:"Issuer"`

	// 凭证摘要
	// 注意：此字段可能返回 null，表示取不到有效值。
	Digest *string `json:"Digest,omitempty" name:"Digest"`

	// 凭证签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Signature *string `json:"Signature,omitempty" name:"Signature"`

	// 更新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeStamp *uint64 `json:"TimeStamp,omitempty" name:"TimeStamp"`
}

type FunctionArg struct {

	// CPT ID
	CptId *uint64 `json:"CptId,omitempty" name:"CptId"`

	// 签发者 did
	Issuer *string `json:"Issuer,omitempty" name:"Issuer"`

	// 过期时间
	ExpirationDate *string `json:"ExpirationDate,omitempty" name:"ExpirationDate"`

	// 声明
	ClaimJson *string `json:"ClaimJson,omitempty" name:"ClaimJson"`
}

type GetAuthorityIssuerRequest struct {
	*tchttp.BaseRequest

	// tdid
	Did *string `json:"Did,omitempty" name:"Did"`
}

func (r *GetAuthorityIssuerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthorityIssuerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAuthorityIssuerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetAuthorityIssuerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 区块链网络id
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 区块链群组id
		GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

		// 权威机构did
		Did *string `json:"Did,omitempty" name:"Did"`

		// 机构备注信息
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 注册时间
		RegisterTime *string `json:"RegisterTime,omitempty" name:"RegisterTime"`

		// 认证时间
		RecognizeTime *string `json:"RecognizeTime,omitempty" name:"RecognizeTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAuthorityIssuerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthorityIssuerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDidDocumentRequest struct {
	*tchttp.BaseRequest

	// tdid
	Did *string `json:"Did,omitempty" name:"Did"`
}

func (r *GetDidDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDidDocumentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// DID文档
		Document *string `json:"Document,omitempty" name:"Document"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDidDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Proof struct {

	// 创建时间
	Created *int64 `json:"Created,omitempty" name:"Created"`

	// 创建着did
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// salt值
	SaltJson *string `json:"SaltJson,omitempty" name:"SaltJson"`

	// 签名
	SignatureValue *string `json:"SignatureValue,omitempty" name:"SignatureValue"`

	// type类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

type SetCredentialStatusRequest struct {
	*tchttp.BaseRequest

	// 凭证状态
	CredentialStatus *CredentialStatus `json:"CredentialStatus,omitempty" name:"CredentialStatus"`
}

func (r *SetCredentialStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCredentialStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetCredentialStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetCredentialStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetCredentialStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCredentialStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransactionArg struct {

	// 凭证did
	InvokerTDid *string `json:"InvokerTDid,omitempty" name:"InvokerTDid"`
}

type VerifyCredentialRequest struct {
	*tchttp.BaseRequest

	// 参数集合
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitempty" name:"FunctionArg"`
}

func (r *VerifyCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VerifyCredentialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否验证成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 验证返回码
		VerifyCode *uint64 `json:"VerifyCode,omitempty" name:"VerifyCode"`

		// 验证消息
		VerifyMessage *string `json:"VerifyMessage,omitempty" name:"VerifyMessage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyFunctionArg struct {

	// CPT ID
	CptId *uint64 `json:"CptId,omitempty" name:"CptId"`

	// issuer did
	Issuer *string `json:"Issuer,omitempty" name:"Issuer"`

	// 过期时间
	ExpirationDate *int64 `json:"ExpirationDate,omitempty" name:"ExpirationDate"`

	// 声明
	ClaimJson *string `json:"ClaimJson,omitempty" name:"ClaimJson"`

	// 颁发时间
	IssuanceDate *int64 `json:"IssuanceDate,omitempty" name:"IssuanceDate"`

	// context值
	Context *string `json:"Context,omitempty" name:"Context"`

	// id值
	Id *string `json:"Id,omitempty" name:"Id"`

	// 签名值
	Proof *Proof `json:"Proof,omitempty" name:"Proof"`

	// type值
	Type []*string `json:"Type,omitempty" name:"Type"`
}
