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

package v20190722

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-22"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateCloudRecordingRequest() (request *CreateCloudRecordingRequest) {
    request = &CreateCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "CreateCloudRecording")
    
    
    return
}

func NewCreateCloudRecordingResponse() (response *CreateCloudRecordingResponse) {
    response = &CreateCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudRecording
// ###接口说明：
//
// 启动云端录制功能，完成房间内的音视频录制，并上传到指定的云存储。您可以通过此 API 接口把TRTC 房间中的每一路音视频流做单独的录制有或者多路视频画面混流一路。
//
// 
//
// ###您可以通过此接口实现如下目标：
//
// * 指定订阅流参数（RecordParams）来指定需要录制的主播的黑名单或者白名单。
//
// * 指定第三方存储的参数（StorageParams）来指定上传到您希望的云存储
//
// * 指定混流模式下的音视频转码详细参数（MixTranscodeParams），包括视频分辨率、视频码率、视频帧率、以及声音质量等
//
// * 指定混流模式各路画面的位置和布局或者也可以指定自动模板的方式来配置。
//
// 
//
// ###关键名词：
//
// * 单流录制：分别录制房间的订阅UserId的音频和视频。录制服务会实时将录制文件（M3U8/TS）上传至云存储。
//
// * 混流录制：将房间内订阅UserId的音视频混录成一个音视频文件，并将录制文件（M3U8/TS）上传至云存储。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_CLOUDSTORAGE = "MissingParameter.CloudStorage"
//  MISSINGPARAMETER_RECORDMODE = "MissingParameter.RecordMode"
//  MISSINGPARAMETER_RECORDPARAMS = "MissingParameter.RecordParams"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_STORAGEPARAMS = "MissingParameter.StorageParams"
//  MISSINGPARAMETER_STREAMTYPE = "MissingParameter.StreamType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
func (c *Client) CreateCloudRecording(request *CreateCloudRecordingRequest) (response *CreateCloudRecordingResponse, err error) {
    return c.CreateCloudRecordingWithContext(context.Background(), request)
}

// CreateCloudRecording
// ###接口说明：
//
// 启动云端录制功能，完成房间内的音视频录制，并上传到指定的云存储。您可以通过此 API 接口把TRTC 房间中的每一路音视频流做单独的录制有或者多路视频画面混流一路。
//
// 
//
// ###您可以通过此接口实现如下目标：
//
// * 指定订阅流参数（RecordParams）来指定需要录制的主播的黑名单或者白名单。
//
// * 指定第三方存储的参数（StorageParams）来指定上传到您希望的云存储
//
// * 指定混流模式下的音视频转码详细参数（MixTranscodeParams），包括视频分辨率、视频码率、视频帧率、以及声音质量等
//
// * 指定混流模式各路画面的位置和布局或者也可以指定自动模板的方式来配置。
//
// 
//
// ###关键名词：
//
// * 单流录制：分别录制房间的订阅UserId的音频和视频。录制服务会实时将录制文件（M3U8/TS）上传至云存储。
//
// * 混流录制：将房间内订阅UserId的音视频混录成一个音视频文件，并将录制文件（M3U8/TS）上传至云存储。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_CLOUDSTORAGE = "MissingParameter.CloudStorage"
//  MISSINGPARAMETER_RECORDMODE = "MissingParameter.RecordMode"
//  MISSINGPARAMETER_RECORDPARAMS = "MissingParameter.RecordParams"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_STORAGEPARAMS = "MissingParameter.StorageParams"
//  MISSINGPARAMETER_STREAMTYPE = "MissingParameter.StreamType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
func (c *Client) CreateCloudRecordingWithContext(ctx context.Context, request *CreateCloudRecordingRequest) (response *CreateCloudRecordingResponse, err error) {
    if request == nil {
        request = NewCreateCloudRecordingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePictureRequest() (request *CreatePictureRequest) {
    request = &CreatePictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "CreatePicture")
    
    
    return
}

func NewCreatePictureResponse() (response *CreatePictureResponse) {
    response = &CreatePictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁新增自定义背景图或水印，可通过此接口上传新的图片素材。无需频繁新增图片的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CHECKCONTENTFAILED = "InvalidParameter.CheckContentFailed"
//  INVALIDPARAMETER_CHECKSUFFIXFAILED = "InvalidParameter.CheckSuffixFailed"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePicture(request *CreatePictureRequest) (response *CreatePictureResponse, err error) {
    return c.CreatePictureWithContext(context.Background(), request)
}

// CreatePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁新增自定义背景图或水印，可通过此接口上传新的图片素材。无需频繁新增图片的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CHECKCONTENTFAILED = "InvalidParameter.CheckContentFailed"
//  INVALIDPARAMETER_CHECKSUFFIXFAILED = "InvalidParameter.CheckSuffixFailed"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePictureWithContext(ctx context.Context, request *CreatePictureRequest) (response *CreatePictureResponse, err error) {
    if request == nil {
        request = NewCreatePictureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePictureResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTroubleInfoRequest() (request *CreateTroubleInfoRequest) {
    request = &CreateTroubleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "CreateTroubleInfo")
    
    
    return
}

func NewCreateTroubleInfoResponse() (response *CreateTroubleInfoResponse) {
    response = &CreateTroubleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTroubleInfo
// 创建异常信息
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDFAIL = "InternalError.BackendFail"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateTroubleInfo(request *CreateTroubleInfoRequest) (response *CreateTroubleInfoResponse, err error) {
    return c.CreateTroubleInfoWithContext(context.Background(), request)
}

// CreateTroubleInfo
// 创建异常信息
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDFAIL = "InternalError.BackendFail"
//  INTERNALERROR_BACKENDTIMEOUT = "InternalError.BackendTimeOut"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateTroubleInfoWithContext(ctx context.Context, request *CreateTroubleInfoRequest) (response *CreateTroubleInfoResponse, err error) {
    if request == nil {
        request = NewCreateTroubleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTroubleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTroubleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudRecordingRequest() (request *DeleteCloudRecordingRequest) {
    request = &DeleteCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DeleteCloudRecording")
    
    
    return
}

func NewDeleteCloudRecordingResponse() (response *DeleteCloudRecordingResponse) {
    response = &DeleteCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCloudRecording
// 成功开启录制后，可以使用此接口来停止录制任务。仅在录制任务进行时有效，录制退出后更新将会返回错误。停止录制成功后不代表文件全部传输完成，如果未完成后台将会继续上传文件，成功后通过事件回调通知客户文件全部传输完成状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudRecording(request *DeleteCloudRecordingRequest) (response *DeleteCloudRecordingResponse, err error) {
    return c.DeleteCloudRecordingWithContext(context.Background(), request)
}

// DeleteCloudRecording
// 成功开启录制后，可以使用此接口来停止录制任务。仅在录制任务进行时有效，录制退出后更新将会返回错误。停止录制成功后不代表文件全部传输完成，如果未完成后台将会继续上传文件，成功后通过事件回调通知客户文件全部传输完成状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudRecordingWithContext(ctx context.Context, request *DeleteCloudRecordingRequest) (response *DeleteCloudRecordingResponse, err error) {
    if request == nil {
        request = NewDeleteCloudRecordingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePictureRequest() (request *DeletePictureRequest) {
    request = &DeletePictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DeletePicture")
    
    
    return
}

func NewDeletePictureResponse() (response *DeletePictureResponse) {
    response = &DeletePictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁删除自定义背景图或水印，可通过此接口删除已上传的图片。无需频繁删除图片的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PICTURENOTFOUND = "InvalidParameter.PictureNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeletePicture(request *DeletePictureRequest) (response *DeletePictureResponse, err error) {
    return c.DeletePictureWithContext(context.Background(), request)
}

// DeletePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁删除自定义背景图或水印，可通过此接口删除已上传的图片。无需频繁删除图片的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PICTURENOTFOUND = "InvalidParameter.PictureNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeletePictureWithContext(ctx context.Context, request *DeletePictureRequest) (response *DeletePictureResponse, err error) {
    if request == nil {
        request = NewDeletePictureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePictureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAbnormalEventRequest() (request *DescribeAbnormalEventRequest) {
    request = &DescribeAbnormalEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeAbnormalEvent")
    
    
    return
}

func NewDescribeAbnormalEventResponse() (response *DescribeAbnormalEventResponse) {
    response = &DescribeAbnormalEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAbnormalEvent
// 查询SDKAppID下用户的异常体验事件，返回异常体验ID与可能产生异常体验的原因。可查询15天内数据，查询起止时间不超过1个小时。支持跨天查询。异常体验ID映射见：https://cloud.tencent.com/document/product/647/44916
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEEXPIRE = "InvalidParameter.StartTimeExpire"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeAbnormalEvent(request *DescribeAbnormalEventRequest) (response *DescribeAbnormalEventResponse, err error) {
    return c.DescribeAbnormalEventWithContext(context.Background(), request)
}

// DescribeAbnormalEvent
// 查询SDKAppID下用户的异常体验事件，返回异常体验ID与可能产生异常体验的原因。可查询15天内数据，查询起止时间不超过1个小时。支持跨天查询。异常体验ID映射见：https://cloud.tencent.com/document/product/647/44916
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEEXPIRE = "InvalidParameter.StartTimeExpire"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeAbnormalEventWithContext(ctx context.Context, request *DescribeAbnormalEventRequest) (response *DescribeAbnormalEventResponse, err error) {
    if request == nil {
        request = NewDescribeAbnormalEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAbnormalEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAbnormalEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallDetailRequest() (request *DescribeCallDetailRequest) {
    request = &DescribeCallDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCallDetail")
    
    
    return
}

func NewDescribeCallDetailResponse() (response *DescribeCallDetailResponse) {
    response = &DescribeCallDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCallDetail
// 查询指定时间内的用户列表及用户通话质量数据，可查询14天内数据。DataType 不为null，查询起止时间不超过1个小时，查询用户不超过6个，支持跨天查询。DataType，UserIds为null时，查询起止时间不超过4个小时， 默认查询6个用户，同时支持每页查询100以内用户个数（PageSize不超过100）。接口用于查询质量问题，不推荐作为计费使用。
//
// **注意**：该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeCallDetail(request *DescribeCallDetailRequest) (response *DescribeCallDetailResponse, err error) {
    return c.DescribeCallDetailWithContext(context.Background(), request)
}

// DescribeCallDetail
// 查询指定时间内的用户列表及用户通话质量数据，可查询14天内数据。DataType 不为null，查询起止时间不超过1个小时，查询用户不超过6个，支持跨天查询。DataType，UserIds为null时，查询起止时间不超过4个小时， 默认查询6个用户，同时支持每页查询100以内用户个数（PageSize不超过100）。接口用于查询质量问题，不推荐作为计费使用。
//
// **注意**：该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeCallDetailWithContext(ctx context.Context, request *DescribeCallDetailRequest) (response *DescribeCallDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCallDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRecordingRequest() (request *DescribeCloudRecordingRequest) {
    request = &DescribeCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCloudRecording")
    
    
    return
}

func NewDescribeCloudRecordingResponse() (response *DescribeCloudRecordingResponse) {
    response = &DescribeCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudRecording
// 成功开启录制后，可以使用此接口来查询录制状态。仅在录制任务进行时有效，录制退出后查询将会返回错误。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRecording(request *DescribeCloudRecordingRequest) (response *DescribeCloudRecordingResponse, err error) {
    return c.DescribeCloudRecordingWithContext(context.Background(), request)
}

// DescribeCloudRecording
// 成功开启录制后，可以使用此接口来查询录制状态。仅在录制任务进行时有效，录制退出后查询将会返回错误。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRecordingWithContext(ctx context.Context, request *DescribeCloudRecordingRequest) (response *DescribeCloudRecordingResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRecordingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDetailEventRequest() (request *DescribeDetailEventRequest) {
    request = &DescribeDetailEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeDetailEvent")
    
    
    return
}

func NewDescribeDetailEventResponse() (response *DescribeDetailEventResponse) {
    response = &DescribeDetailEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDetailEvent
// 查询用户某次通话内的进退房，视频开关等详细事件。可查询14天内数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
func (c *Client) DescribeDetailEvent(request *DescribeDetailEventRequest) (response *DescribeDetailEventResponse, err error) {
    return c.DescribeDetailEventWithContext(context.Background(), request)
}

// DescribeDetailEvent
// 查询用户某次通话内的进退房，视频开关等详细事件。可查询14天内数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
func (c *Client) DescribeDetailEventWithContext(ctx context.Context, request *DescribeDetailEventRequest) (response *DescribeDetailEventResponse, err error) {
    if request == nil {
        request = NewDescribeDetailEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDetailEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDetailEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExternalTrtcMeasureRequest() (request *DescribeExternalTrtcMeasureRequest) {
    request = &DescribeExternalTrtcMeasureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeExternalTrtcMeasure")
    
    
    return
}

func NewDescribeExternalTrtcMeasureResponse() (response *DescribeExternalTrtcMeasureResponse) {
    response = &DescribeExternalTrtcMeasureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExternalTrtcMeasure
// 获取Trtc的用量统计数据。走计费渠道二期 只允许查两天的数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
func (c *Client) DescribeExternalTrtcMeasure(request *DescribeExternalTrtcMeasureRequest) (response *DescribeExternalTrtcMeasureResponse, err error) {
    return c.DescribeExternalTrtcMeasureWithContext(context.Background(), request)
}

// DescribeExternalTrtcMeasure
// 获取Trtc的用量统计数据。走计费渠道二期 只允许查两天的数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
func (c *Client) DescribeExternalTrtcMeasureWithContext(ctx context.Context, request *DescribeExternalTrtcMeasureRequest) (response *DescribeExternalTrtcMeasureResponse, err error) {
    if request == nil {
        request = NewDescribeExternalTrtcMeasureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExternalTrtcMeasure require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExternalTrtcMeasureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHistoryScaleRequest() (request *DescribeHistoryScaleRequest) {
    request = &DescribeHistoryScaleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeHistoryScale")
    
    
    return
}

func NewDescribeHistoryScaleResponse() (response *DescribeHistoryScaleResponse) {
    response = &DescribeHistoryScaleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHistoryScale
// 可查询sdkqppid 每天的房间数和用户数，每分钟1次，可查询最近14天的数据。当天未结束，无法查到当天的房间数与用户数。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeHistoryScale(request *DescribeHistoryScaleRequest) (response *DescribeHistoryScaleResponse, err error) {
    return c.DescribeHistoryScaleWithContext(context.Background(), request)
}

// DescribeHistoryScale
// 可查询sdkqppid 每天的房间数和用户数，每分钟1次，可查询最近14天的数据。当天未结束，无法查到当天的房间数与用户数。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeHistoryScaleWithContext(ctx context.Context, request *DescribeHistoryScaleRequest) (response *DescribeHistoryScaleResponse, err error) {
    if request == nil {
        request = NewDescribeHistoryScaleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHistoryScale require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHistoryScaleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePictureRequest() (request *DescribePictureRequest) {
    request = &DescribePictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribePicture")
    
    
    return
}

func NewDescribePictureResponse() (response *DescribePictureResponse) {
    response = &DescribePictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁查找自定义背景图或水印信息，可通过此接口查找已上传的图片信息。无需频繁查找图片信息的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中查看。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribePicture(request *DescribePictureRequest) (response *DescribePictureResponse, err error) {
    return c.DescribePictureWithContext(context.Background(), request)
}

// DescribePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁查找自定义背景图或水印信息，可通过此接口查找已上传的图片信息。无需频繁查找图片信息的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中查看。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribePictureWithContext(ctx context.Context, request *DescribePictureRequest) (response *DescribePictureResponse, err error) {
    if request == nil {
        request = NewDescribePictureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePictureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordStatisticRequest() (request *DescribeRecordStatisticRequest) {
    request = &DescribeRecordStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRecordStatistic")
    
    
    return
}

func NewDescribeRecordStatisticResponse() (response *DescribeRecordStatisticResponse) {
    response = &DescribeRecordStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordStatistic
// 查询云端录制计费时长。
//
// 
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 日结后付费将于次日上午推送账单，建议次日上午9点以后再来查询前一天的用量。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPID = "InvalidParameter.AppId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRecordStatistic(request *DescribeRecordStatisticRequest) (response *DescribeRecordStatisticResponse, err error) {
    return c.DescribeRecordStatisticWithContext(context.Background(), request)
}

// DescribeRecordStatistic
// 查询云端录制计费时长。
//
// 
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 日结后付费将于次日上午推送账单，建议次日上午9点以后再来查询前一天的用量。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPID = "InvalidParameter.AppId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRecordStatisticWithContext(ctx context.Context, request *DescribeRecordStatisticRequest) (response *DescribeRecordStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeRecordStatisticRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomInformationRequest() (request *DescribeRoomInformationRequest) {
    request = &DescribeRoomInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRoomInformation")
    
    
    return
}

func NewDescribeRoomInformationResponse() (response *DescribeRoomInformationResponse) {
    response = &DescribeRoomInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoomInformation
// 查询sdkappid下的房间列表。默认返回10条通话，一次最多返回100条通话。可查询14天内的数据。
//
// **注意**：该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMNUM = "MissingParameter.RoomNum"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeRoomInformation(request *DescribeRoomInformationRequest) (response *DescribeRoomInformationResponse, err error) {
    return c.DescribeRoomInformationWithContext(context.Background(), request)
}

// DescribeRoomInformation
// 查询sdkappid下的房间列表。默认返回10条通话，一次最多返回100条通话。可查询14天内的数据。
//
// **注意**：该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMNUM = "MissingParameter.RoomNum"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeRoomInformationWithContext(ctx context.Context, request *DescribeRoomInformationRequest) (response *DescribeRoomInformationResponse, err error) {
    if request == nil {
        request = NewDescribeRoomInformationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoomInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomInformationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrtcMcuTranscodeTimeRequest() (request *DescribeTrtcMcuTranscodeTimeRequest) {
    request = &DescribeTrtcMcuTranscodeTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTrtcMcuTranscodeTime")
    
    
    return
}

func NewDescribeTrtcMcuTranscodeTimeResponse() (response *DescribeTrtcMcuTranscodeTimeResponse) {
    response = &DescribeTrtcMcuTranscodeTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrtcMcuTranscodeTime
// 查询旁路转码计费时长。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 日结后付费将于次日上午推送账单，建议次日上午9点以后再来查询前一天的用量。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcMcuTranscodeTime(request *DescribeTrtcMcuTranscodeTimeRequest) (response *DescribeTrtcMcuTranscodeTimeResponse, err error) {
    return c.DescribeTrtcMcuTranscodeTimeWithContext(context.Background(), request)
}

// DescribeTrtcMcuTranscodeTime
// 查询旁路转码计费时长。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 日结后付费将于次日上午推送账单，建议次日上午9点以后再来查询前一天的用量。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcMcuTranscodeTimeWithContext(ctx context.Context, request *DescribeTrtcMcuTranscodeTimeRequest) (response *DescribeTrtcMcuTranscodeTimeResponse, err error) {
    if request == nil {
        request = NewDescribeTrtcMcuTranscodeTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrtcMcuTranscodeTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrtcMcuTranscodeTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserInformationRequest() (request *DescribeUserInformationRequest) {
    request = &DescribeUserInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeUserInformation")
    
    
    return
}

func NewDescribeUserInformationResponse() (response *DescribeUserInformationResponse) {
    response = &DescribeUserInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserInformation
// 查询指定时间内的用户列表，可查询14天内数据，查询起止时间不超过4小时。默认每页查询6个用户，支持每页最大查询100个用户PageSize不超过100）。
//
// **注意**：该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUserInformation(request *DescribeUserInformationRequest) (response *DescribeUserInformationResponse, err error) {
    return c.DescribeUserInformationWithContext(context.Background(), request)
}

// DescribeUserInformation
// 查询指定时间内的用户列表，可查询14天内数据，查询起止时间不超过4小时。默认每页查询6个用户，支持每页最大查询100个用户PageSize不超过100）。
//
// **注意**：该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUserInformationWithContext(ctx context.Context, request *DescribeUserInformationRequest) (response *DescribeUserInformationResponse, err error) {
    if request == nil {
        request = NewDescribeUserInformationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserInformationResponse()
    err = c.Send(request, response)
    return
}

func NewDismissRoomRequest() (request *DismissRoomRequest) {
    request = &DismissRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DismissRoom")
    
    
    return
}

func NewDismissRoomResponse() (response *DismissRoomResponse) {
    response = &DismissRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DismissRoom
// 接口说明：把房间所有用户从房间移出，解散房间。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoom(request *DismissRoomRequest) (response *DismissRoomResponse, err error) {
    return c.DismissRoomWithContext(context.Background(), request)
}

// DismissRoom
// 接口说明：把房间所有用户从房间移出，解散房间。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoomWithContext(ctx context.Context, request *DismissRoomRequest) (response *DismissRoomResponse, err error) {
    if request == nil {
        request = NewDismissRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DismissRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDismissRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDismissRoomByStrRoomIdRequest() (request *DismissRoomByStrRoomIdRequest) {
    request = &DismissRoomByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DismissRoomByStrRoomId")
    
    
    return
}

func NewDismissRoomByStrRoomIdResponse() (response *DismissRoomByStrRoomIdResponse) {
    response = &DismissRoomByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DismissRoomByStrRoomId
// 接口说明：把房间所有用户从房间移出，解散房间。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoomByStrRoomId(request *DismissRoomByStrRoomIdRequest) (response *DismissRoomByStrRoomIdResponse, err error) {
    return c.DismissRoomByStrRoomIdWithContext(context.Background(), request)
}

// DismissRoomByStrRoomId
// 接口说明：把房间所有用户从房间移出，解散房间。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoomByStrRoomIdWithContext(ctx context.Context, request *DismissRoomByStrRoomIdRequest) (response *DismissRoomByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewDismissRoomByStrRoomIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DismissRoomByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDismissRoomByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewMeasureTrtcMcuExternalRequest() (request *MeasureTrtcMcuExternalRequest) {
    request = &MeasureTrtcMcuExternalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "MeasureTrtcMcuExternal")
    
    
    return
}

func NewMeasureTrtcMcuExternalResponse() (response *MeasureTrtcMcuExternalResponse) {
    response = &MeasureTrtcMcuExternalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MeasureTrtcMcuExternal
// 查询旁路转码计费时长。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过2天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 日结后付费将于次日上午推送账单，建议次日上午9点以后再来查询前一天的用量。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) MeasureTrtcMcuExternal(request *MeasureTrtcMcuExternalRequest) (response *MeasureTrtcMcuExternalResponse, err error) {
    return c.MeasureTrtcMcuExternalWithContext(context.Background(), request)
}

// MeasureTrtcMcuExternal
// 查询旁路转码计费时长。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过2天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 日结后付费将于次日上午推送账单，建议次日上午9点以后再来查询前一天的用量。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) MeasureTrtcMcuExternalWithContext(ctx context.Context, request *MeasureTrtcMcuExternalRequest) (response *MeasureTrtcMcuExternalResponse, err error) {
    if request == nil {
        request = NewMeasureTrtcMcuExternalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MeasureTrtcMcuExternal require credential")
    }

    request.SetContext(ctx)
    
    response = NewMeasureTrtcMcuExternalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudRecordingRequest() (request *ModifyCloudRecordingRequest) {
    request = &ModifyCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "ModifyCloudRecording")
    
    
    return
}

func NewModifyCloudRecordingResponse() (response *ModifyCloudRecordingResponse) {
    response = &ModifyCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCloudRecording
// 成功开启录制后，可以使用此接口来更新录制任务。仅在录制任务进行时有效，录制退出后更新将会返回错误。更新操作是全量覆盖，并不是增量更新的模式，也就是说每次更新都需要携带全量的信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudRecording(request *ModifyCloudRecordingRequest) (response *ModifyCloudRecordingResponse, err error) {
    return c.ModifyCloudRecordingWithContext(context.Background(), request)
}

// ModifyCloudRecording
// 成功开启录制后，可以使用此接口来更新录制任务。仅在录制任务进行时有效，录制退出后更新将会返回错误。更新操作是全量覆盖，并不是增量更新的模式，也就是说每次更新都需要携带全量的信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudRecordingWithContext(ctx context.Context, request *ModifyCloudRecordingRequest) (response *ModifyCloudRecordingResponse, err error) {
    if request == nil {
        request = NewModifyCloudRecordingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPictureRequest() (request *ModifyPictureRequest) {
    request = &ModifyPictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "ModifyPicture")
    
    
    return
}

func NewModifyPictureResponse() (response *ModifyPictureResponse) {
    response = &ModifyPictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁修改自定义背景图或水印素材，可通过此接口修改已上传的图片。无需频繁修改图片素材的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyPicture(request *ModifyPictureRequest) (response *ModifyPictureResponse, err error) {
    return c.ModifyPictureWithContext(context.Background(), request)
}

// ModifyPicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁修改自定义背景图或水印素材，可通过此接口修改已上传的图片。无需频繁修改图片素材的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyPictureWithContext(ctx context.Context, request *ModifyPictureRequest) (response *ModifyPictureResponse, err error) {
    if request == nil {
        request = NewModifyPictureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPictureResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserRequest() (request *RemoveUserRequest) {
    request = &RemoveUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "RemoveUser")
    
    
    return
}

func NewRemoveUserResponse() (response *RemoveUserResponse) {
    response = &RemoveUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveUser
// 接口说明：将用户从房间移出，适用于主播/房主/管理员踢人等场景。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUser(request *RemoveUserRequest) (response *RemoveUserResponse, err error) {
    return c.RemoveUserWithContext(context.Background(), request)
}

// RemoveUser
// 接口说明：将用户从房间移出，适用于主播/房主/管理员踢人等场景。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUserWithContext(ctx context.Context, request *RemoveUserRequest) (response *RemoveUserResponse, err error) {
    if request == nil {
        request = NewRemoveUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserByStrRoomIdRequest() (request *RemoveUserByStrRoomIdRequest) {
    request = &RemoveUserByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "RemoveUserByStrRoomId")
    
    
    return
}

func NewRemoveUserByStrRoomIdResponse() (response *RemoveUserByStrRoomIdResponse) {
    response = &RemoveUserByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveUserByStrRoomId
// 接口说明：将用户从房间移出，适用于主播/房主/管理员踢人等场景。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUserByStrRoomId(request *RemoveUserByStrRoomIdRequest) (response *RemoveUserByStrRoomIdResponse, err error) {
    return c.RemoveUserByStrRoomIdWithContext(context.Background(), request)
}

// RemoveUserByStrRoomId
// 接口说明：将用户从房间移出，适用于主播/房主/管理员踢人等场景。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUserByStrRoomIdWithContext(ctx context.Context, request *RemoveUserByStrRoomIdRequest) (response *RemoveUserByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewRemoveUserByStrRoomIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUserByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewStartMCUMixTranscodeRequest() (request *StartMCUMixTranscodeRequest) {
    request = &StartMCUMixTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "StartMCUMixTranscode")
    
    
    return
}

func NewStartMCUMixTranscodeResponse() (response *StartMCUMixTranscodeResponse) {
    response = &StartMCUMixTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartMCUMixTranscode
// 接口说明：启动云端混流，并指定混流画面中各路画面的布局位置。
//
// 
//
// TRTC 的一个房间中可能会同时存在多路音视频流，您可以通过此 API 接口，通知腾讯云服务端将多路视频画面合成一路，并指定每一路画面的位置，同时将多路声音进行混音，最终形成一路音视频流，以便用于录制和直播观看。房间销毁后混流自动结束。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// - 设置最终直播流的画质和音质，包括视频分辨率、视频码率、视频帧率、以及声音质量等。
//
// - 设置各路画面的位置和布局，您只需要在启动时设置一次，排版引擎会自动完成后续的画面排布。
//
// - 设置录制文件名，用于二次回放。
//
// - 设置 CDN 直播流 ID，用于在 CDN 进行直播观看。
//
// 
//
// 目前已经支持了如下几种布局模板：
//
// - 悬浮模板：第一个进入房间的用户的视频画面会铺满整个屏幕，其他用户的视频画面从左下角依次水平排列，显示为小画面，最多4行，每行4个，小画面悬浮于大画面之上。最多支持1个大画面和15个小画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 九宫格模板：所有用户的视频画面大小一致，平分整个屏幕，人数越多，每个画面的尺寸越小。最多支持16个画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 屏幕分享模板：适合视频会议和在线教育场景的布局，屏幕分享（或者主讲的摄像头）始终占据屏幕左侧的大画面位置，其他用户依次垂直排列于右侧，最多两列，每列最多8个小画面。最多支持1个大画面和15个小画面。若上行分辨率宽高比与画面输出宽高比不一致时，左侧大画面为了保持内容的完整性采用缩放方式处理，右侧小画面采用裁剪方式处理。
//
// - 画中画模板：适用于混合大小两路视频画面和其他用户混音，或者混合一路大画面和其他用户混音的场景。小画面悬浮于大画面之上，可以指定大小画面的用户以及小画面的显示位置，最多支持2个画面。
//
// - 自定义模板：适用于在混流中指定用户的画面位置，或者预设视频画面位置的场景。当预设位置指定用户时，排版引擎会该用户预留位置；当预设位置未指定用户时，排版引擎会根据进房间顺序自动填充。预设位置填满时，不再混合其他用户的画面和声音。自定义模板启用占位图功能时（LayoutParams中的PlaceHolderMode设置成1），在预设位置的用户没有上行视频时可显示对应的占位图（PlaceImageId）。
//
// 
//
// 注意：
//
// 1、**混流转码为收费功能，调用接口将产生云端混流转码费用，详见[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)。**
//
// 2、2020年1月9号及以后创建的应用才能直接调用此接口。2020年1月9日之前创建的应用默认使用云直播的云端混流，如需切换至MCU混流，请[提交工单](https://console.cloud.tencent.com/workorder/category)寻求帮助。
//
// 3、客户端混流和服务端混流不能混用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEORIGHTALIGN = "InvalidParameter.MainVideoRightAlign"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUBLISHCDNURLS = "InvalidParameter.PublishCdnUrls"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscode(request *StartMCUMixTranscodeRequest) (response *StartMCUMixTranscodeResponse, err error) {
    return c.StartMCUMixTranscodeWithContext(context.Background(), request)
}

// StartMCUMixTranscode
// 接口说明：启动云端混流，并指定混流画面中各路画面的布局位置。
//
// 
//
// TRTC 的一个房间中可能会同时存在多路音视频流，您可以通过此 API 接口，通知腾讯云服务端将多路视频画面合成一路，并指定每一路画面的位置，同时将多路声音进行混音，最终形成一路音视频流，以便用于录制和直播观看。房间销毁后混流自动结束。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// - 设置最终直播流的画质和音质，包括视频分辨率、视频码率、视频帧率、以及声音质量等。
//
// - 设置各路画面的位置和布局，您只需要在启动时设置一次，排版引擎会自动完成后续的画面排布。
//
// - 设置录制文件名，用于二次回放。
//
// - 设置 CDN 直播流 ID，用于在 CDN 进行直播观看。
//
// 
//
// 目前已经支持了如下几种布局模板：
//
// - 悬浮模板：第一个进入房间的用户的视频画面会铺满整个屏幕，其他用户的视频画面从左下角依次水平排列，显示为小画面，最多4行，每行4个，小画面悬浮于大画面之上。最多支持1个大画面和15个小画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 九宫格模板：所有用户的视频画面大小一致，平分整个屏幕，人数越多，每个画面的尺寸越小。最多支持16个画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 屏幕分享模板：适合视频会议和在线教育场景的布局，屏幕分享（或者主讲的摄像头）始终占据屏幕左侧的大画面位置，其他用户依次垂直排列于右侧，最多两列，每列最多8个小画面。最多支持1个大画面和15个小画面。若上行分辨率宽高比与画面输出宽高比不一致时，左侧大画面为了保持内容的完整性采用缩放方式处理，右侧小画面采用裁剪方式处理。
//
// - 画中画模板：适用于混合大小两路视频画面和其他用户混音，或者混合一路大画面和其他用户混音的场景。小画面悬浮于大画面之上，可以指定大小画面的用户以及小画面的显示位置，最多支持2个画面。
//
// - 自定义模板：适用于在混流中指定用户的画面位置，或者预设视频画面位置的场景。当预设位置指定用户时，排版引擎会该用户预留位置；当预设位置未指定用户时，排版引擎会根据进房间顺序自动填充。预设位置填满时，不再混合其他用户的画面和声音。自定义模板启用占位图功能时（LayoutParams中的PlaceHolderMode设置成1），在预设位置的用户没有上行视频时可显示对应的占位图（PlaceImageId）。
//
// 
//
// 注意：
//
// 1、**混流转码为收费功能，调用接口将产生云端混流转码费用，详见[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)。**
//
// 2、2020年1月9号及以后创建的应用才能直接调用此接口。2020年1月9日之前创建的应用默认使用云直播的云端混流，如需切换至MCU混流，请[提交工单](https://console.cloud.tencent.com/workorder/category)寻求帮助。
//
// 3、客户端混流和服务端混流不能混用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEORIGHTALIGN = "InvalidParameter.MainVideoRightAlign"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUBLISHCDNURLS = "InvalidParameter.PublishCdnUrls"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscodeWithContext(ctx context.Context, request *StartMCUMixTranscodeRequest) (response *StartMCUMixTranscodeResponse, err error) {
    if request == nil {
        request = NewStartMCUMixTranscodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartMCUMixTranscode require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartMCUMixTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewStartMCUMixTranscodeByStrRoomIdRequest() (request *StartMCUMixTranscodeByStrRoomIdRequest) {
    request = &StartMCUMixTranscodeByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "StartMCUMixTranscodeByStrRoomId")
    
    
    return
}

func NewStartMCUMixTranscodeByStrRoomIdResponse() (response *StartMCUMixTranscodeByStrRoomIdResponse) {
    response = &StartMCUMixTranscodeByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartMCUMixTranscodeByStrRoomId
// 接口说明：启动云端混流，并指定混流画面中各路画面的布局位置。
//
// 
//
// TRTC 的一个房间中可能会同时存在多路音视频流，您可以通过此 API 接口，通知腾讯云服务端将多路视频画面合成一路，并指定每一路画面的位置，同时将多路声音进行混音，最终形成一路音视频流，以便用于录制和直播观看。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// - 设置最终直播流的画质和音质，包括视频分辨率、视频码率、视频帧率、以及声音质量等。
//
// - 设置各路画面的位置和布局，您只需要在启动时设置一次，排版引擎会自动完成后续的画面排布。
//
// - 设置录制文件名，用于二次回放。
//
// - 设置 CDN 直播流 ID，用于在 CDN 进行直播观看。
//
// 
//
// 目前已经支持了如下几种布局模板：
//
// - 悬浮模板：第一个进入房间的用户的视频画面会铺满整个屏幕，其他用户的视频画面从左下角依次水平排列，显示为小画面，最多4行，每行4个，小画面悬浮于大画面之上。最多支持1个大画面和15个小画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 九宫格模板：所有用户的视频画面大小一致，平分整个屏幕，人数越多，每个画面的尺寸越小。最多支持16个画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 屏幕分享模板：适合视频会议和在线教育场景的布局，屏幕分享（或者主讲的摄像头）始终占据屏幕左侧的大画面位置，其他用户依次垂直排列于右侧，最多两列，每列最多8个小画面。最多支持1个大画面和15个小画面。若上行分辨率宽高比与画面输出宽高比不一致时，左侧大画面为了保持内容的完整性采用缩放方式处理，右侧小画面采用裁剪方式处理。
//
// - 画中画模板：适用于混合大小两路视频画面和其他用户混音，或者混合一路大画面和其他用户混音的场景。小画面悬浮于大画面之上，可以指定大小画面的用户以及小画面的显示位置。
//
// - 自定义模板：适用于在混流中指定用户的画面位置，或者预设视频画面位置的场景。当预设位置指定用户时，排版引擎会该用户预留位置；当预设位置未指定用户时，排版引擎会根据进房间顺序自动填充。预设位置填满时，不再混合其他用户的画面和声音。自定义模板启用占位图功能时（LayoutParams中的PlaceHolderMode设置成1），在预设位置的用户没有上行视频时可显示对应的占位图（PlaceImageId）。
//
// 
//
// 注意：
//
// 1、**混流转码为收费功能，调用接口将产生云端混流转码费用，详见[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)。**
//
// 2、2020年1月9号及以后创建的应用才能直接调用此接口。2020年1月9日之前创建的应用默认使用云直播的云端混流，如需切换至MCU混流，请[提交工单](https://console.cloud.tencent.com/workorder/category)寻求帮助。
//
// 3、客户端混流和服务端混流不能混用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUBLISHCDNURLS = "InvalidParameter.PublishCdnUrls"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscodeByStrRoomId(request *StartMCUMixTranscodeByStrRoomIdRequest) (response *StartMCUMixTranscodeByStrRoomIdResponse, err error) {
    return c.StartMCUMixTranscodeByStrRoomIdWithContext(context.Background(), request)
}

// StartMCUMixTranscodeByStrRoomId
// 接口说明：启动云端混流，并指定混流画面中各路画面的布局位置。
//
// 
//
// TRTC 的一个房间中可能会同时存在多路音视频流，您可以通过此 API 接口，通知腾讯云服务端将多路视频画面合成一路，并指定每一路画面的位置，同时将多路声音进行混音，最终形成一路音视频流，以便用于录制和直播观看。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// - 设置最终直播流的画质和音质，包括视频分辨率、视频码率、视频帧率、以及声音质量等。
//
// - 设置各路画面的位置和布局，您只需要在启动时设置一次，排版引擎会自动完成后续的画面排布。
//
// - 设置录制文件名，用于二次回放。
//
// - 设置 CDN 直播流 ID，用于在 CDN 进行直播观看。
//
// 
//
// 目前已经支持了如下几种布局模板：
//
// - 悬浮模板：第一个进入房间的用户的视频画面会铺满整个屏幕，其他用户的视频画面从左下角依次水平排列，显示为小画面，最多4行，每行4个，小画面悬浮于大画面之上。最多支持1个大画面和15个小画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 九宫格模板：所有用户的视频画面大小一致，平分整个屏幕，人数越多，每个画面的尺寸越小。最多支持16个画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 屏幕分享模板：适合视频会议和在线教育场景的布局，屏幕分享（或者主讲的摄像头）始终占据屏幕左侧的大画面位置，其他用户依次垂直排列于右侧，最多两列，每列最多8个小画面。最多支持1个大画面和15个小画面。若上行分辨率宽高比与画面输出宽高比不一致时，左侧大画面为了保持内容的完整性采用缩放方式处理，右侧小画面采用裁剪方式处理。
//
// - 画中画模板：适用于混合大小两路视频画面和其他用户混音，或者混合一路大画面和其他用户混音的场景。小画面悬浮于大画面之上，可以指定大小画面的用户以及小画面的显示位置。
//
// - 自定义模板：适用于在混流中指定用户的画面位置，或者预设视频画面位置的场景。当预设位置指定用户时，排版引擎会该用户预留位置；当预设位置未指定用户时，排版引擎会根据进房间顺序自动填充。预设位置填满时，不再混合其他用户的画面和声音。自定义模板启用占位图功能时（LayoutParams中的PlaceHolderMode设置成1），在预设位置的用户没有上行视频时可显示对应的占位图（PlaceImageId）。
//
// 
//
// 注意：
//
// 1、**混流转码为收费功能，调用接口将产生云端混流转码费用，详见[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)。**
//
// 2、2020年1月9号及以后创建的应用才能直接调用此接口。2020年1月9日之前创建的应用默认使用云直播的云端混流，如需切换至MCU混流，请[提交工单](https://console.cloud.tencent.com/workorder/category)寻求帮助。
//
// 3、客户端混流和服务端混流不能混用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUBLISHCDNURLS = "InvalidParameter.PublishCdnUrls"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscodeByStrRoomIdWithContext(ctx context.Context, request *StartMCUMixTranscodeByStrRoomIdRequest) (response *StartMCUMixTranscodeByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewStartMCUMixTranscodeByStrRoomIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartMCUMixTranscodeByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartMCUMixTranscodeByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewStopMCUMixTranscodeRequest() (request *StopMCUMixTranscodeRequest) {
    request = &StopMCUMixTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "StopMCUMixTranscode")
    
    
    return
}

func NewStopMCUMixTranscodeResponse() (response *StopMCUMixTranscodeResponse) {
    response = &StopMCUMixTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopMCUMixTranscode
// 接口说明：结束云端混流
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscode(request *StopMCUMixTranscodeRequest) (response *StopMCUMixTranscodeResponse, err error) {
    return c.StopMCUMixTranscodeWithContext(context.Background(), request)
}

// StopMCUMixTranscode
// 接口说明：结束云端混流
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscodeWithContext(ctx context.Context, request *StopMCUMixTranscodeRequest) (response *StopMCUMixTranscodeResponse, err error) {
    if request == nil {
        request = NewStopMCUMixTranscodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopMCUMixTranscode require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopMCUMixTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewStopMCUMixTranscodeByStrRoomIdRequest() (request *StopMCUMixTranscodeByStrRoomIdRequest) {
    request = &StopMCUMixTranscodeByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "StopMCUMixTranscodeByStrRoomId")
    
    
    return
}

func NewStopMCUMixTranscodeByStrRoomIdResponse() (response *StopMCUMixTranscodeByStrRoomIdResponse) {
    response = &StopMCUMixTranscodeByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopMCUMixTranscodeByStrRoomId
// 接口说明：结束云端混流
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscodeByStrRoomId(request *StopMCUMixTranscodeByStrRoomIdRequest) (response *StopMCUMixTranscodeByStrRoomIdResponse, err error) {
    return c.StopMCUMixTranscodeByStrRoomIdWithContext(context.Background(), request)
}

// StopMCUMixTranscodeByStrRoomId
// 接口说明：结束云端混流
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscodeByStrRoomIdWithContext(ctx context.Context, request *StopMCUMixTranscodeByStrRoomIdRequest) (response *StopMCUMixTranscodeByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewStopMCUMixTranscodeByStrRoomIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopMCUMixTranscodeByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopMCUMixTranscodeByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}
