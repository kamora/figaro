package figaro

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kamora/morphium"
	"os"
)

func Init() error {
	return godotenv.Load()
}

func Ensure[T morphium.Morphable](key string) (T, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return morphium.Default[T](), fmt.Errorf("undefined value for %s", key)
	}

	return morphium.Morph[T](value)
}

func Obtain[T morphium.Morphable](key string, def T) (T, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}

	converted, err := morphium.Morph[T](value)
	if err != nil {
		return def, fmt.Errorf("unmorphable value for %s: %w", key, err)
	}

	return converted, nil
}

func Collect[T morphium.Morphable](keys ...string) ([]T, error) {
	result := make([]T, len(keys))

	for i, key := range keys {
		value, err := Ensure[T](key)

		if err != nil {
			return nil, err
		}

		result[i] = value
	}

	return result, nil
}
