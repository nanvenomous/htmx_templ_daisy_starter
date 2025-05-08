package models

import "net/http"

type HttpStatus int

var (
	StatusContinue           HttpStatus = http.StatusContinue           // 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols HttpStatus = http.StatusSwitchingProtocols // 101 // RFC 9110, 15.2.2
	StatusProcessing         HttpStatus = http.StatusProcessing         // 102 // RFC 2518, 10.1
	StatusEarlyHints         HttpStatus = http.StatusEarlyHints         // 103 // RFC 8297

	StatusOK                   HttpStatus = http.StatusOK                   // 200 // RFC 9110, 15.3.1
	StatusCreated              HttpStatus = http.StatusCreated              // 201 // RFC 9110, 15.3.2
	StatusAccepted             HttpStatus = http.StatusAccepted             // 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo HttpStatus = http.StatusNonAuthoritativeInfo // 203 // RFC 9110, 15.3.4
	StatusNoContent            HttpStatus = http.StatusNoContent            // 204 // RFC 9110, 15.3.5
	StatusResetContent         HttpStatus = http.StatusResetContent         // 205 // RFC 9110, 15.3.6
	StatusPartialContent       HttpStatus = http.StatusPartialContent       // 206 // RFC 9110, 15.3.7
	StatusMultiStatus          HttpStatus = http.StatusMultiStatus          // 207 // RFC 4918, 11.1
	StatusAlreadyReported      HttpStatus = http.StatusAlreadyReported      // 208 // RFC 5842, 7.1
	StatusIMUsed               HttpStatus = http.StatusIMUsed               // 226 // RFC 3229, 10.4.1

	StatusMultipleChoices   HttpStatus = http.StatusMultipleChoices   // 300 // RFC 9110, 15.4.1
	StatusMovedPermanently  HttpStatus = http.StatusMovedPermanently  // 301 // RFC 9110, 15.4.2
	StatusFound             HttpStatus = http.StatusFound             // 302 // RFC 9110, 15.4.3
	StatusSeeOther          HttpStatus = http.StatusSeeOther          // 303 // RFC 9110, 15.4.4
	StatusNotModified       HttpStatus = http.StatusNotModified       // 304 // RFC 9110, 15.4.5
	StatusUseProxy          HttpStatus = http.StatusUseProxy          // 305 // RFC 9110, 15.4.6
	_                       HttpStatus = 306                          // RFC 9110, 15.4.7 (Unused)
	StatusTemporaryRedirect HttpStatus = http.StatusTemporaryRedirect // 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect HttpStatus = http.StatusPermanentRedirect // 308 // RFC 9110, 15.4.9

	StatusBadRequest                   HttpStatus = http.StatusBadRequest                   // 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 HttpStatus = http.StatusUnauthorized                 // 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              HttpStatus = http.StatusPaymentRequired              // 402 // RFC 9110, 15.5.3
	StatusForbidden                    HttpStatus = http.StatusForbidden                    // 403 // RFC 9110, 15.5.4
	StatusNotFound                     HttpStatus = http.StatusNotFound                     // 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             HttpStatus = http.StatusMethodNotAllowed             // 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                HttpStatus = http.StatusNotAcceptable                // 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            HttpStatus = http.StatusProxyAuthRequired            // 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               HttpStatus = http.StatusRequestTimeout               // 408 // RFC 9110, 15.5.9
	StatusConflict                     HttpStatus = http.StatusConflict                     // 409 // RFC 9110, 15.5.10
	StatusGone                         HttpStatus = http.StatusGone                         // 410 // RFC 9110, 15.5.11
	StatusLengthRequired               HttpStatus = http.StatusLengthRequired               // 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           HttpStatus = http.StatusPreconditionFailed           // 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        HttpStatus = http.StatusRequestEntityTooLarge        // 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            HttpStatus = http.StatusRequestURITooLong            // 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         HttpStatus = http.StatusUnsupportedMediaType         // 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable HttpStatus = http.StatusRequestedRangeNotSatisfiable // 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            HttpStatus = http.StatusExpectationFailed            // 417 // RFC 9110, 15.5.18
	StatusTeapot                       HttpStatus = http.StatusTeapot                       // 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           HttpStatus = http.StatusMisdirectedRequest           // 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          HttpStatus = http.StatusUnprocessableEntity          // 422 // RFC 9110, 15.5.21
	StatusLocked                       HttpStatus = http.StatusLocked                       // 423 // RFC 4918, 11.3
	StatusFailedDependency             HttpStatus = http.StatusFailedDependency             // 424 // RFC 4918, 11.4
	StatusTooEarly                     HttpStatus = http.StatusTooEarly                     // 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              HttpStatus = http.StatusUpgradeRequired              // 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         HttpStatus = http.StatusPreconditionRequired         // 428 // RFC 6585, 3
	StatusTooManyRequests              HttpStatus = http.StatusTooManyRequests              // 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  HttpStatus = http.StatusRequestHeaderFieldsTooLarge  // 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   HttpStatus = http.StatusUnavailableForLegalReasons   // 451 // RFC 7725, 3

	StatusInternalServerError           HttpStatus = http.StatusInternalServerError           // 500 // RFC 9110, 15.6.1
	StatusNotImplemented                HttpStatus = http.StatusNotImplemented                // 501 // RFC 9110, 15.6.2
	StatusBadGateway                    HttpStatus = http.StatusBadGateway                    // 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            HttpStatus = http.StatusServiceUnavailable            // 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                HttpStatus = http.StatusGatewayTimeout                // 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       HttpStatus = http.StatusHTTPVersionNotSupported       // 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         HttpStatus = http.StatusVariantAlsoNegotiates         // 506 // RFC 2295, 8.1
	StatusInsufficientStorage           HttpStatus = http.StatusInsufficientStorage           // 507 // RFC 4918, 11.5
	StatusLoopDetected                  HttpStatus = http.StatusLoopDetected                  // 508 // RFC 5842, 7.2
	StatusNotExtended                   HttpStatus = http.StatusNotExtended                   // 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired HttpStatus = http.StatusNetworkAuthenticationRequired // 511 // RFC 6585, 6
	StatusHTMLTemplatingError           HttpStatus = 512
	StatusPanicUnrecoverableError       HttpStatus = 513
)
