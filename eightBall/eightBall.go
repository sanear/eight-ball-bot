package eightBall

import (
	"math/rand"
	"time"
)

const answersFile string = "answers.prop"

func EightBallAnswer() string {

	rand.Seed(int64(time.Now().Nanosecond()))

	answers := parseAnswers()
	return answers[rand.Intn(len(answers))]
}

func parseAnswers() []string{

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
		"If you like salamanders, then yes. Otherwise, God help you.",
		"Just calm down and have a pizza.",
		"One of these days, Alice.",
		"Sure.",
		"Why not?",
		"In some universes.",
		"Oh, stahp.",
		"No! Panic!",
		"Ask again.",
		"I swear, I've heard this one before.",
		"You know...",
		"By Jove, certainly not!",
		"Bullocks to that.",
		"In some cases.",
		"Do you really think that's appropriate?",
		"In some countries, that question is punishable by hanging.",
		"What?",
		"Who's asking?",
		"Contemplate more deeply, and ask again.",
		"No. Should you drink more? Probably.",
		"If you were in the military, you wouldn't have time for such trivialities.",
		"Do you even lift, brah?",
		"It is a truth universally acknowledged.",
		"Sounds like someone's stressed.",
		"Whoa, now.",
	}
}
