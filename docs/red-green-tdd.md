# Red/green TDD

TDD stands for Test Driven Development. It's a programming style where you ensure every piece of code you write is accompanied by automated tests that demonstrate the code works.

The most disciplined form of TDD is test-first development. You write the automated tests first, confirm that they fail, then iterate on the implementation until the tests pass.

It's important to confirm that the tests fail before implementing the code to make them pass. If you skip that step you risk building a test that passes already, hence failing to exercise and confirm your new implementation.

That's what "red/green" means: the red phase watches the tests fail, then the green phase confirms that they now pass.

Reference:

https://simonwillison.net/guides/agentic-engineering-patterns/red-green-tdd/