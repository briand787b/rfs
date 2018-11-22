package secrets

import (
	"fmt"
)

// PromptSecrets is a helper function that prompts the
// user for each key in the provided map
func PromptSecrets(keys ...string) (map[string]string, error) {
	m := make(map[string]string)
	var tmpV string
	for _, key := range keys {
		fmt.Printf("enter value for %s\n", key)

		if _, err := fmt.Scanln(&tmpV); err != nil {
			return nil, err
		}
		m[key] = tmpV
	}

	return m, nil
}
