#all programs have errors and you need to handle these
#go offers different ways to achieve this
#defer statement ensure a function is run no matter what after all other code has executed
#this can therefore be used even if a function exits due to failure and you still want to close a file for example
#this is usually used for code that needs to run even if an error occurs

#recursion lets a function call itself

#Panic is when a function stops running and the program prints a log message and crashes
#you can call panic yourself and print a message

#Go has a call stack = func called needs to return to func that called it
#panic outputs call stack  which can be useful for debug
#deferred func calls are still made if program panics
#Not ideal as this crashes the program
#should be reserved for impossible situations

#recover stops a program from panicking
#it just returns nil
#if you call recover when its panicking it will stop the panic
#dont call recover in func that is panicking otherwise it already stops executing
#when a func panics deferred calls are completed
#so place recover in a deferred func
#if program is not panicing recover returns nil

#to handle unanticpated panics simply renew panic state