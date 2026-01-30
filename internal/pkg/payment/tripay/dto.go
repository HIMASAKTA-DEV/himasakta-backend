package tripay_payment

type TripayStatus string

const (
	StatusUNPAID  TripayStatus = "UNPAID"
	StatusPAID    TripayStatus = "PAID"
	StatusFAILED  TripayStatus = "FAILED"
	StatusEXPIRED TripayStatus = "EXPIRED"
	StatusREFUND  TripayStatus = "REFUND"
)

type (
	TransactionRequest struct {
		Method        string      `json:"method"`
		MerchantRef   string      `json:"merchant_ref"`
		Amount        int         `json:"amount"`
		CustomerName  string      `json:"customer_name"`
		CustomerEmail string      `json:"customer_email"`
		CustomerPhone string      `json:"customer_phone"`
		OrderItems    []OrderItem `json:"order_items"`
		Signature     string      `json:"signature"`
	}

	OrderItem struct {
		Name     string `json:"name"`
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
	}

	TransactionCallback struct {
		Reference         string  `json:"reference"`
		MerchantRef       string  `json:"merchant_ref"`
		PaymentMethod     string  `json:"payment_method"`
		PaymentMethodCode string  `json:"payment_method_code"`
		TotalAmount       int     `json:"total_amount"`
		FeeMerchant       int     `json:"fee_merchant"`
		FeeCustomer       int     `json:"fee_customer"`
		TotalFee          int     `json:"total_fee"`
		AmountReceived    int     `json:"amount_received"`
		IsClosedPayment   int     `json:"is_closed_payment"`
		Status            string  `json:"status"`
		PaidAt            int     `json:"paid_at"`
		Note              *string `json:"note"`
	}
)
