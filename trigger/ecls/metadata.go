package ecls

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Url string `md:"url,required"` // The websocket server to connect to
}
type HandlerSettings struct {
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
	AsCsv                         string `md:"asCSV,required"`
	AsObject                      string `md:"asObject,required"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ApiKey":                        o.ApiKey,
		"ApiMethod":                     o.ApiMethod,
		"Bytes":                         o.Bytes,
		"CacheHit":                      o.CacheHit,
		"ClientTransferTime":            o.ClientTransferTime,
		"ConnectTime":                   o.ConnectTime,
		"EndpointName":                  o.EndpointName,
		"HttpMethod":                    o.HttpMethod,
		"HttpStatus":                    o.HttpStatus,
		"HttpVersion":                   o.HttpVersion,
		"OauthAccesstoken":              o.OauthAccesstoken,
		"PackageName":                   o.PackageName,
		"PackageUuid":                   o.PackageUuid,
		"PlanName":                      o.PlanName,
		"PlanUuid":                      o.PlanUuid,
		"PreTransferTime":               o.PreTransferTime,
		"QpsThrottleValue":              o.QpsThrottleValue,
		"QuotaValue":                    o.QuotaValue,
		"Referrer":                      o.Referrer,
		"RemoteTotalTime":               o.RemoteTotalTime,
		"RequestHostName":               o.RequestHostName,
		"RequestId":                     o.RequestId,
		"RequestTime":                   o.RequestTime,
		"RequestUuid":                   o.RequestUuid,
		"ResponseString":                o.ResponseString,
		"ServiceDefinitionEndpointUuid": o.ServiceDefinitionEndpointUuid,
		"ServiceId":                     o.ServiceId,
		"ServiceName":                   o.ServiceName,
		"SrcIp":                         o.SrcIp,
		"SslEnabled":                    o.SslEnabled,
		"TotalRequestExecTime":          o.TotalRequestExecTime,
		"TrafficManager":                o.TrafficManager,
		"TrafficManagerErrorCode":       o.TrafficManagerErrorCode,
		"Uri":                           o.Uri,
		"UserAgent":                     o.UserAgent,
		"LogType":                       o.LogType,
		"IngestionTime":                 o.IngestionTime,
		"AsCsv":                         o.AsCsv,
		"AsObject":                      o.AsObject,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.ApiKey, err = coerce.ToString(values["ApiKey"])
	o.ApiMethod, err = coerce.ToString(values["ApiMethod"])
	o.Bytes, err = coerce.ToString(values["Bytes"])
	o.CacheHit, err = coerce.ToString(values["CacheHit"])
	o.ClientTransferTime, err = coerce.ToString(values["ClientTransferTime"])
	o.ConnectTime, err = coerce.ToString(values["ConnectTime"])
	o.EndpointName, err = coerce.ToString(values["EndpointName"])
	o.HttpMethod, err = coerce.ToString(values["HttpMethod"])
	o.HttpStatus, err = coerce.ToString(values["HttpStatus"])
	o.HttpVersion, err = coerce.ToString(values["HttpVersion"])
	o.OauthAccesstoken, err = coerce.ToString(values["OauthAccesstoken"])
	o.PackageName, err = coerce.ToString(values["PackageName"])
	o.PackageUuid, err = coerce.ToString(values["PackageUuid"])
	o.PlanName, err = coerce.ToString(values["PlanName"])
	o.PlanUuid, err = coerce.ToString(values["PlanUuid"])
	o.PreTransferTime, err = coerce.ToString(values["PreTransferTime"])
	o.QpsThrottleValue, err = coerce.ToString(values["QpsThrottleValue"])
	o.QuotaValue, err = coerce.ToString(values["QuotaValue"])
	o.Referrer, err = coerce.ToString(values["Referrer"])
	o.RemoteTotalTime, err = coerce.ToString(values["RemoteTotalTime"])
	o.RequestHostName, err = coerce.ToString(values["RequestHostName"])
	o.RequestId, err = coerce.ToString(values["RequestId"])
	o.RequestTime, err = coerce.ToString(values["RequestTime"])
	o.RequestUuid, err = coerce.ToString(values["RequestUuid"])
	o.ResponseString, err = coerce.ToString(values["ResponseString"])
	o.ServiceDefinitionEndpointUuid, err = coerce.ToString(values["ServiceDefinitionEndpointUuid"])
	o.ServiceId, err = coerce.ToString(values["ServiceId"])
	o.ServiceName, err = coerce.ToString(values["ServiceName"])
	o.SrcIp, err = coerce.ToString(values["SrcIp"])
	o.SslEnabled, err = coerce.ToString(values["SslEnabled"])
	o.TotalRequestExecTime, err = coerce.ToString(values["TotalRequestExecTime"])
	o.TrafficManager, err = coerce.ToString(values["TrafficManager"])
	o.TrafficManagerErrorCode, err = coerce.ToString(values["TrafficManagerErrorCode"])
	o.Uri, err = coerce.ToString(values["Uri"])
	o.UserAgent, err = coerce.ToString(values["UserAgent"])
	o.LogType, err = coerce.ToString(values["LogType"])
	o.IngestionTime, err = coerce.ToString(values["IngestionTime"])
	o.AsCsv, err = coerce.ToString(values["AsCsv"])
	o.AsObject, err = coerce.ToString(values["AsObject"])
	if err != nil {
		return err
	}

	return nil
}
