// Package htmx provides a set of functions for generating HTML attributes as nodes.
// These attributes can be used in conjunction with the Fiber framework and the htmx library.
package htmx

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/katallaxie/pkg/conv"
	"github.com/katallaxie/pkg/errorx"
)

// JSEventType represents the type of JavaScript event.
type JSEventType string

// String returns the string representation of the JSEventType.
func (e JSEventType) String() string {
	return string(e)
}

// JSEvent is a type that represents JavaScript events.
func JSEvent(e JSEventType, value ...string) Node {
	return Attribute(e.String(), value...)
}

// List of predefined JavaScript event types.
const (
	JSEventTypeClickEvent       JSEventType = "onclick"
	JSEventTypeChangeEvent      JSEventType = "onchange"
	JSEventTypeInputEvent       JSEventType = "oninput"
	JSEventTypeFocusEvent       JSEventType = "onfocus"
	JSEventTypeBlurEvent        JSEventType = "onblur"
	JSEventTypeKeyDownEvent     JSEventType = "onkeydown"
	JSEventTypeKeyUpEvent       JSEventType = "onkeyup"
	JSEventTypeKeyPressEvent    JSEventType = "onkeypress"
	JSEventTypeSubmitEvent      JSEventType = "onsubmit"
	JSEventTypeLoadDomEvent     JSEventType = "onload"
	JSEventTypeLoadEvent        JSEventType = "onload"
	JSEventTypeUnloadEvent      JSEventType = "onunload"
	JSEventTypeResizeEvent      JSEventType = "onresize"
	JSEventTypeScrollEvent      JSEventType = "onscroll"
	JSEventTypeDblClickEvent    JSEventType = "ondblclick"
	JSEventTypeMouseOverEvent   JSEventType = "onmouseover"
	JSEventTypeMouseOutEvent    JSEventType = "onmouseout"
	JSEventTypeMouseMoveEvent   JSEventType = "onmousemove"
	JSEventTypeMouseDownEvent   JSEventType = "onmousedown"
	JSEventTypeMouseUpEvent     JSEventType = "onmouseup"
	JSEventTypeContextMenuEvent JSEventType = "oncontextmenu"
	JSEventTypeDragStartEvent   JSEventType = "ondragstart"
	JSEventTypeDragEvent        JSEventType = "ondrag"
	JSEventTypeDragEnterEvent   JSEventType = "ondragenter"
	JSEventTypeDragLeaveEvent   JSEventType = "ondragleave"
	JSEventTypeDragOverEvent    JSEventType = "ondragover"
	JSEventTypeDropEvent        JSEventType = "ondrop"
	JSEventTypeDragEndEvent     JSEventType = "ondragend"
)

// HxEventType represents the type of htmx event.
type HxEventType string

// String returns the string representation of the HxEventType.
func (e HxEventType) String() string {
	return string(e)
}

// List of predefined htmx event types.
// See: https://htmx.org/reference/#events
const (
	HxEventTypeAbort                      HxEventType = "htmx:abort"
	HxEventTypeAfterLoad                  HxEventType = "htmx:afterLoad"
	HxEventTypeAfterProcessNode           HxEventType = "htmx:afterProcessNode"
	HxEventTypeAfterRequest               HxEventType = "htmx:afterRequest"
	HxEventTypeBeforeHistorySaveEvent     HxEventType = "htmx:beforeHistorySave"
	HxEventTypeBeforeOnAbort              HxEventType = "htmx:beforeOnAbort"
	HxEventTypeBeforeOnConfigRequest      HxEventType = "htmx:beforeOnConfigRequest"
	HxEventTypeBeforeOnConfigSwap         HxEventType = "htmx:beforeOnConfigSwap"
	HxEventTypeBeforeOnConfigTrigger      HxEventType = "htmx:beforeOnConfigTrigger"
	HxEventTypeBeforeOnError              HxEventType = "htmx:beforeOnError"
	HxEventTypeBeforeOnLoad               HxEventType = "htmx:beforeOnLoad"
	HxEventTypeBeforeOnRequest            HxEventType = "htmx:beforeOnRequest"
	HxEventTypeBeforeOnSettle             HxEventType = "htmx:beforeOnSettle"
	HxEventTypeBeforeOnSwap               HxEventType = "htmx:beforeOnSwap"
	HxEventTypeBeforeOnTrigger            HxEventType = "htmx:beforeOnTrigger"
	HxEventTypeBeforeRequest              HxEventType = "htmx:beforeRequest"
	HxEventTypeBeforeSettle               HxEventType = "htmx:beforeSettle"
	HxEventTypeBeforeSwap                 HxEventType = "htmx:beforeSwap"
	HxEventTypeConfigRequest              HxEventType = "htmx:configRequest"
	HxEventTypeConfigSwap                 HxEventType = "htmx:configSwap"
	HxEventTypeConfigTrigger              HxEventType = "htmx:configTrigger"
	HxEventTypeError                      HxEventType = "htmx:error"
	HxEventTypeHistoryCacheErrorEvent     HxEventType = "htmx:historyCacheError"
	HxEventTypeHistoryCacheMissErrorEvent HxEventType = "htmx:historyCacheMissError"
	HxEventTypeHistoryCacheMissEvent      HxEventType = "htmx:historyCacheMiss"
	HxEventTypeHistoryCacheMissLoadEvent  HxEventType = "htmx:historyCacheMissLoad"
	HxEventTypeHistoryRestoreEvent        HxEventType = "htmx:historyRestore"
	HxEventTypeLoad                       HxEventType = "htmx:load"
	HxEventTypeNoSSESourceErrorEvent      HxEventType = "htmx:noSSESourceError"
	HxEventTypeOnAbort                    HxEventType = "htmx:onAbort"
	HxEventTypeOnConfigRequest            HxEventType = "htmx:onConfigRequest"
	HxEventTypeOnConfigSwap               HxEventType = "htmx:onConfigSwap"
	HxEventTypeOnConfigTrigger            HxEventType = "htmx:onConfigTrigger"
	HxEventTypeOnError                    HxEventType = "htmx:onError"
	HxEventTypeOnLoad                     HxEventType = "htmx:onLoad"
	HxEventTypeOnLoadErrorEvent           HxEventType = "htmx:onLoadError"
	HxEventTypeOnRequest                  HxEventType = "htmx:onRequest"
	HxEventTypeOnSettle                   HxEventType = "htmx:onSettle"
	HxEventTypeOnSwap                     HxEventType = "htmx:onSwap"
	HxEventTypeOnTrigger                  HxEventType = "htmx:onTrigger"
	HxEventTypeOobAfterSwapEvent          HxEventType = "htmx:oobAfterSwap"
	HxEventTypeOobBeforeSwapEvent         HxEventType = "htmx:oobBeforeSwap"
	HxEventTypePromptEvent                HxEventType = "htmx:prompt"
	HxEventTypePushedIntoHistoryEvent     HxEventType = "htmx:pushedIntoHistory"
	HxEventTypeRequest                    HxEventType = "htmx:request"
	HxEventTypeResponseErrorEvent         HxEventType = "htmx:responseError"
	HxEventTypeSendErrorEvent             HxEventType = "htmx:sendError"
	HxEventTypeSettle                     HxEventType = "htmx:settle"
	HxEventTypeSseAfterMessageEvent       HxEventType = "htmx:sseAfterMessage"
	HxEventTypeSseBeforeMessageEvent      HxEventType = "htmx:sseBeforeMessage"
	HxEventTypeSseClosedEvent             HxEventType = "htmx:sseClose"
	HxEventTypeSseConnectedEvent          HxEventType = "htmx:sseOpen"
	HxEventTypeSseConnectingEvent         HxEventType = "htmx:sseConnecting"
	HxEventTypeSseErrorEvent              HxEventType = "htmx:sseError"
	HxEventTypeSSEErrorEvent              HxEventType = "htmx:sseError"
	HxEventTypeSSEOpenEvent               HxEventType = "htmx:sseOpen"
	HxEventTypeSwap                       HxEventType = "htmx:swap"
	HxEventTypeSwapErrorEvent             HxEventType = "htmx:swapError"
	HxEventTypeTargetErrorEvent           HxEventType = "htmx:targetError"
	HxEventTypeTimeoutEvent               HxEventType = "htmx:timeout"
	HxEventTypeTrigger                    HxEventType = "htmx:trigger"
	HxEventTypeValidationFailedEvent      HxEventType = "htmx:validation:failed"
	HxEventTypeValidationHaltedEvent      HxEventType = "htmx:validation:halted"
	HxEventTypeValidationValidateEvent    HxEventType = "htmx:validation:validate"
	HxEventTypeXhrAbortEvent              HxEventType = "htmx:xhr:abort"
	HxEventTypeXhrLoadEndEvent            HxEventType = "htmx:xhr:loadend"
	HxEventTypeXhrLoadStartEvent          HxEventType = "htmx:xhr:loadstart"
	HxEventTypeXhrProgressEvent           HxEventType = "htmx:xhr:progress"
	HxEventTypeOobErrorNoTargetEvent      HxEventType = "htmx:oobErrorNoTarget"
)

type HxAttribute string // https://htmx.org/reference/#attributes

func (a HxAttribute) String() string {
	return string(a)
}

const (
	HxAttributeBoost       HxAttribute = "hx-boost"
	HxAttributeGet         HxAttribute = "hx-get"
	HxAttributePost        HxAttribute = "hx-post"
	HxAttributePushURL     HxAttribute = "hx-push-url"
	HxAttributeTarget      HxAttribute = "hx-target"
	HxAttributeSelect      HxAttribute = "hx-select"
	HxAttributeSelectOob   HxAttribute = "hx-select-oob"
	HxAttributeSwap        HxAttribute = "hx-swap"
	HxAttributeSwapOob     HxAttribute = "hx-swap-oob"
	HxAttributeTrigger     HxAttribute = "hx-trigger"
	HxAttributeConfirm     HxAttribute = "hx-confirm"
	HxAttributePrompt      HxAttribute = "hx-prompt"
	HxAttributeDelete      HxAttribute = "hx-delete"
	HxAttributeOn          HxAttribute = "hx-on"
	HxAttributePut         HxAttribute = "hx-put"
	HxAttributePatch       HxAttribute = "hx-patch"
	HxAttributeIndicator   HxAttribute = "hx-indicator"
	HxAttributeEncoding    HxAttribute = "hx-encoding"
	HxAttributeExt         HxAttribute = "hx-ext"
	HxAttributeTarget404   HxAttribute = "hx-target-404"
	HxAttributeTarget403   HxAttribute = "hx-target-403"
	HxAttributeTarget401   HxAttribute = "hx-target-401"
	HxAttributeTarget500   HxAttribute = "hx-target-500"
	HxAttributeTarget5xx   HxAttribute = "hx-target-5xx"
	HxAttributeTarget4xx   HxAttribute = "hx-target-4xx"
	HxAttributeTargetError HxAttribute = "hx-target-error"
	HxAttributeTarget50x   HxAttribute = "hx-target-50x"
	HxAttributeDisable     HxAttribute = "hx-disable"
	HxAttributeDisabledElt HxAttribute = "hx-disabled-elt"
	HxAttributeValidate    HxAttribute = "hx-validate"
	HxAttributeInclude     HxAttribute = "hx-include"
	HxAttributeHeaders     HxAttribute = "hx-headers"
	HxAttributeVars        HxAttribute = "hx-vars"
	HxAttributeSync        HxAttribute = "hx-sync"
	HxAttributeParams      HxAttribute = "hx-params"
	HxAttributeVals        HxAttribute = "hx-vals"
)

// HxBoost sets the hx-boost attribute to enable or disable boosting.
func HxBoost(v bool) Node {
	return Attribute(HxAttributeBoost.String(), conv.String(v))
}

// HxGet sets the hx-get attribute to specify the URL for GET requests.
func HxGet(url string) Node {
	return Attribute(HxAttributeGet.String(), url)
}

// HxPost sets the hx-post attribute to specify the URL for POST requests.
func HxPost(url string) Node {
	return Attribute(HxAttributePost.String(), url)
}

// HxPushURL sets the hx-push-url attribute to enable or disable URL pushing.
func HxPushURL(v bool) Node {
	return Attribute(HxAttributeBoost.String(), conv.String(v))
}

// HxTarget sets the hx-target attribute to specify the target element for the response.
func HxTarget(target string) Node {
	return Attribute(HxAttributeTarget.String(), target)
}

// HxSelect sets the hx-select attribute to specify the target element for selection.
func HxSelect(target string) Node {
	return Attribute(HxAttributeSelect.String(), target)
}

// HxSelectOob sets the hx-select-oob attribute to specify the target element for out-of-band selection.
func HxSelectOob(target string) Node {
	return Attribute(HxAttributeSelectOob.String(), target)
}

// HxSwap sets the hx-swap attribute to specify the target element for swapping.
func HxSwap(target string) Node {
	return Attribute(HxAttributeSwap.String(), target)
}

// HxSwapOob sets the hx-swap-oob attribute to specify the target element for out-of-band swapping.
func HxSwapOob(target string) Node {
	return Attribute(HxAttributeSwapOob.String(), target)
}

const (
	TriggerClick        = "click"
	TriggerClickOnce    = TriggerClick + " once"
	TriggerDblClick     = JSEventTypeDblClickEvent
	TriggerKeyUpEnter   = "keyup[keyCode==13]"
	TriggerEnterPressed = TriggerKeyUpEnter
	TriggerBlur         = "blur"
	TriggerEvery1s      = "every:1s"
	TriggerEvery2s      = "every:2s"
	TriggerEvery5s      = "every:5s"
	TriggerEvery10s     = "every:10s"
	TriggerEvery30s     = "every:30s"
	TriggerEvery1m      = "every:1m"
	TriggerLoad         = "onload"
)

// HxTrigger sets the hx-trigger attribute to specify the target element for triggering an event.
func HxTrigger(targets ...string) Node {
	var b strings.Builder
	for i, t := range targets {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(t)
	}

	return Attribute(HxAttributeTrigger.String(), b.String())
}

// HxConfirm sets the hx-confirm attribute to display a confirmation message.
func HxConfirm(msg string) Node {
	return Attribute(HxAttributeConfirm.String(), msg)
}

// HxPrompt sets the hx-prompt attribute to display a prompt message.
func HxPrompt(msg string) Node {
	return Attribute(HxAttributePrompt.String(), msg)
}

// HxDelete sets the hx-delete attribute to specify the URL for DELETE requests.
func HxDelete(url string) Node {
	return Attribute(HxAttributeDelete.String(), url)
}

// HxOn sets the hx-put-{target} attribute to specify the JavaScript code to execute on a PUT request.
func HxOn(target string, js string) Node {
	return Attribute(fmt.Sprintf("hx-on:%s", target), js)
}

// HxPut sets the hx-put attribute to specify the URL for PUT requests.
func HxPut(url string) Node {
	return Attribute(HxAttributePut.String(), url)
}

// HxPatch sets the hx-patch attribute to specify the URL for PATCH requests.
func HxPatch(url string) Node {
	return Attribute(HxAttributePatch.String(), url)
}

// HxIndicator sets the hx-indicator attribute to specify the target element for showing an indicator.
func HxIndicator(target string) Node {
	return Attribute(HxAttributeIndicator.String(), target)
}

// HxEncoding sets the hx-encoding attribute to specify the encoding type for form submission.
func HxEncoding(enc string) Node {
	return Attribute(HxAttributeEncoding.String(), enc)
}

// HxExt sets the hx-ext attribute to specify the file extension for file uploads.
func HxExt(ext string) Node {
	return Attribute(HxAttributeExt.String(), ext)
}

// HxTarget404 sets the hx-target-404 attribute to specify the target element for 404 responses.
func HxTarget404(target string) Node {
	return Attribute(HxAttributeTarget404.String(), target)
}

// HxTarget403 sets the hx-target-403 attribute to specify the target element for 403 responses.
func HxTarget403(target string) Node {
	return Attribute(HxAttributeTarget403.String(), target)
}

// HxTarget401 sets the hx-target-401 attribute to specify the target element for 401 responses.
func HxTarget401(target string) Node {
	return Attribute(HxAttributeTarget401.String(), target)
}

// HxTarget500 sets the hx-target-500 attribute to specify the target element for 500 responses.
func HxTarget500(target string) Node {
	return Attribute(HxAttributeTarget500.String(), target)
}

// HxTarget5xx sets the hx-target-5xx attribute to specify the target element for 5xx responses.
func HxTarget5xx(target string) Node {
	return Attribute(HxAttributeTarget5xx.String(), target)
}

// HxTarget4xx sets the hx-target-4xx attribute to specify the target element for 4xx responses.
func HxTarget4xx(target string) Node {
	return Attribute(HxAttributeTarget4xx.String(), target)
}

// HxTargetError sets the hx-target-error attribute to specify the target element for error responses.
func HxTargetError(target string) Node {
	return Attribute(HxAttributeTargetError.String(), target)
}

// HxTarget50x sets the hx-target-50x attribute to specify the target element for 50x responses.
func HxTarget50x(target string) Node {
	return Attribute(HxAttributeTarget5xx.String(), target)
}

// HxDisable sets the hx-disable attribute to disable htmx functionality.
func HxDisable() Node {
	return Attribute(HxAttributeDisable.String())
}

// HxDisabledElt sets the hx-disable-elt attribute to disable the target element.
func HxDisabledElt(target string) Node {
	return Attribute(HxAttributeDisabledElt.String(), target)
}

// HxValidate sets the hx-validate attribute to enable or disable form validation.
func HxValidate(v bool) Node {
	return Attribute(HxAttributeValidate.String(), conv.String(v))
}

// HxInclude sets the hx-include attribute to specify the target element for inclusion.
func HxInclude(target string) Node {
	return Attribute(HxAttributeInclude.String(), target)
}

// HxHeaders sets the hx-headers attribute to specify the headers for the request.
func HxHeaders(headers map[string]string) Node {
	return Attribute("hx-headers", string(errorx.Ignore(json.Marshal(headers))))
}

// Async sets the async attribute for script elements.
func Async() Node {
	return Attribute("async")
}

// AutoFocus sets the autofocus attribute for form elements.
func AutoFocus() Node {
	return Attribute("autofocus")
}

// AutoPlay sets the autoplay attribute for media elements.
func AutoPlay() Node {
	return Attribute("autoplay")
}

// Checked sets the checked attribute for input elements.
func Checked() Node {
	return Attribute("checked", "checked")
}

// Controls sets the controls attribute for media elements.
func Controls() Node {
	return Attribute("controls")
}

// Defer sets the defer attribute for script elements.
func Defer() Node {
	return Attribute("defer")
}

// Disabled sets the disabled attribute for form elements.
func Disabled() Node {
	return Attribute("disabled", "disabled")
}

// Loop sets the loop attribute for media elements.
func Loop() Node {
	return Attribute("loop")
}

// Multiple sets the multiple attribute for input elements.
func Multiple() Node {
	return Attribute("multiple")
}

// Muted sets the muted attribute for media elements.
func Muted() Node {
	return Attribute("muted")
}

// PlaysInline sets the playsinline attribute for media elements.
func PlaysInline() Node {
	return Attribute("playsinline")
}

// ReadOnly sets the readonly attribute for form elements.
func ReadOnly() Node {
	return Attribute("readonly")
}

// Required sets the required attribute for form elements.
func Required() Node {
	return Attribute("required")
}

// Selected sets the selected attribute for option elements.
func Selected() Node {
	return Attribute("selected")
}

// Accept sets the accept attribute for file input elements.
func Accept(v string) Node {
	return Attribute("accept", v)
}

// Action sets the action attribute for form elements.
func Action(v string) Node {
	return Attribute("action", v)
}

// Alt sets the alt attribute for image elements.
func Alt(v string) Node {
	return Attribute("alt", v)
}

// Aria sets the aria-{name} attribute for elements.
func Aria(name, v string) Node {
	return Attribute("aria-"+name, v)
}

// As sets the as attribute for link elements.
func As(v string) Node {
	return Attribute("as", v)
}

// AutoComplete sets the autocomplete attribute for form elements.
func AutoComplete(v string) Node {
	return Attribute("autocomplete", v)
}

// Charset sets the charset attribute for meta elements.
func Charset(v string) Node {
	return Attribute("charset", v)
}

// Class sets the class attribute for elements.
func Class(v string) Node {
	return Attribute("class", v)
}

// Cols sets the cols attribute for textarea elements.
func Cols(v string) Node {
	return Attribute("cols", v)
}

// ColSpan sets the colspan attribute for table cells.
func ColSpan(v string) Node {
	return Attribute("colspan", v)
}

// Content sets the content attribute for meta elements.
func Content(v string) Node {
	return Attribute("content", v)
}

// DataAttribute sets the data-{name} attribute for elements.
func DataAttribute(name, v string) Node {
	return Attribute("data-"+name, v)
}

// For sets the for attribute for label elements.
func For(v string) Node {
	return Attribute("for", v)
}

// FormAttribute sets the form attribute for elements.
func FormAttribute(v string) Node {
	return Attribute("form", v)
}

// Height sets the height attribute for elements.
func Height(v string) Node {
	return Attribute("height", v)
}

// Href sets the href attribute for anchor elements.
func Href(v string) Node {
	return Attribute("href", v)
}

// ID sets the id attribute for elements.
func ID(v string) Node {
	return Attribute("id", v)
}

// Lang sets the lang attribute for elements.
func Lang(v string) Node {
	return Attribute("lang", v)
}

// Loading sets the loading attribute for elements.
func Loading(v string) Node {
	return Attribute("loading", v)
}

// Max sets the max attribute for input elements.
func Max(v string) Node {
	return Attribute("max", v)
}

// MaxLength sets the maxlength attribute for input elements.
func MaxLength(v string) Node {
	return Attribute("maxlength", v)
}

// Method sets the method attribute for form elements.
func Method(v string) Node {
	return Attribute("method", v)
}

// Min sets the min attribute for input elements.
func Min(v string) Node {
	return Attribute("min", v)
}

// MinLength sets the minlength attribute for input elements.
func MinLength(v string) Node {
	return Attribute("minlength", v)
}

// Name sets the name attribute for elements.
func Name(v string) Node {
	return Attribute("name", v)
}

// Pattern sets the pattern attribute for input elements.
func Pattern(v string) Node {
	return Attribute("pattern", v)
}

// Placeholder sets the placeholder attribute for input elements.
func Placeholder(v string) Node {
	return Attribute("placeholder", v)
}

// Poster sets the poster attribute for video elements.
func Poster(v string) Node {
	return Attribute("poster", v)
}

// Preload sets the preload attribute for media elements.
func Preload(v string) Node {
	return Attribute("preload", v)
}

// Rel sets the rel attribute for link elements.
func Rel(v string) Node {
	return Attribute("rel", v)
}

// Role sets the role attribute for elements.
func Role(v string) Node {
	return Attribute("role", v)
}

// Rows sets the rows attribute for textarea elements.
func Rows(v string) Node {
	return Attribute("rows", v)
}

// RowSpan sets the rowspan attribute for table cells.
func RowSpan(v string) Node {
	return Attribute("rowspan", v)
}

// Src sets the src attribute for elements.
func Src(v string) Node {
	return Attribute("src", v)
}

// SrcSet sets the srcset attribute for elements.
func SrcSet(v string) Node {
	return Attribute("srcset", v)
}

// Step sets the step attribute for input elements.
func Step(v string) Node {
	return Attribute("step", v)
}

// StyleAttribute sets the style attribute for elements.
func StyleAttribute(v string) Node {
	return Attribute("style", v)
}

// TabIndex sets the tabindex attribute for elements.
func TabIndex(v string) Node {
	return Attribute("tabindex", v)
}

// Target sets the target attribute for elements.
func Target(v string) Node {
	return Attribute("target", v)
}

// TitleAttribute sets the title attribute for elements.
func TitleAttribute(v string) Node {
	return Attribute("title", v)
}

// Type sets the type attribute for elements.
func Type(v string) Node {
	return Attribute("type", v)
}

// Value sets the value attribute for elements.
func Value(v string) Node {
	return Attribute("value", v)
}

// Width sets the width attribute for elements.
func Width(v string) Node {
	return Attribute("width", v)
}

// EncType sets the enctype attribute for form elements.
func EncType(v string) Node {
	return Attribute("enctype", v)
}

// OnClick sets the onclick attribute for elements.
func OnClick(v string) Node {
	return Attribute("onclick", v)
}

// Is sets the is attribute for custom elements.
func Is(v string) Node {
	return Attribute("is", v)
}

// Integrity sets the integrity attribute for elements.
func Integrity(v string) Node {
	return Attribute("integrity", v)
}

// NoValidate sets the novalidate attribute for form elements.
func NoValidate() Node {
	return Attribute("novalidate")
}

// CrossOrigin sets the crossorigin attribute for elements.
func CrossOrigin(v string) Node {
	return Attribute("crossorigin", v)
}

// List sets the list attribute for input elements.
func List(v string) Node {
	return Attribute("list", v)
}

// OnePasswordIgnore sets the data-1p-ignore attribute for elements.
func OnePasswordIgnore() Node {
	return Attribute("data-1p-ignore")
}

// ContentEditable sets the contenteditable attribute for elements.
func ContentEditable(v string) Node {
	return Attribute("contenteditable", v)
}
