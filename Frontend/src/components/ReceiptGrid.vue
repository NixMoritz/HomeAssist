<template>
  <div class="receipts-container">
    <!-- Loading State -->
    <div v-if="loading" class="loading-state">Loading receipts...</div>

    <!-- Error State -->
    <div v-else-if="error" class="error-state">
      <p>Error loading receipts: {{ error }}</p>
      <button @click="() => retryFetchReceipts()" class="retry-button">Retry</button>
    </div>

    <!-- Add Receipt Button -->
    <div class="add-receipt-container">
      <button @click="openAddReceiptModal" class="add-receipt-button">Add Receipt</button>
    </div>

    <!-- Add Receipt Modal -->
    <div v-if="showAddReceiptModal" class="modal-overlay">
      <div class="modal">
        <h2>Add New Receipt</h2>
        <div class="modal-content">
          <div class="store-selection">
            <label>Selected Store:</label>
            <div class="selected-store">
              <span v-if="newReceipt.store_id">
                {{ newReceipt.store_name }} - {{ newReceipt.store_branch }}
              </span>
              <span v-else>No store selected</span>
            </div>
            <button @click="openStoreSelection(false)" class="select-store-button">
              Select Store
            </button>
          </div>

          <div class="form-grid">
            <div class="form-group">
              <label for="date">Date and Time:</label>
              <input type="datetime-local" id="datetime" v-model="newReceipt.date_issued" />
            </div>

            <div class="form-group">
              <label for="totalAmount">Total Amount:</label>
              <input
                type="number"
                id="totalAmount"
                v-model.number="newReceipt.total_amount"
                step="0.01"
              />
            </div>

            <div class="form-group">
              <label for="paymentMethod">Payment Method:</label>
              <select id="paymentMethod" v-model="newReceipt.payment_method">
                <option value="CASH">Cash</option>
                <option value="CARD">Card</option>
                <option value="ONLINE">Online</option>
              </select>
            </div>

            <div class="form-group">
              <label for="discountAmount">Discount Amount:</label>
              <input
                type="number"
                id="discountAmount"
                v-model.number="newReceipt.total_discount_amount"
                step="0.01"
              />
            </div>

            <div class="form-group">
              <label for="netAmount">Net Amount:</label>
              <input
                type="number"
                id="netAmount"
                v-model.number="newReceipt.net_amount"
                step="0.01"
              />
            </div>

            <div class="form-group">
              <label for="taxAmount">Tax Amount:</label>
              <input
                type="number"
                id="taxAmount"
                v-model.number="newReceipt.tax_amount"
                step="0.01"
              />
            </div>

            <div class="form-group">
              <label for="receiptType">Receipt Type:</label>
              <select id="receiptType" v-model="newReceipt.receipt_type">
                <option value="PURCHASE">Purchase</option>
                <option value="RETURN">Return</option>
                <option value="REFUND">Refund</option>
              </select>
            </div>

            <div class="form-group">
              <label for="loyaltyCard">Loyalty Card Number:</label>
              <input type="text" id="loyaltyCard" v-model="newReceipt.loyalty_card_number" />
            </div>

            <div class="form-group">
              <label for="cashierName">Cashier Name:</label>
              <input type="text" id="cashierName" v-model="newReceipt.cashier_name" />
            </div>

            <div class="form-group">
              <label for="receiptNumber">Receipt Number:</label>
              <input type="text" id="receiptNumber" v-model="newReceipt.receipt_number" />
            </div>

            <div class="form-group full-width">
              <label for="notes">Notes:</label>
              <textarea id="notes" v-model="newReceipt.notes" rows="3"></textarea>
            </div>

            <div class="form-group full-width">
              <label for="imageUrl">Receipt Image URL:</label>
              <input type="text" id="imageUrl" v-model="newReceipt.image_url" />
            </div>
          </div>
        </div>
        <div class="modal-actions">
          <button @click="addNewReceipt" class="save-button">Add</button>
          <button @click="closeAddReceiptModal" class="cancel-button">Cancel</button>
        </div>
      </div>
    </div>

    <!-- Receipts Grid -->
    <div v-else class="receipts-grid">
      <div v-for="receipt in receipts" :key="receipt.receipt_id" class="receipt-card">
        <div class="receipt-content">
          <h3 class="receipt-title">Receipt #{{ receipt.receipt_id }}</h3>
          <div class="receipt-details">
            <p>
              <strong>Date:</strong> {{ formatDate(receipt.date_issued) }}
              {{ formatTime(receipt.time_issued) }}
            </p>
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
            <button @click="deleteReceipt(receipt.receipt_id)" class="delete-button">Delete</button>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Edit Modal -->
  <div v-if="showModal" class="modal-overlay">
    <div class="modal">
      <h2>Edit Receipt</h2>
      <div class="modal-content">
        <div class="store-selection">
          <label>Selected Store:</label>
          <div class="selected-store">
            <span v-if="editedReceipt.store_id">
              {{ editedReceipt.store_name }} - {{ editedReceipt.store_branch }}
            </span>
            <span v-else>No store selected</span>
          </div>
          <button @click="openStoreSelection(true)" class="select-store-button">
            Select Store
          </button>
        </div>

        <div class="form-grid">
          <div class="form-group">
            <label for="editDate">Date:</label>
            <input type="date" id="editDate" v-model="editedReceipt.date_issued" />
          </div>

          <div class="form-group">
            <label for="editTime">Time:</label>
            <input
              type="time"
              id="editTime"
              :value="formattedTime"
              @input="(e) => updateTime((e.target as HTMLInputElement).value)"
            />
          </div>

          <div class="form-group">
            <label for="editTotalAmount">Total Amount:</label>
            <input
              type="number"
              id="editTotalAmount"
              v-model.number="editedReceipt.total_amount"
              step="0.01"
            />
          </div>

          <div class="form-group">
            <label for="editPaymentMethod">Payment Method:</label>
            <select id="editPaymentMethod" v-model="editedReceipt.payment_method">
              <option value="CASH">Cash</option>
              <option value="CARD">Card</option>
              <option value="ONLINE">Online</option>
            </select>
          </div>

          <div class="form-group">
            <label for="editDiscountAmount">Discount Amount:</label>
            <input
              type="number"
              id="editDiscountAmount"
              v-model.number="editedReceipt.total_discount_amount"
              step="0.01"
            />
          </div>

          <div class="form-group">
            <label for="editNetAmount">Net Amount:</label>
            <input
              type="number"
              id="editNetAmount"
              v-model.number="editedReceipt.net_amount"
              step="0.01"
            />
          </div>

          <div class="form-group">
            <label for="editTaxAmount">Tax Amount:</label>
            <input
              type="number"
              id="editTaxAmount"
              v-model.number="editedReceipt.tax_amount"
              step="0.01"
            />
          </div>

          <div class="form-group">
            <label for="editReceiptType">Receipt Type:</label>
            <select id="editReceiptType" v-model="editedReceipt.receipt_type">
              <option value="PURCHASE">Purchase</option>
              <option value="RETURN">Return</option>
              <option value="REFUND">Refund</option>
            </select>
          </div>

          <div class="form-group">
            <label for="editLoyaltyCard">Loyalty Card Number:</label>
            <input type="text" id="editLoyaltyCard" v-model="editedReceipt.loyalty_card_number" />
          </div>

          <div class="form-group">
            <label for="editCashierName">Cashier Name:</label>
            <input type="text" id="editCashierName" v-model="editedReceipt.cashier_name" />
          </div>

          <div class="form-group">
            <label for="editReceiptNumber">Receipt Number:</label>
            <input type="text" id="editReceiptNumber" v-model="editedReceipt.receipt_number" />
          </div>

          <div class="form-group full-width">
            <label for="editNotes">Notes:</label>
            <textarea id="editNotes" v-model="editedReceipt.notes" rows="3"></textarea>
          </div>

          <div class="form-group full-width">
            <label for="editImageUrl">Receipt Image URL:</label>
            <input type="text" id="editImageUrl" v-model="editedReceipt.image_url" />
          </div>
        </div>
      </div>
      <div class="modal-actions">
        <button @click="saveEditedReceipt" class="save-button">Save</button>
        <button @click="closeModal" class="cancel-button">Cancel</button>
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
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { toRaw } from 'vue'

interface Receipt {
  receipt_id: number
  store_id: number
  store_name: string
  store_branch: string
  store_address: string
  date_issued: string
  time_issued: string
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
  time_issued: '',
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
  if (editedReceipt.value.time_issued) {
    editedReceipt.value.time_issued = new Date(editedReceipt.value.time_issued)
      .toISOString()
      .split('T')[1]
      .slice(0, 5)
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

const updateTime = (time: string) => {
  if (time) {
    editedReceipt.value.time_issued = time
  }
}

const saveEditedReceipt = async () => {
  try {
    // Combine date and time for the backend format
    const dateStr = editedReceipt.value.date_issued
    const timeStr = editedReceipt.value.time_issued
    const combinedDateTime = `${dateStr}T${timeStr}`

    const receiptData = {
      ...toRaw(editedReceipt.value),
      date_issued: combinedDateTime,
      time_issued: combinedDateTime,
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

// Add computed property for formatted time
const formattedTime = computed(() => {
  const time = showModal.value ? editedReceipt.value.time_issued : newReceipt.value.time_issued
  return time.split('T')[1]?.slice(0, 5) || ''
})

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
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.loading-state,
.error-state {
  text-align: center;
  padding: 40px;
  font-size: 1.1rem;
  color: var(--color-text);
}

.error-state p {
  margin-bottom: 1rem;
  color: #dc2626;
}

.receipts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
  padding: 20px 0;
}

.receipt-card {
  background: var(--color-background-soft);
  border: 1px solid var(--color-border);
  border-radius: 8px;
  overflow: hidden;
  transition: transform 0.2s ease-in-out;
}

.receipt-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.receipt-content {
  padding: 15px;
}

.receipt-title {
  margin: 0 0 10px 0;
  font-size: 1.2rem;
  color: var(--color-heading);
  border-bottom: 1px solid var(--color-border);
  padding-bottom: 8px;
}

.receipt-details p {
  margin: 8px 0;
  font-size: 0.9rem;
}

.receipt-image {
  margin: 10px 0;
  text-align: center;
}

.receipt-image img {
  max-width: 100%;
  height: auto;
  border-radius: 4px;
}

.receipt-actions {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-top: 10px;
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

.add-receipt-container {
  text-align: right;
  margin-bottom: 20px;
}

button {
  padding: 8px 16px;
  border-radius: 4px;
  border: none;
  cursor: pointer;
  font-size: 0.9rem;
}

.add-receipt-button {
  background-color: #2563eb;
  color: white;
}

.edit-button {
  background-color: #4ade80;
  color: white;
}

.delete-button {
  background-color: #dc2626;
  color: white;
}

.save-button {
  background-color: #4ade80;
  color: white;
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

@media (max-width: 768px) {
  .receipts-grid {
    grid-template-columns: 1fr;
  }
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
</style>
