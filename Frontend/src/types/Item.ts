export interface Item {
  item_id: number
  item_name: string
  item_price: number
  price_per_unit: number
  units: number
  store_branch: string
  weight: number
  category: string | null
  subcategory: string | null
  is_organic: boolean | null
  brand_name: string | null
  barcode: string | null
  updated_at: string
}
