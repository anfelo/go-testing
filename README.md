## Table Driven Testing

```go
tests := []TestCase{
  {name: "Testing value for: 153", value: 153, expected: true},
  {name: "Testing value for: 370", value: 370, expected: true},
  {name: "Testing value for: 371", value: 371, expected: true},
  {name: "Testing value for: 407", value: 407, expected: true},
}

// Loop over the test table
for _, test := range tests {
  t.Run(test.name, func(t *testing.T) {
    actual := calculator.CalculateIsArmstrong(test.value)
    if test.expected != actual {
      t.Fatalf("Expected %t but got %t", test.expected, actual)
    }
  })
}
```


## Coverage

```cmd
// See test coverage
$ go test ./... --cover


// Create a test profile
$ go test ./... -coverprofile=coverage.out

// View the test profile in the browser
$ go tool cover -html=coverage.out
```

## Test Fixtures

* Test fixtures are effectively files which store some form of state that we need for our tests to run.

Create a subdirectory `testdata` and add the fixture files in some helpful file type. It can be JSON, YAML, or even .go files


## Benchmarking

For the large majority of use-cases, performance is often “good-enough” within your Go applications and you don’t tend to often need to go down the route of benchmarking your code and analyzing performance to any great detail.

There are however, a number of use cases that require absolute peak performance. It is for these use cases that doing some degree of benchmarking your code is incredibly important. You need to be able to understand the performance characteristics and benchmarking is a fantastic way to consistently do that.

* use `testing.B`
* name the test prefixed with **Benchmark**
* run `go test ./... -run=Benchmark -bench=.`


```go
func benchmarkCalculateIsArmstrong(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculator.CalculateIsArmstrong(input)
	}
}

func BenchmarkCalculateIsArmstrong370(b *testing.B) {
	benchmarkCalculateIsArmstrong(370, b)
}

func BenchmarkCalculateIsArmstrong371(b *testing.B) {
	benchmarkCalculateIsArmstrong(371, b)
}

func BenchmarkCalculateIsArmstrong0(b *testing.B) {
	benchmarkCalculateIsArmstrong(0, b)
}
```

## TestMain

As your services become more intricate and complex, so too does your test setup and teardown. Thankfully, this is where the TestMain function can help save the day.

If you need a particular bit of code to execute once within a given test package then you can utilize the TestMain function to handle this execution.

```go
func TestMain(m *testing.M) {
	fmt.Println("Hello World")
	ret := m.Run()
	fmt.Println("Tests have executed")
	os.Exit(ret)
}
```

## Build Tags

We can differentiate types of tests unit, integration... using build tags. By adding "//go:build <tagname>" at the top of the test file we define the type of test. This can be useful when running tests on the CI/CD pipeline.

## Race Conditions

When running code concurrently or with go coroutines we might encouter race conditions. We can run the build with the flag `-race` so go can catch race conditions in our code.