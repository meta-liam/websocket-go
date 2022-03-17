package main

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/meta-liam/websocket-go/v1/utils"
)

func main() {
	port := ":8080"
	fmt.Printf("[INFO]启动服务，端口：%+v\n", port)
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		// fmt.Printf("[INFO1]new connect: %+v\n", r.Header)
		if err != nil {
			// handle error
			fmt.Printf("[ERROR]%+v\n", err)
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				sendMsg := manageMsg(msg)
				if err != nil {
					// handle error
					fmt.Printf("[ERROR2]%+v\n", err)
					break
				}
				err = wsutil.WriteServerMessage(conn, op, sendMsg)
				if err != nil {
					// handle error
					fmt.Printf("[ERROR3]%+v\n", err)
				}
			}
		}()
	}))
}

func manageMsg(b []byte) []byte {
	return utils.ManageMsg(b)
}
