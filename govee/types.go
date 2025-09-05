/*
 The MIT License

 Permission is hereby granted, free of charge, to any person obtaining a copy
 of this software and associated documentation files (the "Software"), to deal
 in the Software without restriction, including without limitation the rights
 to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 copies of the Software, and to permit persons to whom the Software is
 furnished to do so, subject to the following conditions:

 The above copyright notice and this permission notice shall be included in
 all copies or substantial portions of the Software.

 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 THE SOFTWARE.
*/

// Package govee provides a client for Govee APIs
package govee

import (
	"context"
	"net/http"
	"runtime"
)

const (
	version   = "1.0.0"
	userAgent = "govee-go/" + version + " (" + runtime.GOOS + "; " + runtime.GOARCH + ")"

	//Base URL for Govee API
	goveeApiUrl = "https://openapi.api.govee.com/router/api/v1"
	// API Endpoints
	goveeDevices       = "user/devices"
	goveeDeviceState   = "device/state"
	goveeDeviceControl = "device/control"
	goveeDeviceScenes  = "device/scenes"

	//Environment Variables
	envGoveeApiKey = "GOVEE_API_KEY"

	requestLimit = 10000 //10,000 request per day
)

// ClientConfig holds the parameters for creating a new client.
// The only mandatory field is the ApiKey field
type ClientConfig struct {
}

// Client implements the Govee client
type Client struct {
	// Token holds the authorization token for the API.
	// Information on obtaining an API Key can be found at https://developer.govee.com/reference/apply-you-govee-api-key
	ApiKey string

	// HTTPClient is used to make API requests.
	//
	// This can be used to specify a custom TLS configuration
	// (TLSClientConfig), a custom request timeout (Timeout),
	// or other customization as required.
	//
	// It HTTPClient is nil, http.DefaultClient will be used.
	HTTPClient *http.Client

	//Context for the client
	Ctx context.Context
}

// DevicesType represents the type of Govee device
type DevicesType string

// Device types from Govee API documentation
const (
	TypeLight        DevicesType = "devices.types.light"
	TypeAirPurifier  DevicesType = "devices.types.air_purifier"
	TypeThermometer  DevicesType = "devices.types.thermometer"
	TypeSocket       DevicesType = "devices.types.socket"
	TypeSensor       DevicesType = "devices.types.sensor"
	TypeHeater       DevicesType = "devices.types.heater"
	TypeHumidifier   DevicesType = "devices.types.humidifier"
	TypeDehumidifier DevicesType = "devices.types.dehumidifier"
	TypeIceMaker     DevicesType = "devices.types.ice_maker"
	TypeDiffuser     DevicesType = "devices.types.aroma_diffuser"
	TypeBox          DevicesType = "devices.types.box"
)

// DevicesCapabilities represents the capabilities of a Govee device
type DevicesCapabilities string

const (
	// CapabilitiesOnOff represents a powerSwitch with on/off enum options
	CapabilitiesOnOff DevicesCapabilities = "devices.capabilities.on_off"
	// CapabilitiesToggle represents an oscillationToggle,nightlightToggle,gradientToggle,ect with on/off Enum options
	CapabilitiesToggle DevicesCapabilities = "devices.capabilities.toggle"
	// CapabilitiesRange represents brightness,humidity,volume,temperature,ect to set a range number
	CapabilitiesRange DevicesCapabilities = "devices.capabilities.range"
	// CapabilitiesMode represents a nightlightScene,presetScene,gearMode,fanSpeed,ect with enum options
	CapabilitiesMode DevicesCapabilities = "devices.capabilities.mode"
	// CapabilitiesColorSetting represents a colorRgb,colorTemperatureK to set rgb or Kelvin color temperature
	CapabilitiesColorSetting DevicesCapabilities = "devices.capabilities.color_setting"
	// CapabilitiesSegmentColorSetting represents a segmentedBrightness,segmentedColorRgb to set color or brightness on segment
	CapabilitiesSegmentColorSetting DevicesCapabilities = "devices.capabilities.segment_color_setting"
	// CapabilitiesMusicSetting represents a musicMode	set music mode
	CapabilitiesMusicSetting DevicesCapabilities = "devices.capabilities.music_setting"
	// CapabilitiesDynamicScene represents a lightScene,diyScene,snapshot to set the scene,but the options are not static
	CapabilitiesDynamicScene DevicesCapabilities = "devices.capabilities.dynamic_scene"
	// CapabilitiesWorkMode represents a workMode to set the working mode and give it a working value
	CapabilitiesWorkMode DevicesCapabilities = "device.capabilities.work_mode"
	// CapabilitiesTemperatureSetting represents a targetTemperature,sliderTemperature to set the temperature
	CapabilitiesTemperatureSetting DevicesCapabilities = "device.capabilities.temperature_setting"
)

// DataType represents the type of data for a capability parameter or field
type DataType string

const (
	Enum    DataType = "ENUM"
	Integer DataType = "INTEGER"
	Struct  DataType = "STRUCT"
	Array   DataType = "Array"
)

// ApiResponseStatus represents the common fields in all API responses
type ApiResponseStatus struct {
	RequestId string `json:"requestId,omitempty"`
	Code      int    `json:"code"`
	Message   string `json:"msg"`
}

// Option represents an option for a capability parameter or field
type Option struct {
	Name  string `json:"name,omitempty"`
	Value int    `json:"value"`
}

// Range represents the range of values for a capability parameter or field
type Range struct {
	Min       int `json:"min"`
	Max       int `json:"max"`
	Precision int `json:"precision"`
}

// Field represents a field in a STRUCT parameter or capability
type Field struct {
	FieldName string   `json:"fieldName"`
	DataType  DataType `json:"dataType"`
	Options   []Option `json:"options,omitempty"`
	Range     Range    `json:"range,omitempty"`
	Required  bool     `json:"required,omitempty"`
}

// Parameter represents a parameter for a capability
type Parameter struct {
	Unit     string   `json:"unit"`
	DataType DataType `json:"dataType"`
	Options  []Option `json:"options,omitempty"`
	Range    Range    `json:"range,omitempty"`
	Fields   []Field  `json:"fields,omitempty"`
}

// State represents the state of a capability
type State struct {
	Value interface{} `json:"value,omitempty"`
}

// Capability represents a capability of a Govee device
type Capability struct {
	Type       DevicesType `json:"type"`
	Instance   string      `json:"instance"`
	Parameters Parameter   `json:"parameters,omitempty"`
	State      State       `json:"state,omitempty"`
}

// DeviceInfo represents a Govee device
type DeviceInfo struct {
	SKU          string       `json:"sku"`
	Device       string       `json:"device"`
	DeviceName   string       `json:"deviceName"`
	Type         string       `json:"type"`
	Capabilities []Capability `json:"capabilities"`
}

// DiscoveryApiResponse represents the response from the /devices endpoint
type DiscoveryApiResponse struct {
	ApiResponseStatus
	Data []DeviceInfo `json:"data"`
}

// DeviceStateApiResponse represents the response from the /device/state endpoint
type DeviceStateApiResponse struct {
	ApiResponseStatus
	Payload *DeviceInfo `json:"payload"`
}

// DeviceStatePayload represents the payload for the /device/state endpoint
type DeviceStatePayload struct {
	SKU    string `json:"sku"`
	Device string `json:"device"`
}

// RequestPayload represents the request payload for the /device/state endpoint
type RequestPayload struct {
	RequestId string             `json:"requestId"`
	Payload   DeviceStatePayload `json:"payload"`
}
