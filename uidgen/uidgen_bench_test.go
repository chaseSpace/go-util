package uidgen

import "testing"

func BenchmarkImplUIDGenAPIWithUUID_GetOne(b *testing.B) {
	gen := &implUIDGenAPIWithUUID{}

	benchLengths := []int{1, 6, 10, 18}
	for _, length := range benchLengths {
		b.Run("length_"+toString(int64(length)), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if _, err := gen.GetOne(length); err != nil {
					b.Fatalf("GetOne() failed: %v", err)
				}
			}
		})
	}
}
