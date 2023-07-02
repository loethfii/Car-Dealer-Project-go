package helper

import (
	"tugas_akhir_course_net/models"
)

func Error(err error) {
	if err != nil {
		panic(err)
	}
}

func ConvertToResponseCar(car models.Car) models.CarResponse {
	var purchaseFormResponses []models.PurchaseFormResponse

	for _, value := range car.PurchaseForms {
		data := models.PurchaseFormResponse{
			Id:                 value.Id,
			NamaLengkapPembeli: value.NamaLengkapPembeli,
			NomerKTP:           value.NomerKTP,
			AlamatRumah:        value.AlamatRumah,
			NomerDebit:         value.NomerDebit,
			CarId:              value.CarId,
			HarusInden:         value.HarusInden,
			LamaInden:          value.LamaInden,
			CustomPlat:         value.CustomPlat,
			TambahanKit:        value.TambahanKit,
			SalesPeopleId:      value.SalesPeopleId,
		}

		purchaseFormResponses = append(purchaseFormResponses, data)
	}

	var resCar = models.CarResponse{
		Id:            car.Id,
		NamaMobil:     car.NamaMobil,
		TipeMobil:     car.TipeMobil,
		JenisMobil:    car.JenisMobil,
		BahanBakar:    car.BahanBakar,
		Isi_Silinder:  car.Isi_Silinder,
		Warna:         car.Warna,
		Transmisi:     car.Transmisi,
		Harga:         car.Harga,
		Qty:           car.Qty,
		PurchaseForms: purchaseFormResponses,
	}

	return resCar
}

func ConvertToRequestCar(car models.Car) models.CarRequest {
	var resCar = models.CarRequest{
		NamaMobil:    car.NamaMobil,
		TipeMobil:    car.TipeMobil,
		JenisMobil:   car.JenisMobil,
		BahanBakar:   car.BahanBakar,
		Isi_Silinder: car.Isi_Silinder,
		Warna:        car.Warna,
		Transmisi:    car.Transmisi,
		Harga:        car.Harga,
		Qty:          car.Qty,
	}

	return resCar
}

func ConvertToResponseSalesPeople(salesPeople models.SalesPeople) models.SalesPeopleResponse {
	var purchaseFormResponses []models.PurchaseFormResponse

	for _, value := range salesPeople.PurchaseForms {
		data := models.PurchaseFormResponse{
			Id:                 value.Id,
			NamaLengkapPembeli: value.NamaLengkapPembeli,
			NomerKTP:           value.NomerKTP,
			AlamatRumah:        value.AlamatRumah,
			NomerDebit:         value.NomerDebit,
			CarId:              value.CarId,
			HarusInden:         value.HarusInden,
			LamaInden:          value.LamaInden,
			CustomPlat:         value.CustomPlat,
			TambahanKit:        value.TambahanKit,
			SalesPeopleId:      value.SalesPeopleId,
		}

		purchaseFormResponses = append(purchaseFormResponses, data)
	}

	var resSalesPeople = models.SalesPeopleResponse{
		Id:            salesPeople.Id,
		NamaSales:     salesPeople.NamaSales,
		Nip:           salesPeople.Nip,
		NomerTelpon:   salesPeople.NomerTelpon,
		Bagian:        salesPeople.Bagian,
		PurchaseForms: purchaseFormResponses,
	}

	return resSalesPeople
}

func ConvertToReponsePurchaseForm(purchaseForm models.PurchaseForm) models.PurchaseFormResponse {
	var resPurchaseForm = models.PurchaseFormResponse{
		Id:                 purchaseForm.Id,
		NamaLengkapPembeli: purchaseForm.NamaLengkapPembeli,
		NomerKTP:           purchaseForm.NomerKTP,
		AlamatRumah:        purchaseForm.AlamatRumah,
		NomerDebit:         purchaseForm.NomerDebit,
		CarId:              purchaseForm.CarId,
		HarusInden:         purchaseForm.HarusInden,
		LamaInden:          purchaseForm.LamaInden,
		CustomPlat:         purchaseForm.CustomPlat,
		TambahanKit:        purchaseForm.TambahanKit,
		SalesPeopleId:      purchaseForm.SalesPeopleId,
	}

	return resPurchaseForm
}

func ConvertToResponseAndInnerJoin(purchaseForm models.PurchaseForm, car models.Car, salesPeople models.SalesPeople) models.PurchaseFormInnerJoinResponse {
	var carRequest = models.CarRequest{
		NamaMobil:    car.NamaMobil,
		TipeMobil:    car.TipeMobil,
		JenisMobil:   car.JenisMobil,
		BahanBakar:   car.BahanBakar,
		Isi_Silinder: car.Isi_Silinder,
		Warna:        car.Warna,
		Transmisi:    car.Transmisi,
		Harga:        car.Harga,
		Qty:          car.Qty,
	}

	var salesPeopleRequest = models.SalesPeopleRequest{
		NamaSales:   salesPeople.NamaSales,
		Nip:         salesPeople.Nip,
		NomerTelpon: salesPeople.NomerTelpon,
		Bagian:      salesPeople.Bagian,
	}

	var resPurchaseForm = models.PurchaseFormInnerJoinResponse{
		Id:                 purchaseForm.Id,
		NamaLengkapPembeli: purchaseForm.NamaLengkapPembeli,
		NomerKTP:           purchaseForm.NomerKTP,
		AlamatRumah:        purchaseForm.AlamatRumah,
		NomerDebit:         purchaseForm.NomerDebit,
		CarId:              purchaseForm.CarId,
		CarDetail:          carRequest,
		HarusInden:         purchaseForm.HarusInden,
		LamaInden:          purchaseForm.LamaInden,
		CustomPlat:         purchaseForm.CustomPlat,
		TambahanKit:        purchaseForm.TambahanKit,
		SalesPeopleId:      purchaseForm.SalesPeopleId,
		SalesPeopleDetail:  salesPeopleRequest,
	}

	return resPurchaseForm
}

func ConvToPaymentResponse(payment models.Payment) models.PaymentResponse {
	var paymentResoponse = models.PaymentResponse{
		Id:             payment.Id,
		BuktiTransfer:  payment.BuktiTransfer,
		IsConfirm:      payment.IsConfirm,
		PurchaseFormId: payment.PurchaseFormId,
	}

	return paymentResoponse
}

func DataConfirmPayment(payment models.PaymentRequest, findData models.Payment) models.PaymentResponse {
	var newConfirmPaymentResponse = models.PaymentResponse{
		Id:             findData.Id,
		BuktiTransfer:  findData.BuktiTransfer,
		IsConfirm:      payment.IsConfirm,
		PurchaseFormId: findData.PurchaseFormId,
	}

	return newConfirmPaymentResponse
}
