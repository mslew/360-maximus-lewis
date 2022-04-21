let level = 1;
//this function will determine the level and will send the heavy lifting to other functions
function questionCreator(){
    let expr = "";
    let answer = 0; 
    switch(level){
        case 1:
            expr = createLevelOne(9);
            answer = doMath(expr);
            break;
        case 2:
            expr = createLevelTwo(9);
            answer = doMath(expr);
            break;
        case 3:
            expr = createLevelThree(9);
            answer = doMath(expr);
            break;
        case 4:
            expr = createLevelFour(25);
            answer = doMath(expr);
            break;
        case 5:
            expr = createLevelFive(50);
            answer = doMath(expr);
            break;
        case 6:
            expr = createLevelSix(5);
            answer = doMath(expr);
            break;
        case 7:
            expr = createLevelSeven(12);
            answer = doMath(expr);
            break;
        case 8:
            expr = createLevelEight(6);
            answer = doMath(expr);
            break;
        case 9: 
            expr = createLevelNine(12);
            answer = doMath(expr);
            break;
        case 10:
            let newLevel = Math.floor(Math.random() * (9 - 1 + 1) + 1); //generate a newLevel
            console.log(newLevel); 
            expr = questionCreator(newLevel); //recall this function to call the newLevel
            answer = doMath(expr);
            break;
    }
    console.log(level);
    document.getElementById("equationText").innerHTML = expr;
}

//will return a random number in the range of 0 - max. 
function generateRandomInt(max){
    return Math.floor(Math.random() * (max + 1));
}

//Level 1. Addition. Number Ranges: [0-9]
function createLevelOne(max){
    let num1 = generateRandomInt(max).toString();
    let num2 = generateRandomInt(max).toString();
    let expr = num1 + " + " + num2;
    return expr;
}

//Level 2. Subtraction. Number Ranges: [0-9]
function createLevelTwo(max){
    let num1 = generateRandomInt(max);
    let num2 = generateRandomInt(max);
    let expr = "";
    if (num1 < num2){ //include this if statement to avoid negative numbers
        expr = num2.toString() + " - " + num1.toString();
    }else{
        expr = num1.toString() + " - " + num2.toString();
    }
    return expr;
}

//Level 3. Addition and Subtraction. Number Ranges: [0-9]
function createLevelThree(max){
    let choice = Math.floor(Math.random() * (2 - 1 + 1) + 1); //inclusive range. Math.random() * (max - min + 1) + min
    switch(choice){
        case 1:
            return createLevelOne(max);
        case 2:
            return createLevelTwo(max);
    }
}

//Level 4. Addition and Subtraction. Number Ranges: [0-25]
function createLevelFour(max){
    return createLevelThree(max); //reuse this function but with a new max value
}

//Level 5. Addition and Subtraction. Number Ranges: [0-50]
function createLevelFive(max){
    return createLevelThree(max); //reuse this function but with a new max value
}

//Level 6. Multiplication. Number Ranges: [0-5]
function createLevelSix(max){
    let num1 = generateRandomInt(max).toString();
    let num2 = generateRandomInt(max).toString();
    let expr = num1 + " * " + num2;
    return expr;
}

//Level 7. Multiplication. Number Ranges: [0-12]
function createLevelSeven(max){
    return createLevelSix(max); // reuse this function from before but with a new max
}

//Level 8. Division resulting in an integer. Number Ranges: Multiples of [0-6]
function createLevelEight(max){
    let proceed = false;
    let num1 = generateRandomInt(max);
    let num2 = generateRandomInt(max);
    while (proceed == false){ //loop until both numbers are not zero
        if (num1 == 0 || num2 == 0){
            num1 = generateRandomInt(max);
            num2 = generateRandomInt(max);
        }else{
            proceed = true;
        }
    }
    let prod = num1 * num2; //calculate the product b/c we know that anything divided by the product of either number it an integer
    let expr = prod.toString() + " / " + num1.toString();
    return expr;
}

//Level 9. Division resulting in an integer. Number Ranges for Division: Multiples of [0-12]
function createLevelNine(max){
    return createLevelEight(max);
}

//Level 10. All previous levels combined. questionCreator case 10 for this implementation. 


//this function does the math and returns the answer
function doMath(expr){
    let answer = 0;
    let strArray = expr.split(/(\s+)/);
    let num1 = parseInt(strArray[0]);
    let operator = strArray[2];
    let num2 = parseInt(strArray[4]);

    switch(operator){
        case "+":
            answer = num1 + num2;
            break;
        case "-":
            answer = num1 - num2;
            break;
        case "*":
            answer = num1 * num2;
            break;
        case "/":
            answer = num1 / num2;
            break;
    }
    return answer;
}

//checks if user input is valid
function submitChecker(answer){
    answer.trim();
    if(isNaN(answer) == true){
        console.log("Enter a number"); //change this to HTML element later
    }else{
        answer = parseInt(answer);
    }
    return answer;
}

//check if user answer is correct 
function checkAnswer(userAnswer, expr){
    let answer = doMath(expr);
    let isAnswer = false;
    if(answer == parseInt(userAnswer)){
        console.log("correct"); //change this to HTML later
        isAnswer = true;
    }else{
        console.log("incorrect"); //change this to HTML later
        isAnswer = false;
    }
}

//determine the amount of points that needs to be added if a correct answer is chosen
function determinePoints(isAnswer, level){
    let points = 0;
    if(isAnswer){
        points = level * 100
    }
    return points;
}

//function that keeps track of currentscore
function currentScore(currentScore, newScore){
    return (currentScore + newScore);
}

//function that determines the highscore and keeps track of that number
function highScore(highScore, currentScore){
    if (currentScore > highScore){
        return currentScore;
    }else{
        return highScore;
    }
}

//fuunction that displays help message if user is in need of instruction of how the app works
function help(){
    let msg = "This app is a comprehensive math tutor.\n"+
        "This app will progressively get harder.\n"+
        "It will start will start with addition, moving to subtraction, then multiplication and ending with division.\n" +
        "It will end with the final level 10 as a mixture of all levels.\n\n" +
        "Enter the answer to the equation into the box.\n" + 
        "If you are correct the submit button will turn green, if wrong the submit button will turn red.\n"+
        "The arrow button will restart the tutoring process.\n Good Luck!";
    document.getElementById("help").innerHTML = msg; 
}