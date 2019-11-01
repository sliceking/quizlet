# quizlet
a small command line quizzer!

## installation
go install github.com/svwielga4/quizlet

## how to use
Add quizlet to your path
Create a csv file with this format:

    problem,answer
    problem,answer
    problem,answer

Now run quizlet on the file to be quizzed

    quizlet -csv=<file> -time=<seconds>

or from inside the quizlet dir

    go run main.go -csv=<file> -time=<seconds>
