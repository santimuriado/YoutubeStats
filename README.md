# YoutubeStats

YoutubeStats is a very simple project to show how to work with the Youtube API in Golang.
It shows the amount of Subscriptions, View Count and Video Count for a channel providing the
channel id.
Instead of working with a standard http connection and having to request all the time it
uses a websocket connection which updates the numbers every 5 seconds.
It also uses an extremely simple html file with some javascript necessary to use
the websocket connection and show the statistics.

# Requirements

Must have an API key from Google Cloud Platform and also enable Youtube Data API v3.
Should work with Go version 1.13 +.

# Run the Program

+ Clone the repo:

      git clone https://github.com/santimuriado/YoutubeStats.git

+ Generate the go.sum with:

      go mod tidy

+ Set the Env variables:

      export YOUTUBE_KEY="key"
      export CHANNEL_ID="channelId"

+ Run the server:

      go run main.go

+ To see the stats open index.html.
