package main

import (
	"testing"
)

func benchmarkF1(number uint64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f1(number)
	}
}

func benchmarkF2(number uint64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f2(number)
	}
}

func benchmarkF3(number uint64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f3(number)
	}
}
func benchmarkF4(number uint64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f4(number)
	}
}

func BenchmarkF1_F(b *testing.B)        { benchmarkF1(0xF, b) }
func BenchmarkF1_FF(b *testing.B)       { benchmarkF1(0xFF, b) }
func BenchmarkF1_FFF(b *testing.B)      { benchmarkF1(0xFFF, b) }
func BenchmarkF1_FFFF(b *testing.B)     { benchmarkF1(0xFFFF, b) }
func BenchmarkF1_FFFFF(b *testing.B)    { benchmarkF1(0xFFFFF, b) }
func BenchmarkF1_FFFFFF(b *testing.B)   { benchmarkF1(0xFFFFFF, b) }
func BenchmarkF1_FFFFFFF(b *testing.B)  { benchmarkF1(0xFFFFFFF, b) }
func BenchmarkF1_FFFFFFFF(b *testing.B) { benchmarkF1(0xFFFFFFFF, b) }

func BenchmarkF2_F(b *testing.B)        { benchmarkF2(0xF, b) }
func BenchmarkF2_FF(b *testing.B)       { benchmarkF2(0xFF, b) }
func BenchmarkF2_FFF(b *testing.B)      { benchmarkF2(0xFFF, b) }
func BenchmarkF2_FFFF(b *testing.B)     { benchmarkF2(0xFFFF, b) }
func BenchmarkF2_FFFFF(b *testing.B)    { benchmarkF2(0xFFFFF, b) }
func BenchmarkF2_FFFFFF(b *testing.B)   { benchmarkF2(0xFFFFFF, b) }
func BenchmarkF2_FFFFFFF(b *testing.B)  { benchmarkF2(0xFFFFFFF, b) }
func BenchmarkF2_FFFFFFFF(b *testing.B) { benchmarkF2(0xFFFFFFFF, b) }

func BenchmarkF3_F(b *testing.B)        { benchmarkF3(0xF, b) }
func BenchmarkF3_FF(b *testing.B)       { benchmarkF3(0xFF, b) }
func BenchmarkF3_FFF(b *testing.B)      { benchmarkF3(0xFFF, b) }
func BenchmarkF3_FFFF(b *testing.B)     { benchmarkF3(0xFFFF, b) }
func BenchmarkF3_FFFFF(b *testing.B)    { benchmarkF3(0xFFFFF, b) }
func BenchmarkF3_FFFFFF(b *testing.B)   { benchmarkF3(0xFFFFFF, b) }
func BenchmarkF3_FFFFFFF(b *testing.B)  { benchmarkF3(0xFFFFFFF, b) }
func BenchmarkF3_FFFFFFFF(b *testing.B) { benchmarkF3(0xFFFFFFFF, b) }

func BenchmarkF4_F(b *testing.B)        { benchmarkF4(0xF, b) }
func BenchmarkF4_FF(b *testing.B)       { benchmarkF4(0xFF, b) }
func BenchmarkF4_FFF(b *testing.B)      { benchmarkF4(0xFFF, b) }
func BenchmarkF4_FFFF(b *testing.B)     { benchmarkF4(0xFFFF, b) }
func BenchmarkF4_FFFFF(b *testing.B)    { benchmarkF4(0xFFFFF, b) }
func BenchmarkF4_FFFFFF(b *testing.B)   { benchmarkF4(0xFFFFFF, b) }
func BenchmarkF4_FFFFFFF(b *testing.B)  { benchmarkF4(0xFFFFFFF, b) }
func BenchmarkF4_FFFFFFFF(b *testing.B) { benchmarkF4(0xFFFFFFFF, b) }

func BenchmarkF1_0F(b *testing.B)       { benchmarkF1(0xF, b) }
func BenchmarkF1_F0(b *testing.B)       { benchmarkF1(0xF0, b) }
func BenchmarkF1_F00(b *testing.B)      { benchmarkF1(0xF00, b) }
func BenchmarkF1_F000(b *testing.B)     { benchmarkF1(0xF000, b) }
func BenchmarkF1_F0000(b *testing.B)    { benchmarkF1(0xF0000, b) }
func BenchmarkF1_F00000(b *testing.B)   { benchmarkF1(0xF00000, b) }
func BenchmarkF1_F000000(b *testing.B)  { benchmarkF1(0xF000000, b) }
func BenchmarkF1_F0000000(b *testing.B) { benchmarkF1(0xF0000000, b) }

func BenchmarkF2_0F(b *testing.B)       { benchmarkF2(0xF, b) }
func BenchmarkF2_F0(b *testing.B)       { benchmarkF2(0xF0, b) }
func BenchmarkF2_F00(b *testing.B)      { benchmarkF2(0xF00, b) }
func BenchmarkF2_F000(b *testing.B)     { benchmarkF2(0xF000, b) }
func BenchmarkF2_F0000(b *testing.B)    { benchmarkF2(0xF0000, b) }
func BenchmarkF2_F00000(b *testing.B)   { benchmarkF2(0xF00000, b) }
func BenchmarkF2_F000000(b *testing.B)  { benchmarkF2(0xF000000, b) }
func BenchmarkF2_F0000000(b *testing.B) { benchmarkF2(0xF0000000, b) }

func BenchmarkF3_0F(b *testing.B)       { benchmarkF3(0xF, b) }
func BenchmarkF3_F0(b *testing.B)       { benchmarkF3(0xF0, b) }
func BenchmarkF3_F00(b *testing.B)      { benchmarkF3(0xF00, b) }
func BenchmarkF3_F000(b *testing.B)     { benchmarkF3(0xF000, b) }
func BenchmarkF3_F0000(b *testing.B)    { benchmarkF3(0xF0000, b) }
func BenchmarkF3_F00000(b *testing.B)   { benchmarkF3(0xF00000, b) }
func BenchmarkF3_F000000(b *testing.B)  { benchmarkF3(0xF000000, b) }
func BenchmarkF3_F0000000(b *testing.B) { benchmarkF3(0xF0000000, b) }

func BenchmarkF4_0F(b *testing.B)       { benchmarkF4(0xF, b) }
func BenchmarkF4_F0(b *testing.B)       { benchmarkF4(0xF0, b) }
func BenchmarkF4_F00(b *testing.B)      { benchmarkF4(0xF00, b) }
func BenchmarkF4_F000(b *testing.B)     { benchmarkF4(0xF000, b) }
func BenchmarkF4_F0000(b *testing.B)    { benchmarkF4(0xF0000, b) }
func BenchmarkF4_F00000(b *testing.B)   { benchmarkF4(0xF00000, b) }
func BenchmarkF4_F000000(b *testing.B)  { benchmarkF4(0xF000000, b) }
func BenchmarkF4_F0000000(b *testing.B) { benchmarkF4(0xF0000000, b) }
