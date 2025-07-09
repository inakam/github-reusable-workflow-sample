/**
 * フロントエンドアプリケーションのメインファイル
 */

function greetUser(name) {
    return `Hello, ${name}! Welcome to our frontend app.`;
}

function calculateSum(a, b) {
    return a + b;
}

module.exports = {
    greetUser,
    calculateSum
};

console.log(greetUser('World'));