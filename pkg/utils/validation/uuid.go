package validation

import "github.com/google/uuid"

func UUID(u string) bool {
	_, err := uuid.Parse(u)

	return err == nil
}
