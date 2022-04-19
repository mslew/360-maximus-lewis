//questionCreator(level) creates random question based on level, keeps track of level and returns calculated answer
//submitAnswer(correctAnswer) takes in user input answer, checks if correct -> green and next questions shown, if wrong red and allow 3 tries
//doMath(expr)  calculates answer given operator
//help() shows game instrucions (possibly give dynamic hints)

//this function will determine the level and will send the heavy lifting to other functions
function questionCreator(level){
    let max = 0; //max variable to generate random integers. used for generateRandomInt
    let expr = "";
    switch(level){
        case 1:
            max = 9;
            expr = createLevelOne(max);
            return expr;
        case 2:
            max = 9;
            expr = createLevelTwo(max);
            return expr;
        case 3:
            max = 9;
            expr = createLevelThree(max);
            return expr;
        case 4:
            max = 25;
            expr = createLevelFour(max);
            return expr;
        case 5:
            max = 50;
            expr = createLevelFive(max);
            return expr;
        case 6:
            max = 5;
            expr = createLevelSix(max);
            return expr;
        case 7:
            max = 12;
            expr = createLevelSeven(max);
            return expr;
        case 8:
            max = 6;
            expr = createLevelEight(max);
            return expr;
        case 9: 
            max = 12;
            expr = createLevelNine(max);
            return expr;
        case 10:
            let newLevel = Math.floor(Math.random() * (9 - 1 + 1) + 1); //generate a newLevel
            newLevel = 5;
            console.log(newLevel); 
            expr = questionCreator(newLevel); //recall this function to call the newLevel
            return expr;
    }
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
    console.log(answer);
    return answer;
}

function main(){
    let level = 1; //each game will start at level 1
    let score = 0; //each game will start with a score of 0
    let highScore = 0; //highscore will be kept track of
    let expr = "";
    let answer = 0;

    expr = questionCreator(10);
    console.log(expr);
    answer = doMath(expr);
}

main();