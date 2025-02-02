<template>
  <div class="items-container">
    <!-- Loading State -->
    <div v-if="loading" class="loading-state" role="status">Loading items...</div>

    <!-- Error State -->
    <div v-else-if="error" class="error-state" role="alert">
      <p>Error: {{ error }}</p>
      <button @click="retryFetchItems()" class="button-base button-primary" :disabled="processing">
        Retry
      </button>
    </div>

    <!-- Header with Add Button -->
    <div class="page-header">
      <button @click="openAddItemModal" class="button-base button-primary" :disabled="processing">
        Add Item
      </button>
    </div>

    <!-- Main Content -->
    <div class="content-wrapper">
      <!-- Items Grid -->
      <div class="items-grid" role="grid">
        <div
          v-for="item in items"
          :key="item.item_id"
          class="item-card"
          @click="showItemDetails(item)"
          role="gridcell"
        >
          <div class="item-content">
            <h3 class="item-title">{{ item.item_name }}</h3>
            <div class="item-details">
              <p><strong>Category:</strong> {{ getNullableString(item.category) }}</p>
              <p><strong>Price:</strong> {{ formatPrice(item.item_price) }}</p>
              <p><strong>Units:</strong> {{ item.units }}</p>
              <p><strong>Weight:</strong> {{ item.weight }}kg</p>
              <p><strong>Brand:</strong> {{ getNullableString(item.brand_name) }}</p>
            </div>
            <div class="item-actions">
              <button
                @click.stop="editItem(item)"
                class="button-base button-secondary"
                :disabled="processing"
              >
                Edit
              </button>
              <button
                @click.stop="deleteItem(item.item_id)"
                class="button-base button-danger"
                :disabled="processing"
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <ItemModal
    v-if="showModal || showAddItemModal"
    :item="showModal ? editedItem : newItem"
    :isEdit="showModal"
    :processing="processing"
    @save="showModal ? saveEditedItem() : addNewItem()"
    @close="showModal ? closeModal() : closeAddItemModal()"
  />
  <ItemDetailsModal v-if="showDetailsModal" :item="selectedItem" @close="closeDetailsModal" />
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import '@/components/styles/ButtonStyles.vue'
import ItemModal from './ItemModal.vue'
import ItemDetailsModal from './ItemDetailsModal.vue'
import type { Item } from '@/types/Item'

// API configuration
const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080/api'

// State management
const items = ref<Item[]>([])
const loading = ref(true)
const processing = ref(false)
const error = ref<string | null>(null)
const showModal = ref(false)
const showAddItemModal = ref(false)
const showDetailsModal = ref(false)
const selectedItem = ref<Item | null>(null)

// Item templates
const editedItem = ref<Item>({
  item_id: 0,
  item_name: '',
  item_price: 0,
  price_per_unit: 0,
  units: 0,
  store_branch: '',
  weight: 0,
  category: null,
  subcategory: null,
  is_organic: null,
  brand_name: null,
  barcode: null,
  updated_at: '',
})

const newItem = ref<Item>({
  ...editedItem.value,
  units: 1,
  updated_at: new Date().toISOString(),
})

// API calls with improved error handling
const fetchItems = async () => {
  try {
    loading.value = true
    error.value = null
    const response = await fetch(`${API_BASE}/items`)
    if (!response.ok) throw new Error(`HTTP ${response.status}`)
    items.value = await response.json()
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Failed to fetch items'
    console.error('Fetch error:', e)
  } finally {
    loading.value = false
  }
}

const retryFetchItems = async (retries = 3) => {
  for (let attempt = 1; attempt <= retries; attempt++) {
    try {
      await new Promise((resolve) => setTimeout(resolve, attempt * 1000))
      await fetchItems()
      return
    } catch (e) {
      if (attempt === retries) throw e
    }
  }
}

const deleteItem = async (itemId: number) => {
  if (!confirm('Are you sure you want to delete this item?')) return

  try {
    processing.value = true
    const response = await fetch(`${API_BASE}/items/delete?item_id=${itemId}`, {
      method: 'DELETE',
    })
    if (!response.ok) throw new Error(`HTTP ${response.status}`)

    items.value = items.value.filter((item) => item.item_id !== itemId)
  } catch (e) {
    error.value = 'Failed to delete item'
    console.error('Delete error:', e)
  } finally {
    processing.value = false
  }
}

const saveEditedItem = async () => {
  try {
    processing.value = true
    const response = await fetch(`${API_BASE}/items/update`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(editedItem.value),
    })

    if (!response.ok) throw new Error(`HTTP ${response.status}`)
    const updatedItem = await response.json()

    items.value = items.value.map((item) =>
      item.item_id === editedItem.value.item_id ? updatedItem : item,
    )
    showModal.value = false
  } catch (e) {
    error.value = 'Failed to update item'
    console.error('Update error:', e)
  } finally {
    processing.value = false
  }
}

const addNewItem = async () => {
  try {
    processing.value = true
    const response = await fetch(`${API_BASE}/items/create`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newItem.value),
    })

    if (!response.ok) throw new Error(`HTTP ${response.status}`)
    const createdItem = await response.json()

    items.value.push(createdItem)
    showAddItemModal.value = false
  } catch (e) {
    error.value = 'Failed to add item'
    console.error('Add error:', e)
  } finally {
    processing.value = false
  }
}

// UI handlers
const editItem = (item: Item) => {
  editedItem.value = { ...item }
  showModal.value = true
}

const showItemDetails = (item: Item) => {
  selectedItem.value = item
  showDetailsModal.value = true
}

// Utility functions
const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR',
  }).format(price)
}

const getNullableString = (value: string | null | undefined): string => {
  return value ?? '-'
}

const openAddItemModal = () => {
  showAddItemModal.value = true
  newItem.value = {
    item_id: 0,
    item_name: '',
    item_price: 0,
    price_per_unit: 0,
    units: 1,
    store_branch: '',
    weight: 0,
    category: null,
    subcategory: null,
    is_organic: null,
    brand_name: null,
    barcode: null,
    updated_at: new Date().toISOString(),
  }
}

const closeAddItemModal = () => {
  showAddItemModal.value = false
}

const closeModal = () => {
  showModal.value = false
}

const closeDetailsModal = () => {
  showDetailsModal.value = false
  selectedItem.value = null
}

onMounted(() => {
  retryFetchItems()
})
</script>

<style scoped>
/* General Container Styling */
.items-container {
  width: 100vw;
  min-height: 100vh;
  margin-top: 64px;
  box-sizing: border-box;
  overflow-x: hidden;
  position: relative;
  margin-left: -50vw;
  margin-right: -50vw;
}

/* Header */
.page-header {
  width: 100%;
  margin: 0 auto 20px;
  display: flex;
  justify-content: flex-end;
  position: relative;
  z-index: 1;
}

/* Content Wrapper */
.content-wrapper {
  width: 100%;
  margin: 0 auto;
  padding: 0 15%;
  box-sizing: border-box;
}

/* Grid Layout - now using flexbox */
.items-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  padding: 20px 0;
  width: 100%;
}

.item-card {
  flex: 1 1 calc(20% - 16px); /* 20% width for 5 items per row, accounting for gap */
  min-width: 280px; /* Minimum width to prevent too small cards */
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  overflow: hidden;
  transition: all 0.3s ease;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.item-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 12px rgba(0, 0, 0, 0.15);
  border-color: rgba(255, 255, 255, 0.3);
}

/* Item Content */
.item-content {
  padding: 20px;
}

.item-title {
  margin: 0 0 16px;
  font-size: clamp(1rem, 2vw, 1.2rem);
  color: var(--color-heading);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  padding-bottom: 12px;
}

.item-details {
  display: grid;
  gap: 12px;
}

.item-details p {
  margin: 0;
  font-size: clamp(0.9rem, 1.8vw, 1rem);
  line-height: 1.4;
}

.item-details strong {
  color: var(--color-heading);
  font-weight: 600;
}

/* Button Actions */
.item-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 20px;
  padding-top: 16px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

/* Loading & Error States */
.loading-state,
.error-state {
  text-align: center;
  padding: 40px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  margin: 20px auto;
  width: 100%;
  max-width: 1200px;
}

.error-state {
  color: #ef4444;
}

/* RESPONSIVE BREAKPOINTS */
@media (max-width: 1400px) {
  .item-card {
    flex: 1 1 calc(25% - 15px); /* 4 items per row */
  }
}

@media (max-width: 1200px) {
  .item-card {
    flex: 1 1 calc(33.333% - 14px); /* 3 items per row */
  }
}

@media (max-width: 900px) {
  .item-card {
    flex: 1 1 calc(50% - 10px); /* 2 items per row */
  }
}

@media (max-width: 600px) {
  .item-card {
    flex: 1 1 100%; /* 1 item per row */
  }
}
</style>
