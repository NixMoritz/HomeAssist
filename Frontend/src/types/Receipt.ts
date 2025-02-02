export interface Receipt {
  receipt_id: number
  store_id: number
  store_name: string
  store_branch: string
  store_address: string
  date_issued: string
  total_amount: number
  payment_method: string
  total_discount_amount: number
  net_amount: number
  tax_amount: number
  receipt_type: string
  loyalty_card_number: string
  cashier_name: string
  receipt_number: string
  notes: string
  image_url: string
  updated_at: string
}
