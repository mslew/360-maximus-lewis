# Validator With Testing
This is used to validate a .go project. 
So far it will check if there is a README.md file,
LICENSE.md file, if there are tabs in the .go files, 
and counts characters in .go files. 

## Run Validator
go run . in terimanl
This will run automatically, thre is no user input required. It will automatically take the directory and scan files accordingly. This will also give a report of findings.  

## Manually Testing Validator
In order to manually test this code you should understand that there are four main parts that the code is doing. 
    1) It checks if the working directory has a README.md file
    2) It checks if the working directory has a LICENSE.md file
    3) It checks the .go files in the working directory to have correct tab usage
    4) It checks the .go files in the working directory to have the preffered character count per line
The best way one can manually test the validator is to omit or add "errors"
For example, omit the README.md file, and check if the code still runs properly.

## Automatically Test Validator
This code comes with a val_test.go file in it. This will automatically test the functions in the code.
go test . in terminal