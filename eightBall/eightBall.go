package eightBall

import (
	"math/rand"
	"time"
)

const answersFile string = "answers.prop"

func EightBallAnswer() string {

	rand.Seed(int64(time.Now().Nanosecond()))

	answers := parseAnswers(answersFile)
	return answers[rand.Intn(len(answers))]
}

func parseAnswers(filename string) []string{

	return []string {
		"Probably not.",
		"Maybe?",
		"What if it were?",
		"Would that really change anything?",
		"Next question.",
		"Hmm...",
		"Everything I say is false.",
		"Try rolling a die; primes mean 'Nuh uh.'",
		"I'm not sure, but it's definitely worth worrying about.",
		"Keep fretting and ask again.",
		"How important can it be if you're asking Twitter?",
		"Will you be happier if I say yes?",
		"Each day, we die a little more.",
		"Sometimes.",
		"Sometimes.",
		"Sometimes.",
		"Sometimes.",
		"Sometimes.",
		"#NOOOOOOOOOOO",
		"#NOOOOOOOOOOO",
		"#NOOOOOOOOOOO",
		"#NOOOOOOOOOOO",
		"#NOOOOOOOOOOO",
		"#YESSSSSSSSSS",
		"#YESSSSSSSSSS",
		"#YESSSSSSSSSS",
		"#YESSSSSSSSSS",
		"#YESSSSSSSSSS",
		"#YESSSSSSSSSS",		
		"Aye.",
		"Surely not.",
		"Do you really have to ask?",
		"Don't you think you know, you know, in your heart?",
		"Search your feelings; you know it to be true.",
		"... asked the grasshopper of the ant.",
		"Only insofar as you can experience it on this plane of existence, maan.",
		"Oh! I really know this one! It's, uh...",
		"You go, girl!",
		"You already know the answer.",
		"No YOU are.",
		"Does the Pope help cover up an international child-abuse ring?",
		"If you like salamanders, then yes. Otherwise, God help you.",
		"Just calm down and have a pizza.",
		"One of these days, Alice.",
	}
}
