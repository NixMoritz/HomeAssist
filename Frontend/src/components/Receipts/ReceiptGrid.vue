<template>
  <div class="receipts-container">
    <!-- Loading State -->
    <div v-if="loading" class="loading-state">Loading receipts...</div>

    <!-- Error State -->
    <div v-else-if="error" class="error-state">
      <p>Error loading receipts: {{ error }}</p>
      <button @click="() => retryFetchReceipts()" class="retry-button">Retry</button>
    </div>

    <!-- Header with Add Button -->
    <div class="page-header">
      <button @click="openAddReceiptModal" class="add-receipt-button">Add Receipt</button>
    </div>

    <ReceiptModal
      v-if="showModal || showAddReceiptModal"
      :receipt="showModal ? editedReceipt : newReceipt"
      :isEdit="showModal"
      @save="showModal ? saveEditedReceipt() : addNewReceipt()"
      @close="showModal ? closeModal() : closeAddReceiptModal()"
      @openStoreSelect="openStoreSelection(showModal)"
    />

    <!-- Main Content -->
    <div class="content-wrapper">
      <!-- Receipts Grid -->
      <div class="receipts-grid">
        <div v-for="receipt in receipts" :key="receipt.receipt_id" class="receipt-card">
          <div class="receipt-content">
            <h3 class="receipt-title">Receipt #{{ receipt.receipt_id }}</h3>
            <div class="receipt-details">
              <p><strong>Date:</strong> {{ formatDate(receipt.date_issued) }}</p>
              <p>
                <strong>Store:</strong>
                {{
                  receipt.store_id
                    ? `${receipt.store_name} (${receipt.store_branch})`
                    : 'No store selected'
                }}
              </p>
              <p>
                <strong>Address:</strong>
                {{ receipt.store_id ? receipt.store_address : 'No address available' }}
              </p>
              <p><strong>Total Amount:</strong> {{ formatPrice(receipt.total_amount) }}</p>
              <p><strong>Discount:</strong> {{ formatPrice(receipt.total_discount_amount) }}</p>
              <p><strong>Net Amount:</strong> {{ formatPrice(receipt.net_amount) }}</p>
            </div>
            <div class="receipt-image" v-if="receipt.image_url">
              <img :src="receipt.image_url" alt="Receipt Image" />
            </div>
            <div class="receipt-actions">
              <button @click="editReceipt(receipt)" class="edit-button">Edit</button>
              <button @click="deleteReceipt(receipt.receipt_id)" class="delete-button">
                Delete
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Store Selection Modal -->
    <div v-if="showStoreSelectionModal" class="modal-overlay">
      <div class="modal store-list-modal">
        <h2>Select Store</h2>
        <div class="store-list">
          <div class="store-list-header">
            <div class="store-name-col">Store Name</div>
            <div class="store-branch-col">Branch</div>
            <div class="store-address-col">Address</div>
            <div class="store-phone-col">Phone</div>
          </div>
          <div class="store-list-body">
            <div
              v-for="store in stores"
              :key="store.store_id"
              class="store-list-item"
              @click="selectStore(store)"
            >
              <div class="store-name-col">{{ store.store_name }}</div>
              <div class="store-branch-col">{{ store.store_branch }}</div>
              <div class="store-address-col">{{ store.store_address }}</div>
              <div class="store-phone-col">{{ store.store_phone }}</div>
            </div>
          </div>
        </div>
        <div class="modal-actions">
          <button @click="showStoreSelectionModal = false" class="cancel-button">Cancel</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { toRaw } from 'vue'
import ReceiptModal from './ReceiptModal.vue'

interface Receipt {
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

interface Store {
  store_id: number
  unique_uid: string
  store_branch: string
  store_name: string
  store_address: string
  store_phone: string
  updated_at: string
}

const receipts = ref<Receipt[]>([])
const loading = ref(true)
const error = ref<string | null>(null)
const showModal = ref(false)
const showAddReceiptModal = ref(false)
const stores = ref<Store[]>([])

const editedReceipt = ref<Receipt>({
  receipt_id: 0,
  store_id: 0,
  store_name: '',
  store_branch: '',
  store_address: '',
  date_issued: '',
  total_amount: 0,
  payment_method: 'CASH',
  total_discount_amount: 0,
  net_amount: 0,
  tax_amount: 0,
  receipt_type: 'PURCHASE',
  loyalty_card_number: '',
  cashier_name: '',
  receipt_number: '',
  notes: '',
  image_url: '',
  updated_at: '',
})

const newReceipt = ref<Receipt>({
  receipt_id: 0,
  store_id: 0,
  store_name: '',
  store_branch: '',
  store_address: '',
  date_issued: new Date().toISOString().slice(0, 16), // Format: YYYY-MM-DDTHH:mm
  total_amount: 0,
  payment_method: 'CASH',
  total_discount_amount: 0,
  net_amount: 0,
  tax_amount: 0,
  receipt_type: 'PURCHASE',
  loyalty_card_number: '',
  cashier_name: '',
  receipt_number: '',
  notes: '',
  image_url: '',
  updated_at: new Date().toISOString(),
})

// Add new ref for store selection modal
const showStoreSelectionModal = ref(false)
const isEditMode = ref(false)

// Fetch receipts from API
const fetchReceipts = async () => {
  try {
    loading.value = true
    error.value = null
    const response = await fetch('http://localhost:8080/api/receipts')
    if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`)
    receipts.value = await response.json()
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Failed to fetch receipts'
  } finally {
    loading.value = false
  }
}

const retryFetchReceipts = async (retries = 3) => {
  for (let attempt = 1; attempt <= retries; attempt++) {
    try {
      await fetchReceipts()
      return
    } catch (e) {
      if (attempt === retries) throw e
    }
  }
}

// Fetch stores on component mount
const fetchStores = async () => {
  try {
    console.log('Fetching stores...')
    const response = await fetch('http://localhost:8080/api/stores/all', {
      method: 'GET',
      headers: {
        Accept: 'application/json',
      },
    })

    if (!response.ok) {
      const text = await response.text()
      console.error('Store fetch failed:', response.status, text)
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const data = await response.json()
    console.log('Fetched stores:', data)

    if (!Array.isArray(data)) {
      console.error('Expected array of stores, got:', typeof data)
      throw new Error('Invalid store data format')
    }

    stores.value = data
    console.log('Stores updated:', stores.value.length, 'stores')
  } catch (error) {
    console.error('Failed to fetch stores:', error)
    // Retry once after a short delay
    await new Promise((resolve) => setTimeout(resolve, 1000))
    try {
      const response = await fetch('http://localhost:8080/api/stores/all')
      if (!response.ok) throw new Error(`Retry failed: ${response.status}`)
      const data = await response.json()
      stores.value = data
      console.log('Stores loaded on retry:', stores.value.length)
    } catch (retryError) {
      console.error('Store fetch retry failed:', retryError)
    }
  }
}

// Auto-fill store details when a store is selected
const handleStoreSelect = (isEdit = false) => {
  const currentStoreId = isEdit ? editedReceipt.value.store_id : newReceipt.value.store_id
  console.log('Handling store select:', { currentStoreId, availableStores: stores.value })

  const selectedStore = stores.value.find((s) => s.store_id === currentStoreId)
  console.log('Selected store:', selectedStore)

  if (selectedStore) {
    const receipt = isEdit ? editedReceipt.value : newReceipt.value
    receipt.store_name = selectedStore.store_name
    receipt.store_branch = selectedStore.store_branch
    receipt.store_address = selectedStore.store_address
    console.log('Updated receipt with store details:', receipt)
  }
}

// Modal handlers
const openAddReceiptModal = async () => {
  newReceipt.value = {
    receipt_id: 0,
    store_id: 0,
    store_name: '',
    store_branch: '',
    store_address: '',
    date_issued: new Date().toISOString().slice(0, 16), // Format: YYYY-MM-DDTHH:mm
    total_amount: 0,
    payment_method: 'CASH',
    total_discount_amount: 0,
    net_amount: 0,
    tax_amount: 0,
    receipt_type: 'PURCHASE',
    loyalty_card_number: '',
    cashier_name: '',
    receipt_number: '',
    notes: '',
    image_url: '',
    updated_at: new Date().toISOString(),
  }
  await fetchStores()
  console.log('Stores available for add:', stores.value)
  showAddReceiptModal.value = true
}

const closeAddReceiptModal = () => {
  showAddReceiptModal.value = false
}

const handleImageUpload = async (event: Event) => {
  const input = event.target as HTMLInputElement
  if (input.files && input.files[0]) {
    const file = input.files[0]

    // Create a base64 string from the file
    const reader = new FileReader()
    reader.onload = (e) => {
      const receipt = isEditMode.value ? editedReceipt.value : newReceipt.value
      receipt.image_url = e.target?.result as string
    }
    reader.readAsDataURL(file)
  }
}

// CRUD operations
const addNewReceipt = async () => {
  try {
    const receiptData = {
      ...toRaw(newReceipt.value),
      date_issued: newReceipt.value.date_issued, // Already in correct format YYYY-MM-DDTHH:mm
    }

    const response = await fetch('http://localhost:8080/api/receipts/create', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(receiptData),
    })

    if (!response.ok) {
      const errorText = await response.text()
      throw new Error(`Failed to add receipt: ${response.status} - ${errorText}`)
    }

    const createdReceipt = await response.json()
    receipts.value.push(createdReceipt)
    showAddReceiptModal.value = false
    await retryFetchReceipts()
  } catch (error) {
    console.error('Failed to add new receipt:', error)
  }
}

const editReceipt = (receipt: Receipt) => {
  console.log('Editing receipt:', receipt)
  // Create a deep copy of the receipt to avoid reference issues
  editedReceipt.value = JSON.parse(JSON.stringify(receipt))

  // Format dates for input fields
  if (editedReceipt.value.date_issued) {
    editedReceipt.value.date_issued = new Date(editedReceipt.value.date_issued)
      .toISOString()
      .split('T')[0]
  }

  // Ensure store information is properly set
  if (editedReceipt.value.store_id) {
    const selectedStore = stores.value.find((s) => s.store_id === editedReceipt.value.store_id)
    if (selectedStore) {
      editedReceipt.value.store_name = selectedStore.store_name
      editedReceipt.value.store_branch = selectedStore.store_branch
      editedReceipt.value.store_address = selectedStore.store_address
    }
  }

  console.log('Prepared receipt for editing:', editedReceipt.value)
  showModal.value = true
}

const saveEditedReceipt = async () => {
  try {
    const receiptData = {
      ...toRaw(editedReceipt.value),
      date_issued: editedReceipt.value.date_issued, // Already in correct format YYYY-MM-DDTHH:mm
    }

    const response = await fetch('http://localhost:8080/api/receipts/update', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(receiptData),
    })

    if (!response.ok) {
      const errorText = await response.text()
      throw new Error(`Failed to update receipt: ${response.status} - ${errorText}`)
    }

    // Update the local receipt list
    const updatedReceipt = await response.json()
    receipts.value = receipts.value.map((receipt) =>
      receipt.receipt_id === updatedReceipt.receipt_id ? updatedReceipt : receipt,
    )

    showModal.value = false
  } catch (error) {
    console.error('Failed to update receipt:', error)
  }
}

const deleteReceipt = async (receiptId: number) => {
  if (!confirm('Are you sure you want to delete this receipt?')) return

  try {
    const response = await fetch(
      `http://localhost:8080/api/receipts/delete?receipt_id=${receiptId}`,
      {
        method: 'DELETE',
      },
    )

    if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`)

    receipts.value = receipts.value.filter((receipt) => toRaw(receipt).receipt_id !== receiptId)
  } catch (error) {
    console.error('Failed to delete receipt:', error)
  }
}

// Utility functions
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
  })
}

const formatTime = (timeString: string) => {
  const date = new Date(timeString)
  return date.toLocaleTimeString('de-DE', {
    hour: '2-digit',
    minute: '2-digit',
  })
}

const closeModal = () => {
  showModal.value = false
}

// Function to open store selection modal
const openStoreSelection = (forEdit: boolean) => {
  isEditMode.value = forEdit
  showStoreSelectionModal.value = true
}

// Function to handle store selection from the modal
const selectStore = (store: Store) => {
  console.log('Store selected:', store)
  const receipt = isEditMode.value ? editedReceipt.value : newReceipt.value
  receipt.store_id = store.store_id
  receipt.store_name = store.store_name
  receipt.store_branch = store.store_branch
  receipt.store_address = store.store_address
  showStoreSelectionModal.value = false
  console.log('Updated receipt:', receipt)
}

onMounted(async () => {
  console.log('Component mounted, loading data...')
  try {
    await fetchStores()
    await retryFetchReceipts()
  } catch (error) {
    console.error('Error during component mount:', error)
  }
})
</script>

<style scoped>
.receipts-container {
  width: 100%;
  min-height: 100vh;
  padding: 20px;
  margin-top: 64px; /* Add margin to account for topbar height */
}

.page-header {
  width: 65%;
  margin: 0 auto 20px;
  display: flex;
  justify-content: flex-end;
  position: relative;
  z-index: 1; /* Ensure it's above content but below topbar */
}

.content-wrapper {
  width: 65%;
  margin: 0 auto;
}

.add-receipt-button {
  background: rgba(59, 130, 246, 0.2);
  color: #60a5fa;
  padding: 10px 20px;
  border-radius: 8px;
  font-size: 0.95rem;
  border: 1px solid rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  transition: all 0.2s ease;
}

.add-receipt-button:hover {
  background: rgba(59, 130, 246, 0.3);
}

/* Responsive design */
@media (max-width: 1024px) {
  .page-header,
  .content-wrapper {
    width: 80%;
  }
}

@media (max-width: 768px) {
  .page-header,
  .content-wrapper {
    width: 90%;
  }
}

.receipts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 24px;
  padding: 20px 0;
}

.receipt-card {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  overflow: hidden;
  transition: all 0.3s ease;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.receipt-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 12px rgba(0, 0, 0, 0.15);
  border-color: rgba(255, 255, 255, 0.3);
}

.receipt-content {
  padding: 20px;
}

.receipt-title {
  margin: 0 0 16px 0;
  font-size: 1.3rem;
  color: var(--color-heading);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  padding-bottom: 12px;
}

.receipt-details {
  display: grid;
  gap: 12px;
}

.receipt-details p {
  margin: 0;
  font-size: 0.95rem;
  line-height: 1.4;
}

.receipt-details strong {
  color: var(--color-heading);
  font-weight: 600;
}

.receipt-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 20px;
  padding-top: 16px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.edit-button,
.delete-button {
  padding: 8px 16px;
  border-radius: 8px;
  border: none;
  cursor: pointer;
  font-size: 0.9rem;
  transition: all 0.2s ease;
}

.edit-button {
  background: rgba(59, 130, 246, 0.2);
  color: #60a5fa;
}

.delete-button {
  background: rgba(239, 68, 68, 0.2);
  color: #f87171;
}

.edit-button:hover {
  background: rgba(59, 130, 246, 0.3);
}

.delete-button:hover {
  background: rgba(239, 68, 68, 0.3);
}

.receipt-image {
  margin: 16px -20px -20px;
  height: 200px;
  overflow: hidden;
}

.receipt-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  background: var(--color-background-soft);
  padding: 20px;
  border-radius: 8px;
  width: 90%;
  max-width: 500px;
}

.modal-content {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.modal-actions {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.cancel-button {
  background-color: #6b7280;
  color: white;
}

input,
select {
  padding: 8px;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background: var(--color-background);
  color: var(--color-text);
}

label {
  font-weight: bold;
  margin-bottom: 4px;
}

.debug-info {
  background-color: #f3f4f6;
  padding: 8px;
  margin-bottom: 16px;
  border-radius: 4px;
  font-family: monospace;
}

.debug-info p {
  margin: 4px 0;
  font-size: 0.8em;
  color: #374151;
}

.store-list-modal {
  width: 90%;
  max-width: 1000px;
  max-height: 80vh;
  display: flex;
  flex-direction: column;
}

.store-list {
  flex-grow: 1;
  overflow-y: auto;
  margin: 1rem 0;
}

.store-list-header {
  display: grid;
  grid-template-columns: 2fr 2fr 3fr 2fr;
  gap: 1rem;
  padding: 0.5rem;
  background-color: var(--color-background-soft);
  font-weight: bold;
  position: sticky;
  top: 0;
}

.store-list-body {
  overflow-y: auto;
}

.store-list-item {
  display: grid;
  grid-template-columns: 2fr 2fr 3fr 2fr;
  gap: 1rem;
  padding: 0.75rem 0.5rem;
  border-bottom: 1px solid var(--color-border);
  cursor: pointer;
  transition: background-color 0.2s;
}

.store-list-item:hover {
  background-color: var(--color-background-soft);
}

.store-selection {
  margin-bottom: 1rem;
  padding: 1rem;
  border: 1px solid var(--color-border);
  border-radius: 4px;
}

.selected-store {
  margin: 0.5rem 0;
  padding: 0.5rem;
  background-color: var(--color-background-soft);
  border-radius: 4px;
}

.select-store-button {
  background-color: #2563eb;
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  border: none;
  cursor: pointer;
  width: 100%;
  margin-top: 0.5rem;
}

.select-store-button:hover {
  background-color: #1d4ed8;
}

.store-name-col,
.store-branch-col,
.store-address-col,
.store-phone-col {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.form-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1rem;
  margin-top: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-group.full-width {
  grid-column: span 2;
}

label {
  font-weight: 500;
  color: var(--color-text);
}

input,
select,
textarea {
  padding: 0.5rem;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background: var(--color-background);
  color: var(--color-text);
  width: 100%;
}

textarea {
  resize: vertical;
  min-height: 80px;
}

input[type='number'] {
  text-align: right;
}

.modal {
  width: 90%;
  max-width: 800px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-content {
  padding: 1.5rem;
}

/* Add these styles for better form layout */
.store-selection {
  margin-bottom: 2rem;
  padding: 1rem;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background: var(--color-background-soft);
}

.selected-store {
  margin: 0.5rem 0;
  padding: 0.5rem;
  background: var(--color-background);
  border-radius: 4px;
}

/* Make sure these styles are applied to both add and edit modals */
.modal-overlay .modal {
  padding: 0;
}

.modal-actions {
  padding: 1rem;
  border-top: 1px solid var(--color-border);
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  background: var(--color-background);
}

/* Common button styles */
.button-base {
  background-color: #ffffff;
  border: 1px solid #222222;
  border-radius: 8px;
  box-sizing: border-box;
  color: #222222;
  cursor: pointer;
  display: inline-block;
  font-family:
    Circular,
    -apple-system,
    BlinkMacSystemFont,
    Roboto,
    'Helvetica Neue',
    sans-serif;
  font-size: 16px;
  font-weight: 600;
  line-height: 20px;
  margin: 0;
  outline: none;
  padding: 13px 23px;
  position: relative;
  text-align: center;
  text-decoration: none;
  touch-action: manipulation;
  transition:
    box-shadow 0.2s,
    -ms-transform 0.1s,
    -webkit-transform 0.1s,
    transform 0.1s;
  user-select: none;
  -webkit-user-select: none;
  width: auto;
}

.button-base:focus-visible {
  box-shadow:
    #222222 0 0 0 2px,
    rgba(255, 255, 255, 0.8) 0 0 0 4px;
  transition: box-shadow 0.2s;
}

.button-base:active {
  background-color: #f7f7f7;
  border-color: #000000;
  transform: scale(0.96);
}

.button-base:disabled {
  border-color: #dddddd;
  color: #dddddd;
  cursor: not-allowed;
  opacity: 1;
}

/* Apply base styles to all buttons */
.add-receipt-button,
.edit-button,
.delete-button,
.select-store-button {
  @extend .button-base;
}

/* Custom colors for different button types */
.add-receipt-button {
  background-color: #4caf50;
  border-color: #45a049;
  color: white;
}

.edit-button {
  background-color: #2196f3;
  border-color: #1e88e5;
  color: white;
}

.delete-button {
  background-color: #f44336;
  border-color: #e53935;
  color: white;
}

.select-store-button {
  background-color: #9c27b0;
  border-color: #8e24aa;
  color: white;
  width: 100%;
}

/* Hover states */
.add-receipt-button:hover,
.edit-button:hover,
.delete-button:hover,
.select-store-button:hover {
  opacity: 0.9;
}

/* Update receipt actions container */
.receipt-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 20px;
  padding-top: 16px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .button-base {
    padding: 10px 18px;
    font-size: 14px;
  }
}
</style>
