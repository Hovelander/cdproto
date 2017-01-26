// Package rendering provides the Chrome Debugging Protocol
// commands, types, and events for the Chrome Rendering domain.
//
// This domain allows to control rendering of the page.
//
// Generated by the chromedp-gen command.
package rendering

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

// SetShowPaintRectsParams requests that backend shows paint rectangles.
type SetShowPaintRectsParams struct {
	Result bool `json:"result"` // True for showing paint rectangles
}

// SetShowPaintRects requests that backend shows paint rectangles.
//
// parameters:
//   result - True for showing paint rectangles
func SetShowPaintRects(result bool) *SetShowPaintRectsParams {
	return &SetShowPaintRectsParams{
		Result: result,
	}
}

// Do executes Rendering.setShowPaintRects.
func (p *SetShowPaintRectsParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandRenderingSetShowPaintRects, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}

// SetShowDebugBordersParams requests that backend shows debug borders on
// layers.
type SetShowDebugBordersParams struct {
	Show bool `json:"show"` // True for showing debug borders
}

// SetShowDebugBorders requests that backend shows debug borders on layers.
//
// parameters:
//   show - True for showing debug borders
func SetShowDebugBorders(show bool) *SetShowDebugBordersParams {
	return &SetShowDebugBordersParams{
		Show: show,
	}
}

// Do executes Rendering.setShowDebugBorders.
func (p *SetShowDebugBordersParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandRenderingSetShowDebugBorders, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}

// SetShowFPSCounterParams requests that backend shows the FPS counter.
type SetShowFPSCounterParams struct {
	Show bool `json:"show"` // True for showing the FPS counter
}

// SetShowFPSCounter requests that backend shows the FPS counter.
//
// parameters:
//   show - True for showing the FPS counter
func SetShowFPSCounter(show bool) *SetShowFPSCounterParams {
	return &SetShowFPSCounterParams{
		Show: show,
	}
}

// Do executes Rendering.setShowFPSCounter.
func (p *SetShowFPSCounterParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandRenderingSetShowFPSCounter, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}

// SetShowScrollBottleneckRectsParams requests that backend shows scroll
// bottleneck rects.
type SetShowScrollBottleneckRectsParams struct {
	Show bool `json:"show"` // True for showing scroll bottleneck rects
}

// SetShowScrollBottleneckRects requests that backend shows scroll bottleneck
// rects.
//
// parameters:
//   show - True for showing scroll bottleneck rects
func SetShowScrollBottleneckRects(show bool) *SetShowScrollBottleneckRectsParams {
	return &SetShowScrollBottleneckRectsParams{
		Show: show,
	}
}

// Do executes Rendering.setShowScrollBottleneckRects.
func (p *SetShowScrollBottleneckRectsParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandRenderingSetShowScrollBottleneckRects, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}

// SetShowViewportSizeOnResizeParams paints viewport size upon main frame
// resize.
type SetShowViewportSizeOnResizeParams struct {
	Show bool `json:"show"` // Whether to paint size or not.
}

// SetShowViewportSizeOnResize paints viewport size upon main frame resize.
//
// parameters:
//   show - Whether to paint size or not.
func SetShowViewportSizeOnResize(show bool) *SetShowViewportSizeOnResizeParams {
	return &SetShowViewportSizeOnResizeParams{
		Show: show,
	}
}

// Do executes Rendering.setShowViewportSizeOnResize.
func (p *SetShowViewportSizeOnResizeParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandRenderingSetShowViewportSizeOnResize, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}
