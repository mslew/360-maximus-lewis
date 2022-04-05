# Validator With Testing
<p>This is used to validate a .go project.<br>
So far it will check if there is a README.md file,<br>
LICENSE.md file, if there are tabs in the .go files,<br> 
and counts characters in .go files.</p>

## Run Validator
`go run .`
<br>
This will run automatically, thre is no user input required. It will automatically take the directory and scan files accordingly.<br> This will also give a report of findings.</p>  

## Manually Testing Validator
<p>In order to manually test this code you should understand that there are four main parts that the code is doing.<br>
<ol>
    <li> It checks if the working directory has a README.md file</li><br>
    <li> It checks if the working directory has a LICENSE.md file</li><br>
    <li> It checks the .go files in the working directory to have correct tab usage</li><br>
    <li> It checks the .go files in the working directory to have the preffered character count per line</li>
</ol>
The best way one can manually test the validator is to omit or add "errors"<br>
For example, omit the README.md file, and check if the code still runs properly.</p>

## Automatically Test Validator
<p>This code comes with a val_test.go file in it. This will automatically test the functions in the code.</p>

`go test .`
