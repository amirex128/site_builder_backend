package payment_entity

import "time"

type GatewayEntity struct {
	Id                                   string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	SiteId                               string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	Saman_MerchantId                     string    `json:"saman_merchant_id,omitempty" gorm:"column:Saman_MerchantId" faker:"uuid_digit"`
	Saman_Password                       string    `json:"saman_password,omitempty" gorm:"column:Saman_Password" faker:"-"`
	IsActiveSaman                        string    `json:"is_active_saman" gorm:"column:IsActiveSaman" faker:"oneof: active, inactive"`
	Mellat_TerminalId                    int64     `json:"mellat_terminal_id,omitempty" gorm:"column:Mellat_TerminalId" faker:"boundary_start=1000000, boundary_end=9999999"`
	Mellat_UserName                      string    `json:"mellat_user_name,omitempty" gorm:"column:Mellat_UserName" faker:"username"`
	Mellat_UserPassword                  string    `json:"mellat_user_password,omitempty" gorm:"column:Mellat_UserPassword" faker:"-"`
	IsActiveMellat                       string    `json:"is_active_mellat" gorm:"column:IsActiveMellat" faker:"oneof: active, inactive"`
	Parsian_LoginAccount                 string    `json:"parsian_login_account,omitempty" gorm:"column:Parsian_LoginAccount" faker:"username"`
	IsActiveParsian                      string    `json:"is_active_parsian" gorm:"column:IsActiveParsian" faker:"oneof: active, inactive"`
	Pasargad_MerchantCode                string    `json:"pasargad_merchant_code,omitempty" gorm:"column:Pasargad_MerchantCode" faker:"uuid_digit"`
	Pasargad_TerminalCode                string    `json:"pasargad_terminal_code,omitempty" gorm:"column:Pasargad_TerminalCode" faker:"uuid_digit"`
	Pasargad_PrivateKey                  string    `json:"pasargad_private_key,omitempty" gorm:"column:Pasargad_PrivateKey" faker:"-"`
	IsActivePasargad                     string    `json:"is_active_pasargad" gorm:"column:IsActivePasargad" faker:"oneof: active, inactive"`
	IranKish_TerminalId                  string    `json:"iran_kish_terminal_id,omitempty" gorm:"column:IranKish_TerminalId" faker:"uuid_digit"`
	IranKish_AcceptorId                  string    `json:"iran_kish_acceptor_id,omitempty" gorm:"column:IranKish_AcceptorId" faker:"uuid_digit"`
	IranKish_PassPhrase                  string    `json:"iran_kish_pass_phrase,omitempty" gorm:"column:IranKish_PassPhrase" faker:"-"`
	IranKish_PublicKey                   string    `json:"iran_kish_public_key,omitempty" gorm:"column:IranKish_PublicKey" faker:"-"`
	IsActiveIranKish                     string    `json:"is_active_iran_kish" gorm:"column:IsActiveIranKish" faker:"oneof: active, inactive"`
	Melli_TerminalId                     string    `json:"melli_terminal_id,omitempty" gorm:"column:Melli_TerminalId" faker:"uuid_digit"`
	Melli_MerchantId                     string    `json:"melli_merchant_id,omitempty" gorm:"column:Melli_MerchantId" faker:"uuid_digit"`
	Melli_TerminalKey                    string    `json:"melli_terminal_key,omitempty" gorm:"column:Melli_TerminalKey" faker:"-"`
	IsActiveMelli                        string    `json:"is_active_melli" gorm:"column:IsActiveMelli" faker:"oneof: active, inactive"`
	AsanPardakht_MerchantConfigurationId string    `json:"asan_pardakht_merchant_configuration_id,omitempty" gorm:"column:AsanPardakht_MerchantConfigurationId" faker:"uuid_digit"`
	AsanPardakht_UserName                string    `json:"asan_pardakht_user_name,omitempty" gorm:"column:AsanPardakht_UserName" faker:"username"`
	AsanPardakht_Password                string    `json:"asan_pardakht_password,omitempty" gorm:"column:AsanPardakht_Password" faker:"-"`
	AsanPardakht_Key                     string    `json:"asan_pardakht_key,omitempty" gorm:"column:AsanPardakht_Key" faker:"-"`
	AsanPardakht_IV                      string    `json:"asan_pardakht_iv,omitempty" gorm:"column:AsanPardakht_IV" faker:"-"`
	IsActiveAsanPardakht                 string    `json:"is_active_asan_pardakht" gorm:"column:IsActiveAsanPardakht" faker:"oneof: active, inactive"`
	Sepehr_TerminalId                    int64     `json:"sepehr_terminal_id,omitempty" gorm:"column:Sepehr_TerminalId" faker:"boundary_start=1000000, boundary_end=9999999"`
	IsActiveSepehr                       string    `json:"is_active_sepehr" gorm:"column:IsActiveSepehr" faker:"oneof: active, inactive"`
	ZarinPal_MerchantId                  string    `json:"zarin_pal_merchant_id,omitempty" gorm:"column:ZarinPal_MerchantId" faker:"uuid_digit"`
	ZarinPal_AuthorizationToken          string    `json:"zarin_pal_authorization_token,omitempty" gorm:"column:ZarinPal_AuthorizationToken" faker:"-"`
	ZarinPal_IsSandbox                   bool      `json:"zarin_pal_is_sandbox,omitempty" gorm:"column:ZarinPal_IsSandbox" faker:"oneof: true, false"`
	IsActiveZarinPal                     string    `json:"is_active_zarin_pal" gorm:"column:IsActiveZarinPal" faker:"oneof: active, inactive"`
	PayIr_Api                            string    `json:"pay_ir_api,omitempty" gorm:"column:PayIr_Api" faker:"-"`
	PayIr_IsTestAccount                  bool      `json:"pay_ir_is_test_account,omitempty" gorm:"column:PayIr_IsTestAccount" faker:"oneof: true, false"`
	IsActivePayIr                        string    `json:"is_active_pay_ir" gorm:"column:IsActivePayIr" faker:"oneof: active, inactive"`
	IdPay_Api                            string    `json:"id_pay_api,omitempty" gorm:"column:IdPay_Api" faker:"-"`
	IdPay_IsTestAccount                  bool      `json:"id_pay_is_test_account,omitempty" gorm:"column:IdPay_IsTestAccount" faker:"oneof: true, false"`
	IsActiveIdPay                        string    `json:"is_active_id_pay" gorm:"column:IsActiveIdPay" faker:"oneof: active, inactive"`
	YekPay_MerchantId                    string    `json:"yek_pay_merchant_id,omitempty" gorm:"column:YekPay_MerchantId" faker:"uuid_digit"`
	IsActiveYekPay                       string    `json:"is_active_yek_pay" gorm:"column:IsActiveYekPay" faker:"oneof: active, inactive"`
	PayPing_AccessToken                  string    `json:"pay_ping_access_token,omitempty" gorm:"column:PayPing_AccessToken" faker:"-"`
	IsActivePayPing                      string    `json:"is_active_pay_ping" gorm:"column:IsActivePayPing" faker:"oneof: active, inactive"`
	IsActiveParbadVirtual                string    `json:"is_active_parbad_virtual" gorm:"column:IsActiveParbadVirtual" faker:"oneof: active, inactive"`
	UserId                               string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CreatedAt                            time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt                            time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version                              time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted                            bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt                            time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
}

func (GatewayEntity) TableName() string {
	return "Payment.Gateways"
}
