package alexflipnote

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// Achievement returns an "achievement" image
func (client *Client) Achievement(text string, icon int) (interface{}, error) {
	return client.MakeRequest("achievement", url.Values{
		"text": {text},
		"icon": {strconv.Itoa(icon)},
	})
}

// AmIAJoke returns an "am i a joke" image
func (client *Client) AmIAJoke(image *url.URL) (interface{}, error) {
	return client.MakeRequest("amiajoke", url.Values{
		"image": {image.String()},
	})
}

// Bad returns a "bad" image
func (client *Client) Bad(image url.URL) (interface{}, error) {
	return client.MakeRequest("bad", url.Values{
		"image": {image.String()},
	})
}

// Birb returns a "birb" image
func (client *Client) Birb() (interface{}, error) {
	return client.MakeRequest("birb", url.Values{})
}

// Calling returns a "calling" image
func (client *Client) Calling(text string) (interface{}, error) {
	return client.MakeRequest("calling", url.Values{
		"text": {text},
	})
}

// Captcha returns a "captcha" image
func (client *Client) Captcha(text string) (interface{}, error) {
	return client.MakeRequest("captcha", url.Values{
		"text": {text},
	})
}

// Cats returns a "cats" image
func (client *Client) Cats() (interface{}, error) {
	return client.MakeRequest("cats", url.Values{})
}

// Challenge returns a "challenge" image
func (client *Client) Challenge(text string, icon int) (interface{}, error) {
	return client.MakeRequest("achievement", url.Values{
		"text": {text},
		"icon": {strconv.Itoa(icon)},
	})
}

// ColorHex returns information on a hex code
func (client *Client) ColorHex(hex string) (interface{}, error) {
	return client.MakeRequest(fmt.Sprint("color/", hex), url.Values{})
}

// ColorGitHub returns information on the colors of programming languages on GitHub
func (client *Client) ColorGitHub() (interface{}, error) {
	return client.MakeRequest("color/github", url.Values{})
}

// ColorImage returns an image of a color
func (client *Client) ColorImage(hex string) (interface{}, error) {
	return client.MakeRequest(fmt.Sprint("color/image/", hex), url.Values{})
}

// ColorImageGradient returns a image of different gradients of a color
func (client *Client) ColorImageGradient(hex string) (interface{}, error) {
	return client.MakeRequest(fmt.Sprint("color/image/gradient/", hex), url.Values{})
}

// Colourify returns a "colourified" image
func (client *Client) Colourify(image url.URL, c, b string) (interface{}, error) {
	return client.MakeRequest("colourify", url.Values{
		"image": {image.String()},
		"c":     {c},
		"b":     {b},
	})
}

// DidYouMean returns a "did you mean" image
func (client *Client) DidYouMean(top, bottom string) (interface{}, error) {
	return client.MakeRequest("didyoumean", url.Values{
		"top":    {top},
		"bottom": {bottom},
	})
}

// Dogs returns a "dog" image
func (client *Client) Dogs() (interface{}, error) {
	return client.MakeRequest("dogs", url.Values{})
}

// Drake returns a "drake meme" image
func (client *Client) Drake(top, bottom string) (interface{}, error) {
	return client.MakeRequest("drake", url.Values{
		"top":    {top},
		"bottom": {bottom},
	})
}

// Facts returns a "facts" image
func (client *Client) Facts(text string) (interface{}, error) {
	return client.MakeRequest("facts", url.Values{
		"text": {text},
	})
}

// Filter returns a filtered image
func (client *Client) Filter(filter string, image url.URL) (interface{}, error) {
	switch filter {
	case "blue", "b&w", "communist", "deepfry",
		"gay", "invert", "jpegify", "pixelate",
		"magik", "sepia", "snow", "wide":
		return client.MakeRequest(fmt.Sprint("filter/", filter), url.Values{
			"image": {image.String()},
		})
	default:
		return nil, errors.New("invalid filter")
	}
}

// Floor returns a "floor" image
func (client *Client) Floor(image url.URL, text string) (interface{}, error) {
	return client.MakeRequest("floor", url.Values{
		"image": {image.String()},
		"text":  {text},
	})
}

// Fml returns an "fml" line
func (client *Client) Fml() (interface{}, error) {
	return client.MakeRequest("fml", url.Values{})
}

// JokeOverHead returns a "joke over head" image
func (client *Client) JokeOverHead(image url.URL) (interface{}, error) {
	return client.MakeRequest("jokeoverhead", url.Values{
		"image": {image.String()},
	})
}

// Pornhub returns a "pornhub" image
func (client *Client) Pornhub(text, text2 string) (interface{}, error) {
	return client.MakeRequest("pornhub", url.Values{
		"text":  {text},
		"text2": {text2},
	})
}

// SadCat returns a "sad cat" image
func (client *Client) SadCat() (interface{}, error) {
	return client.MakeRequest("sadcat", url.Values{})
}

// Salty returns a "salty" image
func (client *Client) Salty(image url.URL) (interface{}, error) {
	return client.MakeRequest("salty", url.Values{
		"image": {image.String()},
	})
}

// Scroll returns a "scroll" image
func (client *Client) Scroll(text string) (interface{}, error) {
	return client.MakeRequest("scroll", url.Values{
		"image": {text},
	})
}

// Shame returns a "shame" image
func (client *Client) Shame(image url.URL) (interface{}, error) {
	return client.MakeRequest("shame", url.Values{
		"image": {image.String()},
	})
}

// Ship returns a "ship" image
func (client *Client) Ship(user, user2 url.URL) (interface{}, error) {
	return client.MakeRequest("salty", url.Values{
		"user":  {user.String()},
		"user2": {user2.String()},
	})
}

// Supreme returns a "supreme" image
func (client *Client) Supreme(text string, dark, light bool) (interface{}, error) {
	return client.MakeRequest("supreme", url.Values{
		"text":  {text},
		"dark":  {strconv.FormatBool(dark)},
		"light": {strconv.FormatBool(light)},
	})
}

// Trash returns a "trash" image
func (client *Client) Trash(face, trash url.URL) (interface{}, error) {
	return client.MakeRequest("trash", url.Values{
		"face":  {face.String()},
		"trash": {trash.String()},
	})
}

// What returns a "what" image
func (client *Client) What(image url.URL) (interface{}, error) {
	return client.MakeRequest("what", url.Values{
		"image": {image.String()},
	})
}
