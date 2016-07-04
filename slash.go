package slash

import (
      "appengine"
      "appengine/urlfetch"
      "encoding/json"
      "fmt"
      "net/http"
      "net/url"
)

func init() {
      http.HandleFunc("/", handler)
      http.HandleFunc("/reminder", reminder_handler)
}

const (
      incoming_webhook_url string = "https://hooks.slack.com/services/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"
)

type Message struct {
      Channel    string `json:"channel"`
      Text       string `json:"text"`
      Username   string `json:"username"`
      Icon_emoji string `json:"icon_emoji"`
}

func handler(w http.ResponseWriter, r *http.Request) {

      //Read the Request Parameter "command"
      command := r.FormValue("command")

      //Ideally do other checks for tokens/username/etc

      if command == "/afk" {
            fmt.Fprint(w, "Success!!")
      } else {
            fmt.Fprint(w, "Failure")
      }
}

func reminder_handler(w http.ResponseWriter, r *http.Request) {
      c := appengine.NewContext(r)

      //Invoke the Slack Team + Channel Endpoint for incoming Webhook

      m := Message{"#afk", "This is the afk message", "AFK-Bot", ":computer:"}
      b, err := json.Marshal(m)
      if err != nil {
            c.Errorf("%v", err)
            return
      }

      client := urlfetch.Client(c)

      v := url.Values{}
      v.Set("payload", string(b))
      _, err = client.PostForm(incoming_webhook_url, v)
      if err != nil {
            c.Errorf("Exception while posting to Slack Channel : %v", err)
            return
      }
}
