<template>
  <div class="items-container">
    <!-- Loading State -->
    <div v-if="loading" class="loading-state">Loading items...</div>

    <!-- Error State -->
    <div v-else-if="error" class="error-state">
      <p>Error loading items: {{ error }}</p>
      <button @click="retryFetchItems()" class="button-base button-primary">Retry</button>
    </div>

    <!-- Header with Add Button -->
    <div class="page-header">
      <button @click="openAddItemModal" class="button-base button-primary">Add Item</button>
    </div>

    <!-- Main Content -->
    <div class="content-wrapper">
      <!-- Items Grid -->
      <div class="items-grid">
        <div v-for="item in items" :key="item.item_id" class="item-card">
          <div class="item-content">
            <h3 class="item-title">{{ item.item_name }}</h3>
            <div class="item-details">
              <p><strong>Category:</strong> {{ getNullableString(item.category) }}</p>
              <p><strong>Price:</strong> {{ formatPrice(item.unit_price) }}</p>
              <p><strong>Units:</strong> {{ item.units }}</p>
              <p><strong>Weight:</strong> {{ item.weight }}kg</p>
              <p><strong>Brand:</strong> {{ getNullableString(item.brand_name) }}</p>
            </div>
            <div class="item-actions">
              <button @click="editItem(item)" class="button-base button-secondary">Edit</button>
              <button @click="deleteItem(item.item_id)" class="button-base button-danger">
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
    @save="showModal ? saveEditedItem() : addNewItem()"
    @close="showModal ? closeModal() : closeAddItemModal()"
  />
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { toRaw } from 'vue'
import '@/components/styles/ButtonStyles.vue'
import ItemModal from './ItemModal.vue'

interface Item {
  item_id: number
  item_name: string
  unit_price: number
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

const items = ref<Item[]>([])
const loading = ref(true)
const error = ref<string | null>(null)
const showModal = ref(false)
const editedItem = ref<Item>({
  item_id: 0,
  item_name: '',
  unit_price: 0,
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

const showAddItemModal = ref(false)
const newItem = ref<Item>({
  item_id: 0,
  item_name: '',
  unit_price: 0,
  units: 0,
  store_branch: '',
  weight: 0,
  category: null,
  subcategory: null,
  is_organic: null,
  brand_name: null,
  barcode: null,
  updated_at: new Date().toISOString(),
})

const openAddItemModal = () => {
  showAddItemModal.value = true
  newItem.value = {
    item_id: 0,
    item_name: '',
    unit_price: 0,
    units: 0,
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

const addNewItem = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/items/create', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newItem.value),
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const createdItem = await response.json()
    items.value.push(createdItem)
    console.log(`New item added successfully:`, createdItem)
    showAddItemModal.value = false
  } catch (error) {
    console.error('Failed to add new item:', error)
  }
}

const fetchItems = async () => {
  try {
    loading.value = true
    error.value = null

    console.log('Attempting to fetch items from server...')
    const response = await fetch('http://localhost:8080/api/items', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    })

    if (!response.ok) {
      const errorText = await response.text()
      console.error('Server responded with error:', response.status, errorText)
      throw new Error(`Server error: ${response.status} - ${errorText || response.statusText}`)
    }

    const textData = await response.text()
    console.log('Raw server response:', textData)

    try {
      items.value = JSON.parse(textData)
      console.log('Successfully parsed items:', items.value)
    } catch (parseError) {
      console.error('Failed to parse server response:', parseError)
      throw new Error('Invalid server response format')
    }
  } catch (e) {
    console.error('Fetch error details:', e)
    error.value = e instanceof Error ? e.message : 'Network connection failed'
  } finally {
    loading.value = false
  }
}

const retryFetchItems = async (retries = 3) => {
  for (let attempt = 1; attempt <= retries; attempt++) {
    try {
      await fetchItems()
      return // Exit loop on success
    } catch (e) {
      console.warn(`Attempt ${attempt} failed:`, e)
      if (attempt === retries) throw e
    }
  }
}

onMounted(() => {
  retryFetchItems()
})

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR',
  }).format(price)
}

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString('de-DE', {
    weekday: 'short',
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
  })
}

const deleteItem = async (itemId: number) => {
  console.log('Item ID to delete:', itemId)

  if (itemId === undefined || itemId === null) {
    console.error('Invalid item ID')
    return
  }

  if (!confirm('Are you sure you want to delete this item?')) {
    return
  }

  try {
    const response = await fetch(`http://localhost:8080/api/items/delete?item_id=${itemId}`, {
      method: 'DELETE',
    })
    if (!response.ok) {
      console.error(`Failed to delete item, HTTP error! status: ${response.status}`)
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    // Use `toRaw` to get the unwrapped item data
    items.value = items.value.filter((item) => toRaw(item).item_id !== itemId)
    console.log(`Item with ID ${itemId} deleted successfully`)
  } catch (error) {
    console.error('Failed to delete item:', error)
  }
}
const editItem = (item: Item) => {
  editedItem.value = { ...item } // Copy item to editedItem for modal
  showModal.value = true
}
const closeModal = () => {
  showModal.value = false
}

const saveEditedItem = async () => {
  try {
    const response = await fetch(`http://localhost:8080/api/items/update`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(editedItem.value),
    })

    if (!response.ok) {
      console.error(`Failed to update item, HTTP error! status: ${response.status}`)
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const responseData = await response.json()

    items.value = items.value.map((item) =>
      item.item_id === editedItem.value.item_id ? { ...item, ...editedItem.value } : item,
    )

    console.log(
      `Item with ID ${editedItem.value.item_id} updated successfully:`,
      responseData.message,
    )
    showModal.value = false
  } catch (error) {
    console.error('Failed to update item:', error)
  }
}

// Helper function to handle nullable strings
const getNullableString = (value: string | null): string => {
  return value ?? '-'
}
</script>

<style scoped>
/* General Container Styling */
.items-container {
  width: 100vw; /* Full viewport width */
  min-height: 100vh;
  padding: 20px 0; /* Remove horizontal padding */
  margin-top: 64px;
  box-sizing: border-box;
  overflow-x: hidden; /* Prevent horizontal scroll */
  position: relative; /* Ensure proper positioning */
  left: 50%; /* Center the container */
  right: 50%;
  margin-left: -50vw; /* Negative margin to stretch full width */
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
  font-size: clamp(1.2rem, 2vw, 1.5rem);
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
