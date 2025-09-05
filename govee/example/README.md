# Examples

`example.go` provides a sample program that you can run to list the device state for the devices associated with your Govee API Key.  Instructions on how to get your Govee API key can be found [here](https://developer.govee.com/reference/apply-you-govee-api-key).

## Usage
```sh
go build
export GOVEE_API_KEY=<Your API Key>
./example
```

## Output Example

```sh
Model: H5179
Device: 24:36:1C:4B:A6:74:C1:6B
Name: Living Room
Type: devices.types.thermometer
Capability Type: devices.capabilities.online
Instance: online
Value: true
Capability Type: devices.capabilities.property
Instance: sensorTemperature
Value: 69.98
Capability Type: devices.capabilities.property
Instance: sensorHumidity
Value: 58.1

------------------------------------------------------------------------------------------
Model: H7021
Device: 1E:B4:C7:31:38:32:1A:5C
Name: String Lights
Type: devices.types.light
Capability Type: devices.capabilities.online
Instance: online
Value: false
Capability Type: devices.capabilities.on_off
Instance: powerSwitch
Value: 0
Capability Type: devices.capabilities.toggle
Instance: gradientToggle
Value: 
Capability Type: devices.capabilities.range
Instance: brightness
Value: 50
Capability Type: devices.capabilities.segment_color_setting
Instance: segmentedBrightness
Value: 
Capability Type: devices.capabilities.segment_color_setting
Instance: segmentedColorRgb
Value: 
Capability Type: devices.capabilities.color_setting
Instance: colorRgb
Value: 107200
Capability Type: devices.capabilities.color_setting
Instance: colorTemperatureK
Value: 0
Capability Type: devices.capabilities.dynamic_scene
Instance: lightScene
Value: 
Capability Type: devices.capabilities.music_setting
Instance: musicMode
Value: 
Capability Type: devices.capabilities.dynamic_scene
Instance: diyScene
Value: 
```