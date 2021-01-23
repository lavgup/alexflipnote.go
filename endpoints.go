package alexflipnote

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

func (client *Client) Achievement(text string, icon int) ([]byte, error) {
	return client.MakeRequest("achievement", url.Values{
		"text": {text},
		"icon": {strconv.Itoa(icon)},
	})
}

func (client *Client) AmIAJoke(image *url.URL) ([]byte, error) {
	return client.MakeRequest("amiajoke", url.Values{
		"image": {image.String()},
	})
}

func (client *Client) Bad(image url.URL) ([]byte, error) {
	return client.MakeRequest("bad", url.Values{
		"image": {image.String()},
	})
}

func (client *Client) Birb() ([]byte, error) {
	return client.MakeRequest("birb", url.Values{})
}

func (client *Client) Calling(text string) ([]byte, error) {
	return client.MakeRequest("calling", url.Values{
		"text": {text},
	})
}

func (client *Client) Captcha(text string) ([]byte, error) {
	return client.MakeRequest("captcha", url.Values{
		"text": {text},
	})
}

func (client *Client) Cats() ([]byte, error) {
	return client.MakeRequest("cats", url.Values{})
}

func (client *Client) Challenge(text string, icon int) ([]byte, error) {
	return client.MakeRequest("achievement", url.Values{
		"text": {text},
		"icon": {strconv.Itoa(icon)},
	})
}

func (client *Client) ColorHex(hex string) ([]byte, error) {
	return client.MakeRequest(fmt.Sprint("color/", hex), url.Values{})
}

func (client *Client) ColorGitHub() ([]byte, error) {
	return client.MakeRequest("color/github", url.Values{})
}

func (client *Client) ColorImage(hex string) ([]byte, error) {
	return client.MakeRequest(fmt.Sprint("color/image/", hex), url.Values{})
}

func (client *Client) ColorImageGradient(hex string) ([]byte, error) {
	return client.MakeRequest(fmt.Sprint("color/image/gradient/", hex), url.Values{})
}

func (client *Client) Colourify(image url.URL, c, b string) ([]byte, error) {
	return client.MakeRequest("colourify", url.Values{
		"image": {image.String()},
		"c":     {c},
		"b":     {b},
	})
}

func (client *Client) DidYouMean(top, bottom string) ([]byte, error) {
	return client.MakeRequest("didyoumean", url.Values{
		"top":    {top},
		"bottom": {bottom},
	})
}

func (client *Client) Dogs() ([]byte, error) {
	return client.MakeRequest("dogs", url.Values{})
}

func (client *Client) Drake(top, bottom string) ([]byte, error) {
	return client.MakeRequest("drake", url.Values{
		"top":    {top},
		"bottom": {bottom},
	})
}

func (client *Client) Facts(text string) ([]byte, error) {
	return client.MakeRequest("facts", url.Values{
		"text": {text},
	})
}

func (client *Client) Filter(filter string, image url.URL) ([]byte, error) {
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

func (client *Client) Floor(image url.URL, text string) ([]byte, error) {
	return client.MakeRequest("floor", url.Values{
		"image": {image.String()},
		"text":  {text},
	})
}

func (client *Client) Fml() ([]byte, error) {
	return client.MakeRequest("fml", url.Values{})
}

func (client *Client) JokeOverHead(image url.URL) ([]byte, error) {
	return client.MakeRequest("jokeoverhead", url.Values{
		"image": {image.String()},
	})
}

func (client *Client) Pornhub(text, text2 string) ([]byte, error) {
	return client.MakeRequest("pornhub", url.Values{
		"text":  {text},
		"text2": {text2},
	})
}

func (client *Client) SadCat() ([]byte, error) {
	return client.MakeRequest("sadcat", url.Values{})
}

func (client *Client) Salty(image url.URL) ([]byte, error) {
	return client.MakeRequest("salty", url.Values{
		"image": {image.String()},
	})
}

func (client *Client) Scroll(text string) ([]byte, error) {
	return client.MakeRequest("scroll", url.Values{
		"image": {text},
	})
}

func (client *Client) Shame(image url.URL) ([]byte, error) {
	return client.MakeRequest("shame", url.Values{
		"image": {image.String()},
	})
}

func (client *Client) Ship(user, user2 url.URL) ([]byte, error) {
	return client.MakeRequest("salty", url.Values{
		"user":  {user.String()},
		"user2": {user2.String()},
	})
}

func (client *Client) Supreme(text string, dark, light bool) ([]byte, error) {
	return client.MakeRequest("supreme", url.Values{
		"text":  {text},
		"dark":  {strconv.FormatBool(dark)},
		"light": {strconv.FormatBool(light)},
	})
}

func (client *Client) Trash(face, trash url.URL) ([]byte, error) {
	return client.MakeRequest("trash", url.Values{
		"face":  {face.String()},
		"trash": {trash.String()},
	})
}

func (client *Client) What(image url.URL) ([]byte, error) {
	return client.MakeRequest("what", url.Values{
		"image": {image.String()},
	})
}
