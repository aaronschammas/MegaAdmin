package services

import (
	"MegaModa/modelos"

	"gorm.io/gorm"
)

// VoucherService proporciona métodos para interactuar con el modelo Voucher.
type VoucherService struct {
	DB *gorm.DB
}

// NewVoucherService crea una nueva instancia de VoucherService.
func NewVoucherService(db *gorm.DB) *VoucherService {
	return &VoucherService{DB: db}
}

// CrearVoucher agrega un nuevo voucher a la base de datos.
func (s *VoucherService) CrearVoucher(voucher *modelos.Voucher) error {
	return s.DB.Create(voucher).Error
}

// ObtenerVoucher obtiene un voucher por su ID.
func (s *VoucherService) ObtenerVoucher(id int) (*modelos.Voucher, error) {
	var voucher modelos.Voucher
	err := s.DB.First(&voucher, id).Error
	return &voucher, err
}

// ActualizarVoucher actualiza un voucher existente.
func (s *VoucherService) ActualizarVoucher(voucher *modelos.Voucher) error {
	return s.DB.Save(voucher).Error
}

// EliminarVoucher elimina un voucher por su ID.
func (s *VoucherService) EliminarVoucher(id int) error {
	return s.DB.Delete(&modelos.Voucher{}, id).Error
}
