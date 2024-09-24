module github.com/mrevilla1/goTutorial/modules/hello

go 1.23.1

replace rabor.io/greetings => ../greetings

require github.com/mrevilla1/goTutorial/modules/greetings/greetings v0.0.0-00010101000000-000000000000
