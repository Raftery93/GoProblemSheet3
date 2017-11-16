package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

func ElizaResponse() {

	for reader := bufio.NewReader(os.Stdin); ; {
		//userInput := " "

		elizaOutputs := []string{
			"I’m not sure what you’re trying to say. Could you explain it to me?",
			"How does that make you feel?",
			"Why do you say that?",
		}

		match := regexp.MustCompile(`(?i)\bfather\b`)
		//match2 := regexp.MustCompile(`(?i)^(I am|I'm|im) ([^\.\?!]*)`)

		/*
			testInputs := []string{

				"People say I look like both my mother and father.",
				"Father was a teacher.",
				"I was my father’s favourite.",
				"I’m looking forward to the weekend.",
				"My grandfather was French!",
				"I am happy.",
				"I am not happy with your responses.",
				"I am not sure that you understand the effect that your questions are having on me.",
				"I am supposed to just take what you’re saying at face value?",
			}
		*/

		userInput, _ := reader.ReadString('\n')

		userInput = strings.Trim(userInput, "\r\n")

		//fmt.Println(elizaOutputs[rand.Intn(len(userInputs))])
		//strings.ToLower(userInput)

		if match.MatchString(userInput) {
			//if strings.EqualFold(strings.ToLower(userInput), "father") {
			fmt.Println("Why don’t you tell me more about your father?")
		} else {

			//match := regexp.MustCompile(`(?i)^(I am|I'm|im) ([^\.\?!]*)`)
			match := regexp.MustCompile(`(?i)\bi am\b|\bI am\b|\bim\b|\bIm\b|\bI'm\b|\bi'm\b`)

			if match.MatchString(userInput) {
				fmt.Println("laaaaawd")
			} else {

				rand.Seed(time.Now().UTC().UnixNano())
				fmt.Println(elizaOutputs[rand.Intn(len(elizaOutputs))])

			}
		}

		if strings.Compare(strings.ToLower(strings.TrimSpace(userInput)), "quit") == 0 {
			break
		}
	}

}

func main() {

	ElizaResponse()
}
