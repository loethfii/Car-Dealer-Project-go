package helper

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"tugas_akhir_course_net/models"
)

func Error(err error) {
	if err != nil {
		panic(err)
	}
}

// konversi model car to response car
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

// konversi model car ke car request
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

// konversi model sales people ke salesPeopleResponse
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

// konversi model purchase ke purchase form response
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

// konversi model purchase form ke inner join model car, sales people
func ConvertFromPurchaseFormToPurchaseFormResponse(purchaseForm models.PurchaseForm, car models.Car, salesPeople models.SalesPeople) models.PurchaseFormInnerJoinResponse {
	var carResponse = models.CarResponseToPurchaseForm{
		Id:           car.Id,
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

	var salesPeopleResponse = models.SalesPeopleResponseToPurchaseForm{
		Id:          salesPeople.Id,
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
		CarDetail:          carResponse,
		HarusInden:         purchaseForm.HarusInden,
		LamaInden:          purchaseForm.LamaInden,
		CustomPlat:         purchaseForm.CustomPlat,
		TambahanKit:        purchaseForm.TambahanKit,
		SalesPeopleId:      purchaseForm.SalesPeopleId,
		SalesPeopleDetail:  salesPeopleResponse,
		PaymentID:          purchaseForm.Payment.Id,
		BuktiTransfer:      purchaseForm.Payment.BuktiTransfer,
		IsConfirm:          purchaseForm.Payment.IsConfirm,
	}

	return resPurchaseForm
}

// konversi payment ke payment response
func ConvToPaymentResponse(payment models.Payment) models.PaymentResponse {
	var paymentResoponse = models.PaymentResponse{
		Id:             payment.Id,
		BuktiTransfer:  payment.BuktiTransfer,
		IsConfirm:      payment.IsConfirm,
		PurchaseFormId: payment.PurchaseFormId,
	}

	return paymentResoponse
}

// inner join
func ConvToPaymentResponseInnerJoin(payment models.Payment, purchaseForm models.PurchaseForm) models.PaymentInnerJoinPurchaseForm {
	purchaseFormResponse := models.PurchaseFormResponse{
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
	var paymentResoponse = models.PaymentInnerJoinPurchaseForm{
		Id:             payment.Id,
		BuktiTransfer:  payment.BuktiTransfer,
		IsConfirm:      payment.IsConfirm,
		PurchaseFormId: payment.PurchaseFormId,
		Purchaseforms:  purchaseFormResponse,
	}

	return paymentResoponse
}

// konversi payment request ke payment response
func DataConfirmPayment(payment models.PaymentRequest, findData models.Payment) models.PaymentResponse {
	var newConfirmPaymentResponse = models.PaymentResponse{
		Id:             findData.Id,
		BuktiTransfer:  findData.BuktiTransfer,
		IsConfirm:      payment.IsConfirm,
		PurchaseFormId: findData.PurchaseFormId,
	}

	return newConfirmPaymentResponse
}

// genete nama file
func GenerateFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	filename := fmt.Sprintf("%s%s", RandomString(10), ext)
	return filename
}

// Fungsi untuk menghasilkan string acak dengan panjang tertentu
func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
