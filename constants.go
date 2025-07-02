package iyzipay

// Locale constants
const (
	LocaleTR = "tr"
	LocaleEN = "en"
)

// Currency constants
const (
	CurrencyTRY = "TRY"
	CurrencyEUR = "EUR"
	CurrencyUSD = "USD"
	CurrencyIRR = "IRR"
	CurrencyGBP = "GBP"
	CurrencyNOK = "NOK"
	CurrencyRUB = "RUB"
	CurrencyCHF = "CHF"
)

// Payment Group constants
const (
	PaymentGroupProduct      = "PRODUCT"
	PaymentGroupListing      = "LISTING"
	PaymentGroupSubscription = "SUBSCRIPTION"
)

// Basket Item Type constants
const (
	BasketItemTypePhysical = "PHYSICAL"
	BasketItemTypeVirtual  = "VIRTUAL"
)

// Payment Channel constants
const (
	PaymentChannelMobile        = "MOBILE"
	PaymentChannelWeb           = "WEB"
	PaymentChannelMobileWeb     = "MOBILE_WEB"
	PaymentChannelMobileIOS     = "MOBILE_IOS"
	PaymentChannelMobileAndroid = "MOBILE_ANDROID"
	PaymentChannelMobileWindows = "MOBILE_WINDOWS"
	PaymentChannelMobileTablet  = "MOBILE_TABLET"
	PaymentChannelMobilePhone   = "MOBILE_PHONE"
)

// Sub Merchant Type constants
const (
	SubMerchantTypePersonal                    = "PERSONAL"
	SubMerchantTypePrivateCompany              = "PRIVATE_COMPANY"
	SubMerchantTypeLimitedOrJointStockCompany = "LIMITED_OR_JOINT_STOCK_COMPANY"
)

// APM Type constants
const (
	APMTypeSofort  = "SOFORT"
	APMTypeIdeal   = "IDEAL"
	APMTypeQiwi    = "QIWI"
	APMTypeGiropay = "GIROPAY"
)

// Refund Reason constants
const (
	RefundReasonDoublePayment = "double_payment"
	RefundReasonBuyerRequest  = "buyer_request"
	RefundReasonFraud         = "fraud"
	RefundReasonOther         = "other"
)

// Plan Payment Type constants
const (
	PlanPaymentTypeRecurring = "RECURRING"
)

// Subscription Pricing Plan Interval constants
const (
	SubscriptionPricingPlanIntervalDaily   = "DAILY"
	SubscriptionPricingPlanIntervalWeekly  = "WEEKLY"
	SubscriptionPricingPlanIntervalMonthly = "MONTHLY"
	SubscriptionPricingPlanIntervalYearly  = "YEARLY"
)

// Subscription Upgrade Period constants
const (
	SubscriptionUpgradePeriodNow = "NOW"
)

// Subscription Status constants
const (
	SubscriptionStatusExpired   = "EXPIRED"
	SubscriptionStatusUnpaid    = "UNPAID"
	SubscriptionStatusCanceled  = "CANCELED"
	SubscriptionStatusActive    = "ACTIVE"
	SubscriptionStatusPending   = "PENDING"
	SubscriptionStatusUpgraded  = "UPGRADED"
)

// Subscription Initial Status constants
const (
	SubscriptionInitialStatusActive  = "ACTIVE"
	SubscriptionInitialStatusPending = "PENDING"
)

// HTTP Headers
const (
	HeaderRandomString                 = "x-iyzi-rnd"
	HeaderClientVersion                = "x-iyzi-client-version"
	HeaderAuthorization                = "Authorization"
	HeaderAuthorizationFallback        = "Authorization_Fallback"
	HeaderIyziWSV1                     = "IYZWS"
	HeaderIyziWSV2                     = "IYZWSv2"
	ClientVersion                      = "iyzipay-go-1.0.0"
	Separator                          = ":"
	RandomStringSize                   = 8
)

// API Endpoints
const (
	EndpointAPITest                                    = "/payment/test"
	EndpointPaymentAuth                               = "/payment/auth"
	EndpointPaymentAuthBasic                          = "/payment/auth/basic"
	EndpointPaymentPreAuth                            = "/payment/preauth"
	EndpointPaymentPreAuthBasic                       = "/payment/preauth/basic"
	EndpointPaymentPostAuth                           = "/payment/postauth"
	EndpointPaymentPostAuthBasic                      = "/payment/postauth/basic"
	EndpointPaymentDetail                             = "/payment/detail"
	EndpointPaymentCancel                             = "/payment/cancel"
	EndpointPaymentRefund                             = "/payment/refund"
	EndpointPayment3DSecureInitialize                 = "/payment/3dsecure/initialize"
	EndpointPayment3DSecureInitializeBasic            = "/payment/3dsecure/initialize/basic"
	EndpointPayment3DSecureInitializePreAuth          = "/payment/3dsecure/initialize/preauth"
	EndpointPayment3DSecureInitializePreAuthBasic     = "/payment/3dsecure/initialize/preauth/basic"
	EndpointPayment3DSecureAuth                       = "/payment/3dsecure/auth"
	EndpointPayment3DSecureAuthBasic                  = "/payment/3dsecure/auth/basic"
	EndpointCheckoutFormInitializeAuth                = "/payment/iyzipos/checkoutform/initialize/auth/ecom"
	EndpointCheckoutFormInitializePreAuth             = "/payment/iyzipos/checkoutform/initialize/preauth/ecom"
	EndpointCheckoutFormAuthDetail                    = "/payment/iyzipos/checkoutform/auth/ecom/detail"
	EndpointPaymentBinCheck                           = "/payment/bin/check"
	EndpointPaymentInstallment                        = "/payment/iyzipos/installment"
	EndpointPaymentInstallmentHTML                    = "/payment/iyzipos/installment/html/horizontal"
	EndpointCardStorageCard                           = "/cardstorage/card"
	EndpointCardStorageCards                          = "/cardstorage/cards"
	EndpointBKMInitialize                             = "/payment/bkm/initialize"
	EndpointBKMInitializeBasic                        = "/payment/bkm/initialize/basic"
	EndpointBKMAuthDetail                             = "/payment/bkm/auth/detail"
	EndpointBKMAuthDetailBasic                        = "/payment/bkm/auth/detail/basic"
	EndpointAPMInitialize                             = "/payment/apm/initialize"
	EndpointAPMRetrieve                               = "/payment/apm/retrieve"
	EndpointPeccoInitialize                           = "/payment/pecco/initialize"
	EndpointPeccoAuth                                 = "/payment/pecco/auth"
	EndpointSubMerchant                               = "/onboarding/submerchant"
	EndpointSubMerchantDetail                         = "/onboarding/submerchant/detail"
	EndpointPaymentItem                               = "/payment/item"
	EndpointPaymentItemApprove                        = "/payment/iyzipos/item/approve"
	EndpointPaymentItemDisapprove                     = "/payment/iyzipos/item/disapprove"
	EndpointCrossBookingSend                          = "/crossbooking/send"
	EndpointCrossBookingReceive                       = "/crossbooking/receive"
	EndpointRefundToBalance                           = "/payment/refund-to-balance/init"
	EndpointSettlementToBalance                       = "/payment/settlement-to-balance/init"
	EndpointRefundChargedFromMerchant                 = "/payment/iyzipos/refund/merchant/charge"
	EndpointReportingSettlementBounced                = "/reporting/settlement/bounced"
	EndpointReportingSettlementPayoutCompleted       = "/reporting/settlement/payoutcompleted"
	EndpointUniversalCardStorageInitialize            = "/v2/ucs/init"
	EndpointSubscriptionInitialize                    = "/v2/subscription/initialize"
	EndpointSubscriptionInitializeWithCustomer       = "/v2/subscription/initialize/with-customer"
	EndpointSubscriptionCheckoutFormInitialize        = "/v2/subscription/checkoutform/initialize"
	EndpointSubscriptionCheckoutFormRetrieve          = "/v2/subscription/checkoutform/{checkoutFormToken}"
	EndpointSubscriptionCardUpdateInitialize          = "/v2/subscription/card-update/checkoutform/initialize"
	EndpointSubscriptionCardUpdateWithSubscription    = "/v2/subscription/card-update/checkoutform/initialize/with-subscription"
	EndpointSubscriptionPaymentRetry                  = "/v2/subscription/operation/retry"
	EndpointSubscriptionCancel                        = "/v2/subscription/subscriptions/{subscriptionReferenceCode}/cancel"
	EndpointSubscriptionActivate                      = "/v2/subscription/subscriptions/{subscriptionReferenceCode}/activate"
	EndpointSubscriptionUpgrade                       = "/v2/subscription/subscriptions/{subscriptionReferenceCode}/upgrade"
	EndpointSubscriptionRetrieve                      = "/v2/subscription/subscriptions/{subscriptionReferenceCode}"
	EndpointSubscriptionSearch                        = "/v2/subscription/subscriptions"
	EndpointSubscriptionCustomers                     = "/v2/subscription/customers"
	EndpointSubscriptionCustomer                      = "/v2/subscription/customers/{customerReferenceCode}"
	EndpointSubscriptionProducts                      = "/v2/subscription/products"
	EndpointSubscriptionProduct                       = "/v2/subscription/products/{productReferenceCode}"
	EndpointSubscriptionPricingPlans                  = "/v2/subscription/products/{productReferenceCode}/pricing-plans"
	EndpointSubscriptionPricingPlan                   = "/v2/subscription/pricing-plans/{pricingPlanReferenceCode}"
	EndpointSubscriptionPricingPlanRetrieve           = "/v2/subscription/pricing-plans/{pricingPlanReferenceCode}"
)