package obsws

import (
	"errors"
	"time"
)

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// SetCurrentSceneRequest : Switch to the specified scene.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.3-maintenance/docs/generated/protocol.md#setcurrentscene
type SetCurrentSceneRequest struct {
	// Name of the scene to switch to.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	_request  `json:",squash"`
	response  chan SetCurrentSceneResponse
}

// NewSetCurrentSceneRequest returns a new SetCurrentSceneRequest.
func NewSetCurrentSceneRequest(sceneName string) SetCurrentSceneRequest {
	return SetCurrentSceneRequest{
		sceneName,
		_request{
			ID_:   getMessageID(),
			Type_: "SetCurrentScene",
			err:   make(chan error, 1),
		},
		make(chan SetCurrentSceneResponse, 1),
	}
}

// Send sends the request.
func (r *SetCurrentSceneRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.sendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp SetCurrentSceneResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r SetCurrentSceneRequest) Receive() (SetCurrentSceneResponse, error) {
	if !r.sent {
		return SetCurrentSceneResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetCurrentSceneResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetCurrentSceneResponse{}, err
		case <-time.After(receiveTimeout):
			return SetCurrentSceneResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetCurrentSceneRequest) SendReceive(c Client) (SetCurrentSceneResponse, error) {
	if err := r.Send(c); err != nil {
		return SetCurrentSceneResponse{}, err
	}
	return r.Receive()
}

// SetCurrentSceneResponse : Response for SetCurrentSceneRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.3-maintenance/docs/generated/protocol.md#setcurrentscene
type SetCurrentSceneResponse struct {
	_response `json:",squash"`
}

// GetCurrentSceneRequest : Get the current scene's name and source items.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.3-maintenance/docs/generated/protocol.md#getcurrentscene
type GetCurrentSceneRequest struct {
	_request `json:",squash"`
	response chan GetCurrentSceneResponse
}

// NewGetCurrentSceneRequest returns a new GetCurrentSceneRequest.
func NewGetCurrentSceneRequest() GetCurrentSceneRequest {
	return GetCurrentSceneRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "GetCurrentScene",
			err:   make(chan error, 1),
		},
		make(chan GetCurrentSceneResponse, 1),
	}
}

// Send sends the request.
func (r *GetCurrentSceneRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.sendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp GetCurrentSceneResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r GetCurrentSceneRequest) Receive() (GetCurrentSceneResponse, error) {
	if !r.sent {
		return GetCurrentSceneResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetCurrentSceneResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetCurrentSceneResponse{}, err
		case <-time.After(receiveTimeout):
			return GetCurrentSceneResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetCurrentSceneRequest) SendReceive(c Client) (GetCurrentSceneResponse, error) {
	if err := r.Send(c); err != nil {
		return GetCurrentSceneResponse{}, err
	}
	return r.Receive()
}

// GetCurrentSceneResponse : Response for GetCurrentSceneRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.3-maintenance/docs/generated/protocol.md#getcurrentscene
type GetCurrentSceneResponse struct {
	// Name of the currently active scene.
	// Required: Yes.
	Name string `json:"name"`
	// Ordered list of the current scene's source items.
	// Required: Yes.
	Sources   []map[string]interface{} `json:"sources"`
	_response `json:",squash"`
}

// GetSceneListRequest : Get a list of scenes in the currently active profile.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.3-maintenance/docs/generated/protocol.md#getscenelist
type GetSceneListRequest struct {
	_request `json:",squash"`
	response chan GetSceneListResponse
}

// NewGetSceneListRequest returns a new GetSceneListRequest.
func NewGetSceneListRequest() GetSceneListRequest {
	return GetSceneListRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "GetSceneList",
			err:   make(chan error, 1),
		},
		make(chan GetSceneListResponse, 1),
	}
}

// Send sends the request.
func (r *GetSceneListRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.sendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp GetSceneListResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r GetSceneListRequest) Receive() (GetSceneListResponse, error) {
	if !r.sent {
		return GetSceneListResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetSceneListResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetSceneListResponse{}, err
		case <-time.After(receiveTimeout):
			return GetSceneListResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetSceneListRequest) SendReceive(c Client) (GetSceneListResponse, error) {
	if err := r.Send(c); err != nil {
		return GetSceneListResponse{}, err
	}
	return r.Receive()
}

// GetSceneListResponse : Response for GetSceneListRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.3-maintenance/docs/generated/protocol.md#getscenelist
type GetSceneListResponse struct {
	// Name of the currently active scene.
	// Required: Yes.
	CurrentScene string `json:"current-scene"`
	// Ordered list of the current profile's scenes (See `[GetCurrentScene](#getcurrentscene)` for more information).
	// Required: Yes.
	Scenes    []map[string]interface{} `json:"scenes"`
	_response `json:",squash"`
}
