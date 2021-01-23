# alexflipnote.go
A modern and small API wrapper for the [AlexFlipnote API](https://api.alexflipnote.dev). It handles authorization and everything needed to access and use the API.
## Example usage
### Simple
```go
package main
import (
    "fmt"
    alexflipnote "github.com/lavgup/alexflipnote.go"
)
func main() {
    client := alexflipnote.NewClient("TOKEN")
    body, err := client.Facts("butch and sundance are back, baby!")
    if err != nil {
        fmt.Println(err)
    }
    
    // WARNING: This will spam your console
    fmt.Printf("Received bytes: %s", body)
}
```
### Image
```go
package main
import (
    "fmt"
    alexflipnote "github.com/lavgup/alexflipnote.go"
    "net/url"
)
func main() {
    client := alexflipnote.NewClient("TOKEN")
    
    imageUrl := "https://cdn.discordapp.com/avatars/441164156016787486/7a9cc8980bed842503c451efc79b74f7.png"
    parsedUrl, err := url.Parse(imageUrl)
    if err != nil {
        fmt.Println(err)
    }
    body, err := client.AmIAJoke(parsedUrl)
    if err != nil {
        fmt.Println(err)
    }
    // WARNING: This will spam your console
    fmt.Printf("Received bytes: %s", body)
}
```
The image returns are of type `[]byte`, so to send it via discordgo, wrap it with a Buffer e.g `session.ChannelFileSend(id, "image.png", bytes.NewBuffer(image))`