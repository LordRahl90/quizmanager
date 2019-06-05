# Quiz

This reads questions form the adjoining/specified csv file, presents it to the user in form of questions and answer. It also marks the responses and displays the score and the pass percentage. <br />

The questions file can be changed by adding -q flag to the execution

## Setup

* clone https://github.com/LordRahl90/quizmanager/tree/master/readquestion
* run `go run main.go -q questions.csv`

NB: "questions.csv" can be any csv file defined in the order of `"question","answer"`<br />
Also, It is expected that each question spans only a line.

## Credits
 * https://dev.to/adil_w3nomad/gopher-gym-quiz-game-part-1-4lbo
