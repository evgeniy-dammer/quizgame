# "Quiz Game" exercise from Gophercises (by John Calhoun)

Website: https://gophercises.com

GitHub: https://github.com/gophercises

YouTube: https://www.youtube.com/playlist?list=PLVEltXlEeWglGINo25GxVfvSSylLVg4r1

## Part 1
Create a program that will read in a quiz provided via a CSV file and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards. 

The CSV file should default to problems.csv, but the user should be able to customize the filename via a flag. 

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

## Part 2
Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn't wait for the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.
