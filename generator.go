package generator

func GenerateSecret(shadowCount int) Secret {
	shadows := []Shadow{}

	for i := 0; i < shadowCount; i++ {
		shadows = append(shadows, Shadow{float64(i), float64(i)})
	}
	secret := Secret{"Steve Rules", shadows}

	return secret
}

type Secret struct {
	Value   string
	Shadows []Shadow
}

type Shadow struct {
	X float64
	Y float64
}
