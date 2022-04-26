let level = 1;
let lives = 3;
let levelCounter = 0;
let currentS = 0;
let highS = 0;
let GREEN = "#34a853";
let RED = "#dc3545";
let DEFAULT_COLOR = "#333";


function startGame(){
    questionCreator();
    let userAnswerInput = document.querySelector("#answerInput");
    if (userAnswerInput.addEventListener) {
      userAnswerInput.addEventListener("submit", submitAnswer(answer), false);
    } else if (userAnswerInput.attachEvent) {
      userAnswerInput.attachEvent("onsubmit", submitAnswer(answer));
    }
    
}

//resets the game 
function resetGame(){
    document.getElementById("heart3").style.display='block';
    document.getElementById("heart2").style.display='block';
    resetG = true;
    lives = 3;
    currentS = 0;
    highS = 0;
    level = 1;
    levelCounter = 0;
    for (let i = 2; i <= 10; i ++){ //reset level background colors
        let strLevel = i.toString();
        document.getElementById(strLevel).style.backgroundColor = "rgba(184, 20, 20, 0.774)";

    }
    //let a = questionCreator();   
    //submitAnswer(a);  creates two event listeners again, temp sol. answer old question works
}

//this function will determine the level and will send the heavy lifting to other functions
function  questionCreator(){
    checkLevel();
    let expr = "";
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
    document.getElementById("equationText").innerHTML = expr;
    console.log("QCreator answer: " + answer);
    return answer
}

function submitAnswer(answer){
    document.querySelector("#mathForm").addEventListener("submit", function(e) {
        e.stopPropagation(); //waits for response
        e.preventDefault(); //no refresh 
        var userAnswer = document.querySelector("#answerInput").value;
        console.log("answer = " + answer);
    console.log("userAnswer = " + userAnswer);
    var isCorrect = (answer == userAnswer) ? true : false;
    if (isCorrect){
        console.log("correct"); //change this to HTML later
        levelCounter +=  1/5; 
        var roundedL = Math.round(levelCounter * 10) / 10

        if (roundedL > .99 && roundedL == 1)   { //is whole number, handles wonky rounding
            level += levelCounter; 
            levelCounter = 0; rounded = 0;
        }
        //Display correct text on top, set background to green
        document.body.style.backgroundColor = GREEN;
        setTimeout(function() {
            document.body.style.backgroundColor = DEFAULT_COLOR;
        }, 1000);
    document.querySelector("#answerInput").value = "";
    answer =  questionCreator();
    
    }else{
        lives --;
        checkLives();
        console.log("incorrect"); //change this to HTML later
        isAnswer = false;

        // Display incorrect text on top, set background to red
        document.body.style.backgroundColor = RED;
    setTimeout(function() {
        document.body.style.backgroundColor = "#333";
    }, 1000);
    }
    })

function checkLives(){
    if (lives == 2 ){
        document.getElementById("heart3").style.display='none';
        
    }else if( lives ==1){
        document.getElementById("heart2").style.display='none';

    }else {
        document.getElementById("heart1").style.display='none';
        setTimeout(function() {
            alert("GAME OVER\nPress Enter to Play Again!");
            document.location.reload();
        }, 200);  
        }
    }  
    
}
function checkLevel(){
    let strLevel = level.toString();
    document.getElementById(strLevel).style.backgroundColor = "#ffffff";
    console.log("Level: "+ level);
    console.log("LevelCounter: " + levelCounter);
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

//determine the amount of points that needs to be added if a correct answer is chosen
function determinePoints(isAnswer, level){
    let points = 0;
    if(isAnswer){
        points = level * 100
    }
    return points;
}

//function that keeps track of currentscore
function updateCurrentScore(currentScore, newScore){
    return (currentScore + newScore);
}

//function that determines the highscore and keeps track of that number
function determineHighScore(highScore, currentScore){
    if (currentScore > highScore){
        //change highscore to be currentscore
        document.cookie = currentScore;
    }else{
        //leave highscore
    }
}

function readHighScore(highScore){
    let thisScore = document.cookie;
    if (highScore == null){
        document.cookie = highScore;
    }
    return highScore;
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
    setTimeout(function() {
        document.getElementById("help").innerHTML = ""
      }, 10000);  // 10 sec timeout funciton for help, could also refresh on any input
}