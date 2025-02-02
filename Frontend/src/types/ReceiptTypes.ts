export const PAYMENT_METHODS = {
  CASH: 'CASH',
  CARD: 'CARD',
  ONLINE: 'ONLINE',
} as const

export const RECEIPT_TYPES = {
  PURCHASE: 'PURCHASE',
  RETURN: 'RETURN',
  REFUND: 'REFUND',
} as const

export type PaymentMethod = (typeof PAYMENT_METHODS)[keyof typeof PAYMENT_METHODS]
export type ReceiptType = (typeof RECEIPT_TYPES)[keyof typeof RECEIPT_TYPES]
