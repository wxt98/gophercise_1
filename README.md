# Gophercise 1: Quiz Game
Provides simple math quiz questions in the CLI when run.  
Questions and answers supplied from a CSV file.  

## Installation
Requirements:  
* golang 1.18+

## Run
Simply run
```bash
go run main.go
```
Or if you wish to specify the CSV file to be used for the quizzes (default problems.csv), run
```bash
go run main.go --csv CSVFILE_NAME_HERE.csv
```

If you wish to modify the timelimit for the quizzes (default 30s), run
```bash
go run main.go --timelimit TIMELIMIT_IN_SECONDS_HERE
```

Both csv and timelimit flags can be used together