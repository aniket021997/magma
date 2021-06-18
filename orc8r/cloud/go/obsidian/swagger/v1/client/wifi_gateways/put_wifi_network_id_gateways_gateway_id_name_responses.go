// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutWifiNetworkIDGatewaysGatewayIDNameReader is a Reader for the PutWifiNetworkIDGatewaysGatewayIDName structure.
type PutWifiNetworkIDGatewaysGatewayIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWifiNetworkIDGatewaysGatewayIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutWifiNetworkIDGatewaysGatewayIDNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutWifiNetworkIDGatewaysGatewayIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDNameNoContent creates a PutWifiNetworkIDGatewaysGatewayIDNameNoContent with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDNameNoContent() *PutWifiNetworkIDGatewaysGatewayIDNameNoContent {
	return &PutWifiNetworkIDGatewaysGatewayIDNameNoContent{}
}

/*PutWifiNetworkIDGatewaysGatewayIDNameNoContent handles this case with default header values.

Success
*/
type PutWifiNetworkIDGatewaysGatewayIDNameNoContent struct {
}

func (o *PutWifiNetworkIDGatewaysGatewayIDNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}/name][%d] putWifiNetworkIdGatewaysGatewayIdNameNoContent ", 204)
}

func (o *PutWifiNetworkIDGatewaysGatewayIDNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWifiNetworkIDGatewaysGatewayIDNameDefault creates a PutWifiNetworkIDGatewaysGatewayIDNameDefault with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDNameDefault(code int) *PutWifiNetworkIDGatewaysGatewayIDNameDefault {
	return &PutWifiNetworkIDGatewaysGatewayIDNameDefault{
		_statusCode: code,
	}
}

/*PutWifiNetworkIDGatewaysGatewayIDNameDefault handles this case with default header values.

Unexpected Error
*/
type PutWifiNetworkIDGatewaysGatewayIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put wifi network ID gateways gateway ID name default response
func (o *PutWifiNetworkIDGatewaysGatewayIDNameDefault) Code() int {
	return o._statusCode
}

func (o *PutWifiNetworkIDGatewaysGatewayIDNameDefault) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}/name][%d] PutWifiNetworkIDGatewaysGatewayIDName default  %+v", o._statusCode, o.Payload)
}

func (o *PutWifiNetworkIDGatewaysGatewayIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutWifiNetworkIDGatewaysGatewayIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}