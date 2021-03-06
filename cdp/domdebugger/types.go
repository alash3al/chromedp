package domdebugger

import (
	"errors"

	"github.com/knq/chromedp/cdp/runtime"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// AUTOGENERATED. DO NOT EDIT.

// DOMBreakpointType dOM breakpoint type.
type DOMBreakpointType string

// String returns the DOMBreakpointType as string value.
func (t DOMBreakpointType) String() string {
	return string(t)
}

// DOMBreakpointType values.
const (
	DOMBreakpointTypeSubtreeModified   DOMBreakpointType = "subtree-modified"
	DOMBreakpointTypeAttributeModified DOMBreakpointType = "attribute-modified"
	DOMBreakpointTypeNodeRemoved       DOMBreakpointType = "node-removed"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t DOMBreakpointType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t DOMBreakpointType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *DOMBreakpointType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch DOMBreakpointType(in.String()) {
	case DOMBreakpointTypeSubtreeModified:
		*t = DOMBreakpointTypeSubtreeModified
	case DOMBreakpointTypeAttributeModified:
		*t = DOMBreakpointTypeAttributeModified
	case DOMBreakpointTypeNodeRemoved:
		*t = DOMBreakpointTypeNodeRemoved

	default:
		in.AddError(errors.New("unknown DOMBreakpointType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *DOMBreakpointType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// EventListener object event listener.
type EventListener struct {
	Type            string                `json:"type,omitempty"`            // EventListener's type.
	UseCapture      bool                  `json:"useCapture,omitempty"`      // EventListener's useCapture.
	Passive         bool                  `json:"passive,omitempty"`         // EventListener's passive flag.
	Once            bool                  `json:"once,omitempty"`            // EventListener's once flag.
	ScriptID        runtime.ScriptID      `json:"scriptId,omitempty"`        // Script id of the handler code.
	LineNumber      int64                 `json:"lineNumber,omitempty"`      // Line number in the script (0-based).
	ColumnNumber    int64                 `json:"columnNumber,omitempty"`    // Column number in the script (0-based).
	Handler         *runtime.RemoteObject `json:"handler,omitempty"`         // Event handler function value.
	OriginalHandler *runtime.RemoteObject `json:"originalHandler,omitempty"` // Event original handler function value.
	RemoveFunction  *runtime.RemoteObject `json:"removeFunction,omitempty"`  // Event listener remove function.
}
