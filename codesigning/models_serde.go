package codesigning

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetail.
func (e *ErrorDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
				err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "message":
				err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "target":
				err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "errorDetail", e.ErrorDetail)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "errorDetail":
				err = unpopulate(val, "ErrorDetail", &e.ErrorDetail)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SignRequest.
func (s SignRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "authenticodeHashList", s.AuthenticodeHashList, func() any {
		encodedValue := make([]string, len(s.AuthenticodeHashList))
		for i := 0; i < len(s.AuthenticodeHashList); i++ {
			encodedValue[i] = runtime.EncodeByteArray(s.AuthenticodeHashList[i], runtime.Base64StdFormat)
		}
		return encodedValue
	})
	populateByteArray(objectMap, "digest", s.Digest, func() any {
		return runtime.EncodeByteArray(s.Digest, runtime.Base64StdFormat)
	})
	populateByteArray(objectMap, "fileHashList", s.FileHashList, func() any {
		encodedValue := make([]string, len(s.FileHashList))
		for i := 0; i < len(s.FileHashList); i++ {
			encodedValue[i] = runtime.EncodeByteArray(s.FileHashList[i], runtime.Base64StdFormat)
		}
		return encodedValue
	})
	populate(objectMap, "signatureAlgorithm", s.SignatureAlgorithm)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SignRequest.
func (s *SignRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authenticodeHashList":
			var encodedValue []string
			err = unpopulate(val, "AuthenticodeHashList", &encodedValue)
			if err == nil && len(encodedValue) > 0 {
				s.AuthenticodeHashList = make([][]byte, len(encodedValue))
				for i := 0; i < len(encodedValue) && err == nil; i++ {
					err = runtime.DecodeByteArray(encodedValue[i], &s.AuthenticodeHashList[i], runtime.Base64StdFormat)
				}
			}
			delete(rawMsg, key)
		case "digest":
		if val != nil && string(val) != "null" {
				err = runtime.DecodeByteArray(string(val), &s.Digest, runtime.Base64StdFormat)
		}
			delete(rawMsg, key)
		case "fileHashList":
			var encodedValue []string
			err = unpopulate(val, "FileHashList", &encodedValue)
			if err == nil && len(encodedValue) > 0 {
				s.FileHashList = make([][]byte, len(encodedValue))
				for i := 0; i < len(encodedValue) && err == nil; i++ {
					err = runtime.DecodeByteArray(encodedValue[i], &s.FileHashList[i], runtime.Base64StdFormat)
				}
			}
			delete(rawMsg, key)
		case "signatureAlgorithm":
				err = unpopulate(val, "SignatureAlgorithm", &s.SignatureAlgorithm)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SignStatus.
func (s SignStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "operationId", s.OperationID)
	populateByteArray(objectMap, "signature", s.Signature, func() any {
		return runtime.EncodeByteArray(s.Signature, runtime.Base64StdFormat)
	})
	populateByteArray(objectMap, "signingCertificate", s.SigningCertificate, func() any {
		return runtime.EncodeByteArray(s.SigningCertificate, runtime.Base64StdFormat)
	})
	populate(objectMap, "status", s.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SignStatus.
func (s *SignStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "operationId":
				err = unpopulate(val, "OperationID", &s.OperationID)
			delete(rawMsg, key)
		case "signature":
		if val != nil && string(val) != "null" {
				err = runtime.DecodeByteArray(string(val), &s.Signature, runtime.Base64StdFormat)
		}
			delete(rawMsg, key)
		case "signingCertificate":
		if val != nil && string(val) != "null" {
				err = runtime.DecodeByteArray(string(val), &s.SigningCertificate, runtime.Base64StdFormat)
		}
			delete(rawMsg, key)
		case "status":
				err = unpopulate(val, "Status", &s.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateByteArray[T any](m map[string]any, k string, b []T, convert func() any) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = convert()
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}

