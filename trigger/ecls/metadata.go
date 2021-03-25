package ecls

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Url string `md:"url,required"` // The websocket server to connect to
}
type HandlerSettings struct {
}

type Reply struct {
}

type Output struct {
	ApiKey                        string `md:"api_key,required"`
	ApiMethod                     string `md:"api_method_name"`
	Bytes                         string `md:"bytes,required"`
	CacheHit                      string `md:"cache_hit,required"`
	ClientTransferTime            string `md:"client_transfer_time,required"`
	ConnectTime                   string `md:"connect_time,required"`
	EndpointName                  string `md:"endpoint_name,required"`
	HttpMethod                    string `md:"http_method,required"`
	HttpStatus                    string `md:"http_status_code,required"`
	HttpVersion                   string `md:"http_version,required"`
	OauthAccesstoken              string `md:"oauth_access_token,required"`
	PackageName                   string `md:"package_name,required"`
	PackageUuid                   string `md:"package_uuid,required"`
	PlanName                      string `md:"plan_name,required"`
	PlanUuid                      string `md:"plan_uuid,required"`
	PreTransferTime               string `md:"pre_transfer_time,required"`
	QpsThrottleValue              string `md:"qps_throttle_value,required"`
	QuotaValue                    string `md:"quota_value,required"`
	Referrer                      string `md:"referrer,required"`
	RemoteTotalTime               string `md:"remote_total_time,required"`
	RequestHostName               string `md:"request_host_name,required"`
	RequestId                     string `md:"request_id,required"`
	RequestTime                   string `md:"request_time,required"`
	RequestUuid                   string `md:"request_uuid,required"`
	ResponseString                string `md:"response_string,required"`
	ServiceDefinitionEndpointUuid string `md:"service_definition_endpoint_uuid,required"`
	ServiceId                     string `md:"service_id,required"`
	ServiceName                   string `md:"service_name,required"`
	SrcIp                         string `md:"src_ip,required"`
	SslEnabled                    string `md:"ssl_enabled,required"`
	TotalRequestExecTime          string `md:"total_request_exec_time,required"`
	TrafficManager                string `md:"traffic_manager,required"`
	TrafficManagerErrorCode       string `md:"traffic_manager_error_code,required"`
	Uri                           string `md:"uri,required"`
	UserAgent                     string `md:"user_agent,required"`
	LogType                       string `md:"log_type,required"`
	IngestionTime                 string `md:"ingestion_time,required"`
	OrgUuid                       string `md:"org_uuid"`
	OrgName                       string `md:"org_name"`
	SubOrgUuid                    string `md:"sub_org_uuid"`
	SubOrgName                    string `md:"sub_org_name"`
	AsCsv                         string `md:"asCSV,required"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"api_key":                          o.ApiKey,
		"api_method_name":                  o.ApiMethod,
		"bytes":                            o.Bytes,
		"cache_hit":                        o.CacheHit,
		"client_transfer_time":             o.ClientTransferTime,
		"connect_time":                     o.ConnectTime,
		"endpoint_name":                    o.EndpointName,
		"http_method":                      o.HttpMethod,
		"http_status_code":                 o.HttpStatus,
		"http_version":                     o.HttpVersion,
		"oauth_access_token":               o.OauthAccesstoken,
		"package_name":                     o.PackageName,
		"package_uuid":                     o.PackageUuid,
		"plan_name":                        o.PlanName,
		"plan_uuid":                        o.PlanUuid,
		"pre_transfer_time":                o.PreTransferTime,
		"qps_throttle_value":               o.QpsThrottleValue,
		"quota_value":                      o.QuotaValue,
		"referrer":                         o.Referrer,
		"remote_total_time":                o.RemoteTotalTime,
		"request_host_name":                o.RequestHostName,
		"request_id":                       o.RequestId,
		"request_time":                     o.RequestTime,
		"request_uuid":                     o.RequestUuid,
		"response_string":                  o.ResponseString,
		"service_definition_endpoint_uuid": o.ServiceDefinitionEndpointUuid,
		"service_id":                       o.ServiceId,
		"service_name":                     o.ServiceName,
		"src_ip":                           o.SrcIp,
		"ssl_enabled":                      o.SslEnabled,
		"total_request_exec_time":          o.TotalRequestExecTime,
		"traffic_manager":                  o.TrafficManager,
		"traffic_manager_error_code":       o.TrafficManagerErrorCode,
		"uri":                              o.Uri,
		"user_agent":                       o.UserAgent,
		"log_type":                         o.LogType,
		"ingestion_time":                   o.IngestionTime,
		"org_uuid":                         o.OrgUuid,
		"org_name":                         o.OrgName,
		"sub_org_uuid":                     o.SubOrgUuid,
		"sub_org_name":                     o.SubOrgName,
		"asCSV":                            o.AsCsv,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.ApiKey, err = coerce.ToString(values["api_key"])
	if err != nil {
		return err
	}
	o.ApiMethod, err = coerce.ToString(values["api_method_name"])
	if err != nil {
		return err
	}
	o.Bytes, err = coerce.ToString(values["bytes"])
	if err != nil {
		return err
	}
	o.CacheHit, err = coerce.ToString(values["cache_hit"])
	if err != nil {
		return err
	}
	o.ClientTransferTime, err = coerce.ToString(values["client_transfer_time"])
	if err != nil {
		return err
	}
	o.ConnectTime, err = coerce.ToString(values["connect_time"])
	if err != nil {
		return err
	}
	o.EndpointName, err = coerce.ToString(values["endpoint_name"])
	if err != nil {
		return err
	}
	o.HttpMethod, err = coerce.ToString(values["http_method"])
	if err != nil {
		return err
	}
	o.HttpStatus, err = coerce.ToString(values["http_status_code"])
	if err != nil {
		return err
	}
	o.HttpVersion, err = coerce.ToString(values["http_version"])
	if err != nil {
		return err
	}
	o.OauthAccesstoken, err = coerce.ToString(values["oauth_access_token"])
	if err != nil {
		return err
	}
	o.PackageName, err = coerce.ToString(values["package_name"])
	if err != nil {
		return err
	}
	o.PackageUuid, err = coerce.ToString(values["package_uuid"])
	if err != nil {
		return err
	}
	o.PlanName, err = coerce.ToString(values["plan_name"])
	if err != nil {
		return err
	}
	o.PlanUuid, err = coerce.ToString(values["plan_uuid"])
	if err != nil {
		return err
	}
	o.PreTransferTime, err = coerce.ToString(values["pre_transfer_time"])
	if err != nil {
		return err
	}
	o.QpsThrottleValue, err = coerce.ToString(values["qps_throttle_value"])
	if err != nil {
		return err
	}
	o.QuotaValue, err = coerce.ToString(values["quota_value"])
	if err != nil {
		return err
	}
	o.Referrer, err = coerce.ToString(values["referrer"])
	if err != nil {
		return err
	}
	o.RemoteTotalTime, err = coerce.ToString(values["remote_total_time"])
	if err != nil {
		return err
	}
	o.RequestHostName, err = coerce.ToString(values["request_host_name"])
	if err != nil {
		return err
	}
	o.RequestId, err = coerce.ToString(values["request_id"])
	if err != nil {
		return err
	}
	o.RequestTime, err = coerce.ToString(values["request_time"])
	if err != nil {
		return err
	}
	o.RequestUuid, err = coerce.ToString(values["request_uuid"])
	if err != nil {
		return err
	}
	o.ResponseString, err = coerce.ToString(values["response_string"])
	if err != nil {
		return err
	}
	o.ServiceDefinitionEndpointUuid, err = coerce.ToString(values["service_definition_endpoint_uuid"])
	if err != nil {
		return err
	}
	o.ServiceId, err = coerce.ToString(values["service_id"])
	if err != nil {
		return err
	}
	o.ServiceName, err = coerce.ToString(values["service_name"])
	if err != nil {
		return err
	}
	o.SrcIp, err = coerce.ToString(values["src_ip"])
	if err != nil {
		return err
	}
	o.SslEnabled, err = coerce.ToString(values["ssl_enabled"])
	if err != nil {
		return err
	}
	o.TotalRequestExecTime, err = coerce.ToString(values["total_request_exec_time"])
	if err != nil {
		return err
	}
	o.TrafficManager, err = coerce.ToString(values["traffic_manager"])
	if err != nil {
		return err
	}
	o.TrafficManagerErrorCode, err = coerce.ToString(values["traffic_manager_error_code"])
	if err != nil {
		return err
	}
	o.Uri, err = coerce.ToString(values["uri"])
	if err != nil {
		return err
	}
	o.UserAgent, err = coerce.ToString(values["user_agent"])
	if err != nil {
		return err
	}
	o.LogType, err = coerce.ToString(values["log_type"])
	if err != nil {
		return err
	}
	o.IngestionTime, err = coerce.ToString(values["ingestion_time"])
	if err != nil {
		return err
	}
	o.OrgUuid, err = coerce.ToString(values["org_uuid"])
	if err != nil {
		return err
	}
	o.OrgName, err = coerce.ToString(values["org_name"])
	if err != nil {
		return err
	}
	o.SubOrgUuid, err = coerce.ToString(values["sub_org_uuid"])
	if err != nil {
		return err
	}
	o.SubOrgName, err = coerce.ToString(values["sub_org_name"])
	if err != nil {
		return err
	}
	o.AsCsv, err = coerce.ToString(values["asCSV"])
	if err != nil {
		return err
	}

	return nil
}
