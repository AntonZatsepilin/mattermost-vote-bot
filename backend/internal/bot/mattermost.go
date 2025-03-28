package bot

import (
	"github.com/AntonZatsepilin/mattermost-vote-bot.git/internal/service"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/sirupsen/logrus"
)

type MattermostBot struct {
    Client *model.Client4
    WebSocket *model.WebSocketClient
    User    *model.User
    Service *service.Service
    Logger  *logrus.Logger
}

func NewBot(serverURL, token string) (*MattermostBot, error) {
    client := model.NewAPIv4Client(serverURL)
    client.SetToken(token)
    
    user, resp := client.GetMe("")
    if resp.Error != nil {
        return nil, resp.Error
    }
    
    ws, err := model.NewWebSocketClient4(serverURL, token)
    if err != nil {
        return nil, err
    }
    
    return &MattermostBot{
        Client: client,
        WebSocket: ws,
        User:    user,
        Logger:  logrus.New(),
    }, nil
}

func (b *MattermostBot) Start() {
    b.WebSocket.Listen()
    for {
        select {
        case event := <-b.WebSocket.EventChannel:
            b.handleEvent(event)
        }
    }
}

func (b *MattermostBot) handleEvent(event *model.WebSocketEvent) {
    if event.EventType() == model.WEBSOCKET_EVENT_POSTED {
        b.handleMessage(event.GetData()["post"].(string))
    }
}