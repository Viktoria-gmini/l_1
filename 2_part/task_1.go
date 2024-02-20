package main

import (
	"strings"
	"testing"
)

//  самый эффективный способ конкатенации строк?

func BenchmarkStringBuilder(b *testing.B) {
	str1 := "Hello, "
	str2 := "World!"

	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.WriteString(str1)
		sb.WriteString(str2)
		_ = sb.String()
	}
}

// strings.Builder - это более эффективный способ для конкатенации строк,
// особенно при работе с большим количеством строк.
// Этот метод минимизирует количество аллокаций памяти,
// что делает его предпочтительным выбором для высокопроизводительных операций.
