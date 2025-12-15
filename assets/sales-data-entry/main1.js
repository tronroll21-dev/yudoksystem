function updateTotals() {

const Machine1CashCount = document.getElementById("Machine1CashCount");
const Machine1CashAmount = document.getElementById("Machine1CashAmount");
const Machine1SettleCount = document.getElementById("Machine1SettleCount");
const Machine1SettleAmount = document.getElementById("Machine1SettleAmount");
const Machine1UnsettledCount = document.getElementById("Machine1UnsettledCount");
const Machine1UnsettledAmount = document.getElementById("Machine1UnsettledAmount");
const Machine1QrCount = document.getElementById("Machine1QrCount");
const Machine1QrAmount = document.getElementById("Machine1QrAmount");
const Machine1QrSettleCount = document.getElementById("Machine1QrSettleCount");
const Machine1QrSettleAmount = document.getElementById("Machine1QrSettleAmount");
const Machine1ECount = document.getElementById("Machine1ECount");
const Machine1EAmount = document.getElementById("Machine1EAmount");
const Machine1ESettleCount = document.getElementById("Machine1ESettleCount");
const Machine1ESettleAmount = document.getElementById("Machine1ESettleAmount");
const Machine2CashCount = document.getElementById("Machine2CashCount");
const Machine2CashAmount = document.getElementById("Machine2CashAmount");
const Machine2SettleCount = document.getElementById("Machine2SettleCount");
const Machine2SettleAmount = document.getElementById("Machine2SettleAmount");
const Machine2UnsettledCount = document.getElementById("Machine2UnsettledCount");
const Machine2UnsettledAmount = document.getElementById("Machine2UnsettledAmount");
const Machine2QrCount = document.getElementById("Machine2QrCount");
const Machine2QrAmount = document.getElementById("Machine2QrAmount");
const Machine2QrSettleCount = document.getElementById("Machine2QrSettleCount");
const Machine2QrSettleAmount = document.getElementById("Machine2QrSettleAmount");
const Machine2ECount = document.getElementById("Machine2ECount");
const Machine2EAmount = document.getElementById("Machine2EAmount");
const Machine2ESettleCount = document.getElementById("Machine2ESettleCount");
const Machine2ESettleAmount = document.getElementById("Machine2ESettleAmount");
const Machine3CashCount = document.getElementById("Machine3CashCount");
const Machine3CashAmount = document.getElementById("Machine3CashAmount");
const Machine3SettleCount = document.getElementById("Machine3SettleCount");
const Machine3SettleAmount = document.getElementById("Machine3SettleAmount");
const Machine3UnsettledCount = document.getElementById("Machine3UnsettledCount");
const Machine3UnsettledAmount = document.getElementById("Machine3UnsettledAmount");
const Machine3QrCount = document.getElementById("Machine3QrCount");
const Machine3QrAmount = document.getElementById("Machine3QrAmount");
const Machine3QrSettleCount = document.getElementById("Machine3QrSettleCount");
const Machine3QrSettleAmount = document.getElementById("Machine3QrSettleAmount");
const Machine3ECount = document.getElementById("Machine3ECount");
const Machine3EAmount = document.getElementById("Machine3EAmount");
const Machine3ESettleCount = document.getElementById("Machine3ESettleCount");
const Machine3ESettleAmount = document.getElementById("Machine3ESettleAmount");
const Machine4CashCount = document.getElementById("Machine4CashCount");
const Machine4CashAmount = document.getElementById("Machine4CashAmount");
const Machine4SettleCount = document.getElementById("Machine4SettleCount");
const Machine4SettleAmount = document.getElementById("Machine4SettleAmount");
const Machine4UnsettledCount = document.getElementById("Machine4UnsettledCount");
const Machine4UnsettledAmount = document.getElementById("Machine4UnsettledAmount");
const Machine4QrCount = document.getElementById("Machine4QrCount");
const Machine4QrAmount = document.getElementById("Machine4QrAmount");
const Machine4QrSettleCount = document.getElementById("Machine4QrSettleCount");
const Machine4QrSettleAmount = document.getElementById("Machine4QrSettleAmount");
const Machine4ECount = document.getElementById("Machine4ECount");
const Machine4EAmount = document.getElementById("Machine4EAmount");
const Machine4ESettleCount = document.getElementById("Machine4ESettleCount");
const Machine4ESettleAmount = document.getElementById("Machine4ESettleAmount");
const Machine5CashCount = document.getElementById("Machine5CashAmount");
const Machine5CashAmount = document.getElementById("Machine5CashAmount");
const Machine5SettleCount = document.getElementById("Machine5SettleCount");
const Machine5SettleAmount = document.getElementById("Machine5SettleAmount");
const Machine5UnsettledCount = document.getElementById("Machine5UnsettledCount");
const Machine5UnsettledAmount = document.getElementById("Machine5UnsettledAmount");
const Machine5QrCount = document.getElementById("Machine5QrCount");
const Machine5QrAmount = document.getElementById("Machine5QrAmount");
const Machine5QrSettleCount = document.getElementById("Machine5QrSettleCount");
const Machine5QrSettleAmount = document.getElementById("Machine5QrSettleAmount");
const Machine5ECount = document.getElementById("Machine5ECount");
const Machine5EAmount = document.getElementById("Machine5EAmount");
const Machine5ESettleCount = document.getElementById("Machine5ESettleCount");
const Machine5ESettleAmount = document.getElementById("Machine5ESettleAmount");

const CashCountTotal = document.getElementById("CashCountTotal");
const CashAmountTotal = document.getElementById("CashAmountTotal");
const SettleCountTotal = document.getElementById("SettleCountTotal");
const SettleAmountTotal = document.getElementById("SettleAmountTotal");
const UnsettledCountTotal = document.getElementById("UnsettledCountTotal");
const UnsettledAmountTotal = document.getElementById("UnsettledAmountTotal");
const QrCountTotal = document.getElementById("QrCountTotal");
const QrAmountTotal = document.getElementById("QrAmountTotal");
const QrSettleCountTotal = document.getElementById("QrSettleCountTotal");
const QrSettleAmountTotal = document.getElementById("QrSettleAmountTotal");
const ECountTotal = document.getElementById("ECountTotal");
const EAmountTotal = document.getElementById("EAmountTotal");
const ESettleCountTotal = document.getElementById("ESettleCountTotal");
const ESettleAmountTotal = document.getElementById("ESettleAmountTotal");

CashCountTotal = 
Machine1CashCount.valueAsNumber +
Machine2CashCount.valueAsNumber +
Machine3CashCount.valueAsNumber +
Machine4CashCount.valueAsNumber +
Machine5CashCount.valueAsNumber;
CashAmountTotal = 
Machine1CashAmount.valueAsNumber +
Machine2CashAmount.valueAsNumber +
Machine3CashAmount.valueAsNumber +
Machine4CashAmount.valueAsNumber +
Machine5CashAmount.valueAsNumber; 
SettleCountTotal = 
Machine1SettleCount.valueAsNumber +
Machine2SettleCount.valueAsNumber +
Machine3SettleCount.valueAsNumber +
Machine4SettleCount.valueAsNumber +
Machine5SettleCount.valueAsNumber; 
SettleAmountTotal = 
Machine1SettleAmount.valueAsNumber +
Machine2SettleAmount.valueAsNumber +
Machine3SettleAmount.valueAsNumber +
Machine4SettleAmount.valueAsNumber +
Machine5SettleAmount.valueAsNumber;  
UnsettledCountTotal = 
Machine1UnsettledCount.valueAsNumber +
Machine2UnsettledCount.valueAsNumber +
Machine3UnsettledCount.valueAsNumber +
Machine4UnsettledCount.valueAsNumber +
Machine5UnsettledCount.valueAsNumber;  
UnsettledAmountTotal = 
Machine1UnsettledAmount.valueAsNumber +
Machine2UnsettledAmount.valueAsNumber +
Machine3UnsettledAmount.valueAsNumber +
Machine4UnsettledAmount.valueAsNumber +
Machine5UnsettledAmount.valueAsNumber;   
QrCountTotal =  
Machine1QrCount.valueAsNumber +
Machine2QrCount.valueAsNumber +
Machine3QrCount.valueAsNumber +
Machine4QrCount.valueAsNumber +
Machine5QrCount.valueAsNumber;  
QrAmountTotal =  
Machine1QrAmount.valueAsNumber +
Machine2QrAmount.valueAsNumber +
Machine3QrAmount.valueAsNumber +
Machine4QrAmount.valueAsNumber +
Machine5QrAmount.valueAsNumber;   
QrSettleCountTotal =  
Machine1QrSettleCount.valueAsNumber +
Machine2QrSettleCount.valueAsNumber +
Machine3QrSettleCount.valueAsNumber +
Machine4QrSettleCount.valueAsNumber +
Machine5QrSettleCount.valueAsNumber;   
QrSettleAmountTotal =  
Machine1QrSettleAmount.valueAsNumber +
Machine2QrSettleAmount.valueAsNumber +
Machine3QrSettleAmount.valueAsNumber +
Machine4QrSettleAmount.valueAsNumber +
Machine5QrSettleAmount.valueAsNumber;   
ECountTotal =   
Machine1ECount.valueAsNumber +
Machine2ECount.valueAsNumber +
Machine3ECount.valueAsNumber +
Machine4ECount.valueAsNumber +
Machine5ECount.valueAsNumber;  
EAmountTotal =  
Machine1EAmount.valueAsNumber +
Machine2EAmount.valueAsNumber +
Machine3EAmount.valueAsNumber +
Machine4EAmount.valueAsNumber +
Machine5EAmount.valueAsNumber;   
ESettleCountTotal =  
Machine1ESettleCount.valueAsNumber +
Machine2ESettleCount.valueAsNumber +
Machine3ESettleCount.valueAsNumber +
Machine4ESettleCount.valueAsNumber +
Machine5ESettleCount.valueAsNumber;   
ESettleAmountTotal =  
Machine1ESettleAmount.valueAsNumber +
Machine2ESettleAmount.valueAsNumber +
Machine3ESettleAmount.valueAsNumber +
Machine4ESettleAmount.valueAsNumber +
Machine5ESettleAmount.valueAsNumber;
}



