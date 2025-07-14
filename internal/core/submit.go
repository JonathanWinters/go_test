package core

import "context"

func HandleSubmit(ctx context.Context, submitRequest SubmitRequest) (submitResponse SubmitResponse, err error) {

	return
}

// 1. Maps must be retangular
// 2. Maps may not be large than 100 in any dimenion
// 3. Map spaces may not use values other the number 0, 1, 2, 3, or 4.
func ValidateSubmission() {

}
