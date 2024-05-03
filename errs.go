package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	E100 = "E-100"
	E400 = "E-400"
	E300 = "E-300"
	E301 = "E-301"
	E302 = "E-302"
	E303 = "E-303"
	E404 = "E-404"
	E409 = "E-409"
	E429 = "E-429"
	E500 = "E-500"
)

type AppError struct {
	StatusCode string      `json:"status_code,omitempty"`
	ErrorCode  string      `json:"error_code"`
	Message    string      `json:"message"`
	Timestamp  time.Time   `json:"timestamp,omitempty"`
	Details    interface{} `json:"details,omitempty"`
}

const (
	ErrBadRequest           = "BAD_REQUEST"
	ErrUnauthorized         = "UNAUTHORIZED"
	ErrNotFound             = "NOT_FOUND"
	ErrInternalServerError  = "INTERNAL_SERVER_ERROR"
	ErrForbidden            = "FORBIDDEN"
	ErrInvalidInput         = "INVALID_INPUT"
	ErrDatabaseError        = "DATABASE_ERROR"
	ErrFileError            = "FILE_ERROR"
	ErrAuthenticationFailed = "AUTHENTICATION_FAILED"
	ErrValidationFailed     = "VALIDATION_FAILED"
	ErrEmailError           = "EMAIL_ERROR"
	ErrInvalidPassword      = "INVALID_PASSWORD"
	ErrPaymentError         = "PAYMENT_ERROR"
	ErrNetworkError         = "NETWORK_ERROR"
	ErrTimeout              = "TIMEOUT"
	ErrRateLimitExceeded    = "RATE_LIMIT_EXCEEDED"
	ErrServiceUnavailable   = "SERVICE_UNAVAILABLE"
	ErrConflict             = "CONFLICT"
	ErrUnprocessableEntity  = "UNPROCESSABLE_ENTITY"
	ErrTooManyRequests      = "TOO_MANY_REQUESTS"
	ErrInvalidCredentials   = "INVALID_CREDENTIALS"
	ErrInvalidToken         = "INVALID_TOKEN"
	ErrExpiredToken         = "EXPIRED_TOKEN"
	ErrInvalidRequest       = "INVALID_REQUEST"
	ErrInvalidFileType      = "INVALID_FILE_TYPE"
	ErrFileTooLarge         = "FILE_TOO_LARGE"
	ErrReadError            = "READ_ERROR"
	ErrWriteError           = "WRITE_ERROR"
	ErrDataCorrupted        = "DATA_CORRUPTED"
	ErrInvalidDateFormat    = "INVALID_DATE_FORMAT"
	ErrInvalidAmount        = "INVALID_AMOUNT"
	ErrInvalidCurrency      = "INVALID_CURRENCY"
	ErrDuplicateEntry       = "DUPLICATE_ENTRY"
	ErrGeneralError         = "GENERAL_ERROR"
	ErrUnexpectedError      = "UNEXPECTED_ERROR"
	ErrUnknown              = "UNKNOWN"
	ErrSystemError          = "SYSTEM_ERROR"
	ErrGatewayTimeout       = "GATEWAY_TIMEOUT"
)

// Error messages map
var ErrorMessages = map[string]string{
	ErrBadRequest:           "Bad Request",
	ErrUnauthorized:         "Unauthorized Access",
	ErrNotFound:             "Resource Not Found",
	ErrInternalServerError:  "Internal Server Error",
	ErrForbidden:            "Forbidden Access",
	ErrInvalidInput:         "Invalid Input",
	ErrDatabaseError:        "Database Error",
	ErrFileError:            "File Operation Error",
	ErrAuthenticationFailed: "Authentication Failed",
	ErrValidationFailed:     "Validation Failed",
	ErrEmailError:           "Email Error",
	ErrInvalidPassword:      "Invalid Password",
	ErrPaymentError:         "Payment Error",
	ErrNetworkError:         "Network Error",
	ErrTimeout:              "Timeout",
	ErrRateLimitExceeded:    "Rate Limit Exceeded",
	ErrServiceUnavailable:   "Service Unavailable",
	ErrConflict:             "Conflict",
	ErrUnprocessableEntity:  "Unprocessable Entity",
	ErrTooManyRequests:      "Too Many Requests",
	ErrInvalidCredentials:   "Invalid Credentials",
	ErrInvalidToken:         "Invalid Token",
	ErrExpiredToken:         "Expired Token",
	ErrInvalidRequest:       "Invalid Request",
	ErrInvalidFileType:      "Invalid File Type",
	ErrFileTooLarge:         "File Too Large",
	ErrReadError:            "Read Error",
	ErrWriteError:           "Write Error",
	ErrDataCorrupted:        "Data Corrupted",
	ErrInvalidDateFormat:    "Invalid Date Format",
	ErrInvalidAmount:        "Invalid Amount",
	ErrInvalidCurrency:      "Invalid Currency",
	ErrDuplicateEntry:       "Duplicate Entry",
	ErrGeneralError:         "General Error",
	ErrUnexpectedError:      "Unexpected Error",
	ErrUnknown:              "Unknown",
	ErrSystemError:          "System Error",
	ErrGatewayTimeout:       "Gateway Timeout",
}

func NewError(errorCode string) *AppError {
	return NewAppError(AppError{ErrorCode: errorCode, Timestamp: time.Now(), StatusCode: E400})
}

// NewAppError creates a new application error with the given code
func NewAppError(p AppError) *AppError {
	message, exists := ErrorMessages[p.ErrorCode]

	if !exists {
		codeWords := strings.Split(p.ErrorCode, "_")
		if len(codeWords) == 1 && codeWords[0] == p.ErrorCode {
			spaced := regexp.MustCompile(`([a-z0-9])([A-Z])`).ReplaceAllString(p.ErrorCode, "${1} ${2}")
			if spaced != "" {
				message = strings.ToLower(spaced)
			} else {
				message = "Unknown error"
			}
		} else {
			message = strings.Join(codeWords, " ")
		}
	}
	if p.Message == "" {
		p.Message = message
	}
	if p.StatusCode == "" {
		p.StatusCode = E400
	}
	if p.Timestamp.IsZero() {
		p.Timestamp = time.Now()
	}
	return &AppError{
		StatusCode: p.StatusCode,
		ErrorCode:  p.ErrorCode,
		Message:    p.Message,
		Timestamp:  p.Timestamp,
		Details:    p.Details,
	}
}

// Error returns the error message of the AppError
func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode, e.Message)
}

// GetMessage returns the error message of the AppError
func (e *AppError) GetMessage() string {
	return e.Message
}

// GetCode returns the error code of the AppError
func (e *AppError) GetCode() string {
	return e.ErrorCode
}

func (e *AppError) GetStatusCode() string {
	return e.StatusCode
}

func (e *AppError) GetErrorDetails() interface{} {
	return e.Details
}
func (e *AppError) GetTimestamp() time.Time {
	return e.Timestamp
}

func (e *AppError) GetResponse() AppError {
	return AppError{
		ErrorCode: e.ErrorCode,
		Message:   e.Message,
		Timestamp: e.Timestamp,
		Details:   e.Details,
	}
}
