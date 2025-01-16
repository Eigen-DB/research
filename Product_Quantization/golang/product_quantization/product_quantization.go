package product_quantization

import (
	"math"
	"math/rand/v2"
)

func GenerateRandomVector(dim int) []float32 {
	vec := make([]float32, dim)
	for i := 0; i < dim; i++ {
		vec[i] = rand.Float32()*200 - 100
	}
	return vec
}

func computeL2Distance(u []float32, v []float32) float64 {
	dist := 0.0
	for i := 0; i < len(u); i++ {
		dist += math.Pow(float64(u[i]-v[i]), 2)
	}
	return math.Sqrt(dist)
}

func generateSubVectors(m int, vector []float32) [][]float32 {
	subVectors := make([][]float32, m)
	D_ := len(vector) / m
	for i := 0; i < m; i++ {
		subVectors[i] = vector[i*D_ : (i+1)*D_]
	}
	return subVectors
}

func generateCentroids(k int, m int, D_ int) [][][]float32 {
	if k%m != 0 {
		panic("k must be divisible by m")
	}

	centroids := make([][][]float32, m)
	k_ := k / m
	for i := range m {
		centroids[i] = make([][]float32, k_)
		for j := range k_ {
			centroids[i][j] = make([]float32, D_)
			centroids[i][j] = GenerateRandomVector(D_)
		}
	}
	return centroids
}

func compressVector(subVectors [][]float32, centroids [][][]float32) []int {
	compressed := make([]int, len(subVectors))
	for i := range len(subVectors) {
		minDist := math.MaxFloat64
		for j := range len(centroids[i]) {
			dist := computeL2Distance(subVectors[i], centroids[i][j])
			if dist < minDist {
				minDist = dist
				compressed[i] = j
			}
		}
	}
	return compressed
}

func DecompressVector(compressed []int, centroids [][][]float32) []float32 {
	decompressed := make([]float32, 0)
	for i := range len(compressed) {
		decompressed = append(decompressed, centroids[i][compressed[i]]...)
	}
	return decompressed
}

func CompressManyVectors(num_vectors int, dim int, vectors [][]float32) [][]int {
	m := 4
	subVectors := make([][][]float32, num_vectors)
	for i := range num_vectors {
		subVectors[i] = generateSubVectors(m, vectors[i])
		//fmt.Println(subVectors[i])
	}

	k := int(math.Pow(2, 10))
	centroids := generateCentroids(k, m, dim/m)
	//fmt.Println(centroids)

	compressedVectors := make([][]int, num_vectors)
	for i := range num_vectors {
		compressedVectors[i] = compressVector(subVectors[i], centroids)
		//fmt.Println(compressedVectors[i])
	}

	return compressedVectors
}
