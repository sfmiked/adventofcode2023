# Advent of Code 2023

This year, after much debate I will be using golang for the annual [advent of code](https://adventofcode.com/2023) competition. 

While this is a repeat of 2020, golang and it's ecosystem seems to have moved on quite a bit since then.

Of the features I know about, generics is the one that I expect will be the most useful.

While I have in previous years created test cases, this year I've been challenged to stick to a more tdd-focused approach (rather than simply test to execute). 

## Badges
![Build Status](https://github.com/sfmiked/adventofcode2023/actions/workflows/go.yml/badge.svg)
[![codecov](https://codecov.io/gh/sfmiked/adventofcode2023/graph/badge.svg?token=63MRY6V67B)](https://codecov.io/gh/sfmiked/adventofcode2023)


## setup steps

### command line
```
# create go.mod file
go mod init github.com/miketzian/adventofcode2023

# add testify dependency
go get github.com/stretchr/testify

# add cobra dependency for cli
go get -u github.com/spf13/cobra@latest
```

### devcontainer

1. In vscode - Add Dev Container Configuration Files
2. Choose 'Go' option
3. No options selected

### github 

1. In actions, choose [https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go](go action)
2. In test workflow step, add coverage option (by default no coverage report)
3. Login to codecov*, get the secret, create the repository secret
4. Add the codecov workflow action

*note! codecov.io wants to know our email address. They're not allowed to know this, so I had to create another dummy github account with an email address they're allowed to know. 