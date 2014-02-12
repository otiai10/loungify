package controllers

import (
    "code.google.com/p/go.net/websocket"
    "github.com/robfig/revel"

    "fmt"
)

type WebSocket struct {
    *revel.Controller
}

func (c WebSocket)Socket(ws *websocket.Conn) revel.Result {

    // communication between goroutine & goroutine
    simpleChan := make(chan string)

    // goroutine: Listener
    go func() {
        var message string
        for {
            err := websocket.Message.Receive(ws, &message)
            // 何かしらエラーが起きたら
            if err != nil {
                // err === "EOF"
                // fmt.Printf("err?%T\n", err)
                // fmt.Println("チャネルを閉じる")
                close(simpleChan)
                // fmt.Println("ルーチンを終わる")
                return
            }
            // エラーが無ければ、チャネルを通す
            simpleChan <- message
        }
    }()

    // ループ: Broadcaster
    for {
        // チャネルからstringが来るのをまっている
        // ここで同期処理は一時待機になる
        fmt.Println("Waiting for channel...")
        message, stillOpen := <-simpleChan
        // チャネルから送信されてきたら再開するはず
        fmt.Println("Received from channel!")
        if !stillOpen {
            // チャネルが閉じているので、これ以上なにもしない
            return nil
        }
        // stillOpenなので、ウェブソケットで送る
        websocket.JSON.Send(ws, &map[string]string{"received":message})
    }

    return nil
}
