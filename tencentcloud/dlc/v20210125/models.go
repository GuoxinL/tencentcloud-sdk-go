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

package v20210125

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddUsersToWorkGroupRequest struct {
	*tchttp.BaseRequest

	// 要操作的工作组和用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitempty" name:"AddInfo"`
}

func (r *AddUsersToWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersToWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUsersToWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddUsersToWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUsersToWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersToWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachUserPolicyRequest struct {
	*tchttp.BaseRequest

	// 用户Id，和子用户uin相同，需要先使用CreateUser接口创建用户。可以使用DescribeUsers接口查看。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
}

func (r *AttachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachWorkGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 要绑定的策略集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
}

func (r *AttachWorkGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachWorkGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachWorkGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachWorkGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachWorkGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachWorkGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindWorkGroupsToUserRequest struct {
	*tchttp.BaseRequest

	// 绑定的用户和工作组信息
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitempty" name:"AddInfo"`
}

func (r *BindWorkGroupsToUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindWorkGroupsToUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindWorkGroupsToUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindWorkGroupsToUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindWorkGroupsToUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindWorkGroupsToUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CSV struct {

	// 压缩格式，["Snappy", "Gzip", "None"选一]。
	CodeCompress *string `json:"CodeCompress,omitempty" name:"CodeCompress"`

	// CSV序列化及反序列化数据结构。
	CSVSerde *CSVSerde `json:"CSVSerde,omitempty" name:"CSVSerde"`

	// 标题行，默认为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadLines *int64 `json:"HeadLines,omitempty" name:"HeadLines"`

	// 格式，默认值为CSV
	// 注意：此字段可能返回 null，表示取不到有效值。
	Format *string `json:"Format,omitempty" name:"Format"`
}

type CSVSerde struct {

	// CSV序列化转义符，默认为"\\"，最长8个字符，如 Escape: "/\"
	Escape *string `json:"Escape,omitempty" name:"Escape"`

	// CSV序列化字段域符，默认为"'"，最长8个字符, 如 Quote: "\""
	Quote *string `json:"Quote,omitempty" name:"Quote"`

	// CSV序列化分隔符，默认为"\t"，最长8个字符, 如 Separator: "\t"
	Separator *string `json:"Separator,omitempty" name:"Separator"`
}

type CancelTaskRequest struct {
	*tchttp.BaseRequest

	// 任务Id，全局唯一
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *CancelTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Column struct {

	// 列名称，不区分大小写，最大支持25个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 列类型，支持如下类型定义:
	// string|tinyint|smallint|int|bigint|boolean|float|double|decimal|timestamp|date|binary|array<data_type>|map<primitive_type, data_type>|struct<col_name : data_type [COMMENT col_comment], ...>|uniontype<data_type, data_type, ...>。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 对该类的注释。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 表示整个 numeric 的长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Precision *int64 `json:"Precision,omitempty" name:"Precision"`

	// 表示小数部分的长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scale *int64 `json:"Scale,omitempty" name:"Scale"`

	// 是否为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nullable *string `json:"Nullable,omitempty" name:"Nullable"`

	// 字段位置，小的在前
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *int64 `json:"Position,omitempty" name:"Position"`

	// 字段创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 字段修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type CreateDatabaseRequest struct {
	*tchttp.BaseRequest

	// 数据库基础信息
	DatabaseInfo *DatabaseInfo `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`
}

func (r *CreateDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseInfo")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生成的建库执行语句对象。
		Execution *Execution `json:"Execution,omitempty" name:"Execution"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportTaskRequest struct {
	*tchttp.BaseRequest

	// 数据来源，lakefsStorage、taskResult
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 导出任务输入配置
	InputConf []*KVPair `json:"InputConf,omitempty" name:"InputConf"`

	// 导出任务输出配置
	OutputConf []*KVPair `json:"OutputConf,omitempty" name:"OutputConf"`

	// 目标数据源的类型，目前支持导出到cos
	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`
}

func (r *CreateExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputType")
	delete(f, "InputConf")
	delete(f, "OutputConf")
	delete(f, "OutputType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateImportTaskRequest struct {
	*tchttp.BaseRequest

	// 数据来源，cos
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 输入配置
	InputConf []*KVPair `json:"InputConf,omitempty" name:"InputConf"`

	// 输出配置
	OutputConf []*KVPair `json:"OutputConf,omitempty" name:"OutputConf"`

	// 目标数据源的类型，目前支持导入到托管存储，即lakefsStorage
	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`
}

func (r *CreateImportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputType")
	delete(f, "InputConf")
	delete(f, "OutputConf")
	delete(f, "OutputType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateImportTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateScriptRequest struct {
	*tchttp.BaseRequest

	// 脚本名称，最大不能超过255个字符。
	ScriptName *string `json:"ScriptName,omitempty" name:"ScriptName"`

	// base64编码后的sql语句
	SQLStatement *string `json:"SQLStatement,omitempty" name:"SQLStatement"`

	// 脚本描述， 不能超过50个字符
	ScriptDesc *string `json:"ScriptDesc,omitempty" name:"ScriptDesc"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

func (r *CreateScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptName")
	delete(f, "SQLStatement")
	delete(f, "ScriptDesc")
	delete(f, "DatabaseName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateScriptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSparkAppRequest struct {
	*tchttp.BaseRequest

	// spark应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 1代表spark jar应用，2代表spark streaming应用
	AppType *int64 `json:"AppType,omitempty" name:"AppType"`

	// 执行spark作业的数据引擎
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// spark应用的执行入口
	AppFile *string `json:"AppFile,omitempty" name:"AppFile"`

	// 执行spark作业的角色ID
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// spark作业driver资源规格大小, 可取small,medium,large,xlarge
	AppDriverSize *string `json:"AppDriverSize,omitempty" name:"AppDriverSize"`

	// spark作业executor资源规格大小, 可取small,medium,large,xlarge
	AppExecutorSize *string `json:"AppExecutorSize,omitempty" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitempty" name:"AppExecutorNums"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// 是否本地上传，可去cos,lakefs
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// spark jar作业时的主类
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitempty" name:"AppConf"`

	// 是否本地上传，包含cos,lakefs
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// spark jar作业依赖jars，以逗号分隔
	AppJars *string `json:"AppJars,omitempty" name:"AppJars"`

	// 是否本地上传，包含cos,lakefs
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// spark作业依赖资源，以逗号分隔
	AppFiles *string `json:"AppFiles,omitempty" name:"AppFiles"`

	// spark作业命令行参数
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// 只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitempty" name:"MaxRetries"`

	// 数据源名
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// pyspark：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// pyspark：python依赖, 除py文件外，还支持zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`
}

func (r *CreateSparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "AppType")
	delete(f, "DataEngine")
	delete(f, "AppFile")
	delete(f, "RoleArn")
	delete(f, "AppDriverSize")
	delete(f, "AppExecutorSize")
	delete(f, "AppExecutorNums")
	delete(f, "Eni")
	delete(f, "IsLocal")
	delete(f, "MainClass")
	delete(f, "AppConf")
	delete(f, "IsLocalJars")
	delete(f, "AppJars")
	delete(f, "IsLocalFiles")
	delete(f, "AppFiles")
	delete(f, "CmdArgs")
	delete(f, "MaxRetries")
	delete(f, "DataSource")
	delete(f, "IsLocalPythonFiles")
	delete(f, "AppPythonFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSparkAppResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSparkAppTaskRequest struct {
	*tchttp.BaseRequest

	// spark作业名
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// spark作业的命令行参数，以空格分隔；一般用于周期性调用使用
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`
}

func (r *CreateSparkAppTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobName")
	delete(f, "CmdArgs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkAppTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSparkAppTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批Id
		BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

		// 任务Id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSparkAppTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStoreLocationRequest struct {
	*tchttp.BaseRequest

	// 计算结果存储cos路径，如：cosn://bucketname/
	StoreLocation *string `json:"StoreLocation,omitempty" name:"StoreLocation"`
}

func (r *CreateStoreLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStoreLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StoreLocation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStoreLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateStoreLocationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStoreLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStoreLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTableRequest struct {
	*tchttp.BaseRequest

	// 数据表配置信息
	TableInfo *TableInfo `json:"TableInfo,omitempty" name:"TableInfo"`
}

func (r *CreateTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生成的建表执行语句对象。
		Execution *Execution `json:"Execution,omitempty" name:"Execution"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest

	// 计算任务，该参数中包含任务类型及其相关配置信息
	Task *Task `json:"Task,omitempty" name:"Task"`

	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 默认数据源名称。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 数据引擎名称，不填提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
}

func (r *CreateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Task")
	delete(f, "DatabaseName")
	delete(f, "DatasourceConnectionName")
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTasksInOrderRequest struct {
	*tchttp.BaseRequest

	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 数据源名称，默认为COSDataCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`
}

func (r *CreateTasksInOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksInOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Tasks")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTasksInOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTasksInOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本批次提交的任务的批次Id
		BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

		// 任务Id集合，按照执行顺序排列
		TaskIdSet []*string `json:"TaskIdSet,omitempty" name:"TaskIdSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTasksInOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksInOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTasksRequest struct {
	*tchttp.BaseRequest

	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 计算引擎名称，不填任务提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
}

func (r *CreateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Tasks")
	delete(f, "DatasourceConnectionName")
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本批次提交的任务的批次Id
		BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

		// 任务Id集合，按照执行顺序排列
		TaskIdSet []*string `json:"TaskIdSet,omitempty" name:"TaskIdSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserRequest struct {
	*tchttp.BaseRequest

	// 需要授权的子用户uin，可以通过腾讯云控制台右上角 → “账号信息” → “账号ID进行查看”。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户描述信息，方便区分不同用户
	UserDescription *string `json:"UserDescription,omitempty" name:"UserDescription"`

	// 绑定到用户的权限集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`

	// 用户类型。ADMIN：管理员 COMMON：一般用户。当用户类型为管理员的时候，不能设置权限集合和绑定的工作组集合，管理员默认拥有所有权限。该参数不填默认为COMMON
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// 绑定到用户的工作组ID集合。
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitempty" name:"WorkGroupIds"`

	// 用户别名，字符长度小50
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserDescription")
	delete(f, "PolicySet")
	delete(f, "UserType")
	delete(f, "WorkGroupIds")
	delete(f, "UserAlias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkGroupRequest struct {
	*tchttp.BaseRequest

	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitempty" name:"WorkGroupName"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitempty" name:"WorkGroupDescription"`

	// 工作组绑定的鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`

	// 需要绑定到工作组的用户Id集合
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *CreateWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupName")
	delete(f, "WorkGroupDescription")
	delete(f, "PolicySet")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 工作组Id，全局唯一
		WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataFormat struct {

	// 文本格式，TextFile。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextFile *TextFile `json:"TextFile,omitempty" name:"TextFile"`

	// 文本格式，CSV。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CSV *CSV `json:"CSV,omitempty" name:"CSV"`

	// 文本格式，Json。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Json *Other `json:"Json,omitempty" name:"Json"`

	// Parquet格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parquet *Other `json:"Parquet,omitempty" name:"Parquet"`

	// ORC格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ORC *Other `json:"ORC,omitempty" name:"ORC"`

	// AVRO格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AVRO *Other `json:"AVRO,omitempty" name:"AVRO"`
}

type DatabaseInfo struct {

	// 数据库名称，长度0~128，支持数字、字母下划线，不允许数字大头，统一转换为小写。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据库描述信息，长度 0~500。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 数据库属性列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`

	// 数据库cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitempty" name:"Location"`
}

type DatabaseResponseInfo struct {

	// 数据库名称。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据库描述信息，长度 0~256。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 允许针对数据库的属性元数据信息进行指定。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`

	// 数据库创建时间戳，单位：s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 数据库更新时间戳，单位：s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type DeleteScriptRequest struct {
	*tchttp.BaseRequest

	// 脚本id，其可以通过DescribeScripts接口提取
	ScriptIds []*string `json:"ScriptIds,omitempty" name:"ScriptIds"`
}

func (r *DeleteScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScriptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除的脚本数量
		ScriptsAffected *int64 `json:"ScriptsAffected,omitempty" name:"ScriptsAffected"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSparkAppRequest struct {
	*tchttp.BaseRequest

	// spark应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *DeleteSparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSparkAppResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest

	// 需要删除的用户的Id
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsersFromWorkGroupRequest struct {
	*tchttp.BaseRequest

	// 要删除的用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitempty" name:"AddInfo"`
}

func (r *DeleteUsersFromWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersFromWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsersFromWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsersFromWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUsersFromWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersFromWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkGroupRequest struct {
	*tchttp.BaseRequest

	// 要删除的工作组Id集合
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitempty" name:"WorkGroupIds"`
}

func (r *DeleteWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 模糊匹配，库名关键字。
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 数据源唯名称，该名称可以通过DescribeDatasourceConnection接口查询到。默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 排序字段，当前版本仅支持按库名排序
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序类型：false：降序（默认）、true：升序
	Asc *bool `json:"Asc,omitempty" name:"Asc"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "KeyWord")
	delete(f, "DatasourceConnectionName")
	delete(f, "Sort")
	delete(f, "Asc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据库对象列表。
		DatabaseList []*DatabaseResponseInfo `json:"DatabaseList,omitempty" name:"DatabaseList"`

		// 实例总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScriptsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按字段排序，支持如下字段类型，update-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// script-id - String - （过滤条件）script-id取值形如：157de0d1-26b4-4df2-a2d0-b64afc406c25。
	// script-name-keyword - String - （过滤条件）数据表名称,形如：script-test。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeScriptsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScriptsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScriptsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScriptsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Script列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Scripts []*Script `json:"Scripts,omitempty" name:"Scripts"`

		// 实例总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScriptsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScriptsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSparkAppJobRequest struct {
	*tchttp.BaseRequest

	// spark作业Id，与JobName同时存在时，JobName无效
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// spark作业名
	JobName *string `json:"JobName,omitempty" name:"JobName"`
}

func (r *DescribeSparkAppJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSparkAppJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// spark作业详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Job *SparkJobInfo `json:"Job,omitempty" name:"Job"`

		// 查询的spark作业是否存在
		IsExists *bool `json:"IsExists,omitempty" name:"IsExists"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSparkAppJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSparkAppJobsRequest struct {
	*tchttp.BaseRequest

	// 返回结果按照该字段排序
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 正序或者倒序，例如：desc
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 按照该参数过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 更新时间起始点
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 更新时间截止点
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表限制数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSparkAppJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSparkAppJobsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// spark作业列表详情
		SparkAppJobs []*SparkJobInfo `json:"SparkAppJobs,omitempty" name:"SparkAppJobs"`

		// spark作业总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSparkAppJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSparkAppTasksRequest struct {
	*tchttp.BaseRequest

	// spark作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 分页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSparkAppTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSparkAppTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务结果（该字段已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
		Tasks *TaskResponseInfo `json:"Tasks,omitempty" name:"Tasks"`

		// 任务总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		SparkAppTasks []*TaskResponseInfo `json:"SparkAppTasks,omitempty" name:"SparkAppTasks"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSparkAppTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStoreLocationRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeStoreLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStoreLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStoreLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStoreLocationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回用户设置的结果存储位置路径，如果未设置则返回空字符串：""
	// 注意：此字段可能返回 null，表示取不到有效值。
		StoreLocation *string `json:"StoreLocation,omitempty" name:"StoreLocation"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStoreLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStoreLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTableRequest struct {
	*tchttp.BaseRequest

	// 查询对象表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 查询表所在的数据库名称。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 查询表所在的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`
}

func (r *DescribeTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "DatabaseName")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据表对象
		Table *TableResponseInfo `json:"Table,omitempty" name:"Table"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest

	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 指定查询的数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 起始时间：用于对更新时间的筛选
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 终止时间：用于对更新时间的筛选
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序字段，支持：ModifiedTime（默认）；CreateTime
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序字段，false：降序（默认）；true
	Asc *bool `json:"Asc,omitempty" name:"Asc"`

	// table type，表类型查询,可用值:EXTERNAL_TABLE,INDEX_TABLE,MANAGED_TABLE,MATERIALIZED_VIEW,TABLE,VIEW,VIRTUAL_VIEW
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 筛选字段-表格式：不传（默认）为查全部；LAKEFS：托管表；ICEBERG：非托管iceberg表；HIVE：非托管hive表；OTHER：非托管其它；
	TableFormat *string `json:"TableFormat,omitempty" name:"TableFormat"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "DatasourceConnectionName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Sort")
	delete(f, "Asc")
	delete(f, "TableType")
	delete(f, "TableFormat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据表对象列表。
		TableList []*TableResponseInfo `json:"TableList,omitempty" name:"TableList"`

		// 实例总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest

	// 任务唯一ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回1000行数据。
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询的任务信息，返回为空表示输入任务ID对应的任务不存在。只有当任务状态为成功（2）的时候，才会返回任务的结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskInfo *TaskResultInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,其中task-id支持最大50个过滤个数，其他过滤参数支持的总数不超过5个。
	// task-id - String - （任务ID准确过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字模糊过滤）取值形如：DROP TABLE。
	// task-operator- string （子uin过滤）
	// task-kind - string （任务类型过滤）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time（创建时间，默认）、update-time（更新时间）
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 支持计算资源名字筛选
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务对象列表。
		TaskList []*TaskResponseInfo `json:"TaskList,omitempty" name:"TaskList"`

		// 实例总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest

	// 指定查询的子用户uin，用户需要通过CreateUser接口创建。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 过滤条件，支持如下字段类型，user-type：根据用户类型过滤。user-keyword：根据用户名称过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询到的用户总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 查询到的授权用户信息集合
		UserSet []*UserInfo `json:"UserSet,omitempty" name:"UserSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeViewsRequest struct {
	*tchttp.BaseRequest

	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// view-name - String - （过滤条件）数据表名称,形如：view-001。
	// view-id - String - （过滤条件）view id形如：12342。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 数据库所属的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 排序字段
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序规则
	Asc *bool `json:"Asc,omitempty" name:"Asc"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeViewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeViewsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "DatasourceConnectionName")
	delete(f, "Sort")
	delete(f, "Asc")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeViewsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeViewsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 视图对象列表。
		ViewList []*ViewResponseInfo `json:"ViewList,omitempty" name:"ViewList"`

		// 实例总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeViewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeViewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWorkGroupsRequest struct {
	*tchttp.BaseRequest

	// 查询的工作组Id，不填或填0表示不过滤。
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 过滤条件，当前仅支持按照工作组名称进行模糊搜索。Key为workgroup-name
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`
}

func (r *DescribeWorkGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Sorting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWorkGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 工作组总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 工作组信息集合
		WorkGroupSet []*WorkGroupInfo `json:"WorkGroupSet,omitempty" name:"WorkGroupSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWorkGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachUserPolicyRequest struct {
	*tchttp.BaseRequest

	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
}

func (r *DetachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachWorkGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
}

func (r *DetachWorkGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachWorkGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachWorkGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachWorkGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachWorkGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachWorkGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Execution struct {

	// 自动生成SQL语句。
	SQL *string `json:"SQL,omitempty" name:"SQL"`
}

type Filter struct {

	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑或（OR）关系。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type KVPair struct {

	// 配置的key值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 配置的value值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifySparkAppRequest struct {
	*tchttp.BaseRequest

	// spark应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 1代表spark jar应用，2代表spark streaming应用
	AppType *int64 `json:"AppType,omitempty" name:"AppType"`

	// 执行spark作业的数据引擎
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// spark应用的执行入口
	AppFile *string `json:"AppFile,omitempty" name:"AppFile"`

	// 执行spark作业的角色ID
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// spark作业driver资源规格大小, 可取small,medium,large,xlarge
	AppDriverSize *string `json:"AppDriverSize,omitempty" name:"AppDriverSize"`

	// spark作业executor资源规格大小, 可取small,medium,large,xlarge
	AppExecutorSize *string `json:"AppExecutorSize,omitempty" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitempty" name:"AppExecutorNums"`

	// spark应用Id
	SparkAppId *string `json:"SparkAppId,omitempty" name:"SparkAppId"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// 是否本地上传，可取cos,lakefs
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// spark jar作业时的主类
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitempty" name:"AppConf"`

	// 是否本地上传，可去cos,lakefs
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// spark jar作业依赖jars，以逗号分隔
	AppJars *string `json:"AppJars,omitempty" name:"AppJars"`

	// 是否本地上传，可去cos,lakefs
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// spark作业依赖资源，以逗号分隔
	AppFiles *string `json:"AppFiles,omitempty" name:"AppFiles"`

	// pyspark：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// pyspark：python依赖, 除py文件外，还支持zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`

	// spark作业命令行参数
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// 只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitempty" name:"MaxRetries"`

	// 数据源名
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`
}

func (r *ModifySparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "AppType")
	delete(f, "DataEngine")
	delete(f, "AppFile")
	delete(f, "RoleArn")
	delete(f, "AppDriverSize")
	delete(f, "AppExecutorSize")
	delete(f, "AppExecutorNums")
	delete(f, "SparkAppId")
	delete(f, "Eni")
	delete(f, "IsLocal")
	delete(f, "MainClass")
	delete(f, "AppConf")
	delete(f, "IsLocalJars")
	delete(f, "AppJars")
	delete(f, "IsLocalFiles")
	delete(f, "AppFiles")
	delete(f, "IsLocalPythonFiles")
	delete(f, "AppPythonFiles")
	delete(f, "CmdArgs")
	delete(f, "MaxRetries")
	delete(f, "DataSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySparkAppResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest

	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户描述
	UserDescription *string `json:"UserDescription,omitempty" name:"UserDescription"`
}

func (r *ModifyUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWorkGroupRequest struct {
	*tchttp.BaseRequest

	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitempty" name:"WorkGroupDescription"`
}

func (r *ModifyWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "WorkGroupDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Other struct {

	// 枚举类型，默认值为Json，可选值为[Json, Parquet, ORC, AVRD]之一。
	Format *string `json:"Format,omitempty" name:"Format"`
}

type Partition struct {

	// 分区列名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分区类型。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 对分区的描述。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 隐式分区转换策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transform *string `json:"Transform,omitempty" name:"Transform"`

	// 转换策略参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransformArgs []*string `json:"TransformArgs,omitempty" name:"TransformArgs"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Policy struct {

	// 需要授权的数据库名，填*代表当前Catalog下所有数据库。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别时只允许填空，其他类型下可以任意指定数据库。
	Database *string `json:"Database,omitempty" name:"Database"`

	// 需要授权的数据源名称，管理员级别下只支持填*（代表该级别全部资源）；数据源级别和数据库级别鉴权的情况下，只支持填COSDataCatalog或者*；在数据表级别鉴权下可以填写用户自定义数据源。不填情况下默认为DataLakeCatalog。注意：如果是对用户自定义数据源进行鉴权，DLC能够管理的权限是用户接入数据源的时候提供的账户的子集。
	Catalog *string `json:"Catalog,omitempty" name:"Catalog"`

	// 需要授权的表名，填*代表当前Database下所有表。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别、数据库级别时只允许填空，其他类型下可以任意指定数据表。
	Table *string `json:"Table,omitempty" name:"Table"`

	// 授权的权限操作，对于不同级别的鉴权提供不同操作。管理员权限：ALL，不填默认为ALL；数据连接级鉴权：CREATE；数据库级别鉴权：ALL、CREATE、ALTER、DROP；数据表权限：ALL、SELECT、INSERT、ALTER、DELETE、DROP、UPDATE。注意：在数据表权限下，指定的数据源不为COSDataCatalog的时候，只支持SELECT操作。
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 授权类型，现在支持八种授权类型：ADMIN:管理员级别鉴权 DATASOURCE：数据连接级别鉴权 DATABASE：数据库级别鉴权 TABLE：表级别鉴权 VIEW：视图级别鉴权 FUNCTION：函数级别鉴权 COLUMN：列级别鉴权 ENGINE：数据引擎鉴权。不填默认为管理员级别鉴权。
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 需要授权的函数名，填*代表当前Catalog下所有函数。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别时只允许填空，其他类型下可以任意指定函数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Function *string `json:"Function,omitempty" name:"Function"`

	// 需要授权的视图，填*代表当前Database下所有视图。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别、数据库级别时只允许填空，其他类型下可以任意指定视图。
	// 注意：此字段可能返回 null，表示取不到有效值。
	View *string `json:"View,omitempty" name:"View"`

	// 需要授权的列，填*代表当前所有列。当授权类型为管理员级别时，只允许填“*”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Column *string `json:"Column,omitempty" name:"Column"`

	// 需要授权的数据引擎，填*代表当前所有引擎。当授权类型为管理员级别时，只允许填“*”
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// 用户是否可以进行二次授权。当为true的时候，被授权的用户可以将本次获取的权限再次授权给其他子用户。默认为false
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReAuth *bool `json:"ReAuth,omitempty" name:"ReAuth"`

	// 权限来源，入参不填。USER：权限来自用户本身；WORKGROUP：权限来自绑定的工作组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 授权模式，入参不填。COMMON：普通模式；SENIOR：高级模式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 操作者，入参不填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 权限创建的时间，入参不填
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 权限所属工作组的ID，只有当该权限的来源为工作组时才会有值。即仅当Source字段的值为WORKGROUP时该字段才有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceId *int64 `json:"SourceId,omitempty" name:"SourceId"`

	// 权限所属工作组的名称，只有当该权限的来源为工作组时才会有值。即仅当Source字段的值为WORKGROUP时该字段才有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`
}

type Property struct {

	// 属性key名称。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 属性key对应的value。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SQLTask struct {

	// base64加密后的SQL语句
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// 任务的配置信息
	Config []*KVPair `json:"Config,omitempty" name:"Config"`
}

type Script struct {

	// 脚本Id，长度36字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptId *string `json:"ScriptId,omitempty" name:"ScriptId"`

	// 脚本名称，长度0-25。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptName *string `json:"ScriptName,omitempty" name:"ScriptName"`

	// 脚本描述，长度0-50。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptDesc *string `json:"ScriptDesc,omitempty" name:"ScriptDesc"`

	// 默认关联数据库。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// SQL描述，长度0-10000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLStatement *string `json:"SQLStatement,omitempty" name:"SQLStatement"`

	// 更新时间戳， 单位：ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type SparkJobInfo struct {

	// spark作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// spark作业名
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// spark作业类型，可去1或者2，1表示batch作业， 2表示streaming作业
	JobType *int64 `json:"JobType,omitempty" name:"JobType"`

	// 引擎名
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// 程序包是否本地上传，cos或者lakefs
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// 程序包路径
	JobFile *string `json:"JobFile,omitempty" name:"JobFile"`

	// 角色ID
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// spark作业运行主类
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// 命令行参数，spark作业命令行参数，空格分隔
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// spark原生配置，换行符分隔
	JobConf *string `json:"JobConf,omitempty" name:"JobConf"`

	// 依赖jars是否本地上传，cos或者lakefs
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// spark作业依赖jars，逗号分隔
	JobJars *string `json:"JobJars,omitempty" name:"JobJars"`

	// 依赖文件是否本地上传，cos或者lakefs
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// spark作业依赖文件，逗号分隔
	JobFiles *string `json:"JobFiles,omitempty" name:"JobFiles"`

	// spark作业driver资源大小
	JobDriverSize *string `json:"JobDriverSize,omitempty" name:"JobDriverSize"`

	// spark作业executor资源大小
	JobExecutorSize *string `json:"JobExecutorSize,omitempty" name:"JobExecutorSize"`

	// spark作业executor个数
	JobExecutorNums *int64 `json:"JobExecutorNums,omitempty" name:"JobExecutorNums"`

	// spark流任务最大重试次数
	JobMaxAttempts *int64 `json:"JobMaxAttempts,omitempty" name:"JobMaxAttempts"`

	// spark作业创建者
	JobCreator *string `json:"JobCreator,omitempty" name:"JobCreator"`

	// spark作业创建时间
	JobCreateTime *int64 `json:"JobCreateTime,omitempty" name:"JobCreateTime"`

	// spark作业更新时间
	JobUpdateTime *uint64 `json:"JobUpdateTime,omitempty" name:"JobUpdateTime"`

	// spark作业最近任务ID
	CurrentTaskId *string `json:"CurrentTaskId,omitempty" name:"CurrentTaskId"`

	// spark作业最近运行状态
	JobStatus *int64 `json:"JobStatus,omitempty" name:"JobStatus"`

	// spark流作业统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamingStat *StreamingStatistics `json:"StreamingStat,omitempty" name:"StreamingStat"`

	// 数据源名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// pyspark：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// pyspark：python依赖, 除py文件外，还支持zip/egg等归档格式，多文件以逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`
}

type StreamingStatistics struct {

	// 任务开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 数据接收器数
	Receivers *int64 `json:"Receivers,omitempty" name:"Receivers"`

	// 运行中的接收器数
	NumActiveReceivers *int64 `json:"NumActiveReceivers,omitempty" name:"NumActiveReceivers"`

	// 不活跃的接收器数
	NumInactiveReceivers *int64 `json:"NumInactiveReceivers,omitempty" name:"NumInactiveReceivers"`

	// 运行中的批数
	NumActiveBatches *int64 `json:"NumActiveBatches,omitempty" name:"NumActiveBatches"`

	// 待处理的批数
	NumRetainedCompletedBatches *int64 `json:"NumRetainedCompletedBatches,omitempty" name:"NumRetainedCompletedBatches"`

	// 已完成的批数
	NumTotalCompletedBatches *int64 `json:"NumTotalCompletedBatches,omitempty" name:"NumTotalCompletedBatches"`

	// 平均输入速率
	AverageInputRate *float64 `json:"AverageInputRate,omitempty" name:"AverageInputRate"`

	// 平均等待时长
	AverageSchedulingDelay *float64 `json:"AverageSchedulingDelay,omitempty" name:"AverageSchedulingDelay"`

	// 平均处理时长
	AverageProcessingTime *float64 `json:"AverageProcessingTime,omitempty" name:"AverageProcessingTime"`

	// 平均延时
	AverageTotalDelay *float64 `json:"AverageTotalDelay,omitempty" name:"AverageTotalDelay"`
}

type TableBaseInfo struct {

	// 该数据表所属数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据表名字
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 该数据表所属数据源名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 该数据表备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableComment *string `json:"TableComment,omitempty" name:"TableComment"`

	// 具体类型，表or视图
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据格式类型，hive，iceberg等
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableFormat *string `json:"TableFormat,omitempty" name:"TableFormat"`
}

type TableInfo struct {

	// 数据表配置信息。
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// 数据表格式。每次入参可选如下其一的KV结构，[TextFile，CSV，Json, Parquet, ORC, AVRD]。
	DataFormat *DataFormat `json:"DataFormat,omitempty" name:"DataFormat"`

	// 数据表列信息。
	Columns []*Column `json:"Columns,omitempty" name:"Columns"`

	// 数据表分块信息。
	Partitions []*Partition `json:"Partitions,omitempty" name:"Partitions"`

	// 数据存储路径。当前仅支持cos路径，格式如下：cosn://bucket-name/filepath。
	Location *string `json:"Location,omitempty" name:"Location"`
}

type TableResponseInfo struct {

	// 数据表基本信息。
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// 数据表列信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitempty" name:"Columns"`

	// 数据表分块信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*Partition `json:"Partitions,omitempty" name:"Partitions"`

	// 数据存储路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 数据表属性信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`

	// 数据表更新时间, 单位: ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 数据表创建时间,单位: ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 数据格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputFormat *string `json:"InputFormat,omitempty" name:"InputFormat"`

	// 数据表存储大小（单位：Byte）
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`

	// 数据表行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`
}

type Task struct {

	// SQL查询任务
	SQLTask *SQLTask `json:"SQLTask,omitempty" name:"SQLTask"`

	// Spark SQL查询任务
	SparkSQLTask *SQLTask `json:"SparkSQLTask,omitempty" name:"SparkSQLTask"`
}

type TaskResponseInfo struct {

	// 任务所属Database的名称。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 任务数据量。
	DataAmount *int64 `json:"DataAmount,omitempty" name:"DataAmount"`

	// 任务Id。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 计算时长，单位： ms。
	UsedTime *int64 `json:"UsedTime,omitempty" name:"UsedTime"`

	// 任务输出路径。
	OutputPath *string `json:"OutputPath,omitempty" name:"OutputPath"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务状态：0 初始化， 1 执行中， 2 执行成功，-1 执行失败，-3 已取消。
	State *int64 `json:"State,omitempty" name:"State"`

	// 任务SQL类型，DDL|DML等
	SQLType *string `json:"SQLType,omitempty" name:"SQLType"`

	// 任务SQL语句
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// 结果是否过期。
	ResultExpired *bool `json:"ResultExpired,omitempty" name:"ResultExpired"`

	// 数据影响统计信息。
	RowAffectInfo *string `json:"RowAffectInfo,omitempty" name:"RowAffectInfo"`

	// 任务结果数据表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSet *string `json:"DataSet,omitempty" name:"DataSet"`

	// 失败信息, 例如：errorMessage。该字段已废弃。
	Error *string `json:"Error,omitempty" name:"Error"`

	// 任务执行进度num/100(%)
	Percentage *int64 `json:"Percentage,omitempty" name:"Percentage"`

	// 任务执行输出信息。
	OutputMessage *string `json:"OutputMessage,omitempty" name:"OutputMessage"`

	// 执行SQL的引擎类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务进度明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgressDetail *string `json:"ProgressDetail,omitempty" name:"ProgressDetail"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 计算资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineId *string `json:"DataEngineId,omitempty" name:"DataEngineId"`

	// 执行sql的子uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 计算资源名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// 导入类型是本地导入还是cos
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 导入配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputConf *string `json:"InputConf,omitempty" name:"InputConf"`

	// 数据条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataNumber *int64 `json:"DataNumber,omitempty" name:"DataNumber"`

	// 查询数据能不能下载
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanDownload *bool `json:"CanDownload,omitempty" name:"CanDownload"`

	// 用户别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`

	// spark应用作业名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkJobName *string `json:"SparkJobName,omitempty" name:"SparkJobName"`

	// spark应用作业Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkJobId *string `json:"SparkJobId,omitempty" name:"SparkJobId"`

	// spark应用入口jar文件
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkJobFile *string `json:"SparkJobFile,omitempty" name:"SparkJobFile"`

	// spark ui url
	// 注意：此字段可能返回 null，表示取不到有效值。
	UiUrl *string `json:"UiUrl,omitempty" name:"UiUrl"`
}

type TaskResultInfo struct {

	// 任务唯一ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据源名称，当前任务执行时候选中的默认数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 数据库名称，当前任务执行时候选中的默认数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 当前执行的SQL，一个任务包含一个SQL
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// 执行任务的类型，现在分为DDL、DML、DQL
	SQLType *string `json:"SQLType,omitempty" name:"SQLType"`

	// 任务当前的状态，0：初始化 1：任务运行中 2：任务执行成功 -1：任务执行失败 -3：用户手动终止。只有任务执行成功的情况下，才会返回任务执行的结果
	State *int64 `json:"State,omitempty" name:"State"`

	// 扫描的数据量，单位byte
	DataAmount *int64 `json:"DataAmount,omitempty" name:"DataAmount"`

	// 任务执行耗时，单位秒
	UsedTime *int64 `json:"UsedTime,omitempty" name:"UsedTime"`

	// 任务结果输出的COS桶地址
	OutputPath *string `json:"OutputPath,omitempty" name:"OutputPath"`

	// 任务创建时间，时间戳
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务执行信息，成功时返回success，失败时返回失败原因
	OutputMessage *string `json:"OutputMessage,omitempty" name:"OutputMessage"`

	// 被影响的行数
	RowAffectInfo *string `json:"RowAffectInfo,omitempty" name:"RowAffectInfo"`

	// 结果的schema信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultSchema []*Column `json:"ResultSchema,omitempty" name:"ResultSchema"`

	// 结果信息，反转义后，外层数组的每个元素为一行数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultSet *string `json:"ResultSet,omitempty" name:"ResultSet"`

	// 分页信息，如果没有更多结果数据，nextToken为空
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 任务执行进度num/100(%)
	Percentage *int64 `json:"Percentage,omitempty" name:"Percentage"`

	// 任务进度明细
	ProgressDetail *string `json:"ProgressDetail,omitempty" name:"ProgressDetail"`

	// 控制台展示格式。table：表格展示 text：文本展示
	DisplayFormat *string `json:"DisplayFormat,omitempty" name:"DisplayFormat"`
}

type TasksInfo struct {

	// 任务类型，SQLTask：SQL查询任务。SparkSQLTask：Spark SQL查询任务
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 容错策略。Proceed：前面任务出错/取消后继续执行后面的任务。Terminate：前面的任务出错/取消之后终止后面任务的执行，后面的任务全部标记为已取消。
	FailureTolerance *string `json:"FailureTolerance,omitempty" name:"FailureTolerance"`

	// base64加密后的SQL语句，用";"号分隔每个SQL语句，一次最多提交50个任务。严格按照前后顺序执行
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// 任务的配置信息，当前仅支持SparkSQLTask任务。
	Config []*KVPair `json:"Config,omitempty" name:"Config"`

	// 任务的用户自定义参数信息
	Params []*KVPair `json:"Params,omitempty" name:"Params"`
}

type TextFile struct {

	// 文本类型，本参数取值为TextFile。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 处理文本用的正则表达式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regex *string `json:"Regex,omitempty" name:"Regex"`
}

type UnbindWorkGroupsFromUserRequest struct {
	*tchttp.BaseRequest

	// 解绑的工作组Id和用户Id的关联关系
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitempty" name:"AddInfo"`
}

func (r *UnbindWorkGroupsFromUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindWorkGroupsFromUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindWorkGroupsFromUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnbindWorkGroupsFromUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindWorkGroupsFromUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindWorkGroupsFromUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserIdSetOfWorkGroupId struct {

	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 用户Id集合，和CAM侧Uin匹配
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

type UserInfo struct {

	// 用户Id，和子用户uin相同
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户描述信息，方便区分不同用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDescription *string `json:"UserDescription,omitempty" name:"UserDescription"`

	// 单独给用户绑定的权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`

	// 当前用户的创建者
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 创建时间，格式如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 关联的工作组集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupSet []*WorkGroupMessage `json:"WorkGroupSet,omitempty" name:"WorkGroupSet"`

	// 是否是主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOwner *bool `json:"IsOwner,omitempty" name:"IsOwner"`

	// 用户类型。ADMIN：管理员 COMMON：普通用户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// 用户别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`
}

type UserMessage struct {

	// 用户Id，和CAM侧子用户Uin匹配
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDescription *string `json:"UserDescription,omitempty" name:"UserDescription"`

	// 当前用户的创建者
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 当前用户的创建时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 用户别名
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`
}

type ViewBaseInfo struct {

	// 该视图所属数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 视图名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
}

type ViewResponseInfo struct {

	// 视图基本信息。
	ViewBaseInfo *ViewBaseInfo `json:"ViewBaseInfo,omitempty" name:"ViewBaseInfo"`

	// 视图列信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitempty" name:"Columns"`

	// 视图属性信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`

	// 视图创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 视图更新时间。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type WorkGroupIdSetOfUserId struct {

	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 工作组Id集合
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitempty" name:"WorkGroupIds"`
}

type WorkGroupInfo struct {

	// 查询到的工作组唯一Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitempty" name:"WorkGroupName"`

	// 工作组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupDescription *string `json:"WorkGroupDescription,omitempty" name:"WorkGroupDescription"`

	// 工作组关联的用户数量
	UserNum *int64 `json:"UserNum,omitempty" name:"UserNum"`

	// 工作组关联的用户集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSet []*UserMessage `json:"UserSet,omitempty" name:"UserSet"`

	// 工作组绑定的权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`

	// 工作组的创建人
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 工作组的创建时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type WorkGroupMessage struct {

	// 工作组唯一Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitempty" name:"WorkGroupName"`

	// 工作组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupDescription *string `json:"WorkGroupDescription,omitempty" name:"WorkGroupDescription"`

	// 创建者
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 工作组创建的时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}
