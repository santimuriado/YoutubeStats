package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Response is the JSON structure that the Youtube API has.
type Response struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

//Items stores the ID and the channel stats.
type Items struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

//Stats stores the channel info as view count,subs count and video count.
type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

func GetChannel() (Items, error) {

	var response Response

	//GET request that includes every parameter from the query.
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)

	//Error check.
	if err != nil {
		fmt.Println(err)
		return Items{}, err
	}

	//Define query parameters and their values.
	q := req.URL.Query()

	//Stores env variables which have to be defined previously.
	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	q.Add("id", os.Getenv("CHANNEL_ID"))
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	//Request to the built URL.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Items{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.Status)

	//Read the body from the JSON respone.
	body, _ := ioutil.ReadAll(resp.Body)

	//Unmarshall in a Response struct.
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Items{}, err
	}

	//Only interested in the first item in the Items array.
	return response.Items[0], nil

}
