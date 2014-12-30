package main

import (
	"log"
	"fmt"
	"time"
	"os"
	"io/ioutil"
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

	postLimit := 4
	if len(os.Args) > 2 {
		postLimit, _ = strconv.Atoi(os.Args[2])
	}	
	go runDumbWebService(":" + port)

	anaconda.SetConsumerKey("VhSUdN4bdnVcXvrUcWCK8lFp6")
	anaconda.SetConsumerSecret("WNBwa0gv6Daadu4hOhL3fVK7mFZUlo4aoK67Dpgj7tNWiO30ia")
	api := anaconda.NewTwitterApi("2879998777-QT8IG6Ncr038vFs6MbITTixf85TXhacoqtmXAFX", "LcMWzd9hgXa6DgZFtwDKHdcSDVabClXTN5w4F6mEIQ43k")

	lastIdFile := "lastId.txt"
	lastId, _ := readLastIdFromFile(lastIdFile)
	for {
		postCount := 0
		postCount, _ = answerQuestions(api, postCount, postLimit)
		postCount, lastId, _ = respondToMentions(api, postCount, postLimit, lastId)
		writeLastIdToFile(lastId, lastIdFile)
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

func answerQuestions(api *anaconda.TwitterApi, postCount int, postLimit int) (newPostCount int, err error) {
	// Main functionality; find yes/no questions and answer them
	questions, err := findQuestions(api)
	if err != nil {
		log.Println("ERROR: Failed to get questions!", err)
		return postCount, err
	} else {
		for _, question := range questions {
			if postCount < postLimit {
				log.Printf("@%s asked, '%s'\n", question.User.ScreenName, question.Text)
				reply, err := answerQuestion(api, question)
				if err != nil {
					log.Printf("ERROR: Unable to post reply: %s\n", err)
					return postCount, err
				} else {
					log.Printf("Replied to tweet with: %s\n", reply.Text)
					postCount++
				}
			} else {
				log.Printf("Too many requests have been made this cycle; sleeping for a bit.")
				return postCount, nil
			}
		}
	}
	return postCount, nil
}

func respondToMentions(api *anaconda.TwitterApi, postCount int, postLimit int, lastId string) (newPostCount int, newLastId string, err error) {
	// Also look for mentions and reply to them
	v := url.Values{}
	v.Set("since_id", string(lastId))
	mentions, err := api.GetMentionsTimeline(v)
	if err != nil {
		log.Println("ERROR: Failed to get mentions!", err)
		return postCount, lastId, err
	} else {
		for _, mention := range mentions {
			if postCount < postLimit {
				log.Printf("@%s mentioned me, saying '%s'\n", mention.User.ScreenName, mention.Text)
				reply, err := answerQuestion(api, mention)
				if err != nil {
					log.Printf("ERROR: Unable to post reply! %s\n", err)
					return postCount, lastId, err
				} else {
					lastId = strconv.FormatInt(mention.Id, 10)
					postCount++
					log.Printf("Replied to tweet with id %s, saying '%s'\n", lastId, reply.Text)
				}
			} else {
				log.Printf("Too many requests have been made this cycle; sleeping for a bit.")
				return postCount, lastId, nil
			}
		}
	}
	return postCount, lastId, nil
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
	answer := "@" + question.User.ScreenName + " " + eightBall.EightBallAnswer()
	v := url.Values{}
	v.Set("in_reply_to_status_id", strconv.FormatInt(question.Id, 10))
	return api.PostTweet(answer, v)
}

func writeLastIdToFile(lastId string, file string) (err error) {
	err = ioutil.WriteFile(file, []byte(lastId), 0644)
	if err != nil {
		log.Println("ERROR: failed to write to lastId file.", err)
		return err
	}
	return err
}

func readLastIdFromFile(file string) (lastId string, err error) {
	var lastIdBytes, e = ioutil.ReadFile(file)
	lastId = string(lastIdBytes)
	if e != nil {
		log.Println("ERROR: failed to read from lastId file.", err)
		return "1", e
	}
	return lastId, nil
}
