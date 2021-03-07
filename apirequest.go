package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

//NewRequest create and prepare new RPC Request object
func NewRequest() (req *ApiRequest) {
	req = new(ApiRequest)
	req.Init()
	return
}

//SetMethod set JSON RPC Request field
func (r *ApiRequest) SetMethod(method string) *ApiRequest {
	r.Method = method
	return r
}

func (r *ApiRequest) Init() {
	if r.initComplete {
		return
	}
	r.Jsonrpc = JSON_RPC_VERSION
	r.Params = make(map[string]interface{})
	r.initComplete = true
}

// Client side methods
//
// Prepare JSO-RPC Request params
// *******************

//SetParam set param for 'key' field. Please note, value must be jsonable
func (r *ApiRequest) SetParam(key string, value interface{}) *ApiRequest {
	r.Params[key] = value
	return r
}

//SetParamBool set bool value for 'key' field
func (r *ApiRequest) SetParamBool(key string, value bool) *ApiRequest {
	r.Params[key] = value
	return r
}

//SetParamString set string value for 'key' field
func (r *ApiRequest) SetParamString(key string, value string) *ApiRequest {
	r.Params[key] = value
	return r
}

//SetParamInt set int value for 'key' field
func (r *ApiRequest) SetParamInt(key string, value int) *ApiRequest {
	r.Params[key] = value
	return r
}

//SetParamInt64 set int64 value for 'key' field
func (r *ApiRequest) SetParamInt64(key string, value int64) *ApiRequest {
	r.Params[key] = value
	return r
}

//SetParamFloat64 set float64 value for 'key' field
func (r *ApiRequest) SetParamFloat64(key string, value float64) *ApiRequest {
	r.Params[key] = value
	return r
}

// Server side methods
//
// Method's for get param's from request

//GetParamsCount get total param's count
func (r *ApiRequest) GetParamsCount() int {
	return len(r.Params)
}

//IsParamFound check is key found
func (r *ApiRequest) IsParamFound(key string) (found bool) {
	_, found = r.Params[key]
	return
}


//GetParamBytes get param as []byte slice
func (r *ApiRequest) GetParamBytes(key string) (param []byte, found bool) {
	paramRaw, found := r.Params[key]
	if !found {
		return []byte{}, false
	}
	param, _ = json.Marshal(paramRaw)
	return
}

//GetParamString get 'key' param as string or error if param not set
func (r *ApiRequest) GetParamString(key string) (param string, err error) {
	paramRaw, found := r.getParamRaw(key)
	if !found {
		return "", ErrParamNotFound
	}
	switch paramRaw.(type) {
	case string:
		param = paramRaw.(string)
		return
	default:
		param = fmt.Sprintf("%v", paramRaw)
	}
	return
}

//GetParamString get 'key' param as string or default value, if key not found in request
func (r *ApiRequest) GetParamStringDefault(key, defaultString string) (param string) {
	paramRaw, found := r.getParamRaw(key)
	if !found {
		return defaultString
	}
	switch paramRaw.(type) {
	case string:
		param = paramRaw.(string)
		return
	default:
		param = fmt.Sprintf("%v", paramRaw)
	}
	return
}

//GetParamBool get 'key' param as bool value or error if param not set
func (r *ApiRequest) GetParamBool(key string) (param bool, err error) {
	paramRaw, found := r.getParamRaw(key)
	if !found {
		return false, ErrParamNotFound
	}
	switch paramRaw.(type) {
	case bool:
		param = paramRaw.(bool)
		return
	default:
		switch paramRaw.(type) {
		case string:
			paramStr := paramRaw.(string)
			if paramStr == "true" || paramStr == "1" {
				param = true
			}
		case json.Number:
			paramStr := string(paramRaw.(json.Number))
			if paramStr == "true" || paramStr == "1" {
				param = true
			}
		case float64:
			paramNumeric := int(paramRaw.(float64))
			if paramNumeric > 0 {
				param = true
			}
		}
	}
	return
}

//GetParamBoolDefault get 'key' param as bool value or default value, if key not found in request
func (r *ApiRequest) GetParamBoolDefault(key string, defaultValue bool) (param bool) {
	paramRaw, found := r.getParamRaw(key)
	if !found {
		return defaultValue
	}
	switch paramRaw.(type) {
	case bool:
		param = paramRaw.(bool)
		return
	default:
		switch paramRaw.(type) {
		case string:
			paramStr := paramRaw.(string)
			if paramStr == "true" || paramStr == "1" {
				param = true
			}
		case json.Number:
			paramStr := string(paramRaw.(json.Number))
			if paramStr == "true" || paramStr == "1" {
				param = true
			}
		case float64:
			paramNumeric := int(paramRaw.(float64))
			if paramNumeric > 0 {
				param = true
			}
		}
	}
	return
}

func (r *ApiRequest) GetParamInt(key string) (param int, err error) {
	paramRaw, found := r.getParamRaw(key)
	if !found {
		return 0, ErrParamNotFound
	}
	if paramRaw == nil {
		param = 0
		return
	}
	switch paramRaw.(type) {
	case int:
		param = paramRaw.(int)
		return
	case int64:
		param = int(paramRaw.(int64))
		return
	case float64:
		param = int(paramRaw.(float64))
		return
	case string:
		paramStr := paramRaw.(string)
		return r.parseStringToInt(paramStr)
	case json.Number:
		paramStr := string(paramRaw.(json.Number))
		return r.parseStringToInt(paramStr)
	default:
		return 0, ErrInvalidParamType
	}
	return
}

func (r *ApiRequest) GetParamInt64(key string) (param int64, err error) {
	paramRaw, found := r.getParamRaw(key)
	if !found {
		return 0, ErrParamNotFound
	}
	if paramRaw == nil {
		param = 0
		return
	}
	switch paramRaw.(type) {
	case int64:
		param = paramRaw.(int64)
		return
	case float64:
		param = int64(paramRaw.(float64))
		return
	case string:
		paramStr := paramRaw.(string)
		return r.parseStringToInt64(paramStr)
	case json.Number:
		paramStr := string(paramRaw.(json.Number))
		return r.parseStringToInt64(paramStr)
	default:
		return 0, ErrInvalidParamType
	}
	return
}

func (r *ApiRequest) GetParamFloat64(key string) (param float64, err error) {
	paramRaw, found := r.getParamRaw(key)
	if !found {
		return 0, ErrParamNotFound
	}
	if paramRaw == nil {
		param = 0
		return
	}
	switch paramRaw.(type) {
	case float64:
		param = paramRaw.(float64)
		return
	case int64:
		param = float64(paramRaw.(int64))
		return
	case string:
		paramStr := paramRaw.(string)
		return r.parseStringToFloat64(paramStr)
	case json.Number:
		paramStr := string(paramRaw.(json.Number))
		return r.parseStringToFloat64(paramStr)
	default:
		return 0, ErrInvalidParamType
	}
	return
}

func (r *ApiRequest) ParamKeysWalk(process func(paramKey string)) {
	for key, _ := range r.Params {
		process(key)
	}
}

func (r *ApiRequest) GetParamToObject(key string, object interface{}) (found bool, err error) {
	paramRaw, found := r.Params[key]
	if !found {
		return false, nil
	}
	paramBytes, err := json.Marshal(paramRaw)
	if err != nil {
		return
	}
	err = json.Unmarshal(paramBytes, object)
	return
}

func (r *ApiRequest) ParseParamsToObject(object interface{}) (err error) {
	paramBytes, err := json.Marshal(r.Params)
	if err != nil {
		return
	}
	err = json.Unmarshal(paramBytes, object)
	return
}



func (r *ApiRequest) getParamRaw(key string) (param interface{}, found bool) {
	param, found = r.Params[key]
	return
}

func (r *ApiRequest) parseStringToInt64(param string) (num int64, err error) {
	return strconv.ParseInt(param, 0, 64)
}

func (r *ApiRequest) parseStringToInt(param string) (num int, err error) {
	return strconv.Atoi(param)
}

func (r *ApiRequest) parseStringToFloat64(param string) (num float64, err error) {
	return strconv.ParseFloat(param, 64)
}

func (r *ApiRequest) Json() (data []byte, err error) {
	return json.Marshal(r)
}

func (r *ApiRequest) JsonBytes() (data []byte) {
	data, _ = json.Marshal(r)
	return
}

func (r *ApiRequest) Unjson(data []byte) (err error) {
	d := json.NewDecoder(bytes.NewBuffer(data))
	d.UseNumber()
	err = d.Decode(r)
	return
}

//todo Custom marshal/unmarshal ID
type Number string
