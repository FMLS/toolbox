package sessionmgr

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/FMLS/toolbox/components/tat/api"
	"github.com/gorilla/websocket"
)

type SessionManager struct {
	ApiHost    string
	ApiPort    int
	WsHost     string
	WsPort     int
	Region     string
	AppID      int
	Uin        string
	InstanceID string
	SessionID  string
	ws         *websocket.Conn
}

func NewSessionManager(region, apiHost, wsHost string, apiPort, wsPort, appID int, uin, instanceId string) *SessionManager {
	return &SessionManager{
		Region:     region,
		ApiHost:    apiHost,
		ApiPort:    apiPort,
		WsHost:     wsHost,
		WsPort:     wsPort,
		AppID:      appID,
		Uin:        uin,
		InstanceID: instanceId,
	}
}

func (sm *SessionManager) StartSession() error {
	resp, err := api.NewAPI(sm.ApiHost, sm.ApiPort, sm.AppID, sm.Region, sm.Uin).StartSession(sm.InstanceID)
	if err != nil {
		return fmt.Errorf("StartSession:%w", err)
	}
	log.Printf("SessionID:%s", resp.Response.SessionID)
	sm.SessionID = resp.Response.SessionID

	url := fmt.Sprintf("ws://%s:%d/ws", sm.WsHost, sm.WsPort)
	log.Printf("connect to %s", url)
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("websocket.DefaultDialer.Dial:%w", err)
	}

	sm.ws = ws
	return nil
}

func (sm *SessionManager) SessionAuth() error {
	return sm.ws.WriteJSON(SessionAuth{
		Type: "SessionAuth",
		Data: struct {
			AppID      int `json:"AppId"`
			Uin        string
			SessionID  string `json:"SessionId"`
			InstanceID string `json:"InstanceId"`
		}{
			sm.AppID,
			sm.Uin,
			sm.SessionID,
			sm.InstanceID,
		},
	})
}

func (sm *SessionManager) WaitPtyReady(timeout time.Duration) error {
	if err := sm.ws.SetReadDeadline(time.Now().Add(timeout)); err != nil {
		return fmt.Errorf("sm.ws.SetReadDeadline:%w", err)
	}
	r := PtyReady{}
	if err := sm.ws.ReadJSON(&r); err != nil {
		return fmt.Errorf("sm.ws.ReadJSON:%w", err)
	}
	log.Printf("->%+v", r)
	if r.Type != "PtyReady" {
		return errors.New("receive PtyReady error")
	}
	return nil
}

func (sm *SessionManager) PtyInput(input string, timeout time.Duration) error {
	if err := sm.ws.SetWriteDeadline(time.Now().Add(timeout)); err != nil {
		return fmt.Errorf("sm.ws.SetWriteDeadline:%w", err)
	}
	msg := PtyInput{
		Type: "PtyInput",
		Data: struct {
			SessionID string `json:"SessionId"`
			Input     string
		}{
			sm.SessionID,
			base64.StdEncoding.EncodeToString([]byte(input)),
		},
	}
	if err := sm.ws.WriteJSON(msg); err != nil {
		return fmt.Errorf("sm.ws.WriteJSON:%w", err)
	}
	log.Printf("<-%+v", msg)
	return nil
}

func (sm *SessionManager) PtyStart(timeout time.Duration) error {
	if err := sm.ws.SetWriteDeadline(time.Now().Add(timeout)); err != nil {
		return fmt.Errorf("sm.ws.SetWriteDeadline:%w", err)
	}
	msg := PtyStart{
		Type: "PtyStart",
		Data: struct {
			SessionID string `json:"SessionId"`
			Cols      int
			Rows      int
		}{
			sm.SessionID,
			80,
			40,
		},
	}
	if err := sm.ws.WriteJSON(msg); err != nil {
		return fmt.Errorf("sm.ws.WriteJSON:%w", err)
	}
	log.Printf("<-%+v", msg)
	return nil
}

func (sm *SessionManager) TestConn() error {
	if err := sm.StartSession(); err != nil {
		return fmt.Errorf("sm.StartSession:%w", err)
	}
	if err := sm.SessionAuth(); err != nil {
		return fmt.Errorf("sm.SessionAuth:%w", err)
	}
	if err := sm.PtyStart(time.Second); err != nil {
		return fmt.Errorf("sm.PtyStart:%w", err)
	}
	if err := sm.WaitPtyReady(time.Second); err != nil {
		return fmt.Errorf("sm.WaitPtyReady:%w", err)
	}
	go func() {
		for {
			if err := sm.PtyInput("", time.Second); err != nil {
				log.Printf("sm.PtyInput:%v", err)
				return
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		_ = sm.ws.SetReadDeadline(time.Time{})
		ptyOutput := PtyOutput{}
		if err := sm.ws.ReadJSON(&ptyOutput); err != nil {
			return fmt.Errorf("sm.ws.ReadJSON:%w", err)
		}
		log.Printf("->%+v", ptyOutput)
	}
}
