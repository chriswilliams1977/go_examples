#tests run your code with specific inputs and expect specific outputs
#if outputs match your test passes otherwise it fails

#go uses the testing package
#use the go test command to run the tests

#test file has to
#import testing package
#func name begins with Test
#accept single parama of pointer to testing.T
#report errors using testing.T error

#run using go test with import package path
#go looks in package for files with _test.go
#for example go test github.com/chriswilliams1977/headfirst/playground/packages/testing

#use t.Errorf - to generate a more detailed error message i.e args, return value, value expected
#test command only runs functions with Test at front so you can include helper functions to test
#example refactor repeating t.Error to a new function

#TDD
#Write the test - run it and it will fail
#make it pass - write the func in main code to make test pass
#refactor code - change and improve code know tests will fail if something breaks

#when compiling code go ignores files with _test

#can pass flags to command line to run specific tests or see output
#go test ... -v shows status of each test
#go test ... -run Two only runs tests with Two in name, add -v to see output

#way to reduce repeated tests is to have a struct that hold inout and output
#then pass this to a test