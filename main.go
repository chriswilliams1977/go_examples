package main

import (
	"fmt"
	o "github.com/chriswilliams1977/headfirst/playground/packages/other"
	"time"
)

func main() {
	//vars.OutputTest()
	//types.OutputTest()
	//loops.GuessingGame()
	//functions.FormattedValues()

	/*
		totalPaintNeeded, err := functions.PaintCalculator(4.2, 3.0, 10)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Total paint needed is %.2f liters\n", totalPaintNeeded)
		}
		//totalPaintNeeded += functions.PaintCalculator(4.2, 3.0, 10)

		amount := 3
		pointers.PointerSample(&amount)
		pointers.GetPointerTypes()
		pointers.GetPointerAddress(amount)
		pointers.GetPointerValue(amount)

		var myStringPointer = pointers.ReturnPointer()
		fmt.Println(*myStringPointer)

		var myString string = "Test"
		pointers.PrintPointer(&myString)

		greeting.Hi()
		english.SayHello()
		dutch.SayHello()
		arrays.GetArrayElements()

		samplearray.GetAveragePurchase()
		sampleslice.GetAveragePurchase()
		sampleslice.GetArgs()
		sampleslice.VariadicFunctionInRange(1,100,-12.5, 3.2, 0, 50, 103.5)
		sampleslice.GetAveragePurchaseArgs()

		samplemaps.GetDataUsingSlice()
		maps.CreateMap()
		maps.Status("Alma")
		maps.Status("Rohit")
		maps.Status("Carl")
		maps.CheckZeroValue()
		maps.DeleteMapKey()
		samplemaps.GetDataUsingMap()
	*/
	//structs.GetStruct()
	//Pass struct and amend name locally without using a pointer
	//subnopointer := s.Subscriber{"Chris",2.20,false}
	//s.NewWithOutPointer(subnopointer)
	//fmt.Println(subnopointer)
	//Pass struct and amend name locally without using a pointer
	//subwithpointer := s.Subscriber{"Bob",3.0,true}
	//pass address of pointer using &
	//&  = address of struct
	//s.NewWithPointer(&subwithpointer)
	//fmt.Println(subwithpointer)
	/*
		subscriber1 := s.DefaultSubscriber("Aman Singh")
		s.ApplyDiscount(subscriber1)
		s.PrintInfo(subscriber1)

		subscriber2 := s.DefaultSubscriber("Beth Ryan")
		s.PrintInfo(subscriber2)
		structs.GetStructLiteral()

		subscriber := s.Subscriber{Name:"Bob"}
		//address is an anonymous type
		//address is promoted to outer struct
		//so you can reference address fields from outer struct
		subscriber.Street = "Rydons Lane"
		subscriber.City = "Croydon"
		subscriber.State = "Surrey"
		subscriber.PostalCode = "CR5 1SU"
		fmt.Println(subscriber)

		d.GetFuel()
		g := d.Gallons(1.2)
		g.GallonsReceiverWithParams(10, true)
		g.GallonsReceiverMultiReturns()
		d.GetMyTypeValues()

		//Using Date type
		date := encap.Date{}
		//Call setters
		//setter return an error if invalid year sent
		err := date.SetYear(1979)
		//check for error
		if err != nil{
			log.Fatal(err)
		}
		date.SetMonth(07)
		date.SetDay(29)
		//Call getters
		fmt.Println(date.Year())

		//Now use Event type with Date as anonymous field
		//Remember unexported fields and methods are not promoted
		//create instance of event
		event := encap.Event{}
		event.SetTitle("Mum's Birthday")
		//event returns and error if incorrect year passed
		//access Date promoted method SetYear()
		err = event.SetYear(1946)
		if err != nil{
			log.Fatal(err)
		}
		err = event.SetMonth(02)
		if err != nil{
			log.Fatal(err)
		}
		err = event.SetDay(11)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(event.Title())
		fmt.Println(event.Day())
		fmt.Println(event.Month())
		fmt.Println(event.Year())

		 tapePlayer := samples.TapePlayer{}
		 tapeRecorder := samples.TapeRecorder{}
		 songs := []string{"One", "Fake Plastic Trees"}

		 samples.PlayList(tapePlayer, songs)
		 //calling method with type assertion
		 samples.TryOut(tapeRecorder)

		 //declare var with MyInterface as its type
		 var myInterface  i.MyInterface
		 //assign myInterface to equal MyType as MyType satifies the interface
		 myInterface = i.MyType(5)
		 //call myType methods
		 myInterface.MethodWithoutParams()
		 myInterface.MethodWithParams(127.3)
		 fmt.Println(myInterface.MethodWithReturnValue())

		 fmt.Println(i.CallInterfaceMethod(i.MyType(3)))

		 //Examples of out of the box interfaces you can implement

		 //Error interface
		 //OverheatError has func CheckTemperature that returns am error
		 /*
		 var err error = i.CheckTemperature(121.3, 100)
		 if err != nil {
		 	log.Fatal(err)
		 }


		 //Stringer interface
		fmt.Println(i.Gallons(12.09248342))
		fmt.Println(i.Liters(12.09248342))
		fmt.Println(i.Milliliters(12.09248342))

		i.AcceptAnything(3.152)
		i.AcceptAnything("Hello")
		i.AcceptAnything(true)
		i.AcceptAnything(tapePlayer)


		//pass filename from cli args
		//call go run main.go floats.txt
		numbers, err := f.GetFloats(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		//init var to hold sum
		var sum float64 = 0
		//for each value in numbers slice
		for _, number := range numbers{
			//add to the sum
			sum += number
		}
		fmt.Printf("Sum: %0.2f\n", sum)

		e := f.DeferExample()
		if e != nil {
			log.Fatal(e)
		}

		//use this if you want to get current dir
		//currentDir := f.GetCurrentDirName()
		//defer ensures recover is handled if panic occurred in ScanDir
		defer f.ReportPanic()
		panic("some other issue")
		f.ScanDir("./packages/failures")

		//f.RecursionExample()


	//using goroutines
	//create a channel to get data from goroutines - remember cant use return in func called by goroutine
	//channel holds a Page struct so we can keep url and size together
	pages := make(chan r.Page)
	urls := []string{"https://example.com","https://golang.org","https://golang.org/doc"}
	//for each url in slice
	for _, url := range urls{
		//call GetPageWeight func with url from slice and a channel holding a Page struct
		go r.GetPageWeight(url, pages)
	}
	//go r.GetPageWeight("https://example.com", sizes)
	//go r.GetPageWeight("https://golang.org", sizes)
	//go r.GetPageWeight("https://golang.org/doc", sizes)
	//gives time for goroutines to finish before main goroutine
	//time.Sleep(5 * time.Second) - can use channels to replace this

	//for each url in the slice
	for i := 0; i < len(urls); i++{
		//get the page struct from the channel
		page := <-pages
		//print values of the struct
		fmt.Printf("%s: %d\n",page.URL, page.Size)
	}
	//fmt.Println(<-sizes)
	//fmt.Println(<-sizes)
	//fmt.Println(<-sizes)

	//channel example
	myChannel := make(chan string)
	go r.ChannelExample(myChannel)
	fmt.Println(<-myChannel)

	//sample of blocking
	//it goroutine blocks when it sends value to channel until each received by main goroutine
	//create two channels
	channel1 := make(chan string)
	channel2 := make(chan string)
	//call funcs as goroutines
	//sending operation blocks goroutine until another goroutine executes recieve operation on same channel
	//recieving goroutine blocks goroutine until another goroutine executes send operation on same channel
	//below each goroutine is blocked after sending value until main goroutine receives from it
	//send goroutines
	go r.ChannelABC(channel1)
	go r.ChannelDEF(channel2)
	//blocked goroutines
	//recieve value
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()

	//example of goroutine slowed down
	myChannel1 := make(chan string)
	//send goroutine
	go r.SendGoRoutine(myChannel1)
	//sleep main goroutine sleeps for 5 seconds
	r.ReportNap("recieving goroutine", 5)
	//two recieve operations on the channel
	//these are only recieved once the main goroutine wakes up after the extra 3 seconds it sleeps longer than send goroutine
	fmt.Println(<-myChannel1)
	fmt.Println(<-myChannel1)

	words1 := []string{"apple","orange","pear","banana"}
	words2 := []string{"apple","orange"}
	fmt.Println(t.JoinWithCommas(words1))
	fmt.Println(t.JoinWithCommas(words2))
	 */

	//w.FuncVarExample()

	//if we receive request for url enmding with /hello
	//call ViewHandler
	//http.HandleFunc("/hello",w.EnglishHandler)
	//http.HandleFunc("/salut",w.FrenchHandler)
	//http.HandleFunc("/namaste",w.HindiHandler)
	//http.ListenAndServe starts web server
	//remember net.http package has a small web server to handle requests
	//listen for browser requests and respond to them
	//localhost means web server is running on your machine
	//default port for http requests is 80
	//nil value means requests will be handled using funcs set up in the HandleFunc
	//ListenAndServe runs forever unless and error occurs
	/*
	err := http.ListenAndServe("localhost:8080",nil)
	if err != nil {
		log.Fatal(err)
	}

	//w.TextTemplateExample("test string")
	w.TextTemplateWithData("Template start\nAction: {{.}}\nTemplate end \n","ABC")

	//template example using if
	w.TextTemplateWithData("Start {{if .}}Dot is true!{{end}} finish\n",true)
	w.TextTemplateWithData("Start {{if .}}Dot is true!{{end}} finish\n",false)

	//template example using range
	w.TextTemplateWithData("Before the loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n",[]string{"do","re","mi"})

	//template using structs
	w.TextTemplateWithStruct()

	//call viewhandler when request path is  /guestbook
	http.HandleFunc("/guestbook",w.ViewHandler)
	//call newhandler when request path is  /new
	http.HandleFunc("/guestbook/new",w.NewHandler)
	//call newhandler when request path is  /new
	http.HandleFunc("/guestbook/create",w.CreateHandler)

	err := http.ListenAndServe("localhost:8080",nil)
	//this will never be nil so we dont check it
	//if it fails it will always return an error
	log.Fatal(err)
	*/

	//r.ReadFile()
	//r.WriteFile()
	//r.BitwiseExample()

	//print time program started
	fmt.Println(time.Now())
	//make unbuffered channel
	channel1 := make(chan string)
	//launch sendLetters in new goroutine
	go o.SendLetters(channel1)
	//make main goroutine sleep 5 seconds
	time.Sleep(5 * time.Second)
	//receive a print values with current time
	//first value already waiting to be recieved
	//other goroutines blocked until first letter received
	fmt.Println(<-channel1, time.Now())
	fmt.Println(<-channel1, time.Now())
	fmt.Println(<-channel1, time.Now())
	fmt.Println(<-channel1, time.Now())
	fmt.Println(time.Now())

	//buffered channel speeds this up
	//first value sent to channel and doesnt block until main goroutine recieves it
	//send 3 values to the buffer
	//it only blocks the goroutine on forth letter sent to buffer
	fmt.Println(time.Now())
	//make buffered channel with 3 spaces
	channel2 := make(chan string, 3)
	//launch sendLetters in new goroutine
	go o.SendLetters(channel2)
	//make main goroutine sleep 5 seconds
	time.Sleep(5 * time.Second)
	//receive a print values with current time
	//first value already waiting to be recieved
	//other goroutines blocked until first letter received
	fmt.Println(<-channel2, time.Now())
	fmt.Println(<-channel2, time.Now())
	fmt.Println(<-channel2, time.Now())
	fmt.Println(<-channel2, time.Now())
	fmt.Println(time.Now())
}
