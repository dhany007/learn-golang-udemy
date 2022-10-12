package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before unit test, misal koneksi ke db")
	m.Run()
	fmt.Println("After Unit tes")
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Dhany)",
			request: "Dhany",
		},
		{
			name:    "HelloWorld(Budi)",
			request: "Budi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dhany")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Dhany", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dhany")
		}
	})

	b.Run("Budi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Budi")
		}
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name    string
		request string
		expect  string
	}{
		{
			name:    "HelloWorld(Dhany)",
			request: "Dhany",
			expect:  "Hello Dhany",
		},
		{
			name:    "HelloWorld(Budi)",
			request: "Budi",
			expect:  "Hello Budi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expect, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Dhany", func(t *testing.T) {
		result := HelloWorld("Dhany")
		require.Equal(t, "Hello Dhany", result, "Result must be: 'Hello Dhany'")
	})
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		require.Equal(t, "Hello Budi", result, "Result must be: 'Hello Budi'")
	})
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Dhany")

	if result != "Hello Dhany" {
		// t.Fail() // menggagalkan  tes tanpa menutup tes
		t.Error("Result must be 'Hello Dhany'")
	}

	fmt.Println("HelloWorld Dhany")
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")

	if result != "Hello Budi" {
		// t.FailNow() // menggagalkan tes dan menutup tes
		t.Fatal("Result must be 'Hello Budi'")
	}

	fmt.Println("HelloWorld Budi")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Dhany")

	assert.Equal(t, "Hello Dhany", result, "Result must be: 'Hello Dhany'")
	fmt.Println("Test with assert is done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Dhany")
	require.Equal(t, "Hello Dhany", result, "Result must be: 'Hello Dhany'")
	fmt.Println("Test with require is done")
}

func TesSkip(t *testing.T) {
	if runtime.GOOS == "ubuntu" {
		t.Skip("Can not run in ubuntu")
	}

	result := HelloWorld("Dhany")
	require.Equal(t, "Hello Dhany", result, "Result must be: 'Hello Dhany'")
}
