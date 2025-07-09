/**
 * フロントエンドアプリケーションのテストファイル
 */

const { greetUser, calculateSum } = require('../src/index.js');

function runTests() {
    console.log('Running frontend tests...');
    
    // Test greetUser function
    const greeting = greetUser('Test User');
    if (greeting === 'Hello, Test User! Welcome to our frontend app.') {
        console.log('✅ greetUser test passed');
    } else {
        console.log('❌ greetUser test failed');
        process.exit(1);
    }
    
    // Test calculateSum function
    const sum = calculateSum(2, 3);
    if (sum === 5) {
        console.log('✅ calculateSum test passed');
    } else {
        console.log('❌ calculateSum test failed');
        process.exit(1);
    }
    
    console.log('🎉 All frontend tests passed!');
}

runTests();