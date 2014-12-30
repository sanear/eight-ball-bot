package main

import (
	"log"
	"fmt"
	"time"
	"os"
	"net/url"
	"net/http"
	"strconv"
	"github.com/ChimeraCoder/anaconda"
	"github.com/sanear/eightBallBot/questionAnalyzer"
	"github.com/sanear/eightBallBot/eightBall"
)

func main() {

	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	go runDumbWebService(":" + port)

	anaconda.SetConsumerKey("VhSUdN4bdnVcXvrUcWCK8lFp6")
	anaconda.SetConsumerSecret("WNBwa0gv6Daadu4hOhL3fVK7mFZUlo4aoK67Dpgj7tNWiO30ia")
	api := anaconda.NewTwitterApi("2879998777-QT8IG6Ncr038vFs6MbITTixf85TXhacoqtmXAFX", "LcMWzd9hgXa6DgZFtwDKHdcSDVabClXTN5w4F6mEIQ43k")

	for {
		// Main functionality; find yes/no questions and answer them
		questions, err := findQuestions(api)
		if err != nil {
			log.Println("Search failed!", err)
		} else {
			for _, question := range questions {
				log.Printf("@%s asked, '%s'\n", question.User.ScreenName, question.Text)
				reply, err := answerQuestion(api, question)
				if err != nil {
					log.Printf("ERROR: Unable to post reply: %s\n", err)
				} else {
					log.Printf("Replied to tweet with: %s\n", reply.Text)
				}
			}
		}
		time.Sleep(20 * time.Second)
	}
}

func runDumbWebService(port string) {
	log.Printf("Starting Eight Ball webservice on port %s...\n", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
	log.Println("Stopping Eight Ball webservice...")
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "%s\n", eightBall.EightBallAnswer())
}

func findQuestions(api *anaconda.TwitterApi) (questions []anaconda.Tweet, err error) {
	v := url.Values{}
	v.Set("lang", "en")
	v.Set("count", "100")
	
	searchResult, err := api.GetSearch("?", v)
	if err != nil {
		return nil, err
	}

	for _, tweet := range searchResult.Statuses {
		if questionAnalyzer.IsYesNoQuestion(tweet.Text) {
			questions = append(questions, tweet)
		}
	}

	return questions, nil
}

func answerQuestion(api *anaconda.TwitterApi, question anaconda.Tweet) (reply anaconda.Tweet, err error) {
	answer := ".@" + question.User.ScreenName + " " + eightBall.EightBallAnswer()
	v := url.Values{}
	v.Set("in_reply_to_status_id", strconv.FormatInt(question.Id, 10))
	return api.PostTweet(answer, v)
}
