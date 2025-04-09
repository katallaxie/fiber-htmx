package htmx

import (
	"fmt"
	"strings"
	"time"

	"github.com/katallaxie/pkg/cast"
	"github.com/katallaxie/pkg/conv"
)

var (
	HxDefaultSwapDuration = 0 * time.Millisecond
	HxDefaultSettleDelay  = 20 * time.Millisecond
)

// Swap ...
type Swap struct {
	style       HXSwapStyle
	transition  *bool
	timing      *HxSwapTiming
	scrolling   *HxSwapScrolling
	ignoreTitle *bool
	focusScroll *bool
}

// HxSwapTiming ...
type HxSwapTiming struct {
	mode     HxSwapTimingMode
	duration time.Duration
}

// String ...
func (s *HxSwapTiming) String() string {
	var out string

	out = string(s.mode)

	if s.duration != 0 {
		out += ":" + s.duration.String()
	}

	return out
}

// HxSwapScrolling ...
type HxSwapScrolling struct {
	mode      HxSwapScrollingMode
	target    string
	direction HxSwapDirection
}

// String ...
func (s *HxSwapScrolling) String() string {
	var out string

	out = string(s.mode)

	if s.target != "" {
		out += ":" + s.target
	}

	if s.direction != "" {
		out += ":" + s.direction.String()
	}

	return out
}

// NewSwap ...
func NewSwap() *Swap {
	return &Swap{
		style: HxSwapInnerHTML,
	}
}

// Style ...
func (s *Swap) Style(style HXSwapStyle) *Swap {
	s.style = style
	return s
}

func (s *Swap) setScrolling(mode HxSwapScrollingMode, direction HxSwapDirection, target ...string) *Swap {
	scrolling := &HxSwapScrolling{
		mode:      mode,
		direction: direction,
	}

	if len(target) > 0 {
		scrolling.target = target[0]
	}

	s.scrolling = scrolling
	return s
}

// Scroll ...
func (s *Swap) Scroll(direction HxSwapDirection, target ...string) *Swap {
	return s.setScrolling(HxSwapScrollingScroll, direction, target...)
}

// ScrollTop ...
func (s *Swap) ScrollTop(target ...string) *Swap {
	return s.Scroll(HxSwapDirectionTop, target...)
}

// ScrollBottom ...
func (s *Swap) ScrollBottom(target ...string) *Swap {
	return s.Scroll(HxSwapDirectionBottom, target...)
}

// Show ...
func (s *Swap) Show(direction HxSwapDirection, target ...string) *Swap {
	return s.setScrolling(HxSwapScrollingShow, direction, target...)
}

// ShowTop ...
func (s *Swap) ShowTop(target ...string) *Swap {
	return s.Show(HxSwapDirectionTop, target...)
}

// ShowBottom ...
func (s *Swap) ShowBottom(target ...string) *Swap {
	return s.Show(HxSwapDirectionBottom, target...)
}

func (s *Swap) setTiming(mode HxSwapTimingMode, swap ...time.Duration) *Swap {
	var duration time.Duration

	if len(swap) > 0 {
		duration = swap[0]
	} else {
		switch mode {
		case HxTimingSwap:
			duration = HxDefaultSwapDuration
		case HxTimingSettle:
			duration = HxDefaultSettleDelay
		}
	}

	s.timing = &HxSwapTiming{
		mode:     mode,
		duration: duration,
	}
	return s
}

// Swap modifies the amount of time that htmx will wait after receiving a response to swap the content
func (s *Swap) Swap(swap ...time.Duration) *Swap {
	return s.setTiming(HxTimingSwap, swap...)
}

// Settle modifies the amount of time that htmx will wait after receiving a response to settle the content
func (s *Swap) Settle(swap ...time.Duration) *Swap {
	return s.setTiming(HxTimingSwap, swap...)
}

// Transition ...
func (s *Swap) Transition(transition bool) *Swap {
	s.transition = &transition
	return s
}

// IgnoreTitle ...
func (s *Swap) IgnoreTitle(ignoreTitle bool) *Swap {
	s.ignoreTitle = &ignoreTitle
	return s
}

// FocusScroll ...
func (s *Swap) FocusScroll(focusScroll bool) *Swap {
	s.focusScroll = &focusScroll
	return s
}

// String ...
func (s *Swap) String() string {
	var parts []string

	parts = append(parts, string(s.style))

	if s.scrolling != nil {
		parts = append(parts, s.scrolling.String())
	}

	if s.transition != nil {
		parts = append(parts, fmt.Sprintf("transition:%s", conv.String(cast.Value(s.transition))))
	}

	if s.ignoreTitle != nil {
		parts = append(parts, fmt.Sprintf("ignoreTitle:%s", conv.String(cast.Value(s.ignoreTitle))))
	}

	if s.focusScroll != nil {
		parts = append(parts, fmt.Sprintf("focus-scroll:%s", conv.String(cast.Value(s.focusScroll))))
	}

	if s.timing != nil {
		parts = append(parts, s.timing.String())
	}

	return strings.Join(parts, " ")
}

// HxSwapStyle ...
type HXSwapStyle string

// String ...
func (s HXSwapStyle) String() string {
	return string(s)
}

const (
	HxSwapInnerHTML   HXSwapStyle = "innerHTML"
	HxSwapOuterHTML   HXSwapStyle = "outerHTML"
	HxSwapBeforeBegin HXSwapStyle = "beforebegin"
	HxSwapAfterBegin  HXSwapStyle = "afterbegin"
	HxSwapBeforeEnd   HXSwapStyle = "beforeend"
	HxSwapAfterEnd    HXSwapStyle = "afterend"
	HxSwapDelete      HXSwapStyle = "delete"
	HxSwapNone        HXSwapStyle = "none"
)

// HxSwapScrollingMode ...
type HxSwapScrollingMode string

// String ...
func (s HxSwapScrollingMode) String() string {
	return string(s)
}

const (
	HxSwapScrollingScroll HxSwapScrollingMode = "scroll"
	HxSwapScrollingShow   HxSwapScrollingMode = "show"
)

// HxSwapTimingMode ...
type HxSwapTimingMode string

// String ...
func (s HxSwapTimingMode) String() string {
	return string(s)
}

const (
	HxTimingSwap   HxSwapTimingMode = "swap"
	HxTimingSettle HxSwapTimingMode = "settle"
)

// HxSwapDirection ...
type HxSwapDirection string

// String ...
func (s HxSwapDirection) String() string {
	return string(s)
}

const (
	HxSwapDirectionTop    HxSwapDirection = "top"
	HxSwapDirectionBottom HxSwapDirection = "bottom"
)
