package client

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func PostJSON(url string, payload any) error {
    body, err := json.Marshal(payload)
    if err != nil {
        return err
    }

    _, err = http.Post(
        url,
        "application/json",
        bytes.NewBuffer(body),
    )

    return err
}
