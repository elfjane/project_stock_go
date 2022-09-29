module example/hello

go 1.17
// go get elfjane.com/greetings
replace elfjane.com/greetings => ../../module/greetings

// go get elfjane.com/greetings
require elfjane.com/greetings v0.0.0-00010101000000-000000000000 // indirect
