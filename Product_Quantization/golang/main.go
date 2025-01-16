package main

import (
	"fmt"
	"product_quantization/product_quantization"
)

func main() {
	num_vectors := 5
	vectors := make([][]float32, num_vectors)
	dim := 12
	for i := range num_vectors {
		vectors[i] = product_quantization.GenerateRandomVector(dim)
		fmt.Println(vectors[i])
	}

	compressed_vectors := product_quantization.CompressManyVectors(num_vectors, dim, vectors)
	for _, v := range compressed_vectors {
		fmt.Println(v)
	}
}
