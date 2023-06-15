package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	Sentences []struct {
		Trans string `json:"trans"`
	} `json:"sentences"`
}

func Translate(message string, langue string) string {

	body := api_resquests(message, langue)

	return extract_trans(body)
}

func api_resquests(message string, langue string) []byte {

	client := &http.Client{}

	url := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=%s&hl=%s&dt=t&dt=bd&dj=1&source=input&q=%s", langue, langue, url.QueryEscape(message))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête :", err)
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la requête :", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse :", err)
		return nil
	}

	return body
}

func extract_trans(body []byte) string {

	if body == nil {
		return "daz"
	}

	var data Response
	err := json.Unmarshal(body, &data)
	if err != nil {
		return err.Error()
	}

	trans := data.Sentences[0].Trans

	return trans
}
